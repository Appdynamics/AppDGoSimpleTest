
# Simple Golang Example with the AppDynamics Golang SDK

## Caveats

Linux is assumed for the purposes of this example.

## Install and Setup

### AppDynamics GoLang SDK

Download and extract the AppDynamics Golang SDK. We will use `/opt/appdynamics/src` as the location for the example

### Edit test app with your connection detais.

Edit this section GoAppDSimpleTest.go:

```go
    // Controller
    cfg.Controller.Host = "192.168.1.150"
    cfg.Controller.Port = 8090
    cfg.Controller.UseSSL = false
    cfg.Controller.Account = "customer1"
    cfg.Controller.AccessKey = "4a2b69dd-31138179c"
```

## Run

1. set GOPATH, i.e. export GOPATH=/opt/appdynamics:$GOPATH
2. set LD_LIBRARY_PATH, i.e. LD_LIBRARY_PATH=/opt/appdynamics/src/appdynamics/lib:$LD_LIBRARY_PATH
3. go run GoAppDSimpleTest.go


## Using the Dockerfile

### Prep

1. Copy golang-sdk-x64-linux-<version>.tar.bz2 to this directory
2. Edit GoAppDSimpleTest.go with appropraite controller connection properties for your configuration

### Build

run

```
docker build -t appd/simple-golang-example:0.1 .
```

### Run

run

```
docker run --rm -v `pwd`/logs:/tmp/appd --name gotest1 appd/simple-golang-example:0.1
```

The logs will be written into the `logs` directory so you can monitor and troubleshoot from the host


