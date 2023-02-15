package main

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

const (
	idVendor  = 0x2011
	idProduct = 0x0715
)

func main() {
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Setup Device
	mpk, err := ctx.OpenDeviceWithVIDPID(idVendor, idProduct)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}
	mpk.SetAutoDetach(true)

	fmt.Printf("Device Description:\n%s\n", mpk.Desc)

	cfg, err := mpk.Config(1)
	if err != nil {
		log.Fatalf("Could not get device config: %v", err)
	}

	fmt.Println(cfg.Desc.Interfaces)

	// Create Interface
	intf, done, err := mpk.DefaultInterface()
	if err != nil {
		log.Fatalf("%s.DefaultInterface(): %v", mpk, err)
	}
	defer done()

	fmt.Println(intf.String())

}
