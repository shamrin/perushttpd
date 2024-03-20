// Copyright (c) 2009 The Perushttpd Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"tailscale.com/client/tailscale"
)

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}

var localClient tailscale.LocalClient

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <directory>\n", os.Args[0])
	}
	directory := os.Args[1]

	fileServer := http.FileServer(http.Dir(directory))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.
		fileServer.ServeHTTP(w, r)
	})
	http.Handle("/", handler)

	s := &http.Server{
		TLSConfig: &tls.Config{
			GetCertificate: localClient.GetCertificate,
		},
	}

	// HTTP server for redirecting to HTTPS
	go func() {
		httpListenAddr := ":80"
		log.Printf("Starting HTTP server for HTTPS redirect on %s\n", httpListenAddr)
		log.Fatal(http.ListenAndServe(httpListenAddr, http.HandlerFunc(redirectToHTTPS)))
	}()

	log.Printf("Running TLS server on :443 ...")
	log.Fatal(s.ListenAndServeTLS("", ""))
}
