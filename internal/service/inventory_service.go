package service

import "hot-coffee-sql/internal/dal"

type InventoryService interface {
	AddItemToInventory()
	GetInventory()
	GetInventoryByID()
	UpdateInventoryItem()
	RemoveItemFromInventory()
}

type inventoryService struct {
	repo dal.InventoryRepository
}

func (s *inventoryService) AddItemToInventory() {
	// Implementation here
}

func (s *inventoryService) GetInventory() {
	// Implementation here
}

func (s *inventoryService) GetInventoryByID() {
	// Implementation here
}

func (s *inventoryService) UpdateInventoryItem() {
	// Implementation here
}

func (s *inventoryService) RemoveItemFromInventory() {
	// Implementation here
}

func NewInventoryService() InventoryService {
	return &inventoryService{}
}
