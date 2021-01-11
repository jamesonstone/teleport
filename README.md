# Teleport

Dynamic Reverse Proxy

## Getting Started

```bash
make run

navigate: localhost:8080/[app name]
```

## Overview

The purpose of this repository is to achieve dynamic routing in a reverse proxy.

```text
                         /icebreaker --> http://localhost:4000
client --> |teleport| -->
                         /arctic     --> http://localhost:7000
```

Transport holds a map with two entries:

- `icebreaker` with a reference to its running location `localhost:4000`
- `arctic` with a reference to its running location `localhost:7000`

When a user calls into teleport with `/[app name]` (for example, `http://teleport/icebreaker`), a reverse proxy is generated and set at that URI for `icebreaker`. All calls to `http://teleport/icebeaker` will reverse proxy to `localhost:4000/*`

## Resources and Links

- [minimal reverse proxy](https://gist.github.com/thurt/2ae1be5fd12a3501e7f049d96dc68bb9)
- [more complete proxy](https://github.com/ymedialabs/ReverseProxy/blob/master/main.go)
- [multiple host reverse proxy](https://gist.github.com/ego008/12a0ffe02294ae5fafcac704b5172e54)
- [moxy: multiple host reverse proxy](https://github.com/microsoftarchive/moxy)
- `func (mux *ServeMux) Handle(pattern string, handler Handler) {`
