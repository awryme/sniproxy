package connproxy

// old code
// todo: remove
// func twoWayCopy(remoteConn, localConn net.Conn, logOngoing goticker.Action) error {
// 	const logUpdateAfter = time.Second * 10

// 	group := new(errgroup.Group)
// 	group.SetLimit(2)

// 	ignoreError := new(atomic.Bool)

// 	doCopy := func(s, c net.Conn) error {
// 		id := ulid.Make()
// 		typ := "docopy"
// 		log.Println("goroutine started", typ, id)
// 		defer log.Println("goroutine stopped", typ, id)

// 		defer c.Close()
// 		defer s.Close()
// 		defer ignoreError.Store(true)

// 		_, err := io.Copy(s, c)
// 		if ignoreError.Load() {
// 			return nil
// 		}
// 		return err
// 	}

// 	stopTicker := goticker.Run(logUpdateAfter, logOngoing)
// 	defer stopTicker()

// 	group.Go(func() error {
// 		return doCopy(localConn, remoteConn)
// 	})
// 	group.Go(func() error {
// 		return doCopy(remoteConn, localConn)
// 	})

// 	return group.Wait()
// }

//old code
// todo: remove

// func twoWayCopy(remoteConn, localConn io.ReadWriter, logOngoing goticker.Action) error {
// 	const logUpdateAfter = time.Second * 10

// 	returnedCopies := make(chan error, 2)

// 	doCopy := func(s, c io.ReadWriter, stop chan<- error) {
// 		id := ulid.Make()
// 		typ := "docopy"
// 		log.Println("goroutine started", typ, id)
// 		defer log.Println("goroutine stopped", typ, id)

// 		_, err := io.Copy(s, c)
// 		stop <- err
// 	}

// 	stopTicker := goticker.Run(logUpdateAfter, logOngoing)
// 	defer stopTicker()

// 	go doCopy(localConn, remoteConn, returnedCopies)
// 	go doCopy(remoteConn, localConn, returnedCopies)

// 	err := <-returnedCopies
// 	go func() {
// 		<-returnedCopies
// 		close(returnedCopies)
// 	}()
// 	return err
// }
