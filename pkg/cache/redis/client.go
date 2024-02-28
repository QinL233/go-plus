package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

type Client struct {
	Prefix string
}

func (p *Client) Set(key string, value interface{}, second int) error {
	c := driver.Get()
	defer c.Close()
	switch value.(type) {
	case string:
		_, err := c.Do("SET", p.Prefix+key, value)
		if err != nil {
			return err
		}
	default:
		bytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		_, err = c.Do("SET", p.Prefix+key, bytes)
		if err != nil {
			return err
		}
	}
	_, err := c.Do("EXPIRE", p.Prefix+key, second)
	if err != nil {
		return err
	}
	return nil
}

func (p *Client) Exists(key string) bool {
	c := driver.Get()
	defer c.Close()

	exists, err := redis.Bool(c.Do("EXISTS", p.Prefix+key))
	if err != nil {
		return false
	}

	return exists
}

func (p *Client) Get(key string) ([]byte, error) {
	c := driver.Get()
	defer c.Close()

	reply, err := redis.Bytes(c.Do("GET", p.Prefix+key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (p *Client) GetString(key string) (string, error) {
	c := driver.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("GET", p.Prefix+key))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func (p *Client) GetObject(key string, obj any) error {
	c := driver.Get()
	defer c.Close()

	reply, err := redis.Bytes(c.Do("GET", p.Prefix+key))
	if err != nil {
		return err
	}
	if reply == nil || len(reply) < 1 {
		return nil
	}
	err = json.Unmarshal(reply, &obj)
	if err != nil {
		return err
	}
	return nil
}

func (p *Client) Delete(key string) (bool, error) {
	c := driver.Get()
	defer c.Close()

	return redis.Bool(c.Do("DEL", p.Prefix+key))
}
