package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listenner, err := net.Listen("tcp", ":1234")
	if err != nil {
		return err
	}
	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Printf("accept error %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)

	}
	return nil

}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
