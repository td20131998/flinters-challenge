package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerateRandomCSV(t *testing.T) {
	filename := "test_data.csv"
	totalRecords := 1000000 // 1 million records for testing
	totalCampaign := 50     // 50 unique campaigns

	err := generateRandomCSV(totalCampaign, filename, totalRecords)
	if err != nil {
		fmt.Printf("Error generating CSV: %v\n", err)
		t.Fatalf("Failed to generate CSV: %v", err)
	}

	fmt.Printf("Test CSV generated successfully: %s\n", filename)

	// Verify file exists
	_, err = os.Stat(filename)
	if err != nil {
		t.Fatalf("Generated file not found: %v", err)
	}

	// // Clean up test file after generation (uncomment to remove after test)
	// err = os.Remove(filename)
	// if err != nil {
	// 	fmt.Printf("Error removing test file: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Test file removed successfully: %s\n", filename)
}
