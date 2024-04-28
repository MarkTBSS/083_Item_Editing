package main

import (
	"github.com/MarkTBSS/083_Item_Editing/config"
	"github.com/MarkTBSS/083_Item_Editing/databases"
	"github.com/MarkTBSS/083_Item_Editing/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
