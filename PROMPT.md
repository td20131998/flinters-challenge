## Prompt1
```bash
Goal: Implement generic data struct TopK
Input: 
    - type T generic is item on heap data
    - number k is top k
    - one func define priority(max top k or min top k)
Constraints: passing test cases -> write test file
1. k = 0 -> return error
2. k > total items: case k = 5 and data = [1, 2, 3] and max top k -> [3, 2, 1]
3. Simple increasing: case k = 3 and data = [1, 2, 3, 4, 5] and max top k -> [1, 2, 3]
4. Duplicate values: case k = 3 and data = [5, 5, 5, 1, 4] and max top k -> [5, 5, 5]
5. Random order: case k = 3 and data = [3, 1, 5, 2, 4] and max top k -> [5, 4, 3]
6. All eqals: case k = 2 and data = [2, 2, 2, 2, 2, 2] and max top k -> [2, 2]
7. Negative num: case 3 =3 and data = [-1, -5, -3, -2] and max top k -> [-1, -2, -3]
8. Custome struct: T is struct type
9. Large input: case k = 2, data from 0 -> 1000000 and max top k -> [1000000, 999999]
```



## Prompt2
```bash
Goal: Streaming process csv file and aggregate data
Context: large csv file ~1GB, good performance, low memory
Input: 
    - Type: csv file - passing param --input as input path file
    - Format: 
        campaign_id: string
        date: string (YYYY-MM-DD)
        impressions: int
        clicks: int
        spend: float
        conversions: int
Output: 
    - Type: 2 csv file - passing param --output as output directory
    - Format:
        campaign_id: string
        total_impressions: int
        total_clicks: int
        total_spend: float
        total_conversions: int
        CTR: float
        CPA: float
Steps:
1. Validate params: --input(default is current dir), --output(default is ./result)
2. Validate input: check existed file.
3. Create output dir if not existed.
4. Open file input: make sure close when done.
5. Read input file and aggregate data: processing data in parallel
    - Prepare: 1 channel lines arrays string with buffer 10000, 1 channel results mapping campaign_id to data. 
    - n workers is num of CPU core: start n workers in n gorotines with input is lines, results. In worker processing data from lines channel: parsing data campaign_id, impressions, clicks, spend, conversions -> sum impressions, clicks, spend, conversions group by campaign_id -> collect to local map -> push local map to results channel when done. Should separate to 2 func: parsing data func and processing flow func.
    - 1 gorotine for reading file line by line: after read 1 line then push records to channel lines. Any error not EOF then print out, continue read next line.
    - Waiting for n workers done, then close all prepare channels.
6. Merge data: merge campaign in results to final by sum and groups by campaign again -> output is final is group of campaign already sum
7. TopK: base on final result calculate CTR and CPA for each campaign
    - CTR = total_clicks/total_impressions
    - CPA = total_spend/total_conversions, if conversions = 0, ignore or return null for CPA
   Using TopK datastructure to get Top 10 campaigns with the highest CTR(topCTR) and Top 10 campaigns with the lowest CPA(topCPA).
   Should separate to func with input is final result and output is topCTR, top CPA above.
8. Write results: write func common writeCSV with input is filename, top and output is file csv result write to disk.
    - CSV format: 
        campaign_id: string
        total_impressions: int
        total_clicks: int
        total_spend: float
        total_conversions: int
        CTR: float
        CPA: float
    - Result for topCTR: top10_ctr.csv
    - Result for topCPA: top10_cpa.csv
```

### There are also a few other prompts, but their main purpose is to fix bugs and optimize further, which I won't list here. But basically the 2 prompts above are enough to build the code