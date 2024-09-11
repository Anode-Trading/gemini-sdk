package api

import (
	"net/http"
	"time"
)

func NewHttpClient(timeout int) http.Client {
	return http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}
