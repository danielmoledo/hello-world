package repository

import (
	"github.com/danielmoledo/hello-world/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Create(hw *model.HelloWorld) (*model.HelloWorld, error){
	result := DB.Create(&hw)

	return hw, result.Error
}


