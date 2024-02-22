package util

const REDIS_KEY_TO_GET_UNIQUE_ID = "counter"
const REDIS_KEY_TO_STORE_SHORT_URL = "short-url-details"
const REDIS_KEY_TO_STORE_SHORT_URL_ID = "short-url-index"
const REDIS_KEY_TO_STORE_LONG_URL_ID = "long-url-index"

// default redis connection url.
const REDIS_CONNECTION_URL = "localhost:6379"

// leave empty if no authorization is set
// const REDIS_DB_PASSKEY = "redis1234"

const REDIS_DB_PASSKEY = ""

// default redis database to connect
// redis database index from 0 to 15
const DEFAULT_REDIS_DATABASE = 0

// default pool size by redis per every available CPU
const DEFAULT_POOL_SIZE = 10
