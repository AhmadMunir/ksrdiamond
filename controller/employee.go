package controller

import (
	"encoding/json"
	"kasir/models"
	"net/http"

	"github.com/labstack/echo"
)

// const CSRFTokenHeader = "X-CSRF-Token"
// const CSRFKey = "csrf"

func GetEmployees(c echo.Context) error {
	// println("foo")
	var title = "Dashboard"
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"title": title,
	})
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		result := models.GetEmployee()

		jsonData := map[string]interface{}{
			"status": "sukses",
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
