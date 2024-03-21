# Perushttpd - basic file server

Perushttpd is a basic file server with built-in [Tailscale HTTPS support](https://tailscale.com/kb/1153/enabling-https).

Usage:

```
perushttpd <directory>
```

## Story

> *perus* (colloquial Finnish): usual, normal, basic

[Tailscale](https://tailscale.com/) is perfect to safely share internal web things that are not supposed to be available on the *wild* internet.

Any HTTP serving software is fine, because no one from *wild* internet can get to it for nefarous purposes.

There is only one problem - newer Web APIs [do not work](https://developer.mozilla.org/en-US/docs/Web/Security/Secure_Contexts/features_restricted_to_secure_contexts) outside of "secure context" (HTTPS). Tailscale has built-in support for Let's Encrypt HTTPS certificates, but they have to be renewed every 90 days. Renewing them manually every 3 months is quite error-prone. It's better for software to do it. Caddy web server claims to support Tailscale certificates, but I could never make it work.

Fortunately, it's very easy to write your own *basic* web server, based on advice from [Tailscale blog post](https://tailscale.com/blog/tls-certs) and [servetls.go example](https://github.com/tailscale/tailscale/blob/main/client/tailscale/example/servetls/servetls.go).

