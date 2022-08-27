package mysql

import (
	"github.com/silverswords/sand/models"
)

const (
	mysqlInsertProductTable = iota
)

var (
	productFormatStr = []string{
		`INSERT INTO product (pro_id,version_id,main_title,subtitle,main_picture,price,status,active) 
		VALUES(?,?,?,?,?,?,?,?)`,
	}
)

func InsertProduct(product models.Product) error {
	_, err := db.Exec(productFormatStr[mysqlInsertProductTable], product.ProId,
		product.VersionId, product.MainTitle, product.Subtitle, product.MainPicture,
		product.Price, product.Status, product.Active)
	if err != nil {
		return err
	}

	return nil
}
