package models

import (
	"time"
)

type DataSet struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	Mid          int       `json:"mid"`
	Sid          int       `json:"sid"`
	UserId       int       `json:"userId"`
	CenterId     int       `json:"centerId"`
	NonPublic    bool      `json:"nonPublic"`
	VolumeId     int       `json:"volumeId"`
	ProviderType int       `json:"providerType"`
	SubPath      int       `json:"subPath"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (d DataSet) TableName() string {
	return "tb_dataset"
}
