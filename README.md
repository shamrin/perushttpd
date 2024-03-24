# Perushttpd - basic file server

Perushttpd is a basic file server with built-in [Tailscale HTTPS support](https://tailscale.com/kb/1153/enabling-https).

Usage:

```
perushttpd <directory>
```

## Rationale

> *perus* (colloquial Finnish): usual, normal, basic

[Tailscale](https://tailscale.com/) is perfect to safely share internal web things that are not supposed to be available on the public internet.

Behind Tailscale, any HTTP serving software is fine, because no one from the *wild* internet can get to it for nefarous purposes.

There is only one problem - newer Web APIs [do not work](https://developer.mozilla.org/en-US/docs/Web/Security/Secure_Contexts/features_restricted_to_secure_contexts) outside of "secure context" (HTTPS). Tailscale has built-in support for Let's Encrypt HTTPS certificates, but they have to be renewed every 90 days. Renewing them manually every 3 months is quite error-prone. It's better for software to do it. Before making Perushttpd I've considered the alternatives:

1. Use [Caddy](https://caddyserver.com/) web server that claims to support Tailscale certificates.
2. Run [`tailscale cert`](https://tailscale.com/kb/1080/cli#cert) and configure common web server like Nginx.

Number 1 (Caddy) never worked for me after several attempts. It was also hard to debug.

Number 2 is more configuration and more complexity compared to [less than 50 lines of Perushttpd](https://github.com/shamrin/perushttpd/blob/master/main.go). These lines were easy to write, based on advice from [Tailscale blog post](https://tailscale.com/blog/tls-certs) and [servetls.go example](https://github.com/tailscale/tailscale/blob/main/client/tailscale/example/servetls/servetls.go).

## Can I use Perushttpd on public-facing sites?

It may work just fine, because the heavy lifting is done by battle-tested Go standard `net/http` package. That said, Perushttpd hasn't been tested on public-facing web sites yet. The original purpose is to to run internal software protected by Tailscale, where potential web server bugs are not an issue. Behind Tailscale, strangers can not access your server.
