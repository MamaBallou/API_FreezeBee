package controllers

import (
	"encoding/json"
	"freezebee/api/models"
	"freezebee/api/utils"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
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
        var ingredientData models.Delete_Data
        errIng := json.Unmarshal(body, &ingredientData)
        if errIng == nil {
            // La désérialisation a réussi pour "Ingrédient"

            err := models.DeleteIngredient(ingredientData)
            if err != nil {
                utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
                return
            }
            // Répondre avec les newsletters en tant que JSON
            utils.ToJson(w, "Ingredient Deleted", http.StatusOK)
        }
    case "Modèle":
        var modData models.Delete_Data_Modele
        errMod := json.Unmarshal(body, &modData)
        if errMod == nil {
            // La désérialisation a réussi pour "Modèle"
			err := models.DeleteModele(modData)
            if err != nil {
                utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
                return
            }
            // Répondre avec les newsletters en tant que JSON
            utils.ToJson(w, "Modele Deleted", http.StatusOK)
            return
        }
    case "Fabrication":
        var fabData models.Delete_Data
        errFab := json.Unmarshal(body, &fabData)
        if errFab == nil {
            // La désérialisation a réussi pour "Fabrication"	
			err := models.DeleteFabrication(fabData)
            if err != nil {
                utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
                return
            }
            // Répondre avec les newsletters en tant que JSON
            utils.ToJson(w, "Fabrication Deleted", http.StatusOK)
        }
    default:
        // Valeur inconnue pour "Table", renvoyez une réponse d'erreur
		utils.ToJson(w, "Wrong data", http.StatusUnprocessableEntity)
        return
    }
}