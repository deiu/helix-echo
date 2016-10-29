package main

import (
	"crypto/tls"
	"flag"
	"os"
	"time"

	"github.com/deiu/helix"
	"github.com/labstack/echo/engine/standard"
	"github.com/tylerb/graceful"
)

var (
	conf  = *flag.String("conf", "", "use this configuration file")
	port  = *flag.String("port", "8443", "HTTPS listener address")
	debug = *flag.Bool("debug", false, "output extra logging?")
	root  = *flag.String("root", ".", "path to file storage root")
	cert  = *flag.String("cert", "", "TLS certificate eg. /path/to/cert.pem")
	key   = *flag.String("key", "", "TLS certificate eg. /path/to/key.pem")
)

func init() {
	flag.Parse()
}

func main() {
	// Configure server
	config := helix.NewHelixConfig()
	config.Conf = conf
	if len(config.Conf) > 0 {
		err := config.LoadJSONFile(config.Conf)
		if err != nil {
			println("Error loading config file:", err)
			os.Exit(1)
		}
	}
	// override loaded config with CLI params
	if len(port) > 0 {
		config.Port = port
	}
	if debug {
		config.Debug = debug
	}
	if len(root) > 0 {
		config.Root = root
	}
	if len(cert) > 0 {
		config.Cert = cert
	}
	if len(key) > 0 {
		config.Key = key
	}

	// Create a new server
	e := helix.NewServer(config)
	// Start server
	println("Preparing server...")
	println("Listening on port: " + config.Port)
	println("Loaded certificate from: " + config.Cert)
	println("Loaded key from: " + config.Key)
	// set config values
	std := standard.WithTLS(":"+config.Port, config.Cert, config.Key)
	// use strong crypto
	std.TLSConfig.MinVersion = tls.VersionTLS12
	std.TLSConfig.PreferServerCipherSuites = true
	std.TLSConfig.CurvePreferences = []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256}
	std.TLSConfig.CipherSuites = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	}

	println("Server is listening for connections...")
	// start server
	err := e.Run(std)
	if err != nil {
		println("Error starting server:", err.Error())
		os.Exit(1)
	}
	std.SetHandler(e)
	graceful.ListenAndServe(std.Server, 10*time.Second)
}
