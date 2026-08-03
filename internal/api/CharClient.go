package api

import (
	"net/http"
)

// Temporary Client Struct, will move to internal API package later
type Client struct {
        HttpClient      http.Client
}
