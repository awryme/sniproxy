package connproxy

import (
	"io"
	"net"
	"time"

	"github.com/awryme/sniproxy/pkg/goticker"
	"github.com/oklog/run"
)

func twoWayCopy(remoteConn, localConn net.Conn, logOngoing goticker.Action) error {
	const logUpdateAfter = time.Second * 10

	group := new(run.Group)

	closeConns := func(err error) {
		localConn.Close()
		remoteConn.Close()
	}

	stopTicker := goticker.Run(logUpdateAfter, logOngoing)
	defer stopTicker()

	group.Add(makeCopier(localConn, remoteConn), closeConns)
	group.Add(makeCopier(remoteConn, localConn), closeConns)

	return group.Run()
}

func makeCopier(s, c net.Conn) func() error {
	return func() error {
		_, err := io.Copy(s, c)
		return err
	}
}
