package model

//http://jinzhu.me/gorm/models.html#conventions
import (
	"fmt"

	"qpgame/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// import _ "github.com/jinzhu/gorm/dialects/postgres"
	// import _ "github.com/jinzhu/gorm/dialects/sqlite"
	// import _ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	//db, err = sql.Open("mysql", "tardis:aa00qq99@tcp(rdsqe3uffzu3qeyo.mysql.rds.aliyuncs.com:3306)/fangjia")
	//Open("mysql", "root:sqlmima@/lf_game?charset=utf8&parseTime=True&loc=Local")
	conn_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", conf.DB_USER, conf.DB_PWD, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)
	DB, err = gorm.Open("mysql", conn_str)

	fmt.Println(conn_str)
	if err != nil {
		fmt.Println(err)
	} else {
		DB.LogMode(true)
		fmt.Println("Init DB :", &DB)
	}
}

//base model
type BaseModel struct {
	gorm.Model
}
