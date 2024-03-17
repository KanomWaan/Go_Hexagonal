package main

import (
	"fmt"
	"hexagonal/handler"
	"hexagonal/repository"
	"hexagonal/service"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()

	courseRepository := repository.NewCourseRepoDB(db)
	// _ = courseRepository
	courseRepositoryMock := repository.NewCourseRepositoryMock()
	_ = courseRepositoryMock

	courseService := service.NewCourseService(courseRepository)
	courseHandler := handler.NewCourseHandler(courseService)

	router := mux.NewRouter()
	router.HandleFunc("/courses", courseHandler.GetCourses).Methods(http.MethodGet)
	router.HandleFunc("/courses/{id:[0-9]+}", courseHandler.GetCourseByID).Methods(http.MethodGet)

	log.Printf("server is running at port %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)

}

func initConfig() {
	//กำหนดชื่อไฟล์ config และชนิดไฟล์
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	//ตรวจสอบก่อนว่ามี .env ไหม ถ้ามีให้ใช้ ก็ไม่มีไปใช้ config.yaml
	viper.AutomaticEnv()
	// สามารถกำหนด config ได้เอง โดยใส่ไว้หน้า go run main.go เช่น APP_PORT=8080 go run main.go ใช้ _ แทน .
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	//?parseTime=true ใช้สำหรับกรณี time.time
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"))

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		log.Fatal(err)
	}
	// set timeout ในการเชื่อมต่อ db
	db.SetConnMaxLifetime(3 * time.Minute)
	// set max open connection
	db.SetMaxOpenConns(10)
	// set max idle connection
	db.SetMaxIdleConns(10)

	return db
}
