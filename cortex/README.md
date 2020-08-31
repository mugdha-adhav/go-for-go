## Client
Creating a sample client program for parsing Cortex response. 

### Description
- The program is written using standard Go packages only.
- The purpose of this program is to make API calls to Cortex and get the response.
- The response received from Cortex is parsed in a Struct and simply printed on the console.
- All the inputs to the program are taken from command line arguments.
- Basic authentication for Cortex using username and password is supported.

### Usage
```
go run client/main.go --url=<CORTEX_URL> --user=<CORTEX_USERNAME> --password=<CORTEX_PASSWORD>
```