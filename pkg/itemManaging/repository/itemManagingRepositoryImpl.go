package repository

import (
	"github.com/MarkTBSS/083_Item_Editing/databases"
	"github.com/MarkTBSS/083_Item_Editing/entities"
	"github.com/MarkTBSS/083_Item_Editing/pkg/itemManaging/exception"
	"github.com/MarkTBSS/083_Item_Editing/pkg/itemManaging/models"
	"github.com/labstack/echo/v4"
)

type itemMangingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) ItemManagingRepository {
	return &itemMangingRepositoryImpl{db, logger}
}

func (r *itemMangingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)
	err := r.db.Connect().Create(itemEntity).Scan(item).Error
	if err != nil {
		r.logger.Error("Item creating failed:", err.Error())
		return nil, &exception.ItemCreating{}
	}
	return item, nil
}

func (r *itemMangingRepositoryImpl) Editing(itemID uint64, itemEditingReq *models.ItemEditingReq) (uint64, error) {
	err := r.db.Connect().Model(&entities.Item{}).Where(
		"id = ?", itemID,
	).Updates(
		itemEditingReq,
	).Error
	if err != nil {
		r.logger.Error("Editing item failed:", err.Error())
		return 0, &exception.ItemEditing{}
	}
	return itemID, nil
}
