package redis

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool *redis.Pool
}

func New(addr string) (*Redis, error) {
	if addr == "" {
		addr = ":6379"
	}
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}

	return &Redis{
		pool: pool,
	}, nil
}

func (r *Redis) Close() {
	r.pool.Close()
}

func (r *Redis) Set(ctx context.Context, show string, data []byte) error {
	// Retrieve redis connection from the pool, return when finished.
	var c redis.Conn
	c, err := r.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = redis.String(c.Do("SET", show, data))
	return err
}

func (r *Redis) Get(ctx context.Context, name string) ([]byte, error) {
	var c redis.Conn
	c, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}

	defer c.Close()

	resp, _ := redis.Bytes(c.Do("GET", name))
	return resp, nil
}
