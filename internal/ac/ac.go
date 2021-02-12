package ac

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

var (
	livingroomMCUAddr string
	bedroomMCUAddr    string
)

func endpoint(host, path string) string {
	var hosturl, err = url.Parse(fmt.Sprintf("http://%s", host))
	if err != nil {
		panic(err)
	}
	hosturl.Path = path
	return hosturl.String()
}

func Begin(livingroomaddr string, bedroomaddr string) {
	livingroomMCUAddr = livingroomaddr
	bedroomMCUAddr = bedroomaddr
}

func LivingroomOn() error {
	_, err := http.DefaultClient.Get(endpoint(livingroomMCUAddr, "/living-room/ac/on"))
	return err
}

func LivingroomOff() error {
	_, err := http.DefaultClient.Get(endpoint(livingroomMCUAddr, "/living-room/ac/off"))
	return err
}

func BedroomOn() error {
	_, err := http.DefaultClient.Get(endpoint(bedroomMCUAddr, "/bedroom/ac/on"))
	return err
}
func BedroomOff() error {
	_, err := http.DefaultClient.Get(endpoint(bedroomMCUAddr, "/bedroom/ac/off"))
	return err
}

type AirConditioner struct {
	*accessory.Accessory
	Thermostat *service.Thermostat
}

const (
	DefaultTemperature = 24
)

func NewHomekit() (*accessory.Accessory, *accessory.Accessory) {
	livingroomac := accessory.Info{Name: "Livingroom Air Conditioner"}
	livingroomacac := &AirConditioner{
		Accessory:  accessory.New(livingroomac, accessory.TypeAirConditioner),
		Thermostat: service.NewThermostat(),
	}
	livingroomacac.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(i int) {
		switch i {
		case characteristic.TargetHeatingCoolingStateOff:
			LivingroomOff()
		case characteristic.TargetHeatingCoolingStateCool:
			LivingroomOn()
		}
	})
	livingroomacac.Thermostat.CurrentTemperature.SetValue(DefaultTemperature)
	livingroomacac.AddService(livingroomacac.Thermostat.Service)

	bedroomac := accessory.Info{Name: "Bedroom Air Conditioner"}
	bedroomacac := accessory.NewSwitch(bedroomac)
	bedroomacac.Switch.On.OnValueRemoteUpdate(func(b bool) {
		if b {
			BedroomOn()
			return
		}
		BedroomOff()
	})

	return livingroomacac.Accessory, bedroomacac.Accessory
}
