package handler

import (
	opentracing "github.com/opentracing/opentracing-go"
	"io"
	ps "user-service/repository"
	tracer "user-service/tracer"
)

type postServer struct {
	store  *ps.RegisteredUserRepository
	tracer opentracing.Tracer
	closer io.Closer
}

const name = "post_service"

func NewPostServer() (*postServer, error) {
	store, err := ps.New()
	if err != nil {
		return nil, err
	}

	tracer, closer := tracer.Init(name)
	opentracing.SetGlobalTracer(tracer)
	return &postServer{
		store:  store,
		tracer: tracer,
		closer: closer,
	}, nil
}

func (s *postServer) GetTracer() opentracing.Tracer {
	return s.tracer
}

func (s *postServer) GetCloser() io.Closer {
	return s.closer
}

func (s *postServer) CloseTracer() error {
	return s.closer.Close()
}

func (s *postServer) CloseDB() error {
	return s.store.Close()
}
