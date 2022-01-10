package main

import (
	"capstone/app/middlewares"
	"capstone/app/routes"
	roleUsecase "capstone/business/roles"
	roleController "capstone/controllers/roles"
	roleRepo "capstone/driver/database/roles"

	userUsecase "capstone/business/users"
	userController "capstone/controllers/users"
	userRepo "capstone/driver/database/users"

	complexUsecase "capstone/business/complexs"
	complexController "capstone/controllers/complexs"
	complexRepo "capstone/driver/database/complexs"

	facilityUsecase "capstone/business/facilities"
	facilityController "capstone/controllers/facilities"
	facilityRepo "capstone/driver/database/facilities"

	buildingUsecase "capstone/business/buildings"
	buildingController "capstone/controllers/buildings"
	buildingRepo "capstone/driver/database/buildings"

	unitUsecase "capstone/business/units"
	unitController "capstone/controllers/units"
	unitRepo "capstone/driver/database/units"

	"capstone/driver/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug mode")
	}
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&userRepo.User{})
	DB.AutoMigrate(&roleRepo.Role{})
	DB.AutoMigrate(&complexRepo.Complex{})
	DB.AutoMigrate(&facilityRepo.Facility{})
	DB.AutoMigrate(&buildingRepo.Building{})
	DB.AutoMigrate(&unitRepo.Unit{})
}

func main() {
	ConfigDB := mysql.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}

	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString("JWT.secretKey"),
		ExpiresDuration: viper.GetInt("JWT.expired_time"),
	}

	DB := ConfigDB.InitialDB()
	DBMigrate(DB)
	timeoutcontext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()
	
	userRepoInterface := userRepo.NewUserRepository(DB)
	userUseCaseInterface := userUsecase.NewUserUseCase(userRepoInterface, timeoutcontext, &configJWT)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	roleRepoInterface := roleRepo.NewRoleRepository(DB)
	roleUseCaseInterface := roleUsecase.NewRoleUseCase(roleRepoInterface, timeoutcontext)
	roleControllerInterface := roleController.NewRoleController(roleUseCaseInterface)

	complexRepoInterface := complexRepo.NewComplexRepository(DB)
	complexuseCaseInterface := complexUsecase.NewComplexUseCase(complexRepoInterface, timeoutcontext)
	complexControllerInterface := complexController.NewComplexController(complexuseCaseInterface)

	facilityRepoInterface := facilityRepo.NewFacilityRepository(DB)
	facilityUseCaseInterface := facilityUsecase.NewFacilityUseCase(facilityRepoInterface, timeoutcontext)
	facilityContollerInterface := facilityController.NewFacilityController(facilityUseCaseInterface)

	buildingRepoInterface := buildingRepo.NewBuildingRepository(DB)
	buidingUseCaseInterface := buildingUsecase.NewBuildingUseCase(buildingRepoInterface, timeoutcontext)
	buildingControllerInterface := buildingController.NewBuildingController(buidingUseCaseInterface)

	unitRepoInterface := unitRepo.NewUnitRepository(DB)
	unitUseCaseInterface := unitUsecase.NewUnitUseCase(unitRepoInterface, timeoutcontext)
	unitControllerInterface := unitController.NewUnitController(unitUseCaseInterface)


	routesInit := routes.RouteControllerList{
		UserController: *userControllerInterface,
		RoleController: *roleControllerInterface,
		ComplexController: *complexControllerInterface,
		FacilityController: *facilityContollerInterface,
		BuildingController: *buildingControllerInterface,
		UnitController: *unitControllerInterface,

	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}