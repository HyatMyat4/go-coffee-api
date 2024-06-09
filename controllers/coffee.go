package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/takumi/coffee-api/helpers"
	"github.com/takumi/coffee-api/services"
)

var coffee services.Coffee

// Get one cofee by id
func GetCoffeeById(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	result, err := coffee.GetCoffeeById(id)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, result)

}

// Get all coffees
func GetAllCoffees(w http.ResponseWriter, req *http.Request) {
	all, err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envlop{"coffees": all})

}

// Create coffee
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

// Update coffee
func UpdateCoffee(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	var coffeeData services.Coffee

	err := json.NewDecoder(req.Body).Decode(&coffeeData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	result, err := coffee.UpdateCoffee(id, coffeeData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, result)

}

func DeleteCoffee(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	err := coffee.DeleteCoffee(id)

	if err != nil {

		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

}
