package tv

import (
	"fmt"
	"time"

	"github.com/brutella/hc/accessory"
	"github.com/tarm/serial"
)

var (
	MessageOnOff = []byte{0xA1, 0xF1, 0x07, 0x07, 0x02}
	PulseOnOff   = []time.Duration{
		0, 4502018, 4486185, 583176, 1697391, 594894, 1605516, 582915, 1664162, 569582, 535780, 12934702, 540519, 573853, 546613, 579686, 543801, 564530, 544946, 575623, 539374, 586196, 1660568, 589634, 534686, 580050, 537654, 582186, 542863, 574894, 539894, 587707, 535728, 596300, 517863, 569686, 1671349, 575050, 548488, 573905, 1692287, 559373, 1680412, 556613, 1662495, 614946, 1630829, 575727, 1665673, 586508, 1717964, 564477, 46850908, 4547487, 4485456, 586092, 1672704, 565415, 1687859, 556718, 1664786, 578280, 546509, 582133, 546249, 558019, 536353, 584738, 538696, 581092, 543175, 578540, 1662131, 578696, 1666350, 581197, 1686505, 561353, 536404, 585103, 540102, 577030, 542133, 575571, 553644, 570311, 542863, 609842, 515207, 583280, 1676453, 561717, 565103, 556561, 562186, 582030, 541457, 554529, 563801, 558123, 565363, 553800, 565832, 557134, 1658224, 584582, 565311, 555467, 1665724, 575467, 1665933, 592655, 1651870, 581821, 1657912, 562811, 1686766, 554478, 1685152, 617498,
	}
)

func New(name string, ID uint64, irserialport *serial.Port) (*accessory.Television, error) {
	info := accessory.Info{
		Name:         name,
		SerialNumber: "ERIKA-TV-01",
		Manufacturer: "Erika Home",
		Model:        "",
		ID:           ID,
	}
	ac := accessory.NewTelevision(info)
	ac.Television.Active.OnValueRemoteUpdate(func(i int) {
		fmt.Println("tv", i)
		irserialport.Write(MessageOnOff)
		irserialport.Flush()
	})
	ac.Television.RemoteKey.OnValueRemoteUpdate(func(i int) {
		fmt.Println("tvkey", i)

	})
	return ac, nil
}
