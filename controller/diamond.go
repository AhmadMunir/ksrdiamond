package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"kasir/models"
)

type M map[string]interface{}

func RouteInsertDiamond(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var data = M{"title": "Diamond"}
		var tmpl = template.Must(template.ParseFiles(
			"view/diamond/index.html",
			"view/partial/_head.html",
			"view/partial/_sidebar.html",
			"view/partial/_topbar.html",
			"view/partial/_jquery.html",
			"view/partial/_footer.html",
		))

		var name = r.FormValue("name")
		var dm = r.FormValue("dm")
		var i = r.FormValue("price")
		price, _ := strconv.Atoi(i)

		err := models.InsertDiamond(name, dm, price)
		if err != nil {
			fmt.Println(err)
		}

		if err := tmpl.ExecuteTemplate(w, "indexdiamond", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

}

func AddDiamondView(w http.ResponseWriter, r *http.Request) {
	var data = M{"title": "Diamond"}
	var tmpl = template.Must(template.ParseFiles(
		"view/diamond/adddiamond.html",
		"view/partial/_head.html",
		"view/partial/_sidebar.html",
		"view/partial/_topbar.html",
		"view/partial/_jquery.html",
		"view/partial/_footer.html",
	))

	var err = tmpl.ExecuteTemplate(w, "adddiamond", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DiamondView(w http.ResponseWriter, r *http.Request) {
	var data = M{"title": "Diamond"}
	var tmpl = template.Must(template.ParseFiles(
		"view/diamond/index.html",
		"view/partial/_head.html",
		"view/partial/_sidebar.html",
		"view/partial/_topbar.html",
		"view/partial/_jquery.html",
		"view/partial/_footer.html",
	))

	var err = tmpl.ExecuteTemplate(w, "indexdiamond", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func GetAllDiamonds(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var status string
		w.Header().Set("Content-Type", "application/json")
		result, err2 := models.GetAllDiamonds()

		if err2 != nil {
			status = "gagal"
		} else {
			status = "sukses"
		}

		jsonData := map[string]interface{}{
			"status": status,
			"data":   result,
		}

		var data, err = json.Marshal(jsonData)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(data)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func DeleteDiamondById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		// fmt.Println(r.Body)
		payload := struct {
			Id int `json:"id"`
		}{}
		var status string
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			status = "gagal"
			return
		}

		// fmt.Println("id ", payload.Id)

		err := models.DeleteDiamondById(payload.Id)
		if err != nil {
			fmt.Println(err)
			status = "gagal"
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			status = "success"
		}

		w.Header().Set("Content-Type", "application/json")

		jsonData := map[string]interface{}{
			"status": status,
		}

		err2 := json.NewEncoder(w).Encode(jsonData)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdateDiamondById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var status string
		w.Header().Set("Content-Type", "application/json")

		var i = r.FormValue("id")
		id, _ := strconv.Atoi(i)
		var name = r.FormValue("name")
		var dm = r.FormValue("dm")
		var j = r.FormValue("price")
		price, _ := strconv.Atoi(j)

		err := models.UpdateDiamondById(id, price, name, dm)
		if err != nil {
			fmt.Println("gagal")
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
