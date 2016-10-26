#!/bin/sh

export HELIX_PORT=8443
export HELIX_CERT=test_cert.pem
export HELIX_KEY=test_key.pem
export HELIX_LOGGING=true

go run bin/*.go
