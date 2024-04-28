package server

import (
	"github.com/MarkTBSS/083_Item_Editing/pkg/itemShop/controller"
	"github.com/MarkTBSS/083_Item_Editing/pkg/itemShop/repository"
	"github.com/MarkTBSS/083_Item_Editing/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := repository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := service.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := controller.NewItemShopControllerImpl(itemShopService)

	router.GET("/listing", itemShopController.Listing)
}
