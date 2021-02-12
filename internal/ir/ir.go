package ir

import (
	"bufio"
	"fmt"
	"time"

	"github.com/tarm/serial"
)

func NewSender(p *serial.Port) *Sender {
	return &Sender{
		p: p,
	}
}

type Sender struct {
	p *serial.Port
}

func (s *Sender) Burst(pulse []byte, timming []time.Duration) error {
	var err error
	for i, duration := range timming {
		<-time.After(duration)
		fmt.Printf("%s: pulse %x\n", duration, pulse[i])
		_, err = s.p.Write([]byte{pulse[i]})
		if err != nil {
			return err
		}
		err = s.p.Flush()
		if err != nil {
			return err
		}

	}
	return nil
}

func New() {
	c := &serial.Config{Name: "/dev/serial0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(s)
	// reader := hex.NewDecoder(s)
	for {
		var b = make([]byte, 1)
		_, err := reader.Read(b)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%x\n", b)
	}
	// err := rpio.Open()
	// if err != nil {
	// 	panic(err)
	// }
	// rpio.Close()
	// pin := rpio.Pin(4)
	// pin.Input()
	// for {
	// 	fmt.Println(pin.Read())
	// 	time.Sleep(time.Millisecond)
	// }

	// options := serial.OpenOptions{
	// 	PortName:        "/dev/mem",
	// 	BaudRate:        19200,
	// 	DataBits:        8,
	// 	StopBits:        1,
	// 	MinimumReadSize: 4,
	// }

	// Open the port.
	// port, err := serial.Open(options)
	// if err != nil {
	// 	log.Fatalf("serial.Open: %v", err)
	// }
}
