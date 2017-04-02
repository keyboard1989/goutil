package redis

import "github.com/garyburd/redigo/redis"

type Redis struct {
	conn redis.Conn
}

func NewRedis(address string) (Redis, error) {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		return Redis{}, nil
	}

	return Redis{
		conn: conn,
	}, nil
}

func (r *Redis) Get(key string) (string, error) {
	ret, err := r.conn.Do("get", key)
	if err != nil {
		return "", err
	}

	temp, ok := ret.([]byte)
	if ok {
		return string(temp), nil
	}
	return "", nil

}

func (r *Redis) Set(key string, value string) error {
	_, err := r.conn.Do("set", key, value)
	return err
}

func (r *Redis) Close() error {
	return r.conn.Close()
}
