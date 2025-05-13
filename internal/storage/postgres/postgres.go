package postgres

import (
    "database/sql"
    "errors"
    "fmt"
    "github.com/kxddry/random-history-facts/internal/config"
    "github.com/kxddry/random-history-facts/internal/http-server/handlers/post"
    "github.com/kxddry/random-history-facts/internal/lib/factmatcher"
    "github.com/kxddry/random-history-facts/internal/storage"
    _ "github.com/lib/pq"
)

type Storage struct {
    threshold float64
    db        *sql.DB
}

type FactMatcher interface {
    post.FactMatcher
}

func New(cfg *config.Config) (*Storage, error) {
    const op = "storage.postgres.New"
    
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        cfg.Storage.Host,
        cfg.Storage.Port,
        cfg.Storage.User,
        cfg.Storage.Password,
        cfg.Storage.DBName,
        cfg.Storage.SSLMode)
    
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, fmt.Errorf("%s: %w", op, err)
    }
    return &Storage{db: db, threshold: cfg.Threshold}, db.Ping()
}

func (s *Storage) AddFact(matcher factmatcher.Fact_Matcher, fact string) (int64, error) {
    const op = "storage.postgres.AddFact"
    
    threshold := s.threshold
    normalized := matcher.Normalize(fact)
    
    var count int
    
    tx, err := s.db.Begin()
    defer tx.Rollback()
    
    if err != nil {
        return 0, fmt.Errorf("%s: %w", op, err)
    }
    
    err = tx.QueryRow(`SELECT COUNT(*) FROM facts WHERE similarity(normalized_fact, $1) > $2;`,
        normalized, threshold).Scan(&count)
    
    if err != nil {
        return 0, fmt.Errorf("%s: %w", op, err)
    }
    
    if count > 0 {
        return 0, storage.ErrFactAlreadyExists
    }
    
    var id int64
    err = tx.QueryRow(`INSERT INTO facts (fact, normalized_fact) VALUES ($1, $2) RETURNING id;`, fact, normalized).Scan(&id)
    
    if err != nil {
        return 0, fmt.Errorf("%s: %w", op, err)
    }
    
    return id, tx.Commit()
}

func (s *Storage) Fact() (string, error) {
    const op = "storage.postgres.RandomFact"
    query := `SELECT * FROM facts ORDER BY RANDOM() LIMIT 1;`
    
    var txt string
    
    err := s.db.QueryRow(query).Scan(&txt)
    
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return "", storage.ErrNoFacts
        }
        return "", err
    }
    
    return txt, nil
}

func (s *Storage) Close() error {
    return s.db.Close()
}
