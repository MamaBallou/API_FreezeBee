package controllers

import (
	"encoding/json"
	"freezebee/api/models"
	"freezebee/api/utils"
	"net/http"
)

func Patch(w http.ResponseWriter, r *http.Request) {
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
        var ingredientData models.Patch_Data_Ing
        errIng := json.Unmarshal(body, &ingredientData)
        if errIng == nil {
            // La désérialisation a réussi pour "Ingrédient"
            err = models.PatchIngredient(ingredientData)
			if err != nil {
				utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}
			utils.ToJson(w, "Ingredient Updated", http.StatusOK)
        }
    case "Modèle":
        var modData models.Patch_Data_Mod
        errMod := json.Unmarshal(body, &modData)
        if errMod == nil {
            // La désérialisation a réussi pour "Ingrédient"
            err = models.PatchModele(modData)
			if err != nil {
				utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}
			utils.ToJson(w, "Modele Updated", http.StatusOK)
        }
    case "Fabrication":
        var fabData models.Patch_Data_Fab
        errFab := json.Unmarshal(body, &fabData)
        if errFab == nil {
            // La désérialisation a réussi pour "Fabrication"		
			err = models.PatchFabrication(fabData)
			if err != nil {
				utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}
			utils.ToJson(w, "Fabrication Updated", http.StatusOK)
        }
    default:
        // Valeur inconnue pour "Table", renvoyez une réponse d'erreur
		utils.ToJson(w, "Wrong data", http.StatusUnprocessableEntity)
        return
    }
}


