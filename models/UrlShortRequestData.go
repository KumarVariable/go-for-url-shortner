package models

import "github.com/redis/go-redis/v9"

// Represents payload for url shortner
type UrlShortnerData struct {
	Key      string
	ShortUrl string
	LongUrl  string
}

// Represent payload for Short Url request
type ShortUrlRequest struct {
	ShortUrl string `json:"shortUrl"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

var RedisClient *redis.Client

// Represents data transfer for URL shortner
type Payload struct {
	KeyId    string `json:"keyId"`
	ShortUrl string `json:"shortUrl"`
	LongUrl  string `json:"longUrl"`
	Message  string `json:"message"`
}
