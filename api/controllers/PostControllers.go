package controllers

import (
	"encoding/json"
	"freezebee/api/models"
	"freezebee/api/utils"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
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
        var ingredientData models.Post_Data_Ing
        errIng := json.Unmarshal(body, &ingredientData)
        if errIng == nil {
            // La désérialisation a réussi pour "Ingrédient"
			existingIng, err := models.CheckIfIngredientDontExist(ingredientData.Data.Nom)
			if err != nil {
				utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}
			if existingIng.Nom != "" {
				// La newsletter avec ce nom, cet organisateur et cette date existe déjà, retourner une erreur ou un message approprié
				utils.ToJson(w, "This Ingredient is already existed.", http.StatusConflict)
				return
			}

            err = models.PostIngredient(ingredientData)
			if err != nil {
				utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}
			utils.ToJson(w, "Ingredient Created", http.StatusOK)
        }
    case "Modèle":
        var modData models.Post_Data_Mod
        errMod := json.Unmarshal(body, &modData)
        if errMod == nil {
            // La désérialisation a réussi pour "Modèle"
            err := models.CheckIfModeleExists(modData.Data.Nom)
            if err == nil{
                utils.ToJson(w, "This Modele is already existed.", http.StatusConflict)
				return
            }
            models.PostModele(modData)
			utils.ToJson(w, "Modele Created", http.StatusOK)
            return
        }
    case "Fabrication":
        var fabData models.Post_Data_Fab
        errFab := json.Unmarshal(body, &fabData)
        if errFab == nil {
            // La désérialisation a réussi pour "Modèle"
            err := models.CheckIfFabricationExists(fabData.Data.Nom)
            if err == nil{
                utils.ToJson(w, "This Fabrication is already existed.", http.StatusConflict)
				return
            }
            err = models.CheckIfModExists(fabData.Data.Id_Modele)
            if err != nil{
                utils.ToJson(w, "This Modele doesn't exist.", http.StatusConflict)
				return
            }
            models.PostFabrication(fabData)
			utils.ToJson(w, "Fabrication Created", http.StatusOK)
            return
        }
    default:
        // Valeur inconnue pour "Table", renvoyez une réponse d'erreur
		utils.ToJson(w, "Wrong data", http.StatusUnprocessableEntity)
        return
    }
}


