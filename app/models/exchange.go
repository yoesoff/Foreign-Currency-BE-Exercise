package models

import "time"

type Exchange struct {
	Id   uint64 `gorm:"primary_key" json:"id"`
	Date time.Time
	From string  `gorm:"size:50"`
	To   string  `gorm:"size:50"`
	Rate float32 `sql:"type:decimal(10,2);"`
}
