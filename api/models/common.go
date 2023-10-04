package models

type Delete_Data struct {
	Table string `gorm:"size:60;not null" json:"Table"`
	Id    int    `gorm:"size:10;not null" json:"Id"`
}

type Delete_Data_Modele struct {
	Table     string `gorm:"size:60;not null" json:"Table"`
	Id        int    `gorm:"size:10;not null" json:"Id"`
	Id_Modele int    `gorm:"size:10;not null" json:"Id_Modele"`
}