package services

import (
	"github.com/silverswords/sand/models"
	"github.com/silverswords/sand/models/mysql"
)

func InsertProduct(product models.Product) error {
	err := mysql.InsertProduct(product)

	return err
}
