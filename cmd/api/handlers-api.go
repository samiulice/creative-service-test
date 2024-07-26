package main

import (
	"creative-service-test/internal/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	UPLOAD_DIR = "static/private/order-uploads"
)
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id"`
}

func (app *application) MakeOrderHandler(w http.ResponseWriter, r *http.Request) {
	var resp jsonResponse
	if r.Method != http.MethodPost {
		resp.OK = false
		resp.Message = "Invalid request method"
		resp.ID = 405
		app.writeJSON(w, http.StatusMethodNotAllowed, resp)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		resp.OK = false
		resp.Message = "Internal server error"
		resp.ID = 1000 //"Could not parse image"
		app.writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	var order models.Order
	order.ClientName = r.FormValue("client-name")
	order.ClientMobile = r.FormValue("client-mobile")
	order.Division = r.FormValue("division")
	order.District = r.FormValue("district")
	order.Upazila = r.FormValue("upazila")
	order.Area = r.FormValue("area")
	order.Date = r.FormValue("date")
	order.PreferredTime = r.FormValue("preferredtime")
	order.Category = r.FormValue("category")
	order.Service = r.FormValue("service")
	order.ProblemSummary = r.FormValue("problem-summary")

	//Retrive category name
	catSplit := strings.Split(order.Category, "_")
	order.Category = ""
	for _, v := range catSplit{
		order.Category +=  app.titleCase(v)
	}

	// Handling file upload
	file, handler, err := r.FormFile("img_file")
	if err != nil {
		if err != http.ErrMissingFile {
			resp.OK = false
			resp.Message = "Internal server error"
			resp.ID = 1001 //"Could not read the image"
			app.writeJSON(w, http.StatusInternalServerError, resp)
			return
		}
	} else {
		defer file.Close()
		fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename))
		filePath := filepath.Join(UPLOAD_DIR, fileName)
		out, err := os.Create(filePath)
		if err != nil {
			resp.OK = false
			resp.Message = "Internal server error"
			resp.ID = 1002 //"Could not create file"
			app.writeJSON(w, http.StatusInternalServerError, resp)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			resp.OK = false
			resp.Message = "Internal server error"
			resp.ID = 1003 //"Could not save file"
			app.writeJSON(w, http.StatusInternalServerError, resp)
			return
		}
		order.ImgFile = filePath
	}
	

	//Retrive Location ID
	order.LocationID, err = app.DB.GetLocationID(order.Division, order.District, order.Upazila)

	if err != nil {
		app.errorLog.Println(err)
		resp.OK = false
		resp.Message = "Internal server error"
		resp.ID = 2001 //"Could not retrive locationID from the database"
		app.writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	order.ServiceID = 1
	// //Retrive Service ID
	// order.ServiceID, err = app.DB.GetServiceID(order.Category, order.Service)

	// if err != nil {
	// 	app.errorLog.Println(err)
	// 	resp.OK = false
	// 	resp.Message = "Internal server error"
	// 	resp.ID = 2002 //"Could not retrive serviceID from the database"
	// 	app.writeJSON(w, http.StatusInternalServerError, resp)
	// 	return
	// }
	_, err = app.DB.InsertOrder(order)

	if err != nil {
		app.errorLog.Println(err)
		resp.OK = false
		resp.Message = "Internal server error"
		resp.ID = 2000 //"Could not save order to the database"
		app.writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp.OK = true
	resp.Message = "Your order is confirmed"
	resp.ID = 200 //"Could not save order to the database"
	app.writeJSON(w, http.StatusOK, resp)
}
