package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/robertgontarski/hasher/proto"
	"net/http"
	"time"
)

type HasherServiceServer struct {
	Store Store
	proto.UnimplementedHasherServiceServer
}

func NewHasherServiceServer(store Store) *HasherServiceServer {
	return &HasherServiceServer{
		Store: store,
	}
}

func (s *HasherServiceServer) HashEmail(ctx context.Context, r *proto.HashEmailRequest) (*proto.HashResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()

	respch := make(chan HasherResponseCh, 1)
	defer close(respch)

	go func() {
		h := NewEmailHasher()

		hash, err := h.ProcessStringIntoHash(r.Address, nil)
		if err != nil {
			respch <- HasherResponseCh{
				Data: &proto.HashResponse{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("error while hash data: %v", err),
					Data:    nil,
				},
				Err: fmt.Errorf("error while hash data: %v", err),
			}

			return
		}

		if r.Id > 0 {
			if err := s.Store.UpdateHashOnEmailByID(int(r.Id), hash); err != nil {
				respch <- HasherResponseCh{
					Data: &proto.HashResponse{
						Status:  http.StatusInternalServerError,
						Message: fmt.Sprintf("error while save data into db: %v", err),
						Data:    nil,
					},
					Err: fmt.Errorf("error while save data into db: %v", err),
				}
				return
			}
		}

		data, err := json.Marshal(ResponseMap{
			"hash": hash,
		})

		if err != nil {
			respch <- HasherResponseCh{
				Data: &proto.HashResponse{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("error while process hash to json: %v", err),
					Data:    nil,
				},
				Err: fmt.Errorf("error while process hash to json: %v", err),
			}
			return
		}

		respch <- HasherResponseCh{
			Data: &proto.HashResponse{
				Status:  http.StatusOK,
				Message: "success - return hash",
				Data:    data,
			},
			Err: nil,
		}
	}()

	select {
	case <-ctx.Done():
		return &proto.HashResponse{
			Status:  http.StatusRequestTimeout,
			Message: "request timeout",
			Data:    nil,
		}, fmt.Errorf("request timeout")
	case resp := <-respch:
		return resp.Data, resp.Err
	}
}

func (s *HasherServiceServer) HashPhone(ctx context.Context, r *proto.HashPhoneRequest) (*proto.HashResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()

	respch := make(chan HasherResponseCh, 1)
	defer close(respch)

	go func() {
		h := NewPhoneHasher()

		hash, err := h.ProcessStringIntoHash(r.Number, ExtraDataHasher{
			"country_code": r.CountryCode,
		})

		if err != nil {
			respch <- HasherResponseCh{
				Data: &proto.HashResponse{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("error while hash data: %v", err),
					Data:    nil,
				},
				Err: fmt.Errorf("error while hash data: %v", err),
			}
			return
		}

		if r.Id > 0 {
			if err := s.Store.UpdateHashOnPhoneByID(int(r.Id), hash); err != nil {
				respch <- HasherResponseCh{
					Data: &proto.HashResponse{
						Status:  http.StatusInternalServerError,
						Message: fmt.Sprintf("error while save data into db: %v", err),
						Data:    nil,
					},
					Err: fmt.Errorf("error while save data into db: %v", err),
				}
				return
			}
		}

		data, err := json.Marshal(ResponseMap{
			hash: hash,
		})

		if err != nil {
			respch <- HasherResponseCh{
				Data: &proto.HashResponse{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("error while process hash to json: %v", err),
					Data:    nil,
				},
				Err: fmt.Errorf("error while process hash to json: %v", err),
			}
			return
		}

		respch <- HasherResponseCh{
			Data: &proto.HashResponse{
				Status:  http.StatusOK,
				Message: "success - return hash",
				Data:    data,
			},
			Err: nil,
		}
	}()

	select {
	case <-ctx.Done():
		return &proto.HashResponse{
			Status:  http.StatusRequestTimeout,
			Message: "request timeout",
			Data:    nil,
		}, fmt.Errorf("request timeout")
	case resp := <-respch:
		return resp.Data, resp.Err
	}
}

func (s *HasherServiceServer) HashName(ctx context.Context, r *proto.HashNameRequest) (*proto.HashResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()

	respch := make(chan HasherResponseCh, 1)
	defer close(respch)

	go func() {
		h := NewNameHasher()

		hashName, err := h.ProcessStringIntoHash(r.Name, nil)
		if err != nil {
			respch <- HasherResponseCh{
				Data: &proto.HashResponse{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("error while hash name data: %v", err),
					Data:    nil,
				},
				Err: fmt.Errorf("error while hash name data: %v", err),
			}
			return
		}

		hashSurname, err := h.ProcessStringIntoHash(r.Surname, nil)
		if err != nil {
			respch <- HasherResponseCh{
				Data: &proto.HashResponse{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("error while hash surname data: %v", err),
					Data:    nil,
				},
				Err: fmt.Errorf("error while hash surname data: %v", err),
			}
			return
		}

		if r.Id > 0 {
			if err := s.Store.UpdateHashOnNameByID(int(r.Id), hashName, hashSurname); err != nil {
				respch <- HasherResponseCh{
					Data: &proto.HashResponse{
						Status:  http.StatusInternalServerError,
						Message: fmt.Sprintf("error while save data into db: %v", err),
						Data:    nil,
					},
					Err: fmt.Errorf("error while save data into db: %v", err),
				}
				return
			}
		}

		data, err := json.Marshal(ResponseMap{
			"hash_name":    hashName,
			"hash_surname": hashSurname,
		})

		if err != nil {
			respch <- HasherResponseCh{
				Data: &proto.HashResponse{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("error while process hash to json: %v", err),
					Data:    nil,
				},
				Err: fmt.Errorf("error while process hash to json: %v", err),
			}
			return
		}

		respch <- HasherResponseCh{
			Data: &proto.HashResponse{
				Status:  http.StatusOK,
				Message: "success - return hash",
				Data:    data,
			},
			Err: nil,
		}
	}()

	select {
	case <-ctx.Done():
		return &proto.HashResponse{
			Status:  http.StatusRequestTimeout,
			Message: "request timeout",
			Data:    nil,
		}, fmt.Errorf("request timeout")
	case resp := <-respch:
		return resp.Data, resp.Err
	}
}
