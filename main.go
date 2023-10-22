package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

const banner = `
Back-End Service
version %s | OS %s %s %s CPU %v`

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf(banner+"\n", "1.0", runtime.GOOS, runtime.GOARCH, runtime.Version(), runtime.NumCPU())

	//appExitChan := make(chan bool)
	interruptChan := make(chan os.Signal, 1)

	//cmd.Init(appExitChan)

	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	for range interruptChan {
		log.Println("shutting down the system")

		//appExitChan <- true
		//<-appExitChan

		os.Exit(1)
	}
}
