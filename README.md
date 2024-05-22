# gpsroutegen

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/rrgmc/gpsroutegen)

`gpsroutegen` is a Golang package that generates realistic GPS coordinates.

![Mapbox](./images/mapbox.png)

## Installation

```bash
go get github.com/rrgmc/gpsroutegen
```

### Examples

```go
    data := gpsroutegen.Generate(gpsroutegen.NewPoint(55.953251, -3.188267),
        gpsroutegen.GenerateInput{
            Amount: 10,
            BearingRange: gpsroutegen.DirectionRange{
                Min: gpsroutegen.Direction_SSE,
                Max: gpsroutegen.Direction_SSW,
            },
            DistanceRange: gpsroutegen.DataRange{
                Min: 10 * 1000,
                Max: 15 * 1000,
            },
        },
        gpsroutegen.GenerateInput{
            Amount: 20,
            BearingRange: gpsroutegen.DirectionRange{
                Min: gpsroutegen.Direction_W,
                Max: gpsroutegen.Direction_WNW,
            },
            DistanceRange: gpsroutegen.DataRange{
                Min: 1000,
                Max: 1500,
            },
        },
    )

    url, err := data.ToGeoJSONIOUrl()
    if err != nil {
        panic(err)
    }
    fmt.Println("click the URL to see the route in a map")
    fmt.Println(url)
```

```go
    data := gpsroutegen.GenerateRandom(
        gpsroutegen.WithDistance(gpsroutegen.RandRangeInt(3000, 5000)),
        gpsroutegen.WithDirectionChanges(8),
        gpsroutegen.WithStart(gpsroutegen.RandPointNear(gpsroutegen.NewPoint(55.953251, -3.188267), 300.0)))

    url, err := data.ToGeoJSONIOUrl()
    if err != nil {
        panic(err)
    }
    fmt.Println("click the URL to see the route in a map")
    fmt.Println(url)
```

Output data can be tested using the [Mapbox viewer](https://geojson.io/#map=2/0/20). 

## Author

Based on [https://dev.to/daunderworks/create-gps-test-data-in-go-4of6](https://dev.to/daunderworks/create-gps-test-data-in-go-4of6).

Rangel Reale (rangelreale@gmail.com)
