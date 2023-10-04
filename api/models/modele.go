package models

import (
	"errors"
	"fmt"
	"freezebee/api/database"
	"math"
	"gorm.io/gorm"
)

type Post_Data_Mod struct {
	Table string   `gorm:"size:60;not null" json:"Table"`
	Data  Data_Mod `gorm:"size:1000;not null" json:"Data"`
}

type Patch_Data_Mod struct {
	Table string     `gorm:"size:60;not null" json:"Table"`
	Data  Data_Mod_2 `gorm:"size:1000;not null" json:"Data"`
}

type Data_Mod struct {
	Nom         string         `gorm:"size:60;not null" json:"Nom"`
	Description string         `gorm:"size:60;not null" json:"Description"`
	PUHT        float32        `gorm:"size:60;not null" json:"pUHT"`
	Gamme       string         `gorm:"size:60;not null" json:"Gamme"`
	Ingredient  []Data_Ing_Mod `gorm:"size:1000;not null" json:"Ingredient"`
}

type Data_Mod_2 struct {
	Id          int              `gorm:"size:10;not null" json:"Id"`
	Nom         string           `gorm:"size:60;not null" json:"Nom"`
	Description string           `gorm:"size:60;not null" json:"Description"`
	PUHT        float32          `gorm:"size:60;not null" json:"pUHT"`
	Gamme       string           `gorm:"size:60;not null" json:"Gamme"`
	Ingredient  []Data_Ing_Mod_2 `gorm:"size:1000;not null" json:"Ingredient"`
}

type Data_Ing_Mod struct {
	Id_Ingredient int     `gorm:"size:10;not null" json:"Id_Ingredient"`
	Grammage      float32 `gorm:"size:10;not null" json:"Grammage"`
	Data_ModID    uint
}

type Data_Ing_Mod_2 struct {
	Type          string  `gorm:"size:100;not null" json:"Type"`
	Id_Ingredient int     `gorm:"size:10;not null" json:"Id_Ingredient"`
	Grammage      float32 `gorm:"size:10;not null" json:"Grammage"`
	Data_ModID    uint
}


type Modele struct {
	Id          uint32  `gorm:"primary_key;auto_increment"`
	Nom         string  `gorm:"column:nom;type:nvarchar(max);not null"`
	Description string  `gorm:"column:description;type:nvarchar(max);not null"`
	PUHT        float32 `gorm:"column:puht;type:float;not null"`
	Gamme       string  `gorm:"column:gamme;type:nvarchar(max);not null"`
}

type Ingredient struct {
	// Vous devrez ajouter les champs suivants pour la jointure avec Ingredient
	Id_Ingredient          uint32  `gorm:"column:Id_Ingredient"`
	Nom_Ingredient         string  `gorm:"column:nom"`
	Description_Ingredient string  `gorm:"column:description"`
	Grammage               float32 `gorm:"column:grammage"`
}


func getIdAtCreation(nom string, db *gorm.DB) (uint32, error){
	var data database.Modele

	if err := db.Table("dbo.modeles").Where("Nom = ?", nom).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Gestion du cas où l'ingrédient n'est pas trouvé
			return 0, errors.New("Ingrédient non trouvé")
		}
		// Gestion d'autres erreurs éventuelles
		return 0, err
	}
	
	// Vous pouvez maintenant accéder à l'ID de l'ingrédient recherché
	id := data.Id

	return id, nil
}

func CheckIfModeleExists(nom string) error {
    db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }
    defer database.CloseSQLServer(db)

    var count int64

    // Comptez le nombre d'occurrences de l'ingrédient avec l'ID donné
    if err := db.Table("dbo.modeles").Where("nom = ?", nom).Count(&count).Error; err != nil {
        return err // Gérer l'erreur
    }

    if count == 0 {
        return errors.New("Le modele n'existe pas") // L'ingrédient n'existe pas
    }

    return nil 
}

