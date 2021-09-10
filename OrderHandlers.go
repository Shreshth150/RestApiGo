package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var prevOrderID = 0
var orders []Order

func createOrder(w http.ResponseWriter, r *http.Request) {

	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for _, order := range orders {
		if order.OrderID == inputOrderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

func updateOrder(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			var UpdatedOrder Order
			json.NewDecoder(r.Body).Decode(&UpdatedOrder)
			orders = append(orders, UpdatedOrder)
			json.NewEncoder(w).Encode(updateOrder)
			return
		}
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]

	for i, order := range orders {
		if order.OrderID == inputOrderId {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

}
