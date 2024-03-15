package controller

import (
	"assignment3/utill"
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	data := utill.FetchData()
	toMap := utill.ConvertMap(data)

	toMap["water"] = rand.Intn(10) + 1
	toMap["wind"] = rand.Intn(18) + 1

	var status = ""

	jsonData := map[string]any{
		"status": status,
		"wind":   toMap["wind"],
		"water":  toMap["water"],
	}

	if (toMap["wind"].(int) > 15) || (toMap["water"].(int) > 8) {
		status = "Bahaya"
	} else if ((toMap["water"].(int) >= 6) && (toMap["water"].(int) <= 8)) || (toMap["wind"].(int) >= 7 && toMap["wind"].(int) <= 15) {
		status = "Siaga"
	} else if (toMap["water"].(int) <= 5) && (toMap["wind"].(int) <= 6) {
		status = "Aman"
	} else {
		status = "Status tidak diketahui"
	}

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			errs := fmt.Sprintf("Error parse files : %v", err.Error())
			http.Error(w, errs, http.StatusInternalServerError)
		}
		tpl.Execute(w, jsonData)
	} else {
		http.Error(w, "Wrong Method", http.StatusConflict)
	}

}

func GetStatusJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := utill.FetchData()
	toMap := utill.ConvertMap(data)

	toMap["water"] = rand.Intn(10) + 1
	toMap["wind"] = rand.Intn(18) + 1

	var status = ""

	if (toMap["wind"].(int) > 15) || (toMap["water"].(int) > 8) {
		status = "Bahaya"
	} else if ((toMap["water"].(int) >= 6) && (toMap["water"].(int) <= 8)) || (toMap["wind"].(int) >= 7 && toMap["wind"].(int) <= 15) {
		status = "Siaga"
	} else if (toMap["water"].(int) <= 5) && (toMap["wind"].(int) <= 6) {
		status = "Aman"
	} else {
		status = "Status tidak diketahui"
	}

	jsonData := map[string]any{
		"status": status,
		"wind":   toMap["wind"],
		"water":  toMap["water"],
	}

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(jsonData)
	} else {
		http.Error(w, "Wrong Method", http.StatusConflict)
	}

}
