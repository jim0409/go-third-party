package mysqldb

import (
	"time"
)

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	ID        int    `gorm:"column:id;AUTO_INCREMENT"`
	Name      string `gorm:"column:name;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	ID        int    `gorm:"column:id;AUTO_INCREMENT"`
	Name      string `gorm:"column:name;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Users []User `gorm:"many2many:user_languages;"`
}

// CreateRecord is based on usr with multi-languages
func (db *mysqlDBObj) CreateRecord(usr string, lang ...string) error {
	user := User{Name: usr}
	language := make([]Language, len(lang))

	for i, j := range lang {
		language[i] = Language{Name: j}
	}

	user.Languages = language
	return db.DB.Create(&user).Error
}

// QueryRecord is based on usr with multi-languages
func (db *mysqlDBObj) QueryRecord(username string) (*[]Language, error) {
	var err error
	var languages []Language

	user := User{}
	if err = db.DB.Table("user").Where("name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	// // after use username to check user table, User would be filled
	// log.Println(user)

	if err = db.DB.Model(&user).Related(&languages, "Languages").Error; err != nil {
		return nil, err
	} else {
		/*
			assign user value back to languages
			since the origin languages's user value is empty
			it's no need to use j(value)
		*/
		for i, _ := range languages {
			languages[i].Users = []User{user}
		}
	}

	// // after use User{} structure to check languages table, languages would be filled
	// for i, j := range languages {
	// 	log.Printf("%v___%v\n", i, j)
	// }

	return &languages, nil
}
