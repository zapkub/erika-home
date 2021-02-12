package main

import (
	"fmt"
	"log"
	"os"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/mdp/qrterminal/v3"
	"github.com/tarm/serial"
	"github.com/zapkub/erika-home/internal/ac"
	"github.com/zapkub/erika-home/internal/tv"
)

func main() {
	log.Println("Howdy")

	c := &serial.Config{Name: "/dev/serial0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}

	bridge := accessory.NewBridge(accessory.Info{Name: "Bridge", ID: 1})
	livingroomtv, err := tv.New("Livingroom TV", 2, s)
	if err != nil {
		panic(err)
	}

	ac.SetMCUAddr("192.168.1.51")
	livingroomac := accessory.Info{Name: "Livingroom Air Conditoner", ID: 2}
	livingroomacac := accessory.NewSwitch(livingroomac)
	livingroomacac.Switch.On.OnValueRemoteUpdate(func(b bool) {
		if b {
			ac.LivingRoomOn()
			return
		}
		ac.LivingRoomOff()
	})

	// configure the ip transport
	config := hc.Config{Pin: "17293172", StoragePath: ".erika/hc", SetupId: "ERIK", Port: "17293"}
	t, err := hc.NewIPTransport(config, bridge.Accessory, livingroomacac.Accessory, livingroomtv.Accessory)
	if err != nil {
		log.Panic(err)
	}

	xhmrui, err := t.XHMURI()
	if err != nil {
		panic(err)
	}
	fmt.Println(xhmrui)
	uri, _ := t.XHMURI()
	qrterminal.Generate(uri, qrterminal.L, os.Stdout)

	hc.OnTermination(func() {
		<-t.Stop()
		fmt.Println("Server stop")
	})
	t.Start()

}
