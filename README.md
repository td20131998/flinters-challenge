# Campaign Data Processing Pipeline

A high-performance Go application for processing large CSV files (1GB+) containing advertising campaign data. Aggregates metrics by campaign, calculates CTR (Click-Through Rate) and CPA (Cost Per Acquisition), and identifies top performers.

## Setup Instructions

### Prerequisites

- **Go 1.22.5 or higher** (supports generics)
- **macOS/Linux/Windows** (cross-platform)

### Building

```bash
# Clone or navigate to the project directory
cd flinters-challenge

# Build the executable
go build -o flinters-challenge

# Verify build
./flinters-challenge --help
```

## How to Run

### Basic Usage (Default CSV)

```bash
./flinters-challenge
```

This processes `./ad_data.csv` and outputs results to `./result/` directory.

### Custom Input File

```bash
./flinters-challenge --input=/path/to/your/file.csv
```

### Custom Output Directory

```bash
./flinters-challenge --input=data.csv --output=/custom/output/path
```

### Generate Test Data

To generate synthetic test data:

```bash
go test -run TestGenerateRandomCSV -v
```

This generates `test_data.csv` with 1,000,000 random records.

## Docker Usage

### Prerequisites

- **Docker** installed and running
- CSV file available on host machine

### Building Docker Image

```bash
# Build the Docker image
docker build -t flinters-challenge .

# Verify the image was created
docker images | grep flinters-challenge
```

### Running with Docker

#### Process local CSV file

```bash
# Replace /path/to/data with your actual directory containing CSV file
docker run -v /path/to/data:/data \
  -v /path/to/output:/app/result \
  flinters-challenge --input=/data/ad_data.csv
```

#### Example: Process test_data.csv

```bash
# Using current directory's test_data.csv
docker run -v $(pwd):/data \
  -v $(pwd)/result:/app/result \
  flinters-challenge --input=/data/test_data.csv
```

#### Default behavior (if ad_data.csv exists in mounted data volume)

```bash
docker run -v /path/to/data:/data \
  -v /path/to/output:/app/result \
  flinters-challenge
```

### Docker Output

Results will be saved to your mounted output directory:
- `top10_ctr.csv` - Top campaigns by CTR
- `top10_cpa.csv` - Top campaigns by lowest CPA

#### Example output mapping:

```bash
# Input: /Users/yourname/data/ad_data.csv
# Output: /Users/yourname/results/

docker run -v /Users/yourname/data:/data \
  -v /Users/yourname/results:/app/result \
  flinters-challenge --input=/data/ad_data.csv

# Results appear in:
# - /Users/yourname/results/top10_ctr.csv
# - /Users/yourname/results/top10_cpa.csv
```

### Docker Flags Explanation

| Flag | Purpose |
|------|---------|
| `-v /path/host:/path/container` | Mount host directory inside container |
| `$(pwd)` | Current working directory (expands to full path) |
| `--input=/data/file.csv` | CSV file path inside container (use mounted /data) |
| `top10_ctr.csv` | Generated output file in mounted /app/result |

### Docker Image Details

- **Base Image**: `alpine:latest` (minimal size)
- **Go Version**: 1.22.5
- **Binary Size**: Optimized with `-ldflags="-s -w"` for minimal size
- **Multi-stage Build**: Builder stage for compilation, runtime stage for execution

## Output

The program generates two CSV files in the output directory:

- **`top10_ctr.csv`**: Top 10 campaigns by Click-Through Rate (highest CTR)
- **`top10_cpa.csv`**: Top 10 campaigns by Cost Per Acquisition (lowest CPA)

### Output Format

Each file contains:
```
campaign_id,total_impressions,total_clicks,total_spend,total_conversions,ctr,cpa
```

The program generates two CSV files in the output directory:

- **`top10_ctr.csv`**: Top 10 campaigns by Click-Through Rate (highest CTR)
- **`top10_cpa.csv`**: Top 10 campaigns by Cost Per Acquisition (lowest CPA)

