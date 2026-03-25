package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func validateAndCreateDirs(inputFile, outputDir string) (string, error) {
	// Get absolute path for input file
	inputPath, err := filepath.Abs(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for input: %w", err)
	}

	// Check if input file exists
	stat, err := os.Stat(inputPath)
	if err != nil {
		return "", fmt.Errorf("input file not found: %w", err)
	}
	if stat.IsDir() {
		return "", fmt.Errorf("input path is a directory, expected a CSV file")
	}

	// Check if file is a CSV
	if filepath.Ext(inputPath) != ".csv" {
		return "", fmt.Errorf("input file must be a CSV file (got %s)", filepath.Ext(inputPath))
	}

	// Create output directory if it doesn't exist
	outPath, err := filepath.Abs(outputDir)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for output: %w", err)
	}

	err = os.MkdirAll(outPath, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create output directory: %w", err)
	}

	return inputPath, nil
}

func main() {
	startTime := time.Now()
	inputFile := flag.String("input", "./ad_data.csv", "Input CSV file path")
	outputDir := flag.String("output", "./result", "Output directory for results")
	flag.Parse()

	// Print initial stats
	fmt.Println("=== Starting Campaign Data Processing ===")
	printMemStats("START")

	// Start continuous monitoring
	stopMonitor := startMonitoring(1 * time.Second)
	defer func() {
		stopMonitor <- true
		time.Sleep(100 * time.Millisecond)
		printMemStats("END")
	}()

	// Validate and get CSV file path
	csvFile, err := validateAndCreateDirs(*inputFile, *outputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Validation error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Processing file: %s\n", csvFile)
	printMemStats("AFTER_VALIDATION")

	// Process campaign data
	campaignData, err := processCampaignData(csvFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Processing error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Aggregated %d campaigns\n", len(campaignData))
	printMemStats("AFTER_PROCESSING")

	// Calculate metrics
	metrics := calculateMetrics(campaignData)
	fmt.Printf("Calculated metrics for %d campaigns\n", len(metrics))
	printMemStats("AFTER_METRICS")

	// Get Top 10 by CTR
	topCTR := getTopKCTR(metrics, 10)
	fmt.Printf("Found %d top campaigns by CTR\n", len(topCTR))
	printMemStats("AFTER_TOP_CTR")

	// Get Top 10 by lowest CPA
	topCPA := getTopKCPA(metrics, 10)
	fmt.Printf("Found %d top campaigns by lowest CPA\n", len(topCPA))
	printMemStats("AFTER_TOP_CPA")

	// Write results
	ctrFile := filepath.Join(*outputDir, "top10_ctr.csv")
	err = writeCSV(ctrFile, topCTR)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing CTR file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Wrote top CTR results to: %s\n", ctrFile)

	cpaFile := filepath.Join(*outputDir, "top10_cpa.csv")
	err = writeCSV(cpaFile, topCPA)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing CPA file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Wrote top CPA results to: %s\n", cpaFile)
	printMemStats("AFTER_WRITE")

	fmt.Println("=== Done! ===")
	fmt.Printf("Total execution time: %s\n", time.Since(startTime))
}
