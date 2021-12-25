package main

import (
	"fmt"
	"io"
	"os"

	. "github.com/ChaunceyShannon/golanglibs"

	proxyproto "github.com/pires/go-proxyproto"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " [Local Address] [Local Port] [Remote Address] [Remote Port]")
		Os.Exit(0)
	}

	laddr := os.Args[1]
	lport := Int(os.Args[2])
	raddr := os.Args[3]
	rport := Int(os.Args[4])

	Lg.Info("Listen on " + laddr + ":" + Str(lport) + " and forward to " + raddr + ":" + Str(rport))

	for cl := range Socket.TCP.Listen(laddr, lport).Accept() {
		Lg.Trace("New connection from:", cl.Conn.RemoteAddr().String())
		go func(cl *TcpServerSideConn) {
			Try(func() {
				// Lg.Trace("Connect to " + raddr + ":" + Str(rport))
				cr := Socket.TCP.Connect(raddr, rport)

				header := &proxyproto.Header{
					Version:           2, // Change to version 1 if needs
					Command:           proxyproto.PROXY,
					TransportProtocol: proxyproto.TCPv4,
					SourceAddr:        cl.Conn.RemoteAddr(),
					DestinationAddr:   cl.Conn.LocalAddr(),
				}
				// Lg.Trace("Write proxy protocol header")
				_, err := header.WriteTo(cr.Conn)
				Panicerr(err)

				// Lg.Trace("Forwarding data")
				go func() {
					defer cl.Close()
					defer cr.Close()
					io.Copy(cl.Conn, cr.Conn)
				}()
				go func() {
					defer cl.Close()
					defer cr.Close()
					io.Copy(cr.Conn, cl.Conn)
				}()
			}).Except(func(e error) {
				Lg.Error("Error when forwarding:", e)
			})
		}(cl)
	}
}
