package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

var recordPool = sync.Pool{
	New: func() interface{} {
		return make([]string, 6)
	},
}

func parseAndProcessRecord(fields []string, m map[string]*CampaignData) error {
	if len(fields) != 6 {
		return fmt.Errorf("invalid CSV format: expected 6 fields, got %d", len(fields))
	}

	impressions, err := strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid impressions: %w", err)
	}

	clicks, err := strconv.ParseInt(fields[3], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid clicks: %w", err)
	}

	spend, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return fmt.Errorf("invalid spend: %w", err)
	}

	conversions, err := strconv.ParseInt(fields[5], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid conversions: %w", err)
	}

	campaignID := fields[0]

	if data, exists := m[campaignID]; exists {
		data.TotalImpressions += impressions
		data.TotalClicks += clicks
		data.TotalSpend += spend
		data.TotalConversions += conversions
	} else {
		m[campaignID] = &CampaignData{
			CampaignID:       campaignID,
			TotalImpressions: impressions,
			TotalClicks:      clicks,
			TotalSpend:       spend,
			TotalConversions: conversions,
		}
	}

	return nil
}

func processWorker(linesChan <-chan []string, results chan<- map[string]*CampaignData, wg *sync.WaitGroup) {
	defer wg.Done()

	localMap := map[string]*CampaignData{}
	for record := range linesChan {
		if len(record) == 0 {
			continue
		}

		if err := parseAndProcessRecord(record, localMap); err != nil {
			fmt.Printf("Error parsing record: %v\n", err)
			continue
		}

		recordPool.Put(record)
	}

	results <- localMap
}

func processCampaignData(inputFile string) (map[string]*CampaignData, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.ReuseRecord = true

	numWorkers := runtime.NumCPU()

	linesChan := make(chan []string, 10000)
	resultsChan := make(chan map[string]*CampaignData, numWorkers)

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processWorker(linesChan, resultsChan, &wg)
	}

	var counter int64 = 0

	// Read file and send records
	go func() {
		defer close(linesChan)

		// Skip header
		_, err := reader.Read()
		if err != nil {
			fmt.Printf("Error reading header: %v\n", err)
			return
		}

		for {
			line, err := reader.Read()
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Printf("Error reading record: %v\n", err)
				continue
			}

			// Get from pool or create new
			recordCopy := recordPool.Get().([]string)
			copy(recordCopy, line)
			linesChan <- recordCopy
			atomic.AddInt64(&counter, 1)
		}
	}()

	// Collect results and merge
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Merge all results
	finalMap := map[string]*CampaignData{}
	for workerMap := range resultsChan {
		for campaignID, data := range workerMap {
			if existing, exists := finalMap[campaignID]; exists {
				existing.TotalImpressions += data.TotalImpressions
				existing.TotalClicks += data.TotalClicks
				existing.TotalSpend += data.TotalSpend
				existing.TotalConversions += data.TotalConversions
			} else {
				finalMap[campaignID] = data
			}
		}
	}

	fmt.Printf("Total records read: %d\n", atomic.LoadInt64(&counter))

	return finalMap, nil
}
