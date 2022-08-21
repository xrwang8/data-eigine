package models

import "time"

type Model struct {
	ID           int       `gorm:"primary_key" json:"id,omitempty"`
	Name         string    `json:"name"`
	Mid          int       `json:"mid"`
	Sid          int       `json:"sid"`
	UserId       int       `json:"userId"`
	CenterId     int       `json:"centerId"`
	NonPublic    bool      `json:"nonPublic"`
	VolumeId     int       `json:"volumeId,omitempty"`
	ProviderType int       `json:"providerType"`
	SubPath      int       `json:"subPath,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (m Model) TableName() string {
	return "tb_model"
}
