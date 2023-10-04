package database

import (
	"log"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"freezebee/api/middleware_api"
)

type Fabrication struct {
	Id          uint32        `gorm:"primary_key;auto_increment"`
	Nom         string        `gorm:"column:nom;type:nvarchar(max);not null"`
	Description string        `gorm:"column:description;type:nvarchar(max);not null"`
	ModeleID    uint32  
	Etapes_Description string `gorm:"column:etape_description;type:nvarchar(max);not null"`
}

type Modele struct {
	Id          uint32              `gorm:"primary_key;auto_increment"`
	Nom         string              `gorm:"column:nom;type:nvarchar(max);not null"`
	Description string              `gorm:"column:description;type:nvarchar(max);not null"`
	PUHT		float32             `gorm:"column:puht;type:float;not null"`	
	Gamme       string              `gorm:"column:gamme;type:nvarchar(max);not null"`
}

type Ingredient_Modele struct {
    ModeleID      uint32  // ID du modèle
    IngredientID  uint32  // ID de l'ingrédient
    Grammage      float32 `gorm:"column:grammage;type:float;not null"`
}

type Ingredient struct {
	Id          uint32 `gorm:"primary_key;auto_increment"`
	Nom         string `gorm:"column:nom;type:nvarchar(max);not null"`
	Description string `gorm:"column:description;type:nvarchar(max);not null"`
}

func AutoMigrateTables(db *gorm.DB) {
    // Utilisez AutoMigrate pour créer automatiquement les tables si elles n'existent pas
	log.Println("Création des tables")
    db.AutoMigrate(&Ingredient{})
    db.AutoMigrate(&Fabrication{})
    db.AutoMigrate(&Modele{})
    db.AutoMigrate(&Ingredient_Modele{})
}

func ConnectSQLServer() (*gorm.DB, error) {
	// Définissez les informations de connexion
    username := "Xav"
    password := "Cybersecurity1!"
    host := "10.0.2.105"
    port := 1433
    database := middleware_api.GetDatabase()

    // Construisez le DSN (Data Source Name)
    dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        host, username, password, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Impossible de se connecter à la base de données : " + err.Error())
    }

    return db, nil
}

func CloseSQLServer(db *gorm.DB) {
    sqlDB, err := db.DB()
    if err != nil {
        panic("Impossible d'obtenir la connexion SQL sous-jacente : " + err.Error())
    }

    // Fermer la connexion sous-jacente
    if err := sqlDB.Close(); err != nil {
        panic("Impossible de fermer la connexion SQL : " + err.Error())
    }
}

func InitConnectSQLServer(database string) (*gorm.DB, error) {
	// Définissez les informations de connexion
    username := "Xav"
    password := "Cybersecurity1!"
    host := "10.0.2.105"
    port := 1433

    // Construisez le DSN (Data Source Name)
    dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        host, username, password, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Impossible de se connecter à la base de données : " + err.Error())
    }

    return db, nil
}
