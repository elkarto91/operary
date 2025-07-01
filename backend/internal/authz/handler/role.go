package handler

import (
	"encoding/json"
	"net/http"

	"github.com/elkarto91/operary/internal/authz/usecase"
	"github.com/gorilla/mux"
)

type RoleRequest struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	OrgID  string `json:"org_id,omitempty"`
}

func AssignRoleHandler(w http.ResponseWriter, r *http.Request) {
	var req RoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := usecase.AssignRole(req.UserID, req.Role, req.OrgID); err != nil {
		http.Error(w, "Failed to assign role", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetRoleHandler(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)["user"]
	role, orgID, err := usecase.GetRole(user)
	if err != nil {
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}
	resp := RoleRequest{UserID: user, Role: role, OrgID: orgID}
	json.NewEncoder(w).Encode(resp)
}
