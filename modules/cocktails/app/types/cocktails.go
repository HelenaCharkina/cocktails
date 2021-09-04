package types

import (
	"database/sql"
)

type Cocktail struct {
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Recipe  string  `json:"recipe"`
	Rating  float64 `json:"rating"`
	ImgPath string  `json:"img_path"`
}

type DBCocktail struct {
	Id      sql.NullInt64
	Name    sql.NullString
	Recipe  sql.NullString
	Rating  sql.NullFloat64
	ImgPath sql.NullString
}

func (db *DBCocktail) ToType() *Cocktail {
	return &Cocktail{
		Id:      db.Id.Int64,
		Name:    db.Name.String,
		Recipe:  db.Recipe.String,
		Rating:  db.Rating.Float64,
		ImgPath: db.ImgPath.String,
	}
}
