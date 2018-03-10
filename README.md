
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

1. set GOPATH, i.e. export GOPATH=/opt/appdynamics/src/appdynamics:$GOPATH
2. set LD_LIBRARY_PATH, i.e. LD_LIBRARY_PATH=/opt/appdynamics/src/appdynamics/lib:$LD_LIBRARY_PATH
3. go run GoAppDSimpleTest.go


## Using the Dockerfile


