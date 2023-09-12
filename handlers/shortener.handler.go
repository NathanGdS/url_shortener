package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"url_shortener/models"

	"github.com/gorilla/mux"
)

func (h handler) Shortener(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var shortener models.UrlShortener
	json.Unmarshal(body, &shortener)

	// verify if url is defined
	if shortener.Url == "" || shortener.Url == "undefined" || shortener.Url == "nil" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode("Url is required")
		return
	}

	shortener.Create(shortener.Url)

	if result := h.DB.Create(&shortener); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}

func (h handler) GetAllShorteners(w http.ResponseWriter, r *http.Request) {
	var shorteners []models.UrlShortener

	if result := h.DB.Find(&shorteners); result.Error != nil {
		fmt.Println(result.Error)
	}

	// transform to dto
	var response []models.UrlShortenerDto
	var dto models.UrlShortenerDto
	for _, shortener := range shorteners {
		response = append(response, dto.ResponseDto(&shortener))
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h handler) GetShortener(w http.ResponseWriter, r *http.Request) {
	var shortener models.UrlShortener
	vars := mux.Vars(r)
	slug := vars["slug"]

	if result := h.DB.Where("slug = ?", slug).First(&shortener); result.Error != nil {
		fmt.Println(result.Error)
	}

	shortener.Open()

	h.DB.Save(&shortener)

	var response models.UrlShortenerDto
	response = response.ResponseDto(&shortener)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
