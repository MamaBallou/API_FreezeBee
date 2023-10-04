package routes

import (
    "net/http"
    "github.com/gorilla/handlers"
)

// Charge la gestion des en-têtes CORS pour une route HTTP donnée.
func LoadCors(r http.Handler) http.Handler {
    // Définit les en-têtes autorisés.
    headers := handlers.AllowedHeaders([]string{"Origin", "Accept", "Content-Type", "Authorization"})
    // Définit les méthodes HTTP autorisées.
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"})
    // Définit les origines autorisées (ici, n'importe quelle origine).
    origins := handlers.AllowedOrigins([]string{"*"})
    // Autorise l'envoi de cookies et d'informations d'authentification.
    credentials := handlers.AllowCredentials()
    // Définit le temps de mise en cache des pré-vérifications CORS.
    age := handlers.MaxAge(300)
    
    // Applique les paramètres CORS à la route donnée et renvoie le gestionnaire HTTP modifié.
    return handlers.CORS(headers, methods, origins, credentials, age)(r)
}