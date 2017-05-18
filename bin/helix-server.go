package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/deiu/helix-echo"
)

var (
	conf    = flag.String("conf", "", "use this configuration file")
	port    = flag.String("port", "8443", "HTTPS listener address")
	debug   = flag.Bool("debug", false, "output extra logging?")
	logfile = flag.String("log", "", "path to log file")
	root    = flag.String("root", ".", "path to file storage root")
	cert    = flag.String("cert", "", "TLS certificate eg. /path/to/cert.pem")
	key     = flag.String("key", "", "TLS certificate eg. /path/to/key.pem")
)

func init() {
	flag.Parse()
}

func main() {
	var err error
	// Configure server
	config := helix.NewHelixConfig()
	config.Conf = *conf
	if len(config.Conf) > 0 {
		err := config.LoadJSONFile(config.Conf)
		if err != nil {
			println("Error loading config file:", err)
			os.Exit(1)
		}
	}
	// override loaded config with CLI params
	if len(*port) > 0 {
		config.Port = *port
	}
	if *debug {
		config.Debug = *debug
	}
	if len(*logfile) > 0 {
		config.Logfile = *logfile
	}
	if len(*root) > 0 {
		config.Root = *root
	}
	if len(*cert) > 0 {
		config.Cert = *cert
	}
	if len(*key) > 0 {
		config.Key = *key
	}

	// Create a new server
	e := helix.NewServer(config)
	// Start server
	println("Preparing server...")
	println("Loaded certificate from: " + config.Cert)
	println("Loaded key from: " + config.Key)
	log.Printf("%+v\n", config)

	// set config values
	s := &http.Server{
		Addr:    ":" + config.Port,
		Handler: e,
	}

	tlsCfg, err := helix.NewTLSConfig(config.Cert, config.Key)
	if err != nil {
		return
	}
	s.TLSConfig = tlsCfg

	// start server
	e.StartServer(s)
}
