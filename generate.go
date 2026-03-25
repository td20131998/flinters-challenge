package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func generateRandomCSV(totalCampaign int, filename string, totalRecords int) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write header
	header := "campaign_id,date,impressions,clicks,spend,conversions\n"
	_, err = file.WriteString(header)
	if err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	rand.Seed(time.Now().UnixNano())

	// Generate data
	numCampaigns := totalCampaign
	startDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < totalRecords; i++ {
		// Random campaign ID
		campaignID := fmt.Sprintf("CMP%03d", rand.Intn(numCampaigns)+1)

		// Random date within 12 months
		dayOffset := rand.Intn(365)
		date := startDate.AddDate(0, 0, dayOffset).Format("2006-01-02")

		// Random metrics - ensure values never cause rand.Intn(0) panic
		impressions := rand.Intn(100000) + 10000 // 10K to 110K (ensures >= 100)
		maxClicks := impressions / 10
		if maxClicks < 1 {
			maxClicks = 1
		}
		clicks := rand.Intn(maxClicks) + 1

		spend := float64(rand.Intn(100000))/100.0 + 10.0 // 10.0 to 1000.0

		maxConversions := clicks / 5
		if maxConversions < 1 {
			maxConversions = 1
		}
		conversions := rand.Intn(maxConversions) + 1

		// Write record
		record := fmt.Sprintf("%s,%s,%d,%d,%.2f,%d\n",
			campaignID,
			date,
			impressions,
			clicks,
			spend,
			conversions,
		)

		_, err := file.WriteString(record)
		if err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}

		// Progress indicator
		if (i+1)%100000 == 0 {
			fmt.Printf("Generated %d records...\n", i+1)
		}
	}

	fmt.Printf("Successfully generated %d records to %s\n", totalRecords, filename)
	return nil
}
