package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Command TODO...
type Command struct {
	gorm.Model
	Code        string    `selector:"div.usage code" json:"code" gorm:"primary_key"`
	Firmware    string    `selector:"div.meta span.label-success" json:"firmware" gorm:"primary_key"`
	Title       string    `selector:"div.meta h1" json:"title"`
	SourceURL   string    `json:"source"`
	Description string    `selector:"div.long p" json:"description"`
	Version     string    `json:"version"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
	DeletedAt   time.Time `json:"deleted"`
	Usage       []struct {
		Example       string `selector:"div.params table td.arg code" json:"example"`
		ExampleText   string `selector:"div.params table td:nth-child(2) p:first" json:"exampleText"`
		Parameter     string `selector:"div.params table td:nth-child(2) ul > li > code" json:"parameter"`
		ParameterText string `selector:"div.params table td:nth-child(2) ul > li > p" json:"parameterText"`
	}
}
