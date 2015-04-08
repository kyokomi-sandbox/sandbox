package main

import "time"

type User struct {
	ID       int64 `migu:"pk"`
	UserName string
	AgeNum   int64
	RewardID string
	TestFlag bool   `migu:"default:false"`
	CreateAt time.Time `migu:"default:now()"`
}
