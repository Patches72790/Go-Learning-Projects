package main

import (
	//"github.com/patches72790/concurrency/channels"
	clockserver "github.com/patches72790/concurrency/clock_server"
)

func main() {
	//channels.UnbufferedPipeline()

	clockserver.ClockServer()

}
