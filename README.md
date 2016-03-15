haproxy
=======

AlienVault plugin for HAProxy

## HAProxy Configuration

```
global
  [...]
  log           1.2.3.4 syslog info
  log-tag       haproxy
  [...]

defaults
  log     global
  [...]
```

## Usage
This plugin assumes that you define the server names in your backend configuration to be DNS-resolvable.

### Example

```
backend my_backend
  server                foo1.bar.com 192.168.100.10:389 check fall 1 rise 1 inter 60s
  server                foo2.bar.com 192.168.100.11:389 check fall 1 rise 1 inter 60s
```
