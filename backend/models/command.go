package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Command TODO...
type Command struct {
	// Ignore fields with gorm:"-" specified in struct tag.
	// Prevents duplicate columns.
	gorm.Model
	Code             string      `gorm:"unique_index" json:"code"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	SourceURL        string      `json:"source"`
	DocumentationURL string      `json:"documentationUrl"`
	Firmwares        []Firmware  `gorm:"foreignkey:CommandRefer"`
	Parameter        []Parameter `gorm:"foreignkey:ParameterRefer"`
	CreatedAt        time.Time   `gorm:"-" json:"created"`
	UpdatedAt        time.Time   `gorm:"-" son:"updated"`
	DeletedAt        time.Time   `gorm:"-" json:"deleted"`
}

// Firmware TODO...
type Firmware struct {
	gorm.Model
	Name             string    `json:"name"`
	RepRapWikiURL    string    `json:"repRapWikiUrl"`
	DocumentationURL string    `json:"documentationUrl"`
	Features         []Feature `gorm:"foreignkey:FirmwareRefer"`
	CreatedAt        time.Time `gorm:"-" json:"created"`
	UpdatedAt        time.Time `gorm:"-" son:"updated"`
	DeletedAt        time.Time `gorm:"-" json:"deleted"`
}

// Parameter TODO...
type Parameter struct {
	gorm.Model
	Command      Command   `gorm:"foreignkey:CommandRefer" json:"command"`
	CommandRefer uint      // Foreign key to Command
	Parameter    string    `json:"parameter"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `gorm:"-" json:"created"`
	UpdatedAt    time.Time `gorm:"-" son:"updated"`
	DeletedAt    time.Time `gorm:"-" json:"deleted"`
}

// Feature TODO...
type Feature struct {
	gorm.Model
	Command       Command   `gorm:"foreignkey:CommandRefer" json:"command"`
	Firmware      Firmware  `gorm:"foreignkey:FeatureRefer" json:"firmware"`
	CommandRefer  uint      // Foreign key to Command
	FirmwareRefer uint      // Foreign key to Firmware
	Supported     string    `gorm:"default:'Unknown'" json:"supported"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `gorm:"-" json:"created"`
	UpdatedAt     time.Time `gorm:"-" son:"updated"`
	DeletedAt     time.Time `gorm:"-" json:"deleted"`
}
