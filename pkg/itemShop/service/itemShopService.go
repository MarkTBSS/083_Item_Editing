package service

import _models "github.com/MarkTBSS/083_Item_Editing/pkg/itemShop/models"

type ItemShopService interface {
	Listing(itemFilter *_models.ItemFilter) (*_models.ItemResult, error)
}
