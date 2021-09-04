package controllers

import "github.com/revel/revel"

func init() {
	revel.InterceptMethod((*CCocktails).Init, revel.BEFORE)
}