### Output Format

Each file contains:
```
campaign_id,total_impressions,total_clicks,total_spend,total_conversions,ctr,cpa
```

## Libraries Used

| Package | Purpose |
|---------|---------|
| `encoding/csv` | Zero-copy CSV parsing with ReuseRecord |
| `flag` | Command-line argument parsing |
| `fmt` | Formatted output and logging |
| `math` | Mathematical functions (NaN, IsNaN) |
| `os` | File I/O operations |
| `path/filepath` | File path manipulation |
| `runtime` | Memory statistics and performance monitoring |
| `sort` | Sorting campaign results |
| `strconv` | String to numeric conversions |
| `sync` | Synchronization primitives (WaitGroup, Mutex, Pool) |
| `sync/atomic` | Thread-safe atomic operations |
| `time` | Execution timing and duration |

## Performance Metrics

### Processing 1GB CSV File (26,843,544 records)

```
=== Starting Campaign Data Processing ===
[START] Memory: Alloc=102.22 KB | TotalAlloc=102.22 KB | Sys=6.14 MB | GC=0 | Goroutines=1
Processing file: /Users/duongnguyentung82/Project/flinters-challenge/ad_data.csv
[AFTER_VALIDATION] Memory: Alloc=105.71 KB | TotalAlloc=105.71 KB | Sys=6.14 MB | GC=0 | Goroutines=2

[MONITOR] Memory: Alloc=1.85 MB | TotalAlloc=88.23 MB | Sys=11.52 MB | GC=29 | Goroutines=12
[MONITOR] Memory: Alloc=1.52 MB | TotalAlloc=175.31 MB | Sys=11.52 MB | GC=58 | Goroutines=12
[MONITOR] Memory: Alloc=3.24 MB | TotalAlloc=265.74 MB | Sys=11.52 MB | GC=89 | Goroutines=12
[MONITOR] Memory: Alloc=819.48 KB | TotalAlloc=354.64 MB | Sys=11.52 MB | GC=120 | Goroutines=12
[MONITOR] Memory: Alloc=616.42 KB | TotalAlloc=444.01 MB | Sys=11.52 MB | GC=150 | Goroutines=12
[MONITOR] Memory: Alloc=3.18 MB | TotalAlloc=535.49 MB | Sys=11.52 MB | GC=180 | Goroutines=12
[MONITOR] Memory: Alloc=2.49 MB | TotalAlloc=628.88 MB | Sys=11.52 MB | GC=214 | Goroutines=12
[MONITOR] Memory: Alloc=965.45 KB | TotalAlloc=717.07 MB | Sys=11.52 MB | GC=245 | Goroutines=12
[MONITOR] Memory: Alloc=708.82 KB | TotalAlloc=805.93 MB | Sys=11.52 MB | GC=277 | Goroutines=12
[MONITOR] Memory: Alloc=1.67 MB | TotalAlloc=893.78 MB | Sys=11.52 MB | GC=306 | Goroutines=12
[MONITOR] Memory: Alloc=3.00 MB | TotalAlloc=981.72 MB | Sys=11.52 MB | GC=335 | Goroutines=12
[MONITOR] Memory: Alloc=1.60 MB | TotalAlloc=1.05 GB | Sys=11.52 MB | GC=366 | Goroutines=12
[MONITOR] Memory: Alloc=2.20 MB | TotalAlloc=1.13 GB | Sys=15.52 MB | GC=396 | Goroutines=12
[MONITOR] Memory: Alloc=1.29 MB | TotalAlloc=1.21 GB | Sys=15.52 MB | GC=424 | Goroutines=12
[MONITOR] Memory: Alloc=2.42 MB | TotalAlloc=1.29 GB | Sys=15.52 MB | GC=452 | Goroutines=12
[MONITOR] Memory: Alloc=2.26 MB | TotalAlloc=1.37 GB | Sys=15.52 MB | GC=479 | Goroutines=12
[MONITOR] Memory: Alloc=1.38 MB | TotalAlloc=1.45 GB | Sys=15.52 MB | GC=508 | Goroutines=12
[MONITOR] Memory: Alloc=595.87 KB | TotalAlloc=1.53 GB | Sys=15.52 MB | GC=537 | Goroutines=12

Total records read: 26843544
Aggregated 50 campaigns
[AFTER_PROCESSING] Memory: Alloc=549.27 KB | TotalAlloc=1.58 GB | Sys=15.52 MB | GC=555 | Goroutines=2
Calculated metrics for 50 campaigns
[AFTER_METRICS] Memory: Alloc=558.79 KB | TotalAlloc=1.58 GB | Sys=15.52 MB | GC=555 | Goroutines=2
Found 50 top campaigns by CTR
[AFTER_TOP_CTR] Memory: Alloc=574.66 KB | TotalAlloc=1.58 GB | Sys=15.52 MB | GC=555 | Goroutines=2
Found 50 top campaigns by lowest CPA
[AFTER_TOP_CPA] Memory: Alloc=590.52 KB | TotalAlloc=1.58 GB | Sys=15.52 MB | GC=555 | Goroutines=2
Wrote top CTR results to: result/top10_ctr.csv
Wrote top CPA results to: result/top10_cpa.csv
[AFTER_WRITE] Memory: Alloc=609.87 KB | TotalAlloc=1.58 GB | Sys=15.52 MB | GC=555 | Goroutines=2
=== Done! ===
Total execution time: 18.513439292s
[END] Memory: Alloc=610.15 KB | TotalAlloc=1.58 GB | Sys=15.52 MB | GC=555 | Goroutines=1
```

