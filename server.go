package newtest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router

	packUsers []User
}

type User struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
}

func NewServer() *Server {
	srv := &Server{
		mux.NewRouter(),
		[]User{},
	}

	srv.handles()

	return srv
}

func (s *Server) handles() {
	s.HandleFunc("/lists", s.nowShowMeBitch()).Methods("GET")
	s.HandleFunc("/lists", s.createItem()).Methods("POST")
	s.HandleFunc("/somedel/{id}", s.delSome()).Methods("DELETE")

	// http.HandleFunc("/lists", )
}

func (s *Server) createItem() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), 300)
		}

		user.Id = uuid.New()
		s.packUsers = append(s.packUsers, user)
		w.Header().Set("Content-Type", "aplication/json")
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), 300)
		}
	}
}

func (s *Server) nowShowMeBitch() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "aplication/json")
		if err := json.NewEncoder(w).Encode(s.packUsers); err != nil {
			http.Error(w, err.Error(), 300)
		}
	}
}

func (s *Server) delSome() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, _ := uuid.Parse(idStr)
		
		for i, item := range s.packUsers {
			if id == item.Id {
				s.packUsers = append(s.packUsers[:i], s.packUsers[i+1:]... )
			}
		}
	}
}