package models

import (
	"time"
)

type Node1 struct {
	Id           int
	Url          string
	Link         string `gorm:"size:600"`
	Link1        string `gorm:"unique"`
	Ping         int
	AvgSpeed     int
	MaxSpeed     int
	FailCount    int
	SuccessCount int
	UpdateTime   time.Time
	CreateTime   time.Time
}

// func (Nodes) TabelName() string {
// 	return "node1"
// }
