package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type FormC struct {
	Day       string
	Month     string
	Year      string
	Latitude  string
	Longitude string
}
type Coordinates struct {
	gorm.Model
	Date       string
	Latitude   string
	Longitude  string
	TempLow    string
	TempHigh   string
	Fahrenheit bool
}

type CoordinatesService struct {
	db *gorm.DB
}

var newLoggerCo logger.Interface = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,       // Disable color
	},
)

func firstCoordinates(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}
	return err
}

func NewCoordinatesService(connectionInfo string) (*CoordinatesService, error) {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{
		Logger: newLoggerCo,
	})
	if err != nil {
		return nil, fmt.Errorf("DB connection creation failed")
	}
	return &CoordinatesService{
		db: db,
	}, nil
}

func (c *CoordinatesService) CloseConCo() error {
	close, _ := c.db.DB()
	return close.Close()
}

func (c *CoordinatesService) AutoMigrateCo() error {
	if err := c.db.AutoMigrate(&Coordinates{}); err != nil {
		return err
	}
	return nil
}

func (c *CoordinatesService) CreateCo(coordinates *Coordinates) error {
	return c.db.Create(coordinates).Error
}

func (c *CoordinatesService) FreeCoordinates(date, latitude, longitude string, fahrenheit bool) (*Coordinates, error) {
	var freecoo Coordinates
	fc := c.db.Where("latitude = ? AND longitude = ? AND date = ? AND fahrenheit = ?", latitude, longitude, date, fahrenheit)
	err := firstCoordinates(fc, &freecoo)
	if err != nil {
		return nil, err
	}
	return &freecoo, nil
}
