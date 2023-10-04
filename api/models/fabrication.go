package models

import (
	"freezebee/api/database"
	"errors"
	"fmt"
)

type Post_Data_Fab struct {
	Table string   `gorm:"size:60;not null" json:"Table"`
	Data  Data_Fab `gorm:"size:1000;not null" json:"Data"`
}

type Patch_Data_Fab struct {
	Table string           `gorm:"size:60;not null" json:"Table"`
	Data  Patch_Data_Fab_2 `gorm:"size:1000;not null" json:"Data"`
}

type Data_Fab struct {
	Nom                string `gorm:"size:60;not null" json:"Nom"`
	Description        string `gorm:"size:60;not null" json:"Description"`
	Id_Modele          int    `gorm:"size:10;not null" json:"Id_Modele"`
	Etapes_Description string `gorm:"size:60;not null" json:"Etapes_Description"`
}

type Patch_Data_Fab_2 struct {
	Id 			       int	  `gorm:"size:10;not null" json:"Id"`
	Nom                string `gorm:"size:60;not null" json:"Nom"`
	Description        string `gorm:"size:60;not null" json:"Description"`
	Id_Modele          int    `gorm:"size:10;not null" json:"Id_Modele"`
	Etapes_Description string `gorm:"size:60;not null" json:"Etapes_Description"`
}

type Fabrication struct {
    Id                 uint32 `gorm:"primary_key;auto_increment"`
    Nom                string `gorm:"column:nom;type:nvarchar(max);not null"`
    Description        string `gorm:"column:description;type:nvarchar(max);not null"`
    Etapes_Description string `gorm:"column:etape_description;type:nvarchar(max);not null"`
    Id_Modele          uint32 
    Nom_Modele         string
    Description_Modele string
    PUHT               float32
    Gamme              string
}


func CheckIfFabricationExists(nom string) error {
    db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }
    defer database.CloseSQLServer(db)

    var count int64

    // Comptez le nombre d'occurrences de l'ingrédient avec l'ID donné
    if err := db.Table("dbo.fabrications").Where("nom = ?", nom).Count(&count).Error; err != nil {
        return err // Gérer l'erreur
    }

    if count == 0 {
        return errors.New("La fabrication n'existe pas") // L'ingrédient n'existe pas
    }

    return nil 
}

func CheckIfModExists(id int) error{
	db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }

	var count int64

    // Comptez le nombre d'occurrences de l'ingrédient avec l'ID donné
    if err := db.Table("dbo.modeles").Where("id = ?", id).Count(&count).Error; err != nil {
        return err
    }

    if count == 0 {
        return errors.New("Le modele n'existe pas")
    }

	database.CloseSQLServer(db)

	return nil
}

func PostFabrication(dataFab Post_Data_Fab) error {
	db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }

	fabrication := database.Fabrication{
		Nom: dataFab.Data.Nom,
		Description: dataFab.Data.Description,
		ModeleID: uint32(dataFab.Data.Id_Modele),
		Etapes_Description: dataFab.Data.Etapes_Description,
	}

	rs := db.Create(&fabrication)
	if rs.Error != nil {
        return errors.New(fmt.Sprintf("Failed to create fabrication:", rs.Error))
    }

	database.CloseSQLServer(db)

	return nil
}

func GetAllFabrications() ([]map[string]interface{}, error) {
    db, err := database.ConnectSQLServer()
    if err != nil {
        return nil, err
    }
    defer database.CloseSQLServer(db)

    var data []map[string]interface{}

    // Sélectionnez toutes les fabrications de dbo.fabrications
    var fabrications []Fabrication
    result := db.Table("dbo.fabrications").
        Select("dbo.fabrications.id, dbo.fabrications.nom, dbo.fabrications.description, dbo.fabrications.etape_description, dbo.fabrications.modele_id as id_modele, dbo.modeles.nom as nom_modele, dbo.modeles.description as description_modele, dbo.modeles.puht as puht, dbo.modeles.gamme as gamme").
        Joins("LEFT JOIN dbo.modeles ON dbo.fabrications.modele_id = dbo.modeles.id").
        Find(&fabrications)

    if result.Error != nil {
        return nil, result.Error
    }

    // Parcourez les fabrications et ajoutez-les à la liste de données
    for _, fabrication := range fabrications {
        fabricationData := map[string]interface{}{
			"Id":				  fabrication.Id,
            "Nom":                fabrication.Nom,
            "Description":        fabrication.Description,
            "Id_Modèle":          fabrication.Id_Modele,
            "Nom_Modèle":         fabrication.Nom_Modele,
            "Description_Modèle": fabrication.Description_Modele,
            "pUHT":               fabrication.PUHT,
            "Gamme":              fabrication.Gamme,
            "Etapes_Description": fabrication.Etapes_Description,
        }

        // Ajoutez la fabrication à la liste de données
        data = append(data, fabricationData)
    }

    return data, nil
}

func DeleteFabrication(dataFab Delete_Data) error {
	db, err := database.ConnectSQLServer()
	if err != nil {
		return errors.New(err.Error())
	}

	var count int64
    if err := db.Table("dbo.fabrications").Where("id = ?", dataFab.Id).Count(&count).Error; err != nil {
        return errors.New(fmt.Sprintf("Failed to check id in table dbo.fabrications:", err))
    }

    if count == 0 {
        return errors.New("Id not found in dbo.fabrications")
    }
		
	rs := db.Table("dbo.fabrications").Where("id = ?", dataFab.Id).Delete(&dataFab)
	if rs.Error != nil {
		return errors.New(fmt.Sprintf("Failed to delete fabrication : ", rs.Error))
	}

	database.CloseSQLServer(db)

	return nil
}

func PatchFabrication(dataFab Patch_Data_Fab) error {
	db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }

	var count int64
    if err := db.Table("dbo.fabrications").Where("id = ?", dataFab.Data.Id).Count(&count).Error; err != nil {
        return errors.New(fmt.Sprintf("Failed to check id in table dbo.fabrications:", err))
    }

    if count == 0 {
        return errors.New("Id not found in dbo.fabrications")
    }

	if dataFab.Data.Nom != ""{
		rs := db.Table("dbo.fabrications").Where("id = ?", dataFab.Data.Id).Update("nom", dataFab.Data.Nom)
		if rs.Error != nil {
			return errors.New(fmt.Sprintf("Failed to update nom in table dbo.fabrications:", rs.Error))
		}
	}

	if dataFab.Data.Description != ""{
		rs := db.Table("dbo.fabrications").Where("id = ?", dataFab.Data.Id).Update("description", dataFab.Data.Description)
		if rs.Error != nil {
			return errors.New(fmt.Sprintf("Failed to update description in table dbo.fabrications:", rs.Error))
		}
	}

	if dataFab.Data.Id_Modele != 0{
		rs := db.Table("dbo.fabrications").Where("id = ?", dataFab.Data.Id).Update("modele_id", dataFab.Data.Id_Modele)
		if rs.Error != nil {
			return errors.New(fmt.Sprintf("Failed to update modele_id in table dbo.fabrications:", rs.Error))
		}
	}

	if dataFab.Data.Etapes_Description != ""{
		rs := db.Table("dbo.fabrications").Where("id = ?", dataFab.Data.Id).Update("etape_description", dataFab.Data.Etapes_Description)
		if rs.Error != nil {
			return errors.New(fmt.Sprintf("Failed to update etape_description in table dbo.fabrications:", rs.Error))
		}
	}

	database.CloseSQLServer(db)

	return nil
}


