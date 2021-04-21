# go-stein
This Go / Golang client helps you interact with the Stein API.
> Stein is a suite of programs to help you turn any Google Sheet to a database. The core Stein API provides RESTful access to your sheets.
> 
> *[SteinHQ](steinhq.com)*

## Installation
```
$ go get github.com/nasrul21/go-stein
```

## Usage
### Initalize
```go
package main

import (
    "time"
    stein "github.com/nasrul21/go-stein"
)

func main() {
    client := stein.NewClient(
        "https://api.steinhq.com/v1/storages/5cca0542e52a3545102c1665", // STEIN URL 
        // you can set the option nil, if your stein doesn't need authentication
        &stein.Option{
            Username: "user", // your STEIN API username
            Password: "password", // your STEIN API password
            Timeout: 20 * time.Second, // set timeout for Stein Request, default: 15s
        },
    )

    ...
}
```

### Read Data
`client.Read(sheetName, filter, v)`
#### Arguments
| Name          | Description                           | Format                            | Requirement   |
|---------------|---------------------------------------|-----------------------------------|---------------|
| sheetName     | Your sheet name                       | String                            | Required      |
| filter        | Filter / Query data by column name    | map[string]interface{}            | Optional (nil)|
| v             | The result will be given to `v`       | Struct / map[string]interface{}   | Required      |
|               |                                       |                                   |             |
#### Example
```go
// set your filter
filter := map[string]interface{
    "ID": 134,
}
result := &YourStruct{}
status, err := client.Read("Sheet1", filter, &result)

// if you want to get the data without filter
// status, err := client.Read("Sheet1", nil, &result)

if err != nil {
    log.Fatal("Error read from stein: " + err)
    return
}

fmt.Printf("HTTP Status: %d \n", status)
fmt.Printf("Result : %v \n", result)
```

### Insert Data / Add Rows
`client.Insert(sheetName, body) (status, res, err)`

#### Arguments
| Name | Description | Format | Requirement |
|---|---|---|---|
| sheetName | Your sheet's name | String | Required |
| body | The data you want to add | Array of Struct | Required |

#### Return Value
| Name | Description | Format |
|---|---|---|
| status | HTTP Status Response | Number (Status Code) |
| res | Response of updated range | `InsertResponse` |
| err | Error message | error |

#### Example
```go
type Employee struct {
    Name        string `json:"name"`
    Department  string `json:"department"`
    Position    string `json:"position"`
}

employees := []Employee{
    {
        Name:       "Nasrul",
        Department: "Technology",
        Position:   "Software Engineer",
    },
    // add another field of you want to insert multiple rows
    // {
    // 	...
    // },
}

status, res, err := client.Insert("Sheet1", employees)
if err != nil {
    log.Fatal("Error insert data to stein: ", status, err.Error())
    return
}

fmt.Printf("HTTP Status: %d \n", status)
fmt.Printf("Updated Range : %v \n", res)
```

### Update Data
`client.Update(sheetName, set, where) (status, res, err)`

#### Arguments
| Name | Description | Format | Requirement |
|---|---|---|---|
| sheetName | Your sheet's name | String | Required |
| set | The column values to set | {column: value, ...} | Required |
| where | The column values to search for | {column: value, ...} | Optional |

#### Return Value
| Name | Description | Format |
|---|---|---|
| status | HTTP Status Response | Number (Status Code) |
| res | Response of updated range | `UpdateResponse` |
| err | Error message | error |

#### Example
```go
employee := Employee{
    Name:       "Nasrul",
    Department: "Technology",
    Position:   "Software Engineer",
}

where := map[string]interface{}{
    "ID": 156,
}

status, res, err := client.Update("Sheet1", employee, where)
if err != nil {
    log.Fatal("Error update data to stein: ", status, err.Error())
    return
}

fmt.Printf("HTTP Status: %d \n", status)
fmt.Printf("Updated Range : %v \n", res)
```

### Delete Data
`client.Delete(sheetName, condition) (status, res, err)`

#### Arguments
| Name | Description | Format | Requirement |
|---|---|---|---|
| sheetName | Your sheet's name | String | Required |
| condition | The column values to search for | {column: value, ...} | Required |

#### Return Value
| Name | Description | Format |
|---|---|---|
| status | HTTP Status Response | Number (Status Code) |
| res | Response of updated range | `DeleteResponse` |
| err | Error message | error |

#### Example
```go
status, res, err := client.Delete("Sheet1", map[string]interface{}{
    "ID": 156,
})
if err != nil {
    log.Fatal("Error delete data: ", status, err.Error())
    return
}

fmt.Printf("HTTP Status: %d \n", status)
fmt.Printf("Deleted Range : %v \n", res)
```

## License
[MIT licensed](./LICENSE)