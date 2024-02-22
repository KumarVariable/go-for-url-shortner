package models

import "github.com/redis/go-redis/v9"

// variable to hold redis client instance.
// not a good approach
var RedisClient *redis.Client

// Represents data transfer for URL shortner
type Payload struct {
	KeyId      int64  `json:"keyId"`      // unique id - int counter maintained in redis database
	ShortUrl   string `json:"shortUrl"`   // short url created for long url
	ShortUrlId string `json:"shortUrlId"` // unique short url id
	LongUrl    string `json:"longUrl"`    // long url which needs to be shortened
	Message    string `json:"message"`    // a string message
	KeyName    string `json:"keyName"`    // name of key, if available or maintained
}
