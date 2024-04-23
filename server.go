package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/robertgontarski/hasher/proto"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server interface {
	Listen() error
}

type HttpServer struct {
	Ctx   context.Context
	Addr  string
	Store Store
	mux   *http.ServeMux
}

func NewHttpServer(ctx context.Context, addr string, store Store) *HttpServer {
	return &HttpServer{
		Ctx:   ctx,
		Addr:  addr,
		Store: store,
		mux:   http.NewServeMux(),
	}
}

func (s *HttpServer) Listen() error {
	s.mux.HandleFunc("POST /v1/email", s.handleEmail)
	s.mux.HandleFunc("POST /v1/phone", s.handlePhone)
	s.mux.HandleFunc("POST /v1/name", s.handleName)

	srv := &http.Server{
		Addr:    s.Addr,
		Handler: s.mux,
	}

	timeout, err := strconv.Atoi(os.Getenv("HTTP_READ_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("error while convert read timeout: %v", err)
	}

	ctx, cancel := context.WithTimeout(s.Ctx, time.Second*time.Duration(timeout))
	defer cancel()

	errch := make(chan error, 1)
	defer close(errch)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errch <- fmt.Errorf("error while start listen: %v", err)
			return
		}
	}()

	select {
	case err := <-errch:
		return err
	case <-ctx.Done():
		if err := srv.Shutdown(s.Ctx); err != nil {
			return fmt.Errorf("error while shutdown server: %v", err)
		}

		return nil
	}
}

func (s *HttpServer) handleEmail(w http.ResponseWriter, r *http.Request) {
	var msg EmailChangeMessage

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		s.prepareJSONResponse(w, NewResponseEndpoint(
			http.StatusBadRequest,
			fmt.Sprintf("error while decode request body: %v", err),
			nil,
		))

		return
	}

	h := NewEmailHasher()

	hash, err := h.ProcessStringIntoHash(msg.Address, nil)
	if err != nil {
		s.prepareJSONResponse(w, NewResponseEndpoint(
			http.StatusInternalServerError,
			fmt.Sprintf("error while hashing address: %v", err),
			nil,
		))

		return
	}

	if msg.ID > 0 {
		if err := s.Store.UpdateHashOnEmailByID(msg.ID, hash); err != nil {
			s.prepareJSONResponse(w, NewResponseEndpoint(
				http.StatusInternalServerError,
				fmt.Sprintf("error while updating db: %v", err),
				nil,
			))
			return
		}
	}

	s.prepareJSONResponse(w, NewResponseEndpoint(
		http.StatusOK,
		"success - respond hash",
		ResponseMap{
			"hash": hash,
		},
	))
}

func (s *HttpServer) handlePhone(w http.ResponseWriter, r *http.Request) {
	var msg PhoneChangeMessage

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		s.prepareJSONResponse(w, NewResponseEndpoint(
			http.StatusBadRequest,
			fmt.Sprintf("error while decode request body: %v", err),
			nil,
		))

		return
	}

	h := NewPhoneHasher()

	hash, err := h.ProcessStringIntoHash(msg.Number, ExtraDataHasher{
		"country_code": msg.CountryCode,
	})

	if err != nil {
		s.prepareJSONResponse(w, NewResponseEndpoint(
			http.StatusInternalServerError,
			fmt.Sprintf("error while hashing phone number: %v", err),
			nil,
		))
		return
	}

	if msg.ID > 0 {
		if err := s.Store.UpdateHashOnEmailByID(msg.ID, hash); err != nil {
			s.prepareJSONResponse(w, NewResponseEndpoint(
				http.StatusInternalServerError,
				fmt.Sprintf("error while updating db: %v", err),
				nil,
			))
			return
		}
	}

	s.prepareJSONResponse(w, NewResponseEndpoint(
		http.StatusOK,
		"success - return hash",
		ResponseMap{
			"hash": hash,
		},
	))
}

func (s *HttpServer) handleName(w http.ResponseWriter, r *http.Request) {
	var msg NameChangeMessage

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		s.prepareJSONResponse(w, NewResponseEndpoint(
			http.StatusBadRequest,
			fmt.Sprintf("error while decode request body: %v", err),
			nil,
		))

		return
	}

	h := NewNameHasher()

	hashName, err := h.ProcessStringIntoHash(msg.Name, nil)
	if err != nil {
		s.prepareJSONResponse(w, NewResponseEndpoint(
			http.StatusInternalServerError,
			fmt.Sprintf("error while hashing phone number: %v", err),
			nil,
		))
		return
	}

	hashSurname, err := h.ProcessStringIntoHash(msg.Surname, nil)
	if err != nil {
		s.prepareJSONResponse(w, NewResponseEndpoint(
			http.StatusInternalServerError,
			fmt.Sprintf("error while hashing phone number: %v", err),
			nil,
		))
		return
	}

	if msg.ID > 0 {
		if err := s.Store.UpdateHashOnNameByID(msg.ID, hashName, hashSurname); err != nil {
			s.prepareJSONResponse(w, NewResponseEndpoint(
				http.StatusInternalServerError,
				fmt.Sprintf("error while updating db: %v", err),
				nil,
			))
			return
		}
	}

	s.prepareJSONResponse(w, NewResponseEndpoint(
		http.StatusOK,
		"success - return hashes",
		ResponseMap{
			"hash_name":    hashName,
			"hash_surname": hashSurname,
		},
	))
}

func (s *HttpServer) prepareJSONResponse(w http.ResponseWriter, resp *ResponseEndpoint) {
	if resp.Status >= 400 {
		slog.Error(resp.Message)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error(err.Error())
	}
}

type GrpcServer struct {
	Ctx   context.Context
	Addr  string
	Store Store
}

func NewGrpcServer(ctx context.Context, addr string, store Store) *GrpcServer {
	return &GrpcServer{
		Ctx:   ctx,
		Addr:  addr,
		Store: store,
	}
}

func (s *GrpcServer) Listen() error {
	timeout, err := strconv.Atoi(os.Getenv("GRPC_READ_TIMEOUT"))
	if err != nil {
		return fmt.Errorf("error while convert read timeout: %v", err)
	}

	ctx, cancel := context.WithTimeout(s.Ctx, time.Second*time.Duration(timeout))
	defer cancel()

	errch := make(chan error, 1)
	defer close(errch)

	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("error while listen: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterHasherServiceServer(srv, NewHasherServiceServer(s.Store))

	go func() {
		if err := srv.Serve(lis); err != nil {
			errch <- fmt.Errorf("error while serve: %v", err)
		}
	}()

	select {
	case err := <-errch:
		return err
	case <-ctx.Done():
		srv.Stop()
		return nil
	}
}
