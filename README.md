# port-forwarder-with-proxy-protocol-support

For those who want to expose service as NodePort and can not get Client IP address, I wrote a small program to do some help.

Usage & Example

```
$ docker run --rm --network host chaunceyshannon/port-forwarder-with-proxy-protocol-support
Usage: /bin/run [Local Address] [Local Port] [Remote Address] [Remote Port]
$ docker run --rm --network host chaunceyshannon/port-forwarder-with-proxy-protocol-support 0.0.0.0 80 10.103.217.233 80
12-25 06:54:37   1 [INFO] (main.go:24) Listen on 0.0.0.0:80 and forward to 10.103.217.233:80
12-25 06:54:50   1 [TRAC] (main.go:27) New connection from: 192.168.6.1:38886
```

More information: https://shareitnote.com/page/get-client-ip-address-when-server-is-behind-the-tcp-proxy-with-proxy-protocol
