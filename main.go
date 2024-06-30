package main

import ( 
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"unicode"
	"strings"
	"math"
	"strconv"
)

//receipt object
type item struct {
	ShortDescription	string	`json:"shortDescription"`
	Price	string	`json:"price"`
}

type receipt struct {
	Retailer	string	`json:"retailer"`
	PurchaseDate	string	`json:"purchaseDate"`
	PurchaseTime	string 	`json:"purchaseTime"`
	Items		[]item	`json:"items"`
	Total		string 	`json:"total"`
}

type pointStore struct {
	Points	int	`json:"points"`
}

type pointStoreId struct {
	ID	string `json:"id"`
}

var store = make(map[string]pointStore)

func main() {
	router := gin.Default()

	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", getPoints)	
	router.Run("0.0.0.0:8080")
}

//Read receipt body and process points -> store points only
func processReceipt(c *gin.Context) {
	var newReceipt receipt
	
	err := c.BindJSON(&newReceipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	
	
	points := calculatePoints(newReceipt)
	id := uuid.New().String()

	store[id] = pointStore{points}
	c.IndentedJSON(http.StatusCreated, pointStoreId{id})	
}

func calculatePoints(newReceipt receipt) int {
	pointTotal := 0
	// One point for every alphanumeric character in the retailer name.
	pointTotal += countAlphanum(newReceipt.Retailer)

	// 50 points if the total is a round dollar amount with no cents.
	totalFloat, err := strconv.ParseFloat(newReceipt.Total, 64)
	if err != nil{
		return -1
	}
	if math.Mod(totalFloat, 1.0) == 0.0 { 
		pointTotal += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if math.Mod(totalFloat, 0.25) == 0 { 
		pointTotal += 25
	}

	// 5 points for every two items on the receipt.
	pointTotal += 5 * (len(newReceipt.Items) / 2)

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range newReceipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc) % 3 == 0 {
			priceFloat, err := strconv.ParseFloat(item.Price, 64)
			if err != nil{
				return -1
			}
			pointTotal += int(math.Ceil(priceFloat * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd.
	date := newReceipt.PurchaseDate
	dayStr := date[len(date)-2:]
	day, err := strconv.Atoi(dayStr)
	if err != nil{
		return -1
	}

	if day % 2 == 1 {
		pointTotal += 6
	}	

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.	
	time := newReceipt.PurchaseTime
	hour, err := strconv.Atoi(time[:2])
	if err != nil{
		return -1
	}
	

	if hour >= 14 && hour < 16 {
		pointTotal += 10
	}
	return pointTotal
}

func countAlphanum(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func getPoints(c *gin.Context) {
	id := c.Param("id")
	point, ok := store[id]
	if ok {
		c.JSON(http.StatusOK, point)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error":"ID not found"})
	}
}
