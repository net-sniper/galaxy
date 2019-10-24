package options

import (
	"flag"

	"github.com/spf13/pflag"
)

// ServerRunOptions contains the options while running a server
type ServerRunOptions struct {
	Profiling      bool
	Bind           string
	Port           int
	APIPort        int
	Master         string
	KubeConf       string
	Swagger        bool
	LeaderElection LeaderElectionConfiguration
}

var (
	JsonConfigPath string
)

func init() {
	flag.StringVar(&JsonConfigPath, "config", "/etc/galaxy/galaxy-ipam.json", "The json config file location of"+
		" galaxy-ipam")
}

func NewServerRunOptions() *ServerRunOptions {
	opt := &ServerRunOptions{
		Profiling:      true,
		Bind:           "0.0.0.0",
		Port:           9040,
		APIPort:        9041,
		Swagger:        false,
		LeaderElection: DefaultLeaderElectionConfiguration(),
	}
	opt.LeaderElection.LeaderElect = true
	return opt
}

// AddFlags add flags for a specific ASServer to the specified FlagSet
func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&s.Profiling, "profiling", s.Profiling, "Enable profiling via web interface host:port/debug/pprof/")
	fs.StringVar(&s.Bind, "bind", s.Bind, "The IP address on which to listen")
	fs.IntVar(&s.Port, "port", s.Port, "The port on which to serve")
	fs.IntVar(&s.APIPort, "api-port", s.APIPort, "The API port on which to serve")
	fs.StringVar(&s.Master, "master", s.Master, "The address and port of the Kubernetes API server")
	fs.StringVar(&s.KubeConf, "kubeconfig", s.KubeConf, "The kube config file location of APISwitch, used to support TLS")
	fs.BoolVar(&s.Swagger, "swagger", s.Swagger, "Enable swagger via API web interface host:api-port/apidocs.json/")
	BindFlags(&s.LeaderElection, fs)
}
