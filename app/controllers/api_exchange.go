package controllers

import (
	"Foreign-Currency-BE-Exercise/app/models"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/revel/revel"
)

type ApiExchange struct {
	*revel.Controller
}

func (c ApiExchange) PairLast7Days() revel.Result {
	from := c.Params.Route.Get("from")
	to := c.Params.Route.Get("to")

	exchanges := []models.Exchange{}
	result := DB.Where("\"from\" = ?", from).Where("\"to\" = ?", to).Order("id desc").Find(&exchanges)
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record " + from + " - " + to + "  Not Found"))
	}

	return c.RenderJSON(result)
}

// Index
func (c ApiExchange) Index() revel.Result {
	exchanges := []models.Exchange{}

	result := DB.Order("id desc").Find(&exchanges)
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found"))
	}

	return c.RenderJSON(exchanges)
}

// View
func (c ApiExchange) View() revel.Result {
	id := c.Params.Route.Get("id")
	exchanges := []models.Exchange{}
	exchange := DB.Find(&exchanges, id)

	if exchange.Error != nil {
		return c.RenderError(errors.New("Data not found." + exchange.Error.Error()))
	}

	return c.RenderJSON(exchange)
}

// Create
func (c ApiExchange) Create() revel.Result {
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

	c.Response.Status = 201
	return c.RenderJSON(exchange)
}

// Update
func (c ApiExchange) Update() revel.Result {

	id := c.Params.Route.Get("id")
	exchanges := []models.Exchange{}
	exchange := DB.Find(&exchanges, id)

	if exchange.Error != nil {
		return c.RenderError(errors.New("Data not found." + exchange.Error.Error()))
	}

	// layout := "2000-01-29"
	tdate, _ := time.Parse(time.RFC3339, c.Params.Form.Get("Date")+"T00:00:00.000Z") //@TODO: Support multiple a day later
	trate64, rateErr := strconv.ParseFloat(c.Params.Form.Get("Rate"), 32)
	if rateErr != nil {
		fmt.Println("Rate err: ", rateErr)
	}
	trate32 := float32(trate64)
	id64, _ := strconv.ParseInt(id, 10, 64)
	uint64id := uint64(id64)

	exchange2 := models.Exchange{
		Id:   uint64id,
		Date: tdate,
		From: c.Params.Form.Get("From"),
		To:   c.Params.Form.Get("To"),
		Rate: trate32,
	}

	DB.Save(&exchange2)

	c.Response.Status = 201
	return c.RenderJSON(exchange2)
}

// Delete
func (c ApiExchange) Delete() revel.Result {
	id := c.Params.Route.Get("id")
	posts := []models.Exchange{}
	exchange := DB.Delete(&posts, id)
	if exchange.Error != nil {
		return c.RenderError(errors.New("Record Delete failure." + exchange.Error.Error()))
	}

	c.Response.Status = 200
	return c.RenderJSON(exchange)
}
