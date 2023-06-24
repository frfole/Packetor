package packetor

import (
	"context"
	"log"
	"net"
	"time"
)

// Fronter listens to client connections and pass them into procession.
type Fronter struct {
	ctx      context.Context
	dialer   net.Dialer
	backNet  string
	backAddr string
	logger   *log.Logger
	timeout  time.Duration
}

func NewFronter(ctx context.Context, network string, address string, timeout time.Duration, keepAlive time.Duration) Fronter {
	return Fronter{
		ctx: ctx,
		dialer: net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAlive,
			Resolver:  nil,
		},
		backNet:  network,
		backAddr: address,
		logger:   log.Default(),
		timeout:  timeout,
	}
}

// Bind binds Fronter to a new address.
func (f *Fronter) Bind(network string, address string, keepAlive time.Duration) error {
	listenConfig := net.ListenConfig{
		Control:   nil,
		KeepAlive: keepAlive,
	}
	listener, err := listenConfig.Listen(f.ctx, network, address)
	if err != nil {
		return err
	}
	f.logger.Println("Listening on", listener.Addr())
	go f.listenLoop(listener)
	return nil
}

func (f *Fronter) listenLoop(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			f.logger.Println("failed to accept new connection", err)
			break
		}
		go f.handleConnection(conn)
	}
}

func (f *Fronter) handleConnection(fConn net.Conn) {
	defer func(fConn net.Conn) {
		err := fConn.Close()
		if err != nil {
			f.logger.Println("failed while closing front connection", err)
		}
	}(fConn)
	f.logger.Println("New connection", fConn.RemoteAddr())

	bConn, err := f.dialer.DialContext(f.ctx, f.backNet, f.backAddr)
	if err != nil {
		f.logger.Println("failed to dial", err)
		return
	}
	defer func(bConn net.Conn) {
		err := bConn.Close()
		if err != nil {
			f.logger.Println("failed while closing back connection", err)
		}
	}(bConn)
	f.logger.Println("Connection routed", fConn.RemoteAddr(), "->", bConn.RemoteAddr())

	errCh := make(chan error, 1)
	route := Route{
		fronter:  f,
		fConn:    fConn,
		bConn:    bConn,
		errCh:    errCh,
		compress: false,
		state:    0,
	}
	route.Start()
	select {
	case err := <-errCh:
		if err != nil {
			f.logger.Println("Error", err)
		}
	case <-f.ctx.Done():
	}
	return
}
