package config

import (
    "flag"
    "github.com/ilyakaznacheev/cleanenv"
    "log"
    "os"
    "time"
)

type Config struct {
    Env        string     `yaml:"env" env-required:"true"`
    Threshold  float64    `yaml:"threshold" env-required:"true"`
    Storage    Storage    `yaml:"storage" env-required:"true"`
    HTTPServer HTTPServer `yaml:"http_server" env-required:"true"`
}

type Storage struct {
    Host     string `yaml:"host" env-required:"true"`
    Port     int    `yaml:"port" env-required:"true"`
    User     string `yaml:"user" env-required:"true"`
    Password string `yaml:"password" env-required:"true"`
    DBName   string `yaml:"dbname" env-required:"true"`
    SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

type HTTPServer struct {
    Address     string        `yaml:"address" env-required:"true"`
    Timeout     time.Duration `yaml:"timeout" env-required:"true"`
    IdleTimeout time.Duration `yaml:"idle_timeout" env-required:"true"`
}

func MustLoad() *Config {
    path := fetchConfigPath()
    if path == "" {
        log.Fatalln("no config path provided, set CONFIG_PATH or use --config")
    }
    
    if _, err := os.Stat(path); os.IsNotExist(err) {
        log.Fatalln("file does not exist:", path)
    }
    
    var cfg Config
    if err := cleanenv.ReadConfig(path, &cfg); err != nil {
        log.Fatalln(err)
    }
    
    return &cfg
}

func fetchConfigPath() string {
    var path string
    flag.StringVar(&path, "config", "", "path to config *.yaml")
    flag.Parse()
    
    if path == "" {
        path = os.Getenv("CONFIG_PATH")
    }
    
    return path
}
