package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var ()

func main() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	pin4 := rpio.Pin(4)
	pin3 := rpio.Pin(17)
	pin3.Output()
	pin4.Input()

	// var close = irclone.CloneFromPIN(pin4, func(d []time.Duration) {
	// 	templateoutput, err := template.New("dump").Parse(outputtmpl)
	// 	if err != nil {
	// 		fmt.Printf("error: %+v", err)
	// 		os.Exit(2)
	// 	}
	// 	templateoutput.Execute(os.Stdout, templateData{
	// 		Timming: d,
	// 	})
	// })
	// defer close()
	defer func() {
		fmt.Println("close gpio")
		rpio.Close()
	}()
	go func() {

	}()

	pin3.High()
	// for {
	// 	time.Sleep(time.Second * 3)
	// 	for _, t := range ac.OnBurst {
	// 		time.Sleep(t)
	// 		pin3.Toggle()
	// 	}
	// 	pin3.High()
	// }

	// var lastblink rpio.State
	// var elapsed time.Time
	// for {
	// 	<-time.After(time.Microsecond)
	// 	var r = pin4.Read()
	// 	if r != lastblink {
	// 		fmt.Println("blink", r, time.Since(elapsed))
	// 		elapsed = time.Now()
	// 	}
	// 	lastblink = r
	// }

	// c := &serial.Config{Name: "/dev/serial0", Baud: 9600}
	// s, err := serial.OpenPort(c)
	// if err != nil {
	// 	panic(err)
	// }

	// sender := ir.NewSender(s)
	// err = sender.Burst(ac.OnPulses, ac.OnTimming)
	// if err != nil {
	// 	panic(err)
	// }

	// reader := bufio.NewReader(s)
	// var elapsed time.Time
	// var sum time.Duration
	// var prevburst byte

	// var pulses []byte
	// var durations []time.Duration
	// go func() {
	// 	for {
	// 		var b = make([]byte, 1)
	// 		_, err := reader.Read(b)
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		if prevburst == 0x00 && b[0] != 0x00 {
	// 			elapsed = time.Now()
	// 			sum = 0
	// 			fmt.Println("new burst")
	// 		}
	// 		prevburst = b[0]

	// 		if elapsed.IsZero() {
	// 			elapsed = time.Now()
	// 		}

	// 		sum = sum + time.Since(elapsed)
	// 		fmt.Printf("%s: %b %x (%s)\n", time.Since(elapsed), b, b, sum)
	// 		pulses = append(pulses, b[0])
	// 		durations = append(durations, time.Since(elapsed))
	// 		elapsed = time.Now()
	// 	}
	// }()
	// time.Sleep(time.Second * 3)
	// fmt.Println(pulses)

}

type templateData struct {
	Timming []time.Duration
}

func (d templateData) ToTimming() string {

	var timmingstr []string
	for _, dd := range d.Timming {
		timmingstr = append(timmingstr, fmt.Sprintf("%d", dd.Nanoseconds()))
	}

	return fmt.Sprintf("%s", strings.Join(timmingstr, ", "))

}

const outputtmpl = `

var timming = []time.Duration{ {{.ToTimming}} }

`
