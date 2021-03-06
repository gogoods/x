package redis

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"

	"errors"
)

type RedisProxy struct {
	Addr     string
	Password string
	DB       int
	Pool     *redis.Pool
}

func New(addr, password string, maxIdle, maxActive int, idleTimeout time.Duration, db int) (*RedisProxy, error) {
	if addr == "" {
		return nil, errors.New("Invalid parameters 'addr'.")
	}

	r := RedisModel{
		Addr:     addr,
		Password: password,
		DB:       db,
	}
	r.Pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			//			c, err := redis.Dial("tcp", redisConfig.Addr)
			c, err := redis.Dial("tcp", addr, redis.DialPassword(password), redis.DialDatabase(db))

			if err != nil {
				log.Println("[ERROR] Dial redis fail", err)
				return nil, err
			}
			//log.Println("Dial redis succ", addr)
			return c, err
		},
		TestOnBorrow: PingRedis,
	}

	return &r, nil
}

func (r RedisProxy) Close() {
	if r.Pool != nil {
		r.Pool.Close()
		r.Pool = nil
	}
}

func PingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("ping")

	if err != nil {
		log.Println("[ERROR] ping redis fail", err)
	}
	return err
}

func (r RedisProxy) GetConn() redis.Conn {
	return r.Pool.Get()
}

func (r RedisProxy) PubSubConn() redis.PubSubConn {
	return redis.PubSubConn{r.Pool.Get()}
}

func (r RedisProxy) Publish(subject string, content interface{}) (reply interface{}, err error) {
	rc := r.Pool.Get()
	defer rc.Close()

	reply, err = rc.Do("PUBLISH", subject, content)
	return
}
