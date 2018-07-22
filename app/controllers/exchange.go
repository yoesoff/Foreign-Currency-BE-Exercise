package controllers

import (
	"Foreign-Currency-BE-Exercise/app/models"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/revel/revel"
)

type Exchange struct {
	*revel.Controller
}

// Index
func (c Exchange) Index() revel.Result {
	exchanges := []models.Exchange{}

	result := DB.Order("id desc").Find(&exchanges)
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}
	return c.Render(exchanges)
}

// Create
func (c Exchange) Create() revel.Result {
	// layout := "2000-01-29"
	tdate, _ := time.Parse(time.RFC3339, c.Params.Form.Get("Date")+"T00:00:00.000Z") //@TODO: Support multiple a day later
	trate64, rateErr := strconv.ParseFloat(c.Params.Form.Get("Rate"), 32)
	if rateErr != nil {
		fmt.Println("Rate err: ", rateErr)
	}
	trate32 := float32(trate64)

	exchange := models.Exchange{
		Date: tdate,
		From: c.Params.Form.Get("From"),
		To:   c.Params.Form.Get("To"),
		Rate: trate32,
	}

	ret := DB.Create(&exchange)
	if ret.Error != nil {
		return c.RenderError(errors.New("Record Create failure." + ret.Error.Error()))
	}
	return c.Redirect("/exchanges")
}

// Delete
func (c Exchange) Delete() revel.Result {
	id := c.Params.Route.Get("id")
	exchanges := []models.Exchange{}
	ret := DB.Delete(&exchanges, id)
	if ret.Error != nil {
		return c.RenderError(errors.New("Record Delete failure." + ret.Error.Error()))
	}
	return c.Redirect("/exchanges")
}

// Redirect
func (c Exchange) RedirectToExchange() revel.Result {
	return c.Redirect("/exchanges")
}
