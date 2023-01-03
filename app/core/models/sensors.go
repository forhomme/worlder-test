package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Sensors struct {
	Base
	Value string `json:"value"`
	ID1   string `json:"ID1" gorm:"column:id1"`
	ID2   string `json:"ID2" gorm:"column:id2"`
}

func (s *Sensors) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.NewV4()
	return nil
}
