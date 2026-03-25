package main

import (
	"fmt"
	"math"
)

func calculateMetrics(campaignData map[string]*CampaignData) []CampaignMetrics {
	var metrics []CampaignMetrics

	for _, data := range campaignData {
		m := CampaignMetrics{
			CampaignID:       data.CampaignID,
			TotalImpressions: data.TotalImpressions,
			TotalClicks:      data.TotalClicks,
			TotalSpend:       data.TotalSpend,
			TotalConversions: data.TotalConversions,
		}

		// Calculate CTR
		if data.TotalImpressions > 0 {
			m.CTR = float64(data.TotalClicks) / float64(data.TotalImpressions)
		}

		// Calculate CPA
		if data.TotalConversions > 0 {
			m.CPA = data.TotalSpend / float64(data.TotalConversions)
		} else {
			m.CPA = math.NaN()
		}

		metrics = append(metrics, m)
	}

	return metrics
}

func getTopKCTR(metrics []CampaignMetrics, k int) []CampaignMetrics {
	// TopK with highest CTR: less function returns true if a < b (min heap)
	// So we get the k largest CTR values
	topK, err := NewTopK(k, func(a, b CampaignMetrics) bool {
		return a.CTR < b.CTR
	})
	if err != nil {
		fmt.Printf("Error creating TopK for CTR: %v\n", err)
		return []CampaignMetrics{}
	}

	for _, m := range metrics {
		topK.Add(m)
	}

	return topK.Get()
}

func getTopKCPA(metrics []CampaignMetrics, k int) []CampaignMetrics {
	// TopK with lowest CPA: less function returns true if a > b (max heap for lowest values)
	// So we get the k smallest CPA values (excluding NaN)
	topK, err := NewTopK(k, func(a, b CampaignMetrics) bool {
		aVal := a.CPA
		bVal := b.CPA
		// Skip NaN values - treat them as larger than any valid value
		if math.IsNaN(aVal) {
			return false
		}
		if math.IsNaN(bVal) {
			return true
		}
		return aVal > bVal
	})
	if err != nil {
		fmt.Printf("Error creating TopK for CPA: %v\n", err)
		return []CampaignMetrics{}
	}

	for _, m := range metrics {
		// Only add if CPA is valid (not NaN)
		if !math.IsNaN(m.CPA) {
			topK.Add(m)
		}
	}

	return topK.Get()
}
