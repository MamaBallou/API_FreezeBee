package utils

import (
    "net/http"
    "encoding/json"
    "io/ioutil"
)

// Lit le corps de la requête HTTP et renvoie ses données sous forme de bytes.
func BodyParser(r *http.Request) []byte {
    body, _ := ioutil.ReadAll(r.Body)
    return body
}

// Convertit les données en JSON et les envoie comme réponse HTTP avec un code d'état spécifié.
func ToJson(w http.ResponseWriter, data interface{}, statusCode int) {
    // Définit l'en-tête de la réponse pour le type de contenu JSON.
    w.Header().Set("Content-type", "application/json; charset=UTF8")
    // Définit le code d'état HTTP de la réponse.
    w.WriteHeader(statusCode)
    // Encode les données en JSON et les écrit dans la réponse.
    err := json.NewEncoder(w).Encode(data)
    CheckErr(err) // Gérer l'erreur ici
}