package providers

import (
	"cocktails/modules/cocktails/app/mappers"
	"cocktails/modules/cocktails/app/types"
	"cocktails/modules/dbModel"
	"fmt"
)

type PCocktails struct {
	mapper *mappers.MCocktails
}

func New() *PCocktails {
	db := dbModel.GetDB()
	mapper := mappers.New(db)
	return &PCocktails{
		mapper: mapper,
	}
}

func (p *PCocktails) GetAll() (cocktails []*types.Cocktail, err error) {

	cocktails, err = p.mapper.GetAll()
	if err != nil {
		err = fmt.Errorf("PCocktails.GetAll error: %s", err)
		return
	}
	return
}

func (p *PCocktails) Add(cocktail *types.Cocktail) (id int64, err error) {

	id, err = p.mapper.Add(cocktail)
	if err != nil {
		err = fmt.Errorf("PCocktails.Add error: %s", err)
		return
	}
	return
}

func (p *PCocktails) GetById(id int64) (cocktail *types.Cocktail, err error) {

	cocktail, err = p.mapper.GetById(id)
	if err != nil {
		err = fmt.Errorf("PCocktails.GetById error: %s", err)
		return
	}
	return
}