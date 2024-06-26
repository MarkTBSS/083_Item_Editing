package service

import (
	"github.com/MarkTBSS/083_Item_Editing/entities"
	_itemManagingModel "github.com/MarkTBSS/083_Item_Editing/pkg/itemManaging/models"
	_itemManagingRepository "github.com/MarkTBSS/083_Item_Editing/pkg/itemManaging/repository"
	_itemShopModel "github.com/MarkTBSS/083_Item_Editing/pkg/itemShop/models"
	_itemShopRepository "github.com/MarkTBSS/083_Item_Editing/pkg/itemShop/repository"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
	itemShopRepository     _itemShopRepository.ItemShopRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{
		itemManagingRepository,
		itemShopRepository,
	}
}

func (s *itemManagingServiceImpl) Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	item := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}
	itemEntity, err := s.itemManagingRepository.Creating(item)
	if err != nil {
		return nil, err
	}
	return itemEntity.ChangeToItemModel(), nil
}

func (s *itemManagingServiceImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error) {
	updatedItemID, err := s.itemManagingRepository.Editing(itemID, itemEditingReq)
	if err != nil {
		return nil, err
	}
	itemEntity, err := s.itemShopRepository.FindByID(updatedItemID)
	if err != nil {
		return nil, err
	}
	return itemEntity.ChangeToItemModel(), nil
}
