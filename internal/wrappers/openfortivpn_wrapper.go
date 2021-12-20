package wrappers

import (
	"os"
	"os/exec"

	"github.com/mniak/pismo/pkg/pismo"
)

type OpenfortiVPNConfig struct {
	Host        string
	Username    string
	Password    string
	TrustedCert string
}

type OpenfortiVPNWrapper struct {
	config OpenfortiVPNConfig
}

func NewOpenfortiVPN(config OpenfortiVPNConfig) *OpenfortiVPNWrapper {
	return &OpenfortiVPNWrapper{
		config: config,
	}
}

// var (
// 	vpnProcesses     = []*os.Process{}
// 	addVPNProcess    = make(chan *os.Process)
// 	stopVPNProcesses = make(chan interface{})
// )

// func vpnworker() {
// 	for {
// 		select {
// 		case <-stopVPNProcesses:
// 			for _, proc := range vpnProcesses {
// 				if err := proc.Signal(os.Interrupt); err != nil {
// 					log.Printf("Error interrupting process %d: %s\n", proc.Pid, err)
// 					err = proc.Kill()
// 					log.Printf("Process %d killed. %s\n", proc.Pid, err)
// 				} else {
// 					_, err := proc.Wait()
// 					if err != nil {
// 						log.Printf("Could not wait process %d: %s\n", proc.Pid, err)
// 					}
// 				}
// 			}
// 			log.Println("All VPN processes werer stopped")
// 			vpnProcesses = make([]*os.Process, 0)
// 		case p := <-addVPNProcess:
// 			log.Printf("Adding VPN process to list: %d\n", p.Pid)
// 			vpnProcesses = append(vpnProcesses, p)
// 		}
// 	}
// }

// func init() {
// 	go vpnworker()
// }

func (o *OpenfortiVPNWrapper) Connect() (pismo.WaitFunc, pismo.DisconnectFunc, error) {
	cmd := exec.Command("sudo",
		"openfortivpn", o.config.Host,
		"-u", o.config.Username,
		"-p", o.config.Password,
		"--trusted-cert="+o.config.TrustedCert,
		"-vv",
	)

	var err error
	cmd.Stdout = os.Stdout
	// in, err := cmd.StdinPipe()
	// if err != nil {
	// 	return nil, nil, err
	// }

	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}

	// _, err = in.Write([]byte(o.config.Password + "\n"))
	// if err != nil {
	// 	return nil, nil, err
	// }

	return func() error {
			return cmd.Wait()
		},
		func() error {
			return cmd.Process.Signal(os.Interrupt)
		},
		nil
}
