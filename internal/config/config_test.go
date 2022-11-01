package config_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tonnytg/tasklist/internal/config"
	"io/ioutil"
	"os"
	"testing"
)

// Create JSON file with config
func BuildConfig() {
	f := `{
    "name": "test",
    "version": "1.0.0",
    "description": "test",
    "database": {
        "type": "sqlite3",
        "table": "tasks",
        "name": "database_test.sqlite",
        "path": "."
    }
}`
	_ = ioutil.WriteFile("config.json", []byte(f), 0644)
}

// Delete JSON file with config
func DeleteConfig() {
	_ = os.Remove("config.json")
}

func TestConfigReader(t *testing.T) {
	BuildConfig()
	c := config.Reader()
	assert.Equal(t, "test", c.Name)
	assert.Equal(t, "1.0.0", c.Version)
	assert.Equal(t, "test", c.Description)
	assert.Equal(t, "sqlite3", c.Database.Type)
	assert.Equal(t, "tasks", c.Database.Table)
	assert.Equal(t, "database_test.sqlite", c.Database.Name)
	assert.Equal(t, ".", c.Database.Path)

	DeleteConfig()
}
