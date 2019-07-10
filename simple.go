package simple

import (
	"log"
	"time"

	"github.com/go-mach/gm/gm"
)

// SimpleGear .
type SimpleGear struct {
	gm.ConfigurableGear
	canRun bool
}

// UUIDGear .
type UUIDGear struct {
	uuid int
}

//////////// SimpleGear

// NewSimpleGear .
func NewSimpleGear(uname string, config map[string]interface{}) *SimpleGear {
	return &SimpleGear{ConfigurableGear: gm.NewConfigurableGear(uname, config), canRun: false}
}

// Name .
func (sg *SimpleGear) Name() string {
	return sg.UniqueName
}

// Start .
func (sg *SimpleGear) Start(m *gm.Machinery) {
	sg.canRun = true
	go func() {
		for sg.canRun {
			log.Printf("configuration: %v", sg.Config)
			uuidGear := (m.GetGear("uuid-gear"))
			// uGear, ok := (*uuidGear).(*UUIDGear)
			//if ok {
			log.Printf("UUID: %s", uuidGear.Provide())
			//}

			time.Sleep(2 * time.Second)
		}
	}()
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

// Shutdown .
func (sg *SimpleGear) Shutdown() {
	log.Println("simple-gear shut down")
	sg.canRun = false
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

// Shutdown .
func (ug *UUIDGear) Shutdown() {
	log.Println("uuid-gear shut down")
}
