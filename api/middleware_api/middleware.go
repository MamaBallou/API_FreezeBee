package middleware_api

import (
	"freezebee/api/utils"
	"freezebee/config"
	"net/http"
)

var secretKey = config.ApiKey

var apiSecretKey = config.ApiKey

var database = ""

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        apiKey := utils.ExtractApiKey(r)

		if apiKey == ""{
			utils.ToJson(w, "Api Key missing", http.StatusNotAcceptable)
            return
		} else if apiKey != apiSecretKey{
			utils.ToJson(w, "Api Key incorrect", http.StatusUnauthorized)
            return
		}

		jwt := utils.ExtractBearerToken(r)

		if jwt == ""{
			utils.ToJson(w, "JWT missing", http.StatusNotAcceptable)
            return
		}

		// Extrait les revendications JWT de la requête
		claims, err := utils.JwtExtract(r)
		if err != nil {
			utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		if role, ok := claims["user_role"]; ok {
			database = role.(string)
		}

        // Si l'accès est autorisé, passez la requête au gestionnaire de route suivant.
        next.ServeHTTP(w, r)
    })
}

func GetDatabase() string {
	return database
}