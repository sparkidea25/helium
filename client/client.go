package client

import (
	"fmt"
	"time"

	"github.com/sparkidea25/helium/db"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	clientid   uint32    `gorm:"primary_key;auto_increment" json:"clientid"`
	createdAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdat"`
	updatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedat"`
	activeAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"activeat"`
	clientName string    `gorm:"default:CURRENT_TIMESTAMP" json:"clientname"`
}

func Create(name string) *Client {
	c := &Client{clientName: name}
	res := db.Get().Create(&c)
	if res.Error != nil {
		fmt.Println("Failed to create user")
	}
	return c
}

// func Delete(client *Client) error {

// }
