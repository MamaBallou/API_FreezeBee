package api

import (
    "net/http"
    "log"
    "fmt"
    "freezebee/api/routes"
    "freezebee/api/database"
)

// Fonction principale pour démarrer le serveur API.
func Run() {
    InitConnexion()
    // Lance le serveur en écoutant sur le port 8080.
    listen(9200)
}

// Fonction pour démarrer l'écoute du serveur sur un port spécifié.
func listen(p int) {
    port := fmt.Sprintf(":%d", p)
    fmt.Printf("Listening Port %s...\n", port)
    // Crée un routeur pour gérer les différentes routes de l'API.
    r := routes.NewRouter()
    // Lance le serveur en écoutant sur le port spécifié, avec gestion de la stratégie CORS.
    log.Fatal(http.ListenAndServe(port, routes.LoadCors(r)))
}

func InitConnexion(){
    dbList := []string{"prod", "test", "r&d"}
    for i := range dbList {
        db, err := database.InitConnectSQLServer(dbList[i])
        if err != nil {
            panic(err)
        }
        database.AutoMigrateTables(db)
        database.CloseSQLServer(db)
    }
}