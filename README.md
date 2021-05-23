# LogParser
The following GO Application can be used to Parse Logs, it will single out IP Addresses and URLs from the Logs and report on those accordingly.

## Running GO Application and External Packages required
* Ensure that you have golang installed and you should be able to navigate to GoParser folder and run the following command
```cmd
go run .
```
* github.com/pkg/errors for handling errors will need to be imported, run the following command
```cmd
go get github.com/pkg/errors
```
* github.com/pkg/errors for handling App Settings will need to be imported, run the following command
```cmd
 go get github.com/vrischmann/envconfig
```

## Quick Start Guide
Use the below to get instructions on what flags can be passed to the console application, after you compile the application you will be able to use the following instructions to run the LogParser.exe application

The following json file appsettings.json file can be updated to alter the execution of the GO Application
```json
{ 
    "OutputFile": {
        "OutputFilePath":".\\OutputFiles\\Output.txt",
        "AppendToOutputFile":true
    },
    "Output":{
        "NumberToDisplay":3,
        "DisplayUniqueIPs":true,
        "DisplayUniqueURLs":false,
        "SortByKeyAdditionally":true
    },
    "InputFilePath": ".\\TestFiles\\ProvidedExample.log",
    "RegularExpression": "(^(?:[0-9]{1,3}\\.){3}[0-9]{1,3})(?:.+)(?:\"{1}(?:GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE){1}\\s{1}(.+)\\s{1}(?:HTTP/1.1|HTTP/1.0|HTTP/2|HTTP/3){1}\"{1})(?:.+$)"
}
```

* OutputFilePath - if this is left empty an output file will not be created, otherwise an existing directory with the expected file name .txt can be provided
* AppendToOutputFile - true, will ensure that the file is not overwritten but appended to, and false will overwrite the file
* NumberToDisplay - this can be increased to increase the number of IP Hits and URL Occurrences displayed to console
* DisplayUniqueIPs - Include the number of unique IP Addresses in the output
* DisplayUniqueURLs - Include the number of unique URLs in the output
* SortByKeyAdditionally - Will sort the Array of KeyValue pairs by not only hit/occurrences but also IPs/URLs when set to true
* InputFilePath - This is the filepath to the .log file you are looking to process
* RegularExpression - This is used to match line for line the IP Address and URLs, it can not be empty, and it must contain two groups the first identifying the IP Address, and the second identifying the URL