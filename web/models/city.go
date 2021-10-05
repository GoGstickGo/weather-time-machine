package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Form struct {
	Name  string
	Day   string
	Month string
	Year  string
}
type Cities struct {
	gorm.Model
	City        string
	Date        string
	CountryCode string
	TempLow     string
	TempHigh    string
	Fahrenheit  bool
}

type CityService struct {
	db *gorm.DB
}

var newLogger logger.Interface = logger.New(
	log.New(file, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second,  // Slow SQL threshold
		LogLevel:                  logger.Error, // Log level
		IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,        // Disable color
	},
)

func firstCity(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}
	return err
}

func NewCityService(connectionInfo string) (*CityService, error) {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("DB connection creation failed")
	}
	return &CityService{
		db: db,
	}, nil
}

func (c *CityService) CloseCon() error {
	close, _ := c.db.DB()
	return close.Close()
}

func (c *CityService) AutoMigrate() error {
	if err := c.db.AutoMigrate(&Cities{}); err != nil {
		return err
	}
	return nil
}

func (c *CityService) Create(city *Cities) error {
	return c.db.Create(city).Error
}

func (c *CityService) FreeCity(date, city string, fahrenheit bool) (*Cities, error) {
	var freecity Cities
	fc := c.db.Where("city = ? AND date = ? AND fahrenheit = ?", city, date, fahrenheit)
	err := firstCity(fc, &freecity)
	if err != nil {
		return nil, err
	}
	return &freecity, nil
}
