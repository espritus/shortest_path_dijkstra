package main

import (
	"dijkstraproject/dijkstra" // Full import path relative to your module
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type path struct {
	StartPoint       int      `json:"from"`
	EndPoint         int      `json:"to"`
	Path             []string `json:"path"`
	ShortestDistance float32  `json:"distance"`
	PathIndex        []int    `json:"index"`
}

func getPaths(c *gin.Context) {
	startStr := c.Query("start")
	endStr := c.Query("end")

	// Convert start and end strings to integers
	start, err := strconv.Atoi(startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'start' parameter"})
		return
	}

	end, err := strconv.Atoi(endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'end' parameter"})
		return
	}

	// // Run Dijkstra's algorithm
	// distances, previous := dijkstra.DijkstraParallel(start, end)

	// // Get the shortest path
	// shortestPath := dijkstra.GetPath(previous, end)

	shortestPath, shortestDistance, myindex := dijkstra.Initialize(start, end)
	response := struct {
		StartPoint       int      `json:"start"`
		EndPoint         int      `json:"end"`
		Path             []string `json:"path"`
		ShortestDistance float64  `json:"distance"`
		PathIndex        []int    `json:"index"`
	}{
		StartPoint:       start,
		EndPoint:         end,
		Path:             shortestPath,
		ShortestDistance: shortestDistance,
		PathIndex:        myindex,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.GET("/shortestpath", getPaths)
	router.Run("0.0.0.0:8080")
  }
