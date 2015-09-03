package sessionstores

import (
	"errors"
	"log"

	"github.com/samora/ussd-go/Godeps/_workspace/src/github.com/fzzy/radix/redis"
)

// Redis session store. See http://redis.io
type Redis struct {
	address, password string
	client            *redis.Client
}

// NewRedis creates a new pointer to Redis instance.
// First argument is Redis address.
// Second argument is optional Redis password.
func NewRedis(args ...string) *Redis {
	if len(args) == 0 {
		log.Panicln(errors.New("Address is required"))
	}
	address, password := args[0], ""
	if len(args) > 1 {
		password = args[1]
	}
	return &Redis{address: address, password: password}
}

func (s *Redis) Connect() error {
	client, err := redis.Dial("tcp", s.address)
	if err != nil {
		return err
	}
	if s.password != "" {
		err = client.Cmd("AUTH", s.password).Err
		if err != nil {
			return err
		}
	}
	s.client = client
	return err
}

func (s *Redis) SetValue(key, value string) error {
	return s.client.Cmd("SET", key, value).Err
}

func (s *Redis) GetValue(key string) (string, error) {
	return s.client.Cmd("GET", key).Str()
}

func (s *Redis) ValueExists(key string) (bool, error) {
	return s.client.Cmd("EXISTS", key).Bool()
}

func (s *Redis) DeleteValue(key string) error {
	return s.client.Cmd("DEL", key).Err
}

func (s *Redis) HashSetValue(name, key, value string) error {
	return s.client.Cmd("HSET", name, key, value).Err
}

func (s *Redis) HashGetValue(name, key string) (string, error) {
	return s.client.Cmd("HGET", name, key).Str()
}

func (s *Redis) HashValueExists(name, key string) (bool, error) {
	return s.client.Cmd("HEXISTS", name, key).Bool()
}

func (s *Redis) HashDeleteValue(name, key string) error {
	return s.client.Cmd("HDEL", name, key).Err
}

func (s *Redis) HashExists(name string) (bool, error) {
	return s.ValueExists(name)
}

func (s *Redis) HashDelete(name string) error {
	return s.DeleteValue(name)
}

func (s *Redis) Close() error {
	return s.client.Close()
}
