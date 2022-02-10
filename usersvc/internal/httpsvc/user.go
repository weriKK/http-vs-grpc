package httpsvc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"usersvc/internal/grpcapi"
)

type UserDataRequest struct {
	ID int `json:"id"`
}

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

type UserDataResponse struct {
	ID        int       `json:"id"`
	User      User      `json:"user"`
	IPAddress []string  `json:"ip_address"`
	Date      time.Time `json:"date"`
	Cars      []Car     `json:"cars"`
}

func (h *httpsvc) userDataHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var req UserDataRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := User{
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "firstname.lastname@email.moo",
		Gender:    "yes",
	}

	carData, err := h.fetchCarData(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := UserDataResponse{
		ID:        req.ID,
		User:      user,
		IPAddress: []string{"127.0.0.1", "192.168.0.1", "10.10.10.10", "0.0.0.0"},
		Date:      time.Now(),
		Cars:      carData,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type CarDataRequest struct {
	User User `json:"user"`
}

type CarDataResponse struct {
	User User  `json:"user"`
	Cars []Car `json:"cars"`
}

func (h *httpsvc) fetchCarData(user User) ([]Car, error) {

	if h.isGrpcClient {
		return h.fetchCarDataGRPC(user)
	}

	return h.fetchCarDataHTTP(user)
}

func (h *httpsvc) fetchCarDataHTTP(user User) ([]Car, error) {

	req := CarDataRequest{
		User: user,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal CarDataRequest: %w", err)
	}

	resp, err := http.Post(fmt.Sprintf("http://127.0.0.1:%d/car", h.otherport), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("failed to POST CarDataRequest: %w", err)
	}

	defer resp.Body.Close()

	var res CarDataResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("failed to Decode CarDataResponse: %w", err)
	}

	return res.Cars, nil
}

func (h *httpsvc) fetchCarDataGRPC(user User) ([]Car, error) {

	c := grpcapi.NewCarDataServiceClient(h.grpcClientConn)

	req := grpcapi.CarDataRequest{
		User: &grpcapi.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.LastName,
			Gender:    user.Gender,
		},
	}

	res, err := c.GetCarData(context.Background(), &req)
	if err != nil {
		return nil, fmt.Errorf("failed to get car data via gRPC: %w", err)
	}

	cars := []Car{}
	for _, v := range res.Cars {
		cars = append(cars, Car{Make: v.Make, Model: v.Model, Color: v.Color})
	}

	return cars, nil
}
