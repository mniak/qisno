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
	Verbose     bool
}

type OpenfortiVPNWrapper struct {
	config OpenfortiVPNConfig
}

func NewOpenfortiVPN(config OpenfortiVPNConfig) *OpenfortiVPNWrapper {
	return &OpenfortiVPNWrapper{
		config: config,
	}
}

func (o *OpenfortiVPNWrapper) Connect() (pismo.WaitFunc, pismo.DisconnectFunc, error) {
	cmd := exec.Command("sudo",
		"openfortivpn", o.config.Host,
		"-u", o.config.Username,
		// "-p", o.config.Password,
		"--trusted-cert="+o.config.TrustedCert,
	)

	if o.config.Verbose {
		cmd.Args = append(cmd.Args, "-vv")
	}

	var err error
	cmd.Stdout = os.Stdout
	in, err := cmd.StdinPipe()
	if err != nil {
		return nil, nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}

	_, err = in.Write([]byte(o.config.Password + "\n"))
	if err != nil {
		return nil, nil, err
	}

	return func() error {
			return cmd.Wait()
		},
		func() error {
			return cmd.Process.Signal(os.Interrupt)
		},
		nil
}
