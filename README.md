# Perushttpd - basic file server

Perushttpd is a basic file server with built-in [Tailscale HTTPS support](https://tailscale.com/kb/1153/enabling-https).

Usage:

```
perushttpd <directory>
```

Based on advice from [Tailscale blog post](https://tailscale.com/blog/tls-certs) and [servetls.go example](https://github.com/tailscale/tailscale/blob/main/client/tailscale/example/servetls/servetls.go).

> *perus* (colloquial Finnish): usual, normal, basic
