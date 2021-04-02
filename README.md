# toxy

`toxy` is a tiny TCP proxy that can be useful to run as a sidecar to expose services that stubbornly only listens on
`localhost` and where you don't control / change the source.
On the flip-side it can also be used for tricking services to think they are talking to a service on `localhost` but in
reality it's running elsewhere.

## how to use

```shell
$ toxy --help
usage of toxy:
  -f, --forward strings   entry to forward on the format <src-ip>:<src-port>-<dst-ip>:<dst-port>
  -h, --help              you're looking at it
  -v, --version           show version

$ toxy \
  --forward 192.168.0.10:9300-127.0.0.1:9300 \
  --forward 192.168.0.10:9400-127.0.0.1:9400 \
  --forward 192.168.0.10:9500-127.0.0.1:9500 \
  --forward 127.0.0.1:8080-192.168.0.11:8080
```
