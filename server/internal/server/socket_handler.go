package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Scrcpy paths
const (
	Sockets = "/start_sockets/{" + SerialKey + "}"
)

func (s *serverImpl) handleSocketFunctions() {
	http.HandleFunc(Sockets, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.logURL(r)
			s.handleStartSockets(w, r)
			return
		}
		http.Error(w, "use GET method", http.StatusMethodNotAllowed)
	})
}

func (s *serverImpl) handleStartSockets(w http.ResponseWriter, r *http.Request) {
	var serial = r.PathValue(SerialKey)
	if serial == "" {
		http.Error(w, `"serial" param required`, http.StatusBadRequest)
		return
	}

	var started = s.interactor.StartScrcpyServer(serial, s.serverProps.SocketPort)
	if !started {
		var errMsg = fmt.Sprintf("couldn't start scrcpy server for %s", serial)
		http.Error(w, errMsg, http.StatusInternalServerError)
		s.interactor.CloseConnection(serial)
		return
	}

	var data = s.interactor.GetPortsJSON(serial)

	s.setHeaders(w)
	json.NewEncoder(w).Encode(data)

	go s.acceptSocketConnections(serial)
}

func (s *serverImpl) acceptSocketConnections(serial string) {
	doneCtx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()
		s.interactor.AcceptVideoConnections(doneCtx, serial)
	}()

	go func() {
		defer cancel()
		s.interactor.AcceptCvConnection(doneCtx, serial)
	}()

	go func() {
		defer cancel()
		s.interactor.AcceptControlConnection(doneCtx, serial)
	}()

	<-doneCtx.Done()

	s.interactor.CloseConnection(serial)
}
