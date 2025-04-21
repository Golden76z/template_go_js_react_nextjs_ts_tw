package handlers

import (
	"api"
	"models"
	"net/http"
	"utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// * generate your error message
		err := &models.CustomError{
			StatusCode: http.StatusNotFound,
			Message:    "Page Not Found",
		}
		utils.HandleError(w, err.StatusCode, err.Message)
		return
	}

	// Render a basic Home Page
	api.RenderTemplate(w, "layout/index", "page/index", nil)
}
