package cache

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/commons_conf"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"os"
	"sync"
	"time"
)

type RedisStore struct {
	config   *commons_conf.Data_Redis
	DbClient redis.UniversalClient
	ctx      *context.Context
}

var (
	rs   *RedisStore
	once sync.Once
)

type RedisCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewRedisStore(data *commons_conf.Data) *RedisStore {
	redisConfig := data.Redis
	fmt.Println("connecting to redis")

	once.Do(func() {
		credentials := ReadRedisCredential(redisConfig.CredFileLocation)
		if credentials == nil {
			err := errors.New(401, "connection error", "Error fetching redis credentials")
			panic(err)
		}

		addrs := []string{redisConfig.Addr}
		username := credentials.Username
		password := credentials.Password
		poolSize := redisConfig.PoolSize
		if poolSize == 0 {
			poolSize = 1
		}

		opts := &redis.ClusterOptions{
			Addrs:        addrs,
			Password:     password,
			PoolSize:     int(poolSize),
			ReadTimeout:  time.Duration(redisConfig.ReadTimeOutInMs) * time.Millisecond,
			DialTimeout:  time.Duration(redisConfig.DialTimeOutInMs) * time.Millisecond,
			WriteTimeout: time.Duration(redisConfig.WriteTimeOutInMs) * time.Millisecond,
		}

		if username != "" {
			opts.Username = username
		}

		if redisConfig.Tls {
			opts.TLSConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}

		rdb := redis.NewClusterClient(opts)
		var ctx = context.Background()
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			fmt.Printf("Error pinging Redis server: %v\n", err)
		} else {
			fmt.Println("successfully pinged redis")
		}
		rs = &RedisStore{
			config:   redisConfig,
			DbClient: rdb,
			ctx:      &ctx,
		}
	})
	fmt.Println("Connected to Redis")
	return rs
}

func ReadRedisCredential(fileName string) *RedisCredential {
	// read our opened jsonFile as a byte array.
	byteValue, err := os.ReadFile(fileName)

	if err != nil {
		log.Errorf("Error while reading the file for mongo creds")
		return nil
	}

	// we initialize our Users array
	var creds RedisCredential

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &creds)

	if err != nil {
		log.Errorf("Error while marshaling  the file for mongo creds")
		return nil
	}

	return &creds
}
