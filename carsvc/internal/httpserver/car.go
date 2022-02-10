package httpserver

import (
	"encoding/json"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

type Car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Color string `json:"color"`
}

type CarDataRequest struct {
	User User `json:"user"`
}

type CarDataResponse struct {
	User User  `json:"user"`
	Cars []Car `json:"cars"`
}

func (h *httpserver) carDataHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var req CarDataRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cars := []Car{
		{Make: "Honda", Model: "Civic", Color: "Blue"},
		{Make: "Mercedes", Model: "Cla", Color: "Black"},
		{Make: "Ford", Model: "Focus", Color: "Gray"},
	}

	res := CarDataResponse{
		User: req.User,
		Cars: cars,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
