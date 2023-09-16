package websocket

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)





type WebsocketHandler interface {
	HandleConnection(*ConnSession) error
}



type ConnSession struct {
	conn *websocket.Conn

	ctx context.Context
	cancel context.CancelFunc
}

func (c *ConnSession) Close() {
	c.conn.Close()
}

func (c *ConnSession) GetClosedMsgChan() <-chan struct{} {
	return c.ctx.Done()
}

func newSession(ctx context.Context, conn *websocket.Conn) *ConnSession {
	ctx, cancel := context.WithCancel(ctx)

	go func(ctx context.Context) {
		msgType, _, err := conn.ReadMessage()
		if err != nil {
			log.Errorf("Error while reading message: %v", err)
			return
		}

		if msgType == websocket.CloseMessage {

		} else {
			log.Errorf("Unexpected message type: %d", msgType)
		}
	}(ctx)

	return &ConnSession{
		conn: conn,
		ctx: ctx,
		cancel: cancel,
	}
}



type Server struct {
	upgrader websocket.Upgrader
	h WebsocketHandler

	ctx context.Context
	cancel context.CancelFunc
	wg *sync.WaitGroup
}


func (s *Server)handleConnections(w http.ResponseWriter, r *http.Request) {
	_, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading to WebSocket:", err)
		return
	}
	
}


func New(c *resolver.ConfigMap, h WebsocketHandler) (*Server, error) {


	port, err := c.GetStringKey("PORT")
	if err != nil {
		return nil, err
	}

	origin, err := c.GetStringKey("CLIENT_ORIGIN")
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf(":%s", port)

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return (r.Header.Get("Origin") == origin)
		},
	}

	
	failedChan := make(chan error)

	wg := sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		defer wg.Done()
		if err := http.ListenAndServe(addr, nil); err != nil {
			failedChan <- err
		}
	}(ctx)

	select {
	case err := <-failedChan:
		cancel()
		return nil, err		
	default:
		return &Server{
			upgrader: upgrader,
			h: h,
			ctx: ctx,
			cancel: cancel,
			wg: &wg,
		}, nil
	}
}


func (s *Server) Close() {
	s.cancel()
	s.wg.Wait()
}
