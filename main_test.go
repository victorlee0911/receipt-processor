package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// test cases are given
func TestCalculatePoints(t *testing.T) {
	// Test case 1: Target Receipt
	receipt1 := receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
		},
		Total: "35.35",
	}
	expectedPoints1 := 28

	// Test case 2: M&M Corner Market Receipt
	receipt2 := receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}
	expectedPoints2 := 109

	tests := []struct {
		name	string
		data	receipt
		want	int
	}{
		{"Target", receipt1, expectedPoints1},
		{"M&M Corner Market Receipt", receipt2, expectedPoints2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			points := calculatePoints(tc.data)
			assert.Equal(t, points, tc.want)	
		})
	}
}

