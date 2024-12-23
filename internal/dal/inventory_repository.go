package dal

type InventoryRepository interface {
	AddItem()
	RetrieveAllItems()
	RetrieveSpecificItem()
	UpdateItem()
	DeleteItem()
}

type inventoryRepository struct{}

func (i *inventoryRepository) AddItem()              {}
func (i *inventoryRepository) RetrieveAllItems()     {}
func (i *inventoryRepository) RetrieveSpecificItem() {}
func (i *inventoryRepository) UpdateItem()           {}
func (i *inventoryRepository) DeleteItem()           {}

func NewInventoryRepository() InventoryRepository {
	return &inventoryRepository{}
}
