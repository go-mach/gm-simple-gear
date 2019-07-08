package simple

import (
	"fmt"
	"github.com/go-mach/gm/gm"
	"log"
	"time"

)

// SimpleGear .
type SimpleGear struct {
}

// Name .
func (sg *SimpleGear) Name() string {
	return "simple-gear"
}

// Start .
func (sg *SimpleGear) Start(m *gm.Machinery) {
	for {
		fmt.Println("Test")
		time.Sleep(2 * time.Second)
	}
}

// Configure .
func (sg *SimpleGear) Configure(config interface{}) {
	configuration := config.(map[string]interface{})
	log.Printf("%s configured", sg.Name())
	log.Printf("Gear Name: %s\n", configuration["name"])
}
