package xen

import (
	"fmt"
	"github.com/mitchellh/multistep"
	//"github.com/mitchellh/packer/packer"
	"log"
	//"math/rand"
	//"net"
)

// This step adds a NAT port forwarding definition so that SSH is available
// on the guest machine.
//
// Uses:
//
// Produces:
type stepForwardSSH struct{}

var lastIp = 1

func (s *stepForwardSSH) Run(state multistep.StateBag) multistep.StepAction {
	//config := state.Get("config").(*config)
	//ui := state.Get("ui").(packer.Ui)

	//log.Printf("Looking for available SSH port between %d and %d", config.SSHHostPortMin, config.SSHHostPortMax)
	ip := fmt.Sprintf("10.0.0.%d", lastIp)
	state.Put("privateIp", ip)
	log.Printf("lastIp is %d. setting net IP to %s", lastIp, ip)
	lastIp ++

	return multistep.ActionContinue
}

func (s *stepForwardSSH) Cleanup(state multistep.StateBag) {}
