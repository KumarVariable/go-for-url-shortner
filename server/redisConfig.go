package server

// default redis connection url.
const REDIS_CONNECTION_URL = "localhost:6379"

// leave empty if no authorization is set
const REDIS_DB_PASSKEY = "redis1234"

// default redis database to connect
// redis database index from 0 to 15
const DEFAULT_REDIS_DATABASE = 0

// default pool size by redis per every available CPU
const DEFAULT_POOL_SIZE = 10

// Represents Redis Configuration
type RedisConfig struct {
	Address  string
	Password string
	Database int
	PoolSize int
}

// constructor to get fresh instance of RedisConfig struct
// with pre-determined values
func GetRedisConfig() *RedisConfig {

	return &RedisConfig{
		Address:  REDIS_CONNECTION_URL, // default redis connection url.
		Password: REDIS_DB_PASSKEY,     // leave empty if no authorization is set
		Database: DEFAULT_REDIS_DATABASE,
		PoolSize: DEFAULT_POOL_SIZE,
	}

}
