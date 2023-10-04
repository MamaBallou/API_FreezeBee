package utils

import (
    "net/http"
)

func ExtractApiKey(r *http.Request) string {
    headerUserSession := r.Header.Get("ApiKey")
    return headerUserSession
}