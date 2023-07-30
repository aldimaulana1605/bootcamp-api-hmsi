package main

import (
	"fmt"
	"os"

	connectDB "github.com/aldimaulana1605/bootcamp-api-hmsi/ConnectDB"
	"github.com/aldimaulana1605/bootcamp-api-hmsi/modules/customers/customerHandler"
	customerrepository "github.com/aldimaulana1605/bootcamp-api-hmsi/modules/customers/customerRepository"
	"github.com/aldimaulana1605/bootcamp-api-hmsi/modules/customers/customerUsecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {

	err := godotenv.Load("config/.env")

	if ; err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
}

DB_HOST := os.Getenv("DB_HOST")
DB_NAME := os.Getenv("DB_NAME")
DB_PORT := os.Getenv("DB_PORT")
DB_USER := os.Getenv("DB_USER")
DB_PASSWORD := os.Getenv("DB_PASSWORD")
DB_DRIVER := os.Getenv("DB_DRIVER")
PORT := os.Getenv("PORT")

log.Info().Msg("DB_HOST: " + DB_HOST)
log.Info().Msg("DB_PORT: " + DB_PORT)
log.Info().Msg("DB_USER: " + DB_USER)
log.Info().Msg("DB_PASS: " + DB_PASSWORD)
log.Info().Msg("DB_NAME: " + DB_NAME)
log.Info().Msg("PORT: " + PORT)
log.Info().Msg("DB_DRIVER: " + DB_DRIVER)


db, errConn := connectDB.GetConnPostgres(DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_DRIVER)

if errConn!= nil {
    log.Error().Msg(errConn.Error())
    os.Exit(1)
}
fmt.Println("Successfully connected!")

//DB struct initializes
// DB := query.DB{Conn: db}

// //Create Customer
// err = DB.Create(&query.Customers{
// 	Name: "Aldi Maulana Bahari",
// 	Phone: "085950262200",
// 	Email: "aldimaulana1605@gmail.com",
// 	Age: 18,
// })
// if err != nil {
// 	log.Error().Msg(errConn.Error())
// 	os.Exit(1)
// }
// fmt.Println("Insert Data Berhasil")

//Read Customer
// result, err := DB.Read()
// if err != nil {
// 	log.Error().Msg(errConn.Error())
// 	os.Exit(1)
// }
// fmt.Println(result)

// //Update Customer
// err = DB.Update(&query.Customers{
// 	Id: 1,
// 	Name: "Aldi Maulanaa",
//     Phone: "085950262222",
//     Email: "aldimaulana160503@gmail.com",
//     Age: 18,
// })
// if err != nil {
// 	log.Error().Msg(errConn.Error())
// 		os.Exit(1)
// }
// fmt.Printf("Updated Berhasil")

// //Delete Customer
// err = DB.Delete(3)
// if err != nil {
// 	log.Error().Msg(errConn.Error())
// 	os.Exit(1)
// }

// fmt.Printf("Deleted Berhasil")

// inisialisasi router
var  router = gin.Default()

// inisialisasi modules
customerRepo := customerrepository.NewCustomerRepository(db)
customerUC := customerUsecase.NewCustomerUsecase(customerRepo)
customerHandler.NewCustomerHandler(router, customerUC)

log.Info().Msg("Server running on port " + PORT)
router.Run(":"+ PORT)



}


