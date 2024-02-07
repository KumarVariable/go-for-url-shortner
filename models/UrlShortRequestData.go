package models

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
