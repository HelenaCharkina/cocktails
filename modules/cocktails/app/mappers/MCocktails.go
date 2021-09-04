package mappers

import (
	"cocktails/modules/cocktails/app/types"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MCocktails struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *MCocktails {
	return &MCocktails{
		db: db,
	}

}

func (m *MCocktails) GetAll() (cocktails []*types.Cocktail, err error) {

	var (
		dbCocktails []types.DBCocktail
	)

	query := `select id, name, recipe, rating, imgPath from cocktails.cocktails`

	err = m.db.Select(&dbCocktails, query)
	if err != nil {
		err = fmt.Errorf("MCocktails.GetAll db.Select error: %s", err)
		return
	}

	for _, cocktail := range dbCocktails {
		cocktails = append(cocktails, cocktail.ToType())
	}

	return
}

func (m *MCocktails) Add(cocktail *types.Cocktail) (id int64, err error) {

	query := `insert into cocktails.cocktails(name, recipe, rating, imgPath) values ($1, $2, $3, $4) returning id;`

	err = m.db.QueryRow(query, cocktail.Name, cocktail.Recipe, cocktail.Rating, cocktail.ImgPath).Scan(&id)
	if err != nil {
		err = fmt.Errorf("MCocktails.Add Scan error: %s", err)
		return
	}

	return
}

func (m *MCocktails) GetById(id int64) (cocktail *types.Cocktail, err error) {

	var (
		dbCocktail types.DBCocktail
	)

	query := `select id, name, recipe, rating, imgPath from cocktails.cocktails where id = $1`

	err = m.db.Get(&dbCocktail, query, id)
	if err != nil {
		err = fmt.Errorf("MCocktails.GetById db.Get error: %s", err)
		return
	}

	cocktail = dbCocktail.ToType()
	return
}