package main
import (
    "net/http"
)
func Get (url string) (*http.Response, error) {
    resp, err := http.Get(url)
    return resp, err
}
