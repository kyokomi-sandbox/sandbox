package main

type User struct {
	ID       int64 `migu:"pk"`
	UserName string
	AgeNum   int64
	RewardID string
}
