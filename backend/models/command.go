package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Command TODO...
type Command struct {
	gorm.Model
	CreatedAt time.Time `json:"created"` // Auto-updated by GORM.
	UpdatedAt time.Time `json:"updated"` // Auto-updated by GORM.
	DeletedAt time.Time `json:"deleted"` // Auto-updated by GORM.

	Code             string `selector:"div.usage code" json:"code"`
	ShortDescription string `selector:"div.meta h1" json:"shortDescription"`
	LongDescription  string `selector:"" json:"longDescription"`

	FirmwareSupport []struct {
		Firmware  string `selector:"" json:"firmware"`
		Supported string `selector:"" json:"supported"`
	}

	SourceURL        string `json:"source"`
	DocumentationURL string `json:"documentationUrl"`
	Usage            []struct {
		Example       string `selector:"div.params table td.arg code" json:"example"`
		ExampleText   string `selector:"div.params table td:nth-child(2) p:first" json:"exampleText"`
		Parameter     string `selector:"div.params table td:nth-child(2) ul > li > code" json:"parameter"`
		ParameterText string `selector:"div.params table td:nth-child(2) ul > li > p" json:"parameterText"`
	}
}
