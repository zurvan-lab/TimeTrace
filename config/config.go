package config

import (
	"bytes"
	_ "embed"
	"os"

	tte "github.com/zurvan-lab/TimeTrace/utils/errors"
	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var configBytes []byte

type Config struct {
	Name   string `yaml:"name"`
	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
	Users  []User `yaml:"users"`
}

type Server struct {
	IP   string `yaml:"listen"`
	Port string `yaml:"port"`
}

type Log struct {
	Path       string `yaml:"path"`
	Colorful   bool   `yaml:"colorful"`
	Compress   bool   `yaml:"compress"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
	MaxLogSize int    `yaml:"max_log_size"`
}

type User struct {
	Name     string   `yaml:"name"`
	Password string   `yaml:"password"`
	Cmds     []string `yaml:"cmds"`
}

func (conf *Config) BasicCheck() error {
	if len(conf.Users) == 0 {
		return tte.ErrInvalidUsers
	}

	for _, u := range conf.Users {
		allCmds := false

		for _, c := range u.Cmds {
			if c == "*" {
				allCmds = true
			}
		}

		if allCmds && len(u.Cmds) > 1 {
			return tte.ErrSpecificAndAllCommandSameAtTime
		}
	}

	return nil
}

func DefaultConfig() *Config {
	config := &Config{
		Server: Server{
			IP:   "localhost",
			Port: "7070",
		},
		Log: Log{
			Colorful:   true,
			Compress:   true,
			MaxAge:     1,
			MaxBackups: 10,
			MaxLogSize: 10,
			Path:       "log.ttrace",
		},
		Name: "time_trace",
	}

	rootUser := User{
		Name:     "root",
		Password: "super_secret_password",
		Cmds:     []string{"*"},
	}
	config.Users = append(config.Users, rootUser)

	return config
}

func LoadFromFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	err = config.BasicCheck()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (conf *Config) ToYAML() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := yaml.NewEncoder(buf)

	err := encoder.Encode(conf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
