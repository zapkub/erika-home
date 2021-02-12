package main

import (
	"fmt"
	"log"
	"os"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	hclog "github.com/brutella/hc/log"
	"github.com/mdp/qrterminal/v3"
	"github.com/tarm/serial"
	"github.com/zapkub/erika-home/internal/ac"
	"github.com/zapkub/erika-home/internal/fsutil"
	"github.com/zapkub/erika-home/internal/tv"
)

func main() {

	logfile, err := fsutil.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logfile)
	hclog.Info.SetOutput(logfile)
	hclog.Debug.SetOutput(logfile)

	log.Println("erika about to start")
	c := &serial.Config{Name: "/dev/serial0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}

	bridge := accessory.NewBridge(accessory.Info{Name: "Bridge", ID: 1})

	livingroomtv, err := tv.New(s)
	if err != nil {
		panic(err)
	}

	ac.Begin("192.168.1.51", "192.168.1.51")
	livingroomAC, bedroomAC := ac.NewHomekit()

	livingroomtv.ID = 5
	livingroomAC.ID = 10
	bedroomAC.ID = 11

	// configure the ip transport
	config := hc.Config{Pin: "17293172", StoragePath: ".erika/hc", SetupId: "ERIK", Port: "17293"}
	t, err := hc.NewIPTransport(config, bridge.Accessory, livingroomAC, bedroomAC, livingroomtv)
	if err != nil {
		log.Panic(err)
	}

	uri, err := t.XHMURI()
	if err != nil {
		panic(err)
	}
	qrterminal.Generate(uri, qrterminal.L, logfile)

	hc.OnTermination(func() {
		<-t.Stop()
		fmt.Println("Server stop")
	})
	t.Start()

}
