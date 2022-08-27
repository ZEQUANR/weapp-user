package mysql

import (
	"database/sql"

	"github.com/silverswords/sand/sql/mysql"
)

const (
	mysqlCreateProductTable = iota
)

var (
	db *sql.DB

	createFormatStr = []string{
		`CREATE TABLE IF NOT EXISTS mall.product (
			pro_id 		 VARCHAR(64)  NOT NULL DEFAULT "",
			version_id 	 VARCHAR(64)  NOT NULL DEFAULT "",
			main_title 	 VARCHAR(512) NOT NULL DEFAULT "",
			subtitle 	 VARCHAR(200) NOT NULL DEFAULT "",
			main_picture VARCHAR(512) NOT NULL DEFAULT "",
			price 		 DECIMAL(20,2) NOT NULL DEFAULT 0, 
			status 		 INT(6) 	NOT NULL DEFAULT 0,
			active 		 INT(6) 	NOT NULL DEFAULT 0,
			create_date  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			update_date  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY(pro_id, version_id)
		  ) ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8mb4;`,
	}
)

func init() {
	db = mysql.CreateBuilder("./config/sql").Build().Run()

	_, err := db.Exec(createFormatStr[mysqlCreateProductTable])
	if err != nil {
		panic(err)
	}
}