func PostModele(dataMod Post_Data_Mod) error {
	db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }

	puht := dataMod.Data.PUHT

	// Arrondir le PUHT à deux décimales
	roundedPUHT := math.Round(float64(puht)*100) / 100
	// Convertir le résultat en float32
	puhtFloat32 := float32(roundedPUHT)

	modele := database.Modele{
		Nom: dataMod.Data.Nom,
		Description: dataMod.Data.Description,
		PUHT: puhtFloat32,
		Gamme: dataMod.Data.Gamme,
	}

	rs := db.Create(&modele)
	if rs.Error != nil {
        return errors.New(fmt.Sprintf("Failed to create modele:", rs.Error))
    }

	id, err := getIdAtCreation(dataMod.Data.Nom, db)

	for _, ingredient := range dataMod.Data.Ingredient {
		// Vérifiez si l'ingrédient existe déjà
		err := CheckIfIngredientExists(ingredient.Id_Ingredient)
		if err != nil{
			continue
		}
	
		// Créez une nouvelle instance de Ingredient_Modele
		newIngredient := database.Ingredient_Modele{
			ModeleID:      id, 
			IngredientID:  uint32(ingredient.Id_Ingredient),
			Grammage:      ingredient.Grammage,
		}
	
		// Insérez cette instance dans la base de données
		if err := db.Create(&newIngredient).Error; err != nil {
			return err
		}
	}

	database.CloseSQLServer(db)

	return nil
}

func GetAllModels() ([]map[string]interface{}, error) {
    db, err := database.ConnectSQLServer()
    if err != nil {
        return nil, err
    }
    defer database.CloseSQLServer(db)

    var data []map[string]interface{}

    // Sélectionnez tous les modèles de dbo.modeles
    var modeles []Modele
    result := db.Table("dbo.modeles").
        Select("dbo.modeles.id, dbo.modeles.nom, dbo.modeles.description, dbo.modeles.puht, dbo.modeles.gamme").
        Find(&modeles)

    if result.Error != nil {
        return nil, result.Error
    }

    // Parcourez les modèles et récupérez les ingrédients associés
    for _, modele := range modeles {
        modelData := map[string]interface{}{
            "Id":          modele.Id,
            "Nom":         modele.Nom,
            "Description": modele.Description,
            "pUHT":        modele.PUHT,
            "Gamme":       modele.Gamme,
            "Ingredient": []map[string]interface{}{},
        }

        // Sélectionnez les ingrédients associés à ce modèle
        var ingredients []Ingredient
        result := db.Table("dbo.ingredient_modeles").
            Select("dbo.ingredient_modeles.grammage, dbo.ingredients.id as Id_Ingredient, dbo.ingredients.nom, dbo.ingredients.description").
            Joins("INNER JOIN dbo.ingredients ON dbo.ingredient_modeles.ingredient_id = dbo.ingredients.id").
            Where("dbo.ingredient_modeles.modele_id = ?", modele.Id).
            Find(&ingredients)

        if result.Error != nil {
            return nil, result.Error
        }

        // Ajoutez les ingrédients au modèle
        for _, ingredient := range ingredients {
            ingredientData := map[string]interface{}{
                "Id_Ingredient": ingredient.Id_Ingredient,
                "Nom":           ingredient.Nom_Ingredient,
                "Description":   ingredient.Description_Ingredient,
                "Grammage":      ingredient.Grammage,
            }
            modelData["Ingredient"] = append(modelData["Ingredient"].([]map[string]interface{}), ingredientData)
        }

        // Ajoutez le modèle avec ses ingrédients à la liste de données
        data = append(data, modelData)
    }

    return data, nil
}


