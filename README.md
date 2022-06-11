# iot-device-information-go

## Build

```go build -o iot-device-information```

Build for amd64

```go tool dist list | grep amd```

```env GOOS=linux GOARCH=amd64 go build -o iot-device-information-amd64```

Build for arm64

```go tool dist list | grep arm```

```env GOOS=linux GOARCH=arm64 go build -o iot-device-information-arm64```