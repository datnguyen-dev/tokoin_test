package config

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/caarlos0/env"
	"github.com/hashicorp/hcl"
)

// Config is a configuration struct.
type Config struct {
	JSONDb struct {
		OrgPath    string `hcl:"org_path"`
		TicketPath string `hcl:"ticket_path"`
		UserPath   string `hcl:"user_path"`
	} `hcl:"jsondb"`
}

type envConfig struct {
	OrgPath    string `env:"ORG_PATH"`
	TicketPath string `env:"TICKET_PATH"`
	UserPath   string `env:"USER_PATH"`
}

func readFile(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer f.Close()

	cfg, err := read(f)
	if err == nil {
		readEnvConfig(cfg)
	}
	return cfg, err
}

func read(r io.Reader) (*Config, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}

	cfg := &Config{}
	err = hcl.Unmarshal(data, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshal hcl: %v", err)
	}

	return cfg, nil
}

//InitConfig - init config file
func InitConfig(configFile string) (*Config, error) {
	if _, err := os.Stat(configFile); !os.IsNotExist(err) {
		res, err := readFile(configFile)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	cfg, err := initfile()
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(configFile, []byte(cfg), 0666)
	if err != nil {
		return nil, err
	}

	res, err := readFile(configFile)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func initfile() (string, error) {
	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, map[string]interface{}{
		"org_path":  "./resources/organizations.json",
		"tick_path": "./resources/tickets.json",
		"user_path": "./resources/users.json",
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func readEnvConfig(cfg *Config) {
	envCfg := envConfig{}
	err := env.Parse(&envCfg)
	if err != nil {
		return
	}

	// JSONDb
	{
		if envCfg.OrgPath != "" {
			cfg.JSONDb.OrgPath = envCfg.OrgPath
		}
		if envCfg.TicketPath != "" {
			cfg.JSONDb.TicketPath = envCfg.TicketPath
		}
		if envCfg.UserPath != "" {
			cfg.JSONDb.UserPath = envCfg.UserPath
		}
	}
}

var tpl = template.Must(template.New("initial-config").Parse(strings.TrimSpace(`
jsondb {
  org_path    = "{{.org_path}}"
	ticket_path = "{{.tick_path}}"
	user_path = "{{.user_path}}"
}
`)))
