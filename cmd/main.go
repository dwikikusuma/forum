package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mySimpleFprum/internal/configs"
	membershipHandlers "mySimpleFprum/internal/handlers/memberships"
	membershipRepo "mySimpleFprum/internal/repository/memberships"
	membershipService "mySimpleFprum/internal/service/memberships"

	postHandler "mySimpleFprum/internal/handlers/posts"
	postRepository "mySimpleFprum/internal/repository/posts"
	postService "mySimpleFprum/internal/service/posts"
	"mySimpleFprum/pkg/internalsql"
)

func main() {
	r := gin.Default()
	fmt.Println("Hello")

	if err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	); err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cnfg := configs.Get()
	fmt.Println(cnfg)

	db, err := internalsql.Connect(cnfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Failed to Connect With DB")
	}

	membershipRepository := membershipRepo.NewRepository(db)
	membershipSvc := membershipService.NewService(cnfg, membershipRepository)
	membershipHandler := membershipHandlers.NewHandler(r, membershipSvc)

	postRepo := postRepository.NewRepository(db)
	postSvc := postService.NewService(postRepo)
	postHndlr := postHandler.NewHandler(r, postSvc)

	postHndlr.RegisterRoutes()
	membershipHandler.RegisterRoutes()

	_ = r.Run(cnfg.Service.Port)
}
