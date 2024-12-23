package router

import (
	"hot-coffee-sql/internal/handler"
	"hot-coffee-sql/internal/service"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// orderHandler := handler.NewOrderHandler()
	// menuHandler := handler.NewMenuHandler()

	// Order routes
	// mux.HandleFunc("POST /order/", orderHandler.CreateOrder)
	// mux.HandleFunc("GET /order/", orderHandler.RetrieveAllOrders)
	// mux.HandleFunc("GET /order/{id}", orderHandler.RetrieveSepcificOrder)
	// mux.HandleFunc("PUT /order/{id}", orderHandler.UpdateOrder)
	// mux.HandleFunc("DELETE /order/{id}", orderHandler.DeleteOrder)
	// mux.HandleFunc("POST /order/{id}/close", orderHandler.CloseOrder)

	// Menu routes
	// mux.HandleFunc("POST /menu/", handler.GetMenu)
	// mux.HandleFunc("/menu/", handler.GetMenu)
	// mux.HandleFunc("/menu/{id}", handler.GetMenuItemByID)

	// =================
	// Inventory routes
	// =================
	inventorySvc := service.NewInventoryService()
	inventoryHandler := handler.NewInventoryHandler(inventorySvc)

	mux.HandleFunc("POST /inventory/", inventoryHandler.AddItemToInventory)
	mux.HandleFunc("GET /inventory/", inventoryHandler.GetInventory)
	mux.HandleFunc("GET /inventory/{id}", inventoryHandler.GetInventoryByID)
	mux.HandleFunc("PUT /inventory/{id}", inventoryHandler.UpdateInventoryItem)
	mux.HandleFunc("DELETE /inventory/{id}", inventoryHandler.RemoveItemFromInventory)

	// Register reports routes
	// mux.HandleFunc("/reports/total-sales", handler.GetTotalSales)
	// mux.HandleFunc("/reports/popular-items", handler.GetPopularItems)

	return mux
}
