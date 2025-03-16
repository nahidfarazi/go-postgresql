package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nahidfarazi/go-postgresql/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=127.0.0.1 user=postgres password=512609 dbname=goPostgresql port=5432 sslmode=disable"

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db connection fail")
	}
	db.AutoMigrate(&user.User{}, &user.Address{}, &user.Contact{})
}
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var users []user.User
	db.Preload("Address").Preload("Contact").Find(&users)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	param := chi.URLParam(r, "id")
	var users user.User
	if err := db.Preload("Address").Preload("Contact").First(&users, param).Error; err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var user user.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	db.Create(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	param := chi.URLParam(r, "id")
	var user user.User
	if err := db.First(&user, param).Error; err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	db.Save(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	param := chi.URLParam(r, "id")
	var user user.User
	if err := db.First(&user, param).Error; err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}
	db.Delete(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User delete successfully"})
}
