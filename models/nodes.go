package models

import (
	"time"
)

type Nodes struct {
	Id         int
	Url        string
	Link       string `gorm:"unique;size:600"`
	Ping       int
	AvgSpeed   int
	MaxSpeed   int
	FailCount  int
	CreateTime time.Time
	UpdateTime time.Time
}
