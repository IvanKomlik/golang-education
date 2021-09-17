package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type equipmentInterface struct {
	IsDigital       bool
	TypeOfConnector string
}

type equipment struct {
	PartNumber       string             `json:"part_number"`
	Description      string             `json:"description"`
	InterfaceOnBoard equipmentInterface `json:"interface_on_board"` //TODO Need to map[int]equipmentInterface
}

var equipmentList = []equipment{
	{PartNumber: "010110CC", Description: "Some sort of video switch",
		InterfaceOnBoard: equipmentInterface{
			IsDigital:       true,
			TypeOfConnector: "DVI-D",
		},
	},
}

func getEquipment(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, equipmentList)
}

func main() {
	fmt.Println("Test")
	router := gin.Default()
	router.GET("/equipment", getEquipment)
	router.Run(":8080")
}
