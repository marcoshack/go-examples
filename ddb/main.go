package main

import (
	"fmt"
	"time"
)

type Item struct {
	Id          string    `dynamo:"ItemId,hash"`
	Name        string    `dynamo:"Name"`
	Description string    `dynamo:"Description"`
	CreatedAt   time.Time `dynamo:"CreatedAt"`
	UpdateAt    time.Time `dynamo:"UpdatedAt"`
}

type PermissionDefaults struct {
	DefaultState int `dynamo:"DefaultState"`
}

type AccessGroup struct {
	AccessGroupId string                        `dynamo:"AccessGroupId,hash"`
	Roles         []string                      `dynamo:"Roles"`
	Permissions   map[string]PermissionDefaults `dynamo:"Permissions"`
	CreatedAt     time.Time                     `dynamo:"CreatedAt"`
	UpdateAt      time.Time                     `dynamo:"UpdatedAt"`
}

func main() {
	fmt.Println("Hello, world!")
}
