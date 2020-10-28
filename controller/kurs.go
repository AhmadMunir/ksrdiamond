package controller

import (
	"encoding/json"
	"fmt"
	"kasir/models"
	"net/http"
	"strconv"
)

func RouteInsertKurs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var status string
		var i = r.FormValue("rp")
		var j = r.FormValue("coin")
		rp, _ := strconv.Atoi(i)
		coin, _ := strconv.Atoi(j)

		err := models.InsertKurs(rp, coin)
		if err != nil {
			fmt.Println(err)
			status = "gagal"
		} else {
			status = "success"
		}

		jsonData := map[string]interface{}{
			"status": status,
		}

		err2 := json.NewEncoder(w).Encode(jsonData)
		if err2 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetKurs(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var status string

		w.Header().Set("Content-Type", "application/json")

		result, err := models.GetKurs()

		if err != nil {
			status = "gagal"
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			status = "sukses"
		}

		jsonData := map[string]interface{}{
			"status": status,
			"data":   result,
		}

		var data, err2 = json.Marshal(jsonData)
		if err2 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	http.Error(w, "", http.StatusBadRequest)

}