func DeleteModele(dataMod Delete_Data_Modele) error {
	db, err := database.ConnectSQLServer()
	if err != nil {
		return errors.New(err.Error())
	}

	if dataMod.Id != dataMod.Id_Modele{
		return errors.New(fmt.Sprintf("ids not matching", err))
	}

	var count int64
    if err := db.Table("dbo.modeles").Where("id = ?", dataMod.Id).Count(&count).Error; err != nil {
        return errors.New(fmt.Sprintf("Failed to check id in table dbo.modeles:", err))
    }

    if count == 0 {
        return errors.New("Id not found in dbo.modeles")
    }
		
	rs := db.Table("dbo.modeles").Where("id = ?", dataMod.Id).Delete(&dataMod.Id)
	if rs.Error != nil {
		return errors.New(fmt.Sprintf("Failed to delete modele : ", rs.Error))
	}

	rm := db.Table("dbo.ingredient_modeles").Where("modele_id = ?", dataMod.Id_Modele).Delete(&dataMod.Id_Modele)
	if rm.Error != nil {
		return errors.New(fmt.Sprintf("Failed to delete ingredients modele : ", rs.Error))
	}

	database.CloseSQLServer(db)

	return nil
}

func PatchModele(dataMod Patch_Data_Mod) error {
	db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }

	var count int64
    if err := db.Table("dbo.modeles").Where("id = ?", dataMod.Data.Id).Count(&count).Error; err != nil {
        return errors.New(fmt.Sprintf("Failed to check id in table dbo.modeles:", err))
    }

    if count == 0 {
        return errors.New("Id not found in dbo.modeles")
    }

	if dataMod.Data.Nom != ""{
		rs := db.Table("dbo.modeles").Where("id = ?", dataMod.Data.Id).Update("nom", dataMod.Data.Nom)
		if rs.Error != nil {
			return errors.New(fmt.Sprintf("Failed to update nom in table dbo.modeles:", rs.Error))
		}
	}

	if dataMod.Data.Description != ""{
		rs := db.Table("dbo.modeles").Where("id = ?", dataMod.Data.Id).Update("description", dataMod.Data.Description)
		if rs.Error != nil {
			return errors.New(fmt.Sprintf("Failed to update description in table dbo.modeles:", rs.Error))
		}
	}

	if dataMod.Data.PUHT != 0{
		rs := db.Table("dbo.modeles").Where("id = ?", dataMod.Data.Id).Update("puht", dataMod.Data.PUHT)
		if rs.Error != nil {
			return errors.New(fmt.Sprintf("Failed to update puht in table dbo.modeles:", rs.Error))
		}
	}

	if dataMod.Data.Gamme != ""{
		rs := db.Table("dbo.modeles").Where("id = ?", dataMod.Data.Id).Update("gamme", dataMod.Data.Gamme)
		if rs.Error != nil {
			return errors.New(fmt.Sprintf("Failed to update gamme in table dbo.modeles:", rs.Error))
		}
	}

	if len(dataMod.Data.Ingredient) > 0{
		for _, ingredient := range dataMod.Data.Ingredient {
			err := CheckIfIngredientExists(ingredient.Id_Ingredient)
			if err != nil{
				continue
			}

			if ingredient.Type == "Ajout"{
				newIngredient := database.Ingredient_Modele{
					ModeleID:      uint32(dataMod.Data.Id), 
					IngredientID:  uint32(ingredient.Id_Ingredient),
					Grammage:      ingredient.Grammage,
				}

				// Insérez cette instance dans la base de données
				if err := db.Create(&newIngredient).Error; err != nil {
					return err
				}
			}
			if ingredient.Type == "Suppression"{
				rs := db.Table("dbo.ingredient_modeles").Where("modele_id = ? AND ingredient_id = ?", dataMod.Data.Id, ingredient.Id_Ingredient).Delete(&dataMod.Data.Id)
				if rs.Error != nil {
					return errors.New(fmt.Sprintf("Failed to delete ingredient in modele : ", rs.Error))
				}
			}
			if ingredient.Type == "Modification"{
				rs := db.Table("dbo.ingredient_modeles").Where("modele_id = ? AND ingredient_id = ?", dataMod.Data.Id, ingredient.Id_Ingredient).Update("grammage", ingredient.Grammage)
				if rs.Error != nil {
					return errors.New(fmt.Sprintf("Failed to update grammage in table dbo.ingredient_modeles:", rs.Error))
				}
			}
		}
	}

	database.CloseSQLServer(db)

	return nil
}