package api

import (
	"net/http"
	"time"

	
)

// Temporary Client Struct, will move to internal API package later
type Client struct {
        HttpClient      http.Client
        Cache			*CharDataCache.Cache
}

func NewClient() *Client{
	return &Client{
		httpClient: http.Client{},
		Cache: CharDataCache.Cache(5 * time.Second)
	}
} 
