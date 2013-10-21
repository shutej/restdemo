#!/bin/sh
curl -L 'http://golang.org/src/pkg/crypto/tls/generate_cert.go?m=text' \
  | grep -v '// +build ignore' > generate_cert.go
go build
./generate_cert --host=localhost
