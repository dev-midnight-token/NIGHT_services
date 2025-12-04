# Purpose

The code performs simple data format manipulation of live data from Cardano explorers. 

# Service

The codebase provides services listed in the table below.

|Function|Description|Location|
|---|---|---|
|Supply|The total quantity of NIGHT|./tokens.go|
|CirculatingSupply|The total circulating supply of NIGHT|./utxos.go|


# Execution
This codebase uses Google Functions.

Command for local testing of the `Supply` function.
```bash
FUNCTION_TARGET=Supply LOCAL_ONLY=true go run cmd/main.go
```

Command to deploy the `Supply` function to Google cloud. 
``` bash
gcloud functions deploy Supply --region=us-east5 --runtime=go125 --source=. --entry-point=Supply --trigger-http --project=night-allocation
```

Command for local testing of `CirculatingSupply` function
```bash
FUNCTION_TARGET=CirculatingSupply LOCAL_ONLY=true go run cmd/main.go
```

Command to deploy the `CirculatingSupply` function to Google cloud. 
``` bash
gcloud functions deploy CirculatingSupply --region=us-east5 --runtime=go125 --source=. --entry-point=CirculatingSupply --trigger-http --project=night-allocation
```
