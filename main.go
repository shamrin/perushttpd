// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// The servetls program shows how to run an HTTPS server
// using a Tailscale cert via LetsEncrypt.
package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"

	"tailscale.com/client/tailscale"
)

func main() {
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	fileServer := http.FileServer(http.Dir(*directory))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.
		fileServer.ServeHTTP(w, r)
	})

	http.Handle("/", handler)

	s := &http.Server{
		TLSConfig: &tls.Config{
			GetCertificate: tailscale.GetCertificate,
		},
	}
	log.Printf("Running TLS server on :443 ...")
	log.Fatal(s.ListenAndServeTLS("", ""))
}
