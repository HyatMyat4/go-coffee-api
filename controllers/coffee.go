package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/takumi/coffee-api/helpers"
	"github.com/takumi/coffee-api/services"
)

var coffee services.Coffee

// Get coffees
func GetAllCoffees(w http.ResponseWriter, req *http.Request) {
	all, err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envlop{"coffees": all})

}

func CreateCoffee(w http.ResponseWriter, req *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(req.Body).Decode(&coffeeData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, coffeeData)
	result, err := coffee.CreateCoffee(coffeeData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, result)

}
