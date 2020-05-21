# go-usda-nass
Go wrapper for the USDA's NASS service.

More information on the USDA's NASS API can be found in the official [API documentation](https://quickstats.nass.usda.gov/api)

## Installation
Install the library with `go get`:

```bash
go get github.com/jpfyoder/go-usda-nass
```

Add the import statement to your project:

```go
import (
        "github.com/jpfyoder/go-usda-nass"
)
```

## Usage & Examples
### Basic Usage

```go
client := nass.NewClient("API-KEY-GOES-HERE")
fmt.Println(client.ParamValues("sector_desc"))
// [ANIMALS & PRODUCTS CROPS DEMOGRAPHICS ECONOMICS ENVIRONMENTAL]
```

### Query the Database

```go
query := nass.NewQuery(client)
query = query.Filter("commodity_desc", "CATTLE").Filter("statisticcat_desc", "INVENTORY")
query = query.Filter("year", "2010").Filter("agg_level_desc", "NATIONAL")
fmt.Println(query.Count())
// 109
fmt.Println(query.Execute())
// [map[CV (%): Value:2,744,600 agg_level_desc:NATIONAL ...
```

### Applying Operators

```go
query2 := nass.NewQuery(client)
query2 = query2.Filter("commodity_desc", "CORN").Filter("statisticcat_desc", "AREA PLANTED")
query2 = query2.FilterRange("year", "1925", "LT")
fmt.Println(query.Count())
// 86
fmt.Println(query.Execute())
// [map[CV (%): Value:118,200 agg_level_desc:AGRICULTURAL DISTRICT ...
```
