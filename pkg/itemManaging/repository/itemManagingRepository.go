package repository

import (
	"github.com/MarkTBSS/083_Item_Editing/entities"
	"github.com/MarkTBSS/083_Item_Editing/pkg/itemManaging/models"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *models.ItemEditingReq) (uint64, error)
}
