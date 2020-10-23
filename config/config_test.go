package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ConfigFile = "test.conf"
)

func TestInitFile(t *testing.T) {
	defer teadown()
	cnf, err := InitConfig(ConfigFile)
	assert.True(t, err == nil)
	assert.NotEmpty(t, cnf.JSONDb.OrgPath)
	assert.NotEmpty(t, cnf.JSONDb.UserPath)
	assert.NotEmpty(t, cnf.JSONDb.TicketPath)
}

func teadown() {
	os.Remove(ConfigFile)
}
