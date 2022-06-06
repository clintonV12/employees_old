package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	employees "employees/proto"
)

type Employees struct{}

// Return a new handler
func New() *Employees {
	return &Employees{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Employees) Call(ctx context.Context, req *employees.Request, rsp *employees.Response) error {
	log.Info("Received Employees.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Employees) Stream(ctx context.Context, req *employees.StreamingRequest, stream employees.Employees_StreamStream) error {
	log.Infof("Received Employees.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&employees.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Employees) PingPong(ctx context.Context, stream employees.Employees_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&employees.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
