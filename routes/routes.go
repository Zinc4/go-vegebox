package routes

import (
	"mini-project/admin"
	"mini-project/auth"
	"mini-project/cart"
	chatbot "mini-project/chatbox"
	"mini-project/handler"
	"mini-project/middleware"
	"mini-project/order"
	"mini-project/payment"
	"mini-project/product"
	"mini-project/transaction"
	"mini-project/user"
	"mini-project/utils/database"

	"github.com/labstack/echo/v4"
)

func NewRouter(router *echo.Echo) {
	userRepository := user.NewRepository(database.DB)
	adminRepository := admin.NewRepository(database.DB)

	productRepository := product.NewRepository(database.DB)
	cartRepository := cart.NewRepository(database.DB)
	orderRepository := order.NewRepository(database.DB)
	transactionRepository := transaction.NewRepository(database.DB)

	productUsecase := product.NewUsecase(productRepository)
	cartUsecase := cart.NewUsecase(cartRepository)
	orderUsecase := order.NewUsecase(orderRepository)
	paymentUsecase := payment.NewUsecase()
	transactionUsecase := transaction.NewUsecase(transactionRepository, orderRepository, paymentUsecase)

	authUsecase := auth.NewUsecase()
	userUsecase := user.NewUsecase(userRepository)
	adminUsecase := admin.NewUsecase(adminRepository)

	productHandler := handler.NewProductHandler(productUsecase)
	cartHandler := handler.NewCartHandler(cartUsecase, productUsecase)
	orderHandler := handler.NewOrderHandler(orderUsecase, cartUsecase, productUsecase)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase, orderUsecase)

	userHandler := handler.NewUserHandler(userUsecase, authUsecase)
	adminHandler := handler.NewAdminHandler(adminUsecase, authUsecase)

	chatAI := chatbot.NewChatAI()

	api := router.Group("api/v1")

	api.POST("/chatbox", chatAI.HandleChatCompletion)

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/verify-email", userHandler.VerifyEmail)
	api.POST("/resend-otp", userHandler.ResendOTP)

	// user
	api.PATCH("/avatar", middleware.AuthMiddleware(authUsecase, userUsecase, userHandler.UploadAvatar))
	api.PATCH("/profile", middleware.AuthMiddleware(authUsecase, userUsecase, userHandler.UpdateProfile))

	// product
	api.GET("/products", productHandler.GetProducts)
	api.GET("/products/:id", productHandler.GetProductByID)

	// cart
	api.POST("/carts", middleware.AuthMiddleware(authUsecase, userUsecase, cartHandler.NewCart))
	api.GET("/cart/:id", middleware.AuthMiddleware(authUsecase, userUsecase, cartHandler.GetCart))
	api.GET("/viewcarts", middleware.AuthMiddleware(authUsecase, userUsecase, cartHandler.GetCarts))
	api.POST("/cart/:id", middleware.AuthMiddleware(authUsecase, userUsecase, cartHandler.AddProductToCart))
	api.PUT("/cart/:cart_id/product/:product_id", middleware.AuthMiddleware(authUsecase, userUsecase, cartHandler.UpdateCartItem))
	api.DELETE("/cart/:cart_id/product/:product_id", middleware.AuthMiddleware(authUsecase, userUsecase, cartHandler.DeleteProductFromCart))
	api.DELETE("/cart/:id", middleware.AuthMiddleware(authUsecase, userUsecase, cartHandler.DeleteCart))
	api.GET("/cart/:id/checkout", middleware.AuthMiddleware(authUsecase, userUsecase, cartHandler.CheckOut))

	// order
	api.GET("/order/:cart_id", middleware.AuthMiddleware(authUsecase, userUsecase, orderHandler.CreateOrder))
	api.GET("/orders", middleware.AuthMiddleware(authUsecase, userUsecase, orderHandler.GetOrders))
	api.GET("/vieworder/:order_id", middleware.AuthMiddleware(authUsecase, userUsecase, orderHandler.GetOrder))

	// transaction
	api.GET("/transactions/:order_id", middleware.AuthMiddleware(authUsecase, userUsecase, transactionHandler.CreateTransaction))
	api.GET("/transactions", middleware.AuthMiddleware(authUsecase, userUsecase, transactionHandler.GetUserTransaction))
	api.POST("/transactions/payment-callback", transactionHandler.GetPaymentCallback)

	// admin
	api.GET("/users", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.GetAllUsers))
	api.GET("/users/transactions", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.GetAllUserTransactions))
	api.DELETE("/users/:id", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.DeleteUser))
	api.POST("/products", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.CreateProduct))
	api.PUT("/products/:id", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.UpdateProduct))
	api.DELETE("/products/:id", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.DeleteProduct))
	api.POST("/category", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.CreateCategory))
	api.DELETE("/category/:id", middleware.AuthMiddleware(authUsecase, userUsecase, adminHandler.DeleteCategory))

}
