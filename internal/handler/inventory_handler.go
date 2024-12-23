package handler

import (
	"hot-coffee-sql/internal/service"
	"net/http"
)

type InventoryHandler struct {
	InventoryService service.InventoryService
}

func NewInventoryHandler(inventoryService service.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		InventoryService: inventoryService,
	}
}

func (i *InventoryHandler) AddItemToInventory(w http.ResponseWriter, r *http.Request)      {}
func (i *InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request)            {}
func (i *InventoryHandler) GetInventoryByID(w http.ResponseWriter, r *http.Request)        {}
func (i *InventoryHandler) UpdateInventoryItem(w http.ResponseWriter, r *http.Request)     {}
func (i *InventoryHandler) RemoveItemFromInventory(w http.ResponseWriter, r *http.Request) {}
