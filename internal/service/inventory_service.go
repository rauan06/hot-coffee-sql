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

func (s *inventoryService) AddItemToInventory()      {}
func (s *inventoryService) GetInventory()            {}
func (s *inventoryService) GetInventoryByID()        {}
func (s *inventoryService) UpdateInventoryItem()     {}
func (s *inventoryService) RemoveItemFromInventory() {}

func NewInventoryService() InventoryService {
	return &inventoryService{
		repo: dal.NewInventoryRepository(),
	}
}
