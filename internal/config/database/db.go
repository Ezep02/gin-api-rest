package database

import (
	"fmt"

	"github.com/go-api-rest/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func initConfig() {
	viper.SetConfigName(".env") // nombre del archivo de configuración (sin extensión)
	viper.SetConfigType("env")  // tipo de archivo de configuración
	viper.AddConfigPath(".")    // ruta donde buscar el archivo de configuración

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("[ERROR] leyendo el archivo de configuración: %v", err)
	}
}

func NewConnection() *gorm.DB {

	initConfig()

	dbConfig := DataBaseConfig{
		Host:     viper.GetString("DB_HOST"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
		Port:     viper.GetString("DB_PORT"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[DB ERROR]:", r)
		}
	}()

	if err != nil {
		panic("No se pudo establecer una conexion con la base de datos \n")
	}

	// Auto-migrate the schema
	connection.AutoMigrate(&models.User{})

	return connection

}
