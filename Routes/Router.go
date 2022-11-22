package Routes

import (
	"iwogo/auth"
	"iwogo/category"
	"iwogo/handler"
	"iwogo/item"
	"iwogo/middleware"
	"iwogo/pages"
	"iwogo/product"
	"iwogo/settings"
	storebranch "iwogo/storeBranch"
	"iwogo/unit"
	"iwogo/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// REPOSITORY
	userRepository := user.NewRepository(db)
	categoryRepository := category.NewRepository(db)
	productRepository := product.NewRepository(db)
	pageRepository := pages.NewRepository(db)
	settingRepository := settings.NewRepository(db)
	unitRepository := unit.NewRepository(db)
	itemRepository := item.NewRepository(db)
	storeBranchRepository := storebranch.NewRepository(db)

	// SERVICE
	userService := user.NewService(userRepository)
	categoryService := category.NewService(categoryRepository)
	productService := product.NewService(productRepository)
	pageService := pages.NewService(pageRepository)
	settingService := settings.NewService(settingRepository)
	unitService := unit.NewService(unitRepository)
	itemService := item.NewService(itemRepository)
	storeBranchService := storebranch.NewService(storeBranchRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	categoryHandler := handler.NewCategoryHandler(categoryService, authService)
	productHandler := handler.NewProductHandler(productService)
	pageHandler := handler.NewPageHandler(pageService)
	settingHandler := handler.NewSettingHandler(settingService)
	unitHandler := handler.NewUnitHandler(unitService, authService)
	itemHandler := handler.NewItemHandler(itemService, authService)
	storeHandler := handler.NewStoreHandler(storeBranchService, authService)
	branchHandler := handler.NewBranchHandler(storeBranchService, authService)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"POST, GET, OPTIONS, PUT, DELETE"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "X-CSRF-Token"}

	router.Use(cors.New(config))

	// public routes
	router.Static("/storage", "./storage")
	//router.Use(static.Serve("/storage", static.LocalFile("/storage", false)))

	api := router.Group("/api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/user/email/checker", userHandler.ChekEmailAvailability)
	// AFTER LOGIN
	api.POST("/user/change/name", middleware.AuthMiddleware(authService, userService), userHandler.ChangeNameHandler)
	api.GET("/user/detail", middleware.AuthMiddleware(authService, userService), userHandler.FetchUser)
	api.GET("/users", middleware.AuthMiddleware(authService, userService), userHandler.GetAllUsers)
	api.POST("/user/change/password", middleware.AuthMiddleware(authService, userService), userHandler.ChangePassword)
	api.POST("/user/delete", middleware.AuthMiddleware(authService, userService), userHandler.DeleteUser)
	api.POST("/user/detail", middleware.AuthMiddleware(authService, userService), userHandler.GetUserByID)
	api.POST("/user/change", middleware.AuthMiddleware(authService, userService), userHandler.ChangeDetailHandler)

	// CATEGORY
	api.GET("/category", middleware.AuthMiddleware(authService, userService), categoryHandler.GetAllCategory)
	api.GET("/category/:id", categoryHandler.GetCategoryById)
	api.GET("/category/by/:slug", categoryHandler.GetCategoryBySlug)
	api.POST("/category/register", middleware.AuthMiddleware(authService, userService), categoryHandler.RegisterCategory)
	api.PUT("/category/update", middleware.AuthMiddleware(authService, userService), categoryHandler.UpdateCategory)
	api.POST("/category/delete", middleware.AuthMiddleware(authService, userService), categoryHandler.DeleteCategory)

	// PRODUCT
	api.POST("/product", middleware.AuthMiddleware(authService, userService), productHandler.GetAllProduct)
	api.POST("/product/create", middleware.AuthMiddleware(authService, userService), productHandler.CreateProductName)
	api.POST("/product/update", middleware.AuthMiddleware(authService, userService), productHandler.UpdateProduct)
	api.POST("/product/thumb", middleware.AuthMiddleware(authService, userService), productHandler.UpdateThumbProduct)
	api.POST("/product/del", middleware.AuthMiddleware(authService, userService), productHandler.DelProduct)
	api.POST("/product/detail", middleware.AuthMiddleware(authService, userService), productHandler.FindProductByIDHandler)

	// SLIDER
	api.POST("/slider/all", middleware.AuthMiddleware(authService, userService), productHandler.GetAllSliderHanlder)

	api.POST("/slider/del", middleware.AuthMiddleware(authService, userService), productHandler.DelSlider)

	// DISCOUNT
	api.POST("/discount/create", middleware.AuthMiddleware(authService, userService), productHandler.CreateDiscount)

	// SLIDER REALTION
	api.POST("/sliderrelation/create", middleware.AuthMiddleware(authService, userService), productHandler.CreateSliderRelationHandler)
	api.POST("/sliderrelation/get", middleware.AuthMiddleware(authService, userService), productHandler.GetSliderRelationByProductIDHanlder)
	api.POST("/sliderrelation/del", middleware.AuthMiddleware(authService, userService), productHandler.DelSliderRelationHanlder)

	// CATEGORY RELATION
	api.POST("/categoryrelation/create", middleware.AuthMiddleware(authService, userService), productHandler.CreateCategoryRelationHandler)
	api.POST("/categoryrelation/get", middleware.AuthMiddleware(authService, userService), productHandler.FindCategoryRelationHandler)
	api.POST("/categoryrelation/del", middleware.AuthMiddleware(authService, userService), productHandler.DelCategoryRelationHandler)

	// PAGES
	api.POST("/page/all", middleware.AuthMiddleware(authService, userService), pageHandler.FindAllPage)
	api.POST("/page/create", middleware.AuthMiddleware(authService, userService), pageHandler.CreatePage)
	api.POST("/page/id", middleware.AuthMiddleware(authService, userService), pageHandler.FindByid)
	api.POST("/page/del", middleware.AuthMiddleware(authService, userService), pageHandler.DelPage)
	api.POST("/page/update", middleware.AuthMiddleware(authService, userService), pageHandler.UpdatePage)

	// SETTING
	api.POST("/setting/find", middleware.AuthMiddleware(authService, userService), settingHandler.FindByid)
	api.POST("/setting/update", middleware.AuthMiddleware(authService, userService), settingHandler.UpdateSetting)
	api.POST("/setting/update/banner", middleware.AuthMiddleware(authService, userService), settingHandler.UpdateBannerSetting)

	api.POST("/slider/create", middleware.AuthMiddleware(authService, userService), productHandler.CreateSlider)

	// FRONTEND
	api.POST("/front/products", productHandler.GetAllProduct)
	api.POST("/front/search", productHandler.SearchProductHanlder)
	api.POST("/front/products/best", productHandler.GetAllProduct)
	api.POST("/front/product/categ", productHandler.GetProductByCateg)
	api.POST("/front/contact", settingHandler.FindByid)
	api.POST("/front/product/detail", productHandler.FindProductBySlugHandler)
	api.POST("/front/page/slug", pageHandler.FindBySlug)
	api.POST("/front/product/count/views", productHandler.UpdateViewsProductHanlder)
	api.POST("/front/product/slider", productHandler.GetSliderRelationByProductSlugHanlder)
	api.POST("/front/categories", categoryHandler.GetAllCategory)
	api.POST("/front/settings", settingHandler.FindByid)
	// api.GET("/pagination", productHandler.Pagination)

	//----------------------unit api Dasboard----------------------------
	api.GET("/front/unit", middleware.AuthMiddleware(authService, userService), unitHandler.GetAllUnit)
	api.POST("/unit/create", middleware.AuthMiddleware(authService, userService), unitHandler.CreateUnit)
	api.PUT("/unit/update", middleware.AuthMiddleware(authService, userService), unitHandler.UpdateUnit)
	api.POST("/unit/delete", middleware.AuthMiddleware(authService, userService), unitHandler.DeleteUnit)

	//----------------------item api Dasboard-----------------------------
	api.GET("/item/searchAll", middleware.AuthMiddleware(authService, userService), itemHandler.SeachAll)
	api.POST("/item/create", middleware.AuthMiddleware(authService, userService), itemHandler.CreateItem)
	api.PUT("/item/update", middleware.AuthMiddleware(authService, userService), itemHandler.UpdateItem)
	api.DELETE("/item/delete", middleware.AuthMiddleware(authService, userService), itemHandler.DeleteItem)

	//----------------------- Store api Dasboard --------------------------
	api.GET("/store/searchAll", middleware.AuthMiddleware(authService, userService), storeHandler.SeachAll)
	api.POST("/store/create", middleware.AuthMiddleware(authService, userService), storeHandler.CreateStore)
	api.PUT("/store/update", middleware.AuthMiddleware(authService, userService), storeHandler.UpdateStore)

	//----------------------- Branch api Dasboard --------------------------
	api.GET("/branch/searchAll", middleware.AuthMiddleware(authService, userService), branchHandler.SeachAll)
	api.POST("/branch/create", middleware.AuthMiddleware(authService, userService), branchHandler.CreateBranch)
	api.PUT("/branch/update", middleware.AuthMiddleware(authService, userService), branchHandler.UpdateBranch)

	return router
}
