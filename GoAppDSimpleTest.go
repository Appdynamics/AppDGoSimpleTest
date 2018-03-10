package main

import (
	"fmt"
	"os/exec"
	"bytes"
	"log"
	appd "appdynamics"
	"time"
)


func main() {

	fmt.Println("Arch Check")

	// Exec ldd test
	ldd_cmd := exec.Command("ldd", "--version")

	var ldd_out bytes.Buffer
	ldd_cmd.Stdout = &ldd_out

	ldd_err := ldd_cmd.Run();
	if ldd_err != nil {
		log.Fatal(ldd_err);
	}

	fmt.Printf("ldd --version:\n")
	fmt.Printf("%s\n", ldd_out.String())

	// Exec uname test
	uname_cmd := exec.Command("uname", "-a")

	var uname_out bytes.Buffer
	uname_cmd.Stdout = &uname_out

	uname_err := uname_cmd.Run();
	if uname_err != nil {
		log.Fatal(uname_err);
	}

	fmt.Printf("uname -a:\n")
	fmt.Printf("%s\n", uname_out.String())

	// Configure AppD
	cfg := appd.Config{}

	// Controller
	cfg.Controller.Host = "<my-controller-port>"
	cfg.Controller.Port = 8090
	cfg.Controller.UseSSL = false
	cfg.Controller.Account = "customer1"
	cfg.Controller.AccessKey = "<my-access-key>"

	// App Context
	cfg.AppName = "GolangTest1"
	cfg.TierName = "GolangTier1"
	cfg.NodeName = "GolangNode1"

	// misc
	cfg.InitTimeoutMs = 1000

	// init the SDK
	if err := appd.InitSDK(&cfg); err != nil {
		fmt.Printf("Error initializing the AppDynamics SDK\n")
	} else {
		fmt.Printf("Initialized AppDynamics SDK successfully\n")
	}

	// Run some BTs
	maxBtCount := 20000
	btCount := 0

	fmt.Print("Doing something")
	for btCount < maxBtCount {
		// start the "Checkout" transaction
		btHandle := appd.StartBT("MyTestGolangBT", "")

		// do something....
		fmt.Print(".")
		milliseconds := 250
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)

		// end the transaction
		appd.EndBT(btHandle)

	}
	fmt.Print("\n")

	// Stop/Clean up the AppD SDK.
	appd.TerminateSDK()

}
