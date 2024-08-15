package cache

import (
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"

	"github.com/elberthcabrales/movies-api/pkg/config"
)

func TestSetValue(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectSet("example_key", "example_value", 0).SetVal("OK")

	repo := &RedisCache{client: db}
	err := repo.SetValue("example_key", "example_value")

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetValue(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectGet("example_key").SetVal("example_value")

	repo := &RedisCache{client: db}
	value, err := repo.GetValue("example_key")

	assert.NoError(t, err)
	assert.Equal(t, "example_value", value)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetValue_NotFound(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectGet("missing_key").RedisNil()

	repo := &RedisCache{client: db}
	value, err := repo.GetValue("missing_key")

	assert.Error(t, err)
	assert.Equal(t, "", value)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestNewRedisCacheWithOutConfig(t *testing.T) {
	db, mock := redismock.NewClientMock()

	mock.ExpectPing().SetVal("PONG")

	cache := NewRedisCache(db, nil)

	_, err := cache.client.Ping(ctx).Result()

	assert.NoError(t, err)
	assert.NotNil(t, cache)
	assert.Equal(t, db, cache.client)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestNewRedisCacheWithConfig(t *testing.T) {
	cache := NewRedisCache(nil, &config.RedisConfig{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := cache.client.Ping(ctx).Result()

	assert.NoError(t, err)
	assert.NotNil(t, cache)
}
