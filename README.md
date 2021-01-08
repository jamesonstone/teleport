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
                          /app/icebreaker --> http://localhost:4000
client --> |teleport| -->|
                          /app/arctic     --> http://localhost:7000
```



## Resources and Links

- [minimal reverse proxy](https://gist.github.com/thurt/2ae1be5fd12a3501e7f049d96dc68bb9)
