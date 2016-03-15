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
