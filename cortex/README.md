# Client
Creating a Cortex client program for fetching and parsing the response. 

## Description
- The purpose of this program is to make API calls to Cortex and grab the response.
- The response received from Cortex is parsed in a Struct and printed on the console.
- Just standard Go packages are used while writing the client program.
- All the inputs to the program are taken from command line arguments.
- Only basic authentication for Cortex using username and password is supported.

## Command Line Args
---
- **URL:** Valid cortex URL
- **User:** Cortex username for basic authentication
- **Password:** Cortex password for basic authentication
- **Query:** Prometheus query to run. Default value set to **up**

## Usage
---
```
go run client/main.go --url=<CORTEX_URL> --user=<CORTEX_USERNAME> --password=<CORTEX_PASSWORD> --query=<CORTEX_QUERY>
```

## Output
---
```
2020/08/31 20:09:59 Response receieved from Cortex successfully
{ResultType:vector Result:[{Metric:{} Value:[1.598884795671e+09 1]} {Metric:{} Value:[1.598884795671e+09 1]}]}
```
