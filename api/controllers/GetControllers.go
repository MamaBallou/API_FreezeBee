package controllers

import (
	"encoding/json"
	"freezebee/api/models"
	"freezebee/api/utils"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
    body := utils.BodyParser(r)
    
    // Structure pour stocker le JSON désérialisé
    var data map[string]interface{}
    
    err := json.Unmarshal(body, &data)
    if err != nil {
        // La désérialisation a échoué, renvoyez une réponse d'erreur
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }
    
    // Vérifiez la valeur du champ "Table"
    table, ok := data["Table"].(string)
    if !ok {
        // Le champ "Table" n'est pas une chaîne, renvoyez une réponse d'erreur
        utils.ToJson(w, "Missing or invalid 'Table' field", http.StatusUnprocessableEntity)
        return
    }
    
    // Désérialisez le JSON en utilisant la structure appropriée en fonction de la valeur de "Table"
    switch table {
    case "Ingrédient":
        // La désérialisation a réussi pour "Ingrédient"

        ingredients, err := models.GetIngredients()
        if err != nil {
            utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
            return
        }
        // Répondre avec les newsletters en tant que JSON
        utils.ToJson(w, ingredients, http.StatusOK)
    case "Modèle":
        // La désérialisation a réussi pour "Modèle"
        models, err := models.GetAllModels()
        if err != nil {
            utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
            return
        }
        utils.ToJson(w, models, http.StatusOK)
        return
    case "Fabrication":
        var fabData models.Post_Data_Fab
        errFab := json.Unmarshal(body, &fabData)
        if errFab == nil {
            // La désérialisation a réussi pour "Fabrication"		
			fabrication, err := models.GetAllFabrications()
            if err != nil {
                utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
                return
            }
            utils.ToJson(w, fabrication, http.StatusOK)
        }
    default:
        // Valeur inconnue pour "Table", renvoyez une réponse d'erreur
		utils.ToJson(w, "Wrong data", http.StatusUnprocessableEntity)
        return
    }
}


