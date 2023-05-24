package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type constractor struct {
    ID     string  `json:"id"`
    Constractors_name string  `json:"constractors_name"`
}

var constractors = []constractor{
    {
        ID: "1", Constractors_name: "First", 
    },
    {
        ID: "2", Constractors_name: "Second", 
    },
}

func getConstractors(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, constractors)
}

func postConstractor(c *gin.Context) {
    var newConstractor constractor

    if err := c.BindJSON(&newConstractor); err != nil {
        return
    }

    constractors = append(constractors, newConstractor)

    c.IndentedJSON(http.StatusCreated, newConstractor)
}

func getConstractorById(c *gin.Context) {
    id := c.Param("id")

    for _, a := range constractors {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "constractor not found"})
}

func main() {
    router := gin.Default()

    router.POST("/constractors", postConstractor)
    router.GET("/constractors", getConstractors)
    router.GET("/constractors/:id", getConstractorById)
    router.Run("localhost:8080")
}