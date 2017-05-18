package helix

import (
// "log"
// "os"

// "github.com/labstack/echo/middleware"
)

// NewLoggerConfig formats the log to a specific template
// func NewLoggerConfig(logfile string) middleware.LoggerConfig {
// 	cfg := middleware.LoggerConfig{
// 		Output: os.Stdout,
// 		Format: `[${time_rfc3339}] ${method} request for: ${uri} completed in (${latency_human})` + "\n" +
// 			`[${time_rfc3339}] From: ${remote_ip}` + "\n" +
// 			`[${time_rfc3339}] Status: ${status}` + "\n" +
// 			`[${time_rfc3339}] Bytes in: ${bytes_in}` + "\n" +
// 			`[${time_rfc3339}] Bytes out: ${bytes_out}` + "\n",
// 	}
// 	// use file
// 	if len(logfile) > 0 {
// 		output, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		cfg.Output = output
// 	}

// 	return cfg
// }
