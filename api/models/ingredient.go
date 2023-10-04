package models

import (
	"freezebee/api/database"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Post_Data_Ing struct {
	Table string   `gorm:"size:60;not null" json:"Table"`
	Data  Data_Ing `gorm:"size:1000;not null" json:"Data"`
}

type Data_Ing struct {
	Nom         string `gorm:"size:60;not null" json:"Nom"`
	Description string `gorm:"size:60;not null" json:"Description"`
}

type Patch_Data_Ing struct {
	Table string   `gorm:"size:60;not null" json:"Table"`
	Data  Data_Ing_2 `gorm:"size:1000;not null" json:"Data"`
}

type Data_Ing_2 struct {
	Id 			int	   `gorm:"size:10;not null" json:"Id"`
	Nom         string `gorm:"size:60;not null" json:"Nom"`
	Description string `gorm:"size:60;not null" json:"Description"`
}

func CheckIfIngredientDontExist(name string) (Data_Ing, error) {
	db, err := database.ConnectSQLServer()
	if err != nil {
		return Data_Ing{}, errors.New(err.Error())
	}

	var ingredient Data_Ing

	// Recherche la newsletter dans la table "newsletter_dbs" en utilisant les trois critères
	if err := db.Table("dbo.ingredients").Where("nom = ? ", name).First(&ingredient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Data_Ing{}, nil // Newsletter non trouvée
		}
		return Data_Ing{}, err // Autre erreur
	}

	database.CloseSQLServer(db)

	return ingredient, nil
}

func CheckIfIngredientExists(id int) error {
    db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }
    defer database.CloseSQLServer(db)

    var count int64

    // Comptez le nombre d'occurrences de l'ingrédient avec l'ID donné
    if err := db.Table("dbo.ingredients").Where("id = ?", id).Count(&count).Error; err != nil {
        return err // Gérer l'erreur
    }

    if count == 0 {
        return errors.New("L'ingrédient n'existe pas") // L'ingrédient n'existe pas
    }

    return nil 
}


func PostIngredient(dataIng Post_Data_Ing) error {
	db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }

	ingredient := database.Ingredient{
		Nom: dataIng.Data.Nom,
		Description: dataIng.Data.Description,
	}

	rs := db.Create(&ingredient)
	if rs.Error != nil {
        return errors.New(fmt.Sprintf("Failed to create ingredient:", rs.Error))
    }

	database.CloseSQLServer(db)

	return nil
}

func GetIngredients() ([]database.Ingredient, error) {
	db, err := database.ConnectSQLServer()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var ingredient []database.Ingredient

	rs := db.Table("dbo.ingredients").Find(&ingredient)
	if rs.Error != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get all ingredients : ", rs.Error))
	}

	database.CloseSQLServer(db)

	return ingredient, nil
}

func DeleteIngredient(dataIng Delete_Data) error {
	db, err := database.ConnectSQLServer()
	if err != nil {
		return errors.New(err.Error())
	}

	var count int64
    if err := db.Table("dbo.ingredients").Where("id = ?", dataIng.Id).Count(&count).Error; err != nil {
        return errors.New(fmt.Sprintf("Failed to check id in table dbo.ingredients:", err))
    }

    if count == 0 {
        return errors.New("Id not found in dbo.ingredients")
    }
		
	rs := db.Table("dbo.ingredients").Where("id = ?", dataIng.Id).Delete(&dataIng)
	if rs.Error != nil {
		return errors.New(fmt.Sprintf("Failed to delete ingredient : ", rs.Error))
	}

	database.CloseSQLServer(db)

	return nil
}

func PatchIngredient(dataIng Patch_Data_Ing) error {
	db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }

	var count int64
    if err := db.Table("dbo.ingredients").Where("id = ?", dataIng.Data.Id).Count(&count).Error; err != nil {
        return errors.New(fmt.Sprintf("Failed to check id in table dbo.ingredients:", err))
    }

    if count == 0 {
        return errors.New("Id not found in dbo.ingredients")
    }

	if dataIng.Data.Nom != ""{
		rs := db.Table("dbo.ingredients").Where("id = ?", dataIng.Data.Id).Update("nom", dataIng.Data.Nom)
		if rs.Error != nil {
			return errors.New(fmt.Sprintf("Failed to update nom in table dbo.ingredients:", rs.Error))
		}
	}

	if dataIng.Data.Description != ""{
		rs := db.Table("dbo.ingredients").Where("id = ?", dataIng.Data.Id).Update("description", dataIng.Data.Description)
		if rs.Error != nil {
			return errors.New(fmt.Sprintf("Failed to update description in table dbo.ingredients:", rs.Error))
		}
	}

	database.CloseSQLServer(db)

	return nil
}

