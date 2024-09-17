package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd/orion-pay/application"
	"github.com/mfitrahrmd/orion-pay/binding"
	"github.com/mfitrahrmd/orion-pay/model"
	"gorm.io/gorm"
)

func main() {
	app := application.NewApplication()
	app.Setup()
	app.SetupDB()
	r := app.Router.Group("/api")
	{
		r.Static("/public", "./public")
		r = r.Group("/v1")
		{
			users := r.Group("/users")
			{
				users.POST("", func(ctx *gin.Context) {
					var userBinding binding.UserPost
					if err := ctx.ShouldBindJSON(&userBinding); err != nil {
						ctx.JSON(400, map[string]any{
							"message": err.Error(),
						})

						return
					}
					user := userBinding.ToUser()
					if err := app.UserRepository.CreateUser(&user); err != nil {
						ctx.JSON(400, map[string]any{
							"message": err.Error(),
						})

						return
					}
					ctx.JSON(201, user)
				})

				users.GET("", func(ctx *gin.Context) {
					var users []model.User
					if err := app.UserRepository.GetUsers(&users); err != nil {
						fmt.Println(err)
						ctx.JSON(400, map[string]any{
							"message": "bad request",
						})

						return
					}
					ctx.JSON(200, users)
				})

				users.GET("/:id", func(ctx *gin.Context) {
					id := ctx.Param("id")
					intId, err := strconv.Atoi(id)
					if err != nil {
						ctx.JSON(400, map[string]any{
							"message": "bad request",
						})

						return
					}
					user := model.User{
						Model: gorm.Model{
							ID: uint(intId),
						},
					}
					if err := app.UserRepository.GetUser(&user); err != nil {
						ctx.JSON(400, map[string]any{
							"message": "bad request",
						})

						return
					}
					ctx.JSON(200, user)
				})
			}

			wallets := r.Group("/wallets")
			{
				wallets.GET("", func(ctx *gin.Context) {
					var wallets []model.Wallet
					app.Db.Find(&wallets)
					ctx.JSON(200, wallets)
				})
			}

			transfer := r.Group("/transfers")
			{
				transfer.POST("", func(ctx *gin.Context) {
					var transferBinding binding.TransferPost
					if err := ctx.ShouldBindJSON(&transferBinding); err != nil {
						ctx.JSON(400, map[string]any{
							"message": err.Error(),
						})

						return
					}
					transfer := transferBinding.ToTransfer()
					err := app.TransferRepository.CreateTransfer(transfer)
					if err != nil {
						ctx.JSON(400, map[string]any{
							"message": err.Error(),
						})

						return
					}
					ctx.JSON(201, transfer)
				})

				transfer.GET("", func(ctx *gin.Context) {
					var transfers []model.Transfer
					walletID := ctx.Query("walletID")
					if len(walletID) > 0 {
						app.Db.Preload("SentEntry").Preload("ReceivedEntry").Find(&transfers, "wallet_id = ?", walletID)
					} else {
						app.Db.Preload("SentEntry").Preload("ReceivedEntry").Find(&transfers)
					}
					ctx.JSON(200, transfers)
				})
			}
		}
	}
	app.Run()
}
