package corgis

import "github.com/jinzhu/gorm"

type RawVMData struct {
	gorm.Model
	VMName       string  `gorm:"unique"`
	LatencyRead  float64 `gorm:"type:double precision"`
	LatencyWrite float64 `gorm:"type:double precision"`
	IOPSRead     float64 `gorm:"type:double precision"`
	IOPSWrite    float64 `gorm:"type:double precision"`
	ISSSD        bool
}

func (RawVMData) TableName() string {
	return "tiramisu_rawdata"
}

type TiramisuStorage struct {
	VMName          string `gorm:"primary_key;type:varchar(100)"`
	CurrentPool     string `gorm:"type:varchar(10)"`
	AppropriatePool string `gorm:"type:varchar(10)";column:appropiate_pool`
	Notice          int    `gorm:"type:integer"`
}

func (TiramisuStorage) TableName() string {
	return "tiramisu_storage"
}