### Key Metrics

| Metric | Value |
|--------|-------|
| **Processing Time** | **18.51 seconds** |
| **Peak Memory (Alloc)** | **3.24 MB** |
| **Total Memory Allocated** | 1.58 GB (includes garbage-collected temporary objects) |
| **System Memory** | 15.52 MB |
| **GC Cycles** | 555 |
| **Records Processed** | 26,843,544 |
| **Campaigns Aggregated** | 50 |
| **Active Goroutines** | 12 (during processing) |

### Memory Breakdown

- **Alloc**: Actual RAM in use (~610 KB at end) - the real memory footprint
- **TotalAlloc**: Cumulative allocations (~1.58 GB) - includes garbage-collected objects
- **Sys**: System memory reserved from OS (~15.52 MB)

## Optimization Techniques

The pipeline achieves excellent performance through:

1. **Zero-Copy CSV Parsing**: Uses `csv.Reader` with `ReuseRecord=true`
2. **Buffer Pooling**: Recycles `[]string` buffers via `sync.Pool`
3. **Direct Field Parsing**: Avoids intermediate struct allocations
4. **Parallel Workers**: Leverages CPU cores with configurable worker pool

## Development

### Running Tests

```bash
# Run all tests
go test -v

# Run specific tests
go test -run TopK -v
go test -run Generate -v

# Run with coverage
go test -cover
```

### Building with Optimizations

```bash
# Production build (optimized)
go build -ldflags="-s -w" -o flinters-challenge

# Debug build (with symbols)
go build -gcflags="all=-N -l" -o flinters-challenge-debug
```

## Troubleshooting

### CSV File Not Found

```bash
# Check file path
ls -la ./ad_data.csv

# Use absolute path
./flinters-challenge --input=/absolute/path/to/file.csv
```

### Memory/Performance Issues

1. Check available system memory
2. Monitor with `Activity Monitor` (macOS) or `top` (Linux)
3. Verify CSV file format matches expected schema:
   ```
   campaign_id,date,impressions,clicks,spend,conversions
   ```

### Building Issues

```bash
# Update Go modules
go mod tidy

# Force rebuild
go clean
go build -o flinters-challenge
```
