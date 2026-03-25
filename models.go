package main

type CampaignData struct {
	CampaignID       string
	TotalImpressions int64
	TotalClicks      int64
	TotalSpend       float64
	TotalConversions int64
}

type CampaignMetrics struct {
	CampaignID       string
	TotalImpressions int64
	TotalClicks      int64
	TotalSpend       float64
	TotalConversions int64
	CTR              float64
	CPA              float64
}

type CSVRecord struct {
	CampaignID  string
	Date        string
	Impressions int64
	Clicks      int64
	Spend       float64
	Conversions int64
}
