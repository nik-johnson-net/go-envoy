# Enphase Envoy Golang Client

This is a golang Enphase Envoy client for communicating with the Envoy communications gateway. It's built and tested against an IQ Envoy. Other models have not been tested. It communicates directly with the unit via the exposed web interface, **not** the Enphase Enlighten API.

<https://enphase.com/en-us/support/what-envoy>

## Example

```go
import "github.com/nik-johnson-net/go-envoy"

client := envoy.NewClient("192.168.0.201")

// Contains data on Production and Consumption, if equipped.
productionData, err := client.Production()

// Contains information on connected devices
inventoryData, err := client.Inventory()
```

## License

This library is provided under the [MIT License](LICENSE.md)
