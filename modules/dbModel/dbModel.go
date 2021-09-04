package dbModel

import (
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

var dbSession *sqlx.DB

func InitDB() (db *sqlx.DB, err error) {

	if dbSession == nil {
		driver := revel.Config.StringDefault("db.driver", "")
		if driver == "" {
			err = errors.New("в настройках не указан драйвер подключения к БД")
			return
		}
		spec := revel.Config.StringDefault("db.spec", "")
		if spec == "" {
			err = errors.New("в настройках не указаны параметры подключения к БД")
			return
		}

		if dbSession, err = sqlx.Open(driver, spec); err != nil {
			err = errors.New("ошибка подключения к БД: " + err.Error())
			return
		}
		err = dbSession.Ping()
		if err != nil {
			revel.ERROR.Println("DB ping Error", err)
			return nil, err
		}
	}

	return
}

func GetDB() *sqlx.DB {
	return dbSession
}

