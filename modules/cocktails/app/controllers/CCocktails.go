package controllers

import (
	"cocktails/modules/cocktails/app/providers"
	"cocktails/modules/cocktails/app/types"
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"io/ioutil"
)

type CCocktails struct {
	*revel.Controller
	provider *providers.PCocktails
}

func (c *CCocktails) Init() (result revel.Result) {
	c.provider = providers.New()
	return
}

func (c *CCocktails) GetAll() (result revel.Result) {

	cocktails, err := c.provider.GetAll()
	if err != nil {
		return c.RenderError(err)
	}
	return c.RenderJson(cocktails)
}

func (c *CCocktails) Add() (result revel.Result) {

	cocktail, err := c.parse()
	if err != nil {
		return c.RenderError(err) // 400
	}

	id, err := c.provider.Add(cocktail)
	if err != nil {
		return c.RenderError(err)
	}
	return c.RenderJson(id)
}

func (c *CCocktails) GetById(id int64) (result revel.Result) {

	cocktail, err := c.provider.GetById(id)
	if err != nil {
		return c.RenderError(err)
	}
	return c.RenderJson(cocktail)
}

func (c *CCocktails) parse() (cocktail *types.Cocktail, err error) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		err = fmt.Errorf("CCocktails.parse Ошибка чтения тела запроса: %s", err)
		return
	}
	if err = json.Unmarshal(body, &cocktail); err != nil {
		err = fmt.Errorf("CCocktails.parse Ошибка парсинга тела запроса: %s", err)
		return
	}

	return
}