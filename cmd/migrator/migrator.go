package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"strings"

	// drivers
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// USAGE:
// --config=/path/to/config.yaml
// inside config.yaml:
// Storage: host, port, user, password, dbname, sslmode
// Migrations: path
func main() {
	var op string
	flag.StringVar(&op, "operation", "", "operation: up or down")

	cfg := MustLoadMigration()
	pSt := cfg.Storage
	pSt.DBName = "postgres"
	dsn := DataSourceName(pSt)
	link := Link(cfg.Storage)
	err := retry(30, func() error { return EnsureDBexists(cfg.Storage.DBName, dsn) })

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.New("file://"+cfg.Migrations.Path, link)
	if err != nil {
		panic(err)
	}
	switch {
	case op == "" || op == "up":
		if err = m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Println("Nothing to migrate")
				return
			}
			panic(err)
		}
	case op == "down":
		if err = m.Force(1); err != nil {
			panic(err)
		}
		if err = m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Println("Nothing to migrate")
				return
			}
			panic(err)
		}
	default:
		log.Fatalln("Unknown operation:", op)
	}

	log.Println("migration successful")
}

func MustLoadMigration() *MigrationConfig {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}
	return MustLoadMigrationByPath(path)
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res != "" {
		return res
	}
	env := os.Getenv("CONFIG_PATH")
	return env
}

func Link(cfg Storage) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
}

func DataSourceName(cfg Storage) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
}

func EnsureDBexists(dbname, dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	_, err = db.Exec("CREATE DATABASE" + " " + dbname)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}
	return nil
}

type Storage struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	SSLMode  string `yaml:"sslmode" env-default:"enable"`
}

type MigrationConfig struct {
	Storage    Storage    `yaml:"storage" env-required:"true"`
	Migrations Migrations `yaml:"migrations" env-required:"true"`
}

type Migrations struct {
	Path string `yaml:"path" env-required:"true"`
}

func MustLoadMigrationByPath(path string) *MigrationConfig {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file doesn't exist " + path)
	}
	var res MigrationConfig
	if err := cleanenv.ReadConfig(path, &res); err != nil {
		panic(err)
	}
	return &res
}
