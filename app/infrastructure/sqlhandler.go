// ドライバー
// アプリケーションの外部
// interfaces/databaseに依存しているため、アプリケーションの内部が守られている
package infrastructure

import (
	"Yukit02/app/interfaces/database"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	// NOTE: rootユーザーでないとエラーになるため、rootユーザーで実行する
	// セキュリティの観点では好ましくない
	user := "root"
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

// interfaceに定義したメソッドを実装
func (handler *SqlHandler) Create(object interface{}) {
	handler.db.Create(object)
}

func (handler *SqlHandler) FindAll(object interface{}) {
	handler.db.Find(object)
}

func (handler *SqlHandler) DeleteById(object interface{}, id string) {
	handler.db.Delete(object, id)
}
