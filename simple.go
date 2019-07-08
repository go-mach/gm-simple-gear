package simple

import (
	"github.com/go-mach/gm/gm"
	"log"
	"time"
)

// SimpleGear .
type SimpleGear struct {
	gm.BaseGear
}

// Name .
func (sg *SimpleGear) Name() string {
	return "simple-gear"
}

// Start .
func (sg *SimpleGear) Start(m *gm.Machinery) {
	for {
		log.Printf("configuration: %v", sg.Config)
		time.Sleep(2 * time.Second)
	}
}

// Configure .
func (sg *SimpleGear) Configure(config interface{}) {
	sg.Config = config.(map[string]interface{})
	log.Printf("%s configured", sg.Name())
	log.Printf("Gear Name: %s\n", sg.Config["name"])
}
