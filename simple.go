package simple

import (
	"log"
	"time"

	"github.com/go-mach/gm/gm"
)

// SimpleGear .
type SimpleGear struct {
	gm.BaseGear
}

// UUIDGear .
type UUIDGear struct {
	uuid int
}

//////////// SimpleGear

// Name .
func (sg *SimpleGear) Name() string {
	return "simple-gear"
}

// Start .
func (sg *SimpleGear) Start(m *gm.Machinery) {
	for {
		log.Printf("configuration: %v", sg.Config)
		uuidGear := (m.GetGear("uuid-gear"))
		// uGear, ok := (*uuidGear).(*UUIDGear)
		//if ok {
		log.Printf("UUID: %s", uuidGear.Provide())
		//}

		time.Sleep(2 * time.Second)
	}
}

// Configure .
func (sg *SimpleGear) Configure(config interface{}) {
	sg.Config = config.(map[string]interface{})
	log.Printf("%s configured", sg.Name())
	log.Printf("Gear Name: %s\n", sg.Config["name"])
}

// Provide .
func (sg *SimpleGear) Provide() interface{} {
	return nil
}

//////////// UUIDGear

// Name .
func (ug *UUIDGear) Name() string {
	return "uuid-gear"
}

// Start .
func (ug *UUIDGear) Start(m *gm.Machinery) {
	ug.uuid = 1232
}

// Provide .
func (ug *UUIDGear) Provide() interface{} {
	return ug.uuid
}
