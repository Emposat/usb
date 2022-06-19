package usb

import (
	"errors"
	"fmt"
	"os"

	"github.com/Emposat/usb/gusb"
)

type backingUsbfs struct{}

func (b backingUsbfs) getDevNum(d Device) (int, error) {
	// get_connectinfo

	return 0, ErrNotImplemented
}
func (b backingUsbfs) getVendorName(d Device) (string, error) {
	return "", ErrNotImplemented
}
func (b backingUsbfs) getProductName(d Device) (string, error) {
	return "", ErrNotImplemented
}
func (b backingUsbfs) getPort(d Device) (int, error) {
	// hub_portinfo
	// https://elixir.bootlin.com/linux/v3.2/source/drivers/usb/core/hub.c#L1372
	return 0, ErrNotImplemented

}
func (b backingUsbfs) getActiveConfig(d Device) (int, error) {
	// https://github.com/libusb/libusb/blob/93dcb8ed205a4e4cea105c2141fbbbdeac84bb66/libusb/os/linux_usbfs.c#L924
	return 0, ErrNotImplemented

}

func (b backingUsbfs) getSpeed(d Device) (Speed, error) {
	var fh *os.File
	if d.f != nil {
		fh = d.f
	} else if d.Bus <= 0 || d.Device <= 0 {
		return SpeedUnknown, errors.New("unable to determine device speed without being Open, or knowing bus and device numbers")
	} else {
		//grab a file handle ourselves, read only
		f, err := os.OpenFile(fmt.Sprintf("/dev/bus/usb/%03d/%03d", d.Bus, d.Device), os.O_RDONLY, 0644)
		if err != nil {
			return SpeedUnknown, err
		}
		defer f.Close()
		fh = f
	}
	speed, err := gusb.GetSpeed(fh)
	if err != nil {
		return SpeedUnknown, err
	}
	return Speed(speed), nil
}

func (b backingUsbfs) getDriver(d Device, intf int) (string, error) {
	return gusb.GetDriver(d.f, int32(intf))
}

func (b backingUsbfs) setConfiguration(d Device, cfg int) error {
	return ErrNotImplemented
}

func (b backingUsbfs) claim(i Interface) error   { return gusb.Claim(i.d.f, int32(i.ID)) }   // ioctl
func (b backingUsbfs) release(i Interface) error { return gusb.Release(i.d.f, int32(i.ID)) } // ioctl

/* Not universal funcs */
