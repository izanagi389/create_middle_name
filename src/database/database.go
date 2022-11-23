package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"izanagi-portfolio-site.com/model"
)

func godotenvConnect() {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err.Error())
	}
}

func gormConnect() *gorm.DB {

	godotenvConnect()

	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASS")
	PROTOCOL := os.Getenv("MYSQL_PROTOCOL")
	DBNAME := os.Getenv("DBNAME")
	CHARSET := os.Getenv("CHARSET")
	LOC := os.Getenv("LOC")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=" + CHARSET + "&parseTime=true&loc=" + LOC

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// DBの初期化
func Init() {
	db := gormConnect()

	// コネクション解放解放
	defer db.Close()

	// テーブルの初期化
	db.Exec("DROP TABLE mrs")
	db.Exec("DROP TABLE sns")
	db.Exec("DROP TABLE cns")

	db.AutoMigrate(&model.Mr{})
	db.AutoMigrate(&model.SN{})
	db.AutoMigrate(&model.CN{})

	// DB初期値の挿入
	DbInsertSeed(db)
}

func DBGetRandomMrData() model.Mr {
	db := gormConnect()

	defer db.Close()
	var mr model.Mr
	// db.Raw("SELECT * FROM mrs ORDER BY RAND() LIMIT 1").Find(&mr)
	// db.Limit(1).Offset(5).Find(&users)
	db.Order("RAND()").Find(&mr)
	// SELECT * FROM users ORDER BY age desc, name;

	return mr
}
func DBGetRandomSNData() model.SN {
	db := gormConnect()

	defer db.Close()
	var sn model.SN
	// db.Raw("SELECT * FROM mrs ORDER BY RAND() LIMIT 1").Find(&mr)
	// db.Limit(1).Offset(5).Find(&users)
	db.Order("RAND()").Find(&sn)
	// SELECT * FROM users ORDER BY age desc, name;

	return sn
}
func DBGetRandomCNData() model.CN {
	db := gormConnect()

	defer db.Close()
	var cn model.CN
	// db.Raw("SELECT * FROM mrs ORDER BY RAND() LIMIT 1").Find(&mr)
	// db.Limit(1).Offset(5).Find(&users)
	db.Order("RAND()").Find(&cn)
	// SELECT * FROM users ORDER BY age desc, name;

	return cn
}
