package handlers

import (
	"encoding/json"
	"errors"
	"go-api/internal/model/entity"
	"go-api/pkg/config"
	"go-api/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	var products []entity.Product
	err := config.DB.Find(&products).Error
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error"})
		return
	}
	if len(products) == 0 {
		utils.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "data empty"})
		return
	}
	utils.RespondJSON(w, http.StatusOK, products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	vars := mux.Vars(r)
	id := vars["id"]
	err := config.DB.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondJSON(w, http.StatusNotFound, map[string]interface{}{"message": "product not found"})
			return
		}
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error"})
		return
	}
	utils.RespondJSON(w, http.StatusOK, product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]interface{}{"message": "error"})
		return
	}
	config.DB.Create(&product)
	utils.RespondJSON(w, http.StatusCreated, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	vars := mux.Vars(r)
	id := vars["id"]
	err := config.DB.Delete(&product, id).Error

	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error"})
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "success"})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	vars := mux.Vars(r)
	id := vars["id"]
	err := config.DB.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.RespondJSON(w, http.StatusNotFound, map[string]interface{}{"message": "product not found"})
			return
		}
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error"})
		return
	}
	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]interface{}{"message": "error"})
		return
	}
	if err := config.DB.Model(&product).Updates(updates).Error; err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "error updating product"})
		return
	}
	utils.RespondJSON(w, http.StatusOK, product)
}
