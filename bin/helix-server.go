package main

import (
	// "crypto/tls"
	"os"
	"time"

	"github.com/deiu/helix"
	"github.com/labstack/echo/engine/standard"
	"github.com/tylerb/graceful"
)

var (
	SRV_PORT = os.Getenv("HELIX_PORT")
	SRV_CERT = os.Getenv("HELIX_CERT")
	SRV_KEY  = os.Getenv("HELIX_KEY")
)

func main() {
	e := helix.NewServer()
	// Start server
	println("Preparing server...")
	println("Listening on port: " + SRV_PORT)
	println("Loaded certificate from: " + SRV_CERT)
	println("Loaded key from: " + SRV_KEY)
	std := standard.WithTLS(":"+SRV_PORT, SRV_CERT, SRV_KEY)
	println("Server is listening for connections...")
	e.Run(std)
	std.SetHandler(e)
	graceful.ListenAndServe(std.Server, 5*time.Second)
}
