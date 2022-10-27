package handler

import (
	"encoding/json"
	"net/http"

	. "mongorestapi3/internal/model"
	. "mongorestapi3/internal/service"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	service StudentService
}

func NewStudentHandler(service StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.service.GetAllStudents(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, students)
}

func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "ID cannot be empty", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetStudent(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, user)
}

func (h *StudentHandler) InsertStudent(w http.ResponseWriter, r *http.Request) {
	var user Student
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	res, er2 := h.service.InsertStudent(r.Context(), &user)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}

	var msg string
	if res == 1 {
		msg = "The record inserted successfully !!!"
	}
	if res == -1 {
		msg = "The record has not inserted !!!"
	}
	JSON(w, http.StatusOK, msg)
}

func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var user Student
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	if len(user.ID) == 0 {
		user.ID = id
	} else if id != user.ID {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}

	res, er2 := h.service.UpdateStudent(r.Context(), &user)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}

	var msg string
	if res == 0 {
		msg = "There's no record updated !!!"
	}

	if res == 1 {
		msg = "The record is updated !!!"
	}

	JSON(w, http.StatusOK, msg)
}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	res, err := h.service.DeleteStudent(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var msg string
	if res == 0 {
		msg = "No record has found to delete !!!"
	}
	if res == 1 {
		msg = "The record is successful deleted !!!"
	}
	JSON(w, http.StatusOK, msg)
}

func JSON(w http.ResponseWriter, code int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(res)
}
