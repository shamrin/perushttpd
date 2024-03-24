// Copyright (c) 2024 The Perushttpd Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"tailscale.com/client/tailscale"
)

var localClient tailscale.LocalClient

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <directory>\n", os.Args[0])
	}
	directory := os.Args[1]

	go func() {
		log.Printf("%s: running HTTP server that redirects to HTTPS\n", os.Args[0])
		log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
		})))
	}()

	fileServer := http.FileServer(http.Dir(directory))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0
		w.Header().Set("Expires", "0")                                         // Proxies
		fileServer.ServeHTTP(w, r)
	}))
	s := &http.Server{
		TLSConfig: &tls.Config{
			GetCertificate: localClient.GetCertificate,
		},
	}
	log.Printf("%s: running HTTPS server\n", os.Args[0])
	log.Fatal(s.ListenAndServeTLS("", ""))
}
