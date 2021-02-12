package ac

import (
	"fmt"
	"net/http"
	"net/url"
)

var mcuAddr string

func endpoint(path string) string {
	var hosturl, err = url.Parse(fmt.Sprintf("http://%s", mcuAddr))
	if err != nil {
		panic(err)
	}
	hosturl.Path = path
	return hosturl.String()
}

func SetMCUAddr(mcuaddr string) {
	mcuAddr = mcuaddr
}

func LivingRoomOn() error {
	_, err := http.DefaultClient.Get(endpoint("/living-room/ac/on"))
	return err
}

func LivingRoomOff() error {
	_, err := http.DefaultClient.Get(endpoint("/living-room/ac/off"))
	return err
}
