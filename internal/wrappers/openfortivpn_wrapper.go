package wrappers

import (
	"os"
	"os/exec"

	"github.com/mniak/qisno/internal/config"
	"github.com/mniak/qisno/pkg/qisno"
)

type OpenfortiVPNConfig struct {
	Host        string
	Username    string
	Password    string
	TrustedCert string
}

type OpenfortiVPNConfigLoader interface {
	Load() (OpenfortiVPNConfig, error)
}
type passwordManagerVPNConfigLoader struct {
	pwdmgr qisno.PasswordManager
}

func VPNConfigFromPasswordManager(pwdmgr qisno.PasswordManager) OpenfortiVPNConfigLoader {
	return passwordManagerVPNConfigLoader{
		pwdmgr: pwdmgr,
	}
}

func (l passwordManagerVPNConfigLoader) Load() (OpenfortiVPNConfig, error) {
	result := OpenfortiVPNConfig{}
	host, err := l.pwdmgr.Attribute("Pismo VPN", "URL")
	if err != nil {
		return result, err
	}
	username, err := l.pwdmgr.Username("Pismo VPN")
	if err != nil {
		return result, err
	}
	password, err := l.pwdmgr.Password("Provedores/Okta")
	if err != nil {
		return result, err
	}
	trustedCert, err := l.pwdmgr.Attribute("Pismo VPN", "TrustedCert")
	if err != nil {
		return result, err
	}

	return OpenfortiVPNConfig{
		Host:        host,
		Username:    username,
		Password:    password,
		TrustedCert: trustedCert,
	}, nil
}

type inlineVPNConfigLoader struct {
	cfg config.VPNConfig
}

func VPNConfigInline(cfg config.VPNConfig) OpenfortiVPNConfigLoader {
	return inlineVPNConfigLoader{
		cfg: cfg,
	}
}

func (l inlineVPNConfigLoader) Load() (OpenfortiVPNConfig, error) {
	return OpenfortiVPNConfig{
		Host:        l.cfg.Host,
		Username:    l.cfg.Username,
		Password:    l.cfg.Password,
		TrustedCert: l.cfg.TrustedCert,
	}, nil
}

type OpenfortiVPNWrapper struct {
	configLoader OpenfortiVPNConfigLoader
	verbose      bool
}

func NewOpenfortiVPN(configLoader OpenfortiVPNConfigLoader, verbose bool) *OpenfortiVPNWrapper {
	return &OpenfortiVPNWrapper{
		configLoader: configLoader,
		verbose:      verbose,
	}
}

func (o *OpenfortiVPNWrapper) Connect() (qisno.WaitFunc, qisno.DisconnectFunc, error) {
	cfg, err := o.configLoader.Load()
	if err != nil {
		return nil, nil, err
	}
	cmd, err := o.buildCommand(cfg)
	if err != nil {
		return nil, nil, err
	}

	fnWait := func() error {
		err := cmd.Wait()
		return err
	}
	fnKill := func() error {
		err := cmd.Process.Kill()
		return err
	}

	return fnWait, fnKill, nil
}

func (o *OpenfortiVPNWrapper) buildCommand(cfg OpenfortiVPNConfig) (*exec.Cmd, error) {
	cmd := exec.Command("sudo",
		"openfortivpn", cfg.Host,
		"-u", cfg.Username,
		"--trusted-cert="+cfg.TrustedCert,
	)

	if o.verbose {
		cmd.Args = append(cmd.Args, "-vv")
	}

	cmd.Stdout = os.Stdout
	in, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	_, err = in.Write([]byte(cfg.Password + "\n"))
	if err != nil {
		return nil, err
	}

	return cmd, err
}
