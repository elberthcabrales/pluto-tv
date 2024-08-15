package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_WithEnvVars(t *testing.T) {
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("REDIS_PASSWORD", "supersecret")
	os.Setenv("REDIS_DB", "2")

	config := LoadConfig()

	assert.Equal(t, "localhost:6379", config.Addr, "Expected Redis Addr to be localhost:6379")
	assert.Equal(t, "supersecret", config.Password, "Expected Redis Password to be supersecret")
	assert.Equal(t, 2, config.DB, "Expected Redis DB to be 2")

	os.Unsetenv("REDIS_ADDR")
	os.Unsetenv("REDIS_PASSWORD")
	os.Unsetenv("REDIS_DB")
}

func TestLoadConfig_WithDefaultValues(t *testing.T) {
	// Ensure no environment variables are set
	os.Unsetenv("REDIS_ADDR")
	os.Unsetenv("REDIS_PASSWORD")
	os.Unsetenv("REDIS_DB")

	// Load the config
	config := LoadConfig()

	// Assertions
	assert.Equal(t, "localhost:6379", config.Addr, "Expected default Redis Addr to be localhost:6379")
	assert.Equal(t, "", config.Password, "Expected default Redis Password to be an empty string")
	assert.Equal(t, 0, config.DB, "Expected default Redis DB to be 0")
}

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_ENV", "somevalue")

	result := getEnv("TEST_ENV", "defaultvalue")
	assert.Equal(t, "somevalue", result, "Expected getEnv to return the value of TEST_ENV")

	result = getEnv("NON_EXISTENT_ENV", "defaultvalue")
	assert.Equal(t, "defaultvalue", result, "Expected getEnv to return the default value when the env var is not set")

	os.Unsetenv("TEST_ENV")
}

func TestGetEnvAsInt(t *testing.T) {
	os.Setenv("TEST_INT_ENV", "42")

	result := getEnvAsInt("TEST_INT_ENV", 10)
	assert.Equal(t, 42, result, "Expected getEnvAsInt to return the value of TEST_INT_ENV as an integer")

	result = getEnvAsInt("NON_EXISTENT_INT_ENV", 10)
	assert.Equal(t, 10, result, "Expected getEnvAsInt to return the default value when the env var is not set")

	os.Setenv("TEST_INVALID_INT_ENV", "invalid")
	result = getEnvAsInt("TEST_INVALID_INT_ENV", 10)
	assert.Equal(t, 10, result, "Expected getEnvAsInt to return the default value when the env var is not an integer")

	os.Unsetenv("TEST_INT_ENV")
	os.Unsetenv("TEST_INVALID_INT_ENV")
}
