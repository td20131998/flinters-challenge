package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func writeCSV(filename string, metrics []CampaignMetrics) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write header
	header := "campaign_id,total_impressions,total_clicks,total_spend,total_conversions,CTR,CPA\n"
	_, err = file.WriteString(header)
	if err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write data
	for _, m := range metrics {
		cpaStr := "null"
		if !math.IsNaN(m.CPA) {
			cpaStr = strconv.FormatFloat(m.CPA, 'f', 6, 64)
		}

		line := fmt.Sprintf("%s,%d,%d,%.6f,%d,%.6f,%s\n",
			m.CampaignID,
			m.TotalImpressions,
			m.TotalClicks,
			m.TotalSpend,
			m.TotalConversions,
			m.CTR,
			cpaStr,
		)

		_, err := file.WriteString(line)
		if err != nil {
			return fmt.Errorf("failed to write line: %w", err)
		}
	}

	return nil
}
