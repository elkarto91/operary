package usecase

import "github.com/elkarto91/operary/internal/authz/repo"

func AssignRole(userID, role, orgID string) error {
	return repo.UpsertRole(userID, role, orgID)
}

func GetRole(userID string) (string, string, error) {
	r, err := repo.GetRole(userID)
	return r.Role, r.OrgID, err
}
