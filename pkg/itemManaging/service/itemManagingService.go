package service

import (
	_itemManagingModel "github.com/MarkTBSS/083_Item_Editing/pkg/itemManaging/models"
	_itemShopModel "github.com/MarkTBSS/083_Item_Editing/pkg/itemShop/models"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error)
}
