package usecase

import (
	"github.com/elkarto91/operary/internal/configmgr/model"
	"github.com/elkarto91/operary/internal/configmgr/repo"
)

func SetConfig(key, orgID string, value any) error {
	return repo.SetConfig(model.ConfigEntry{Key: key, OrgID: orgID, Value: value})
}

func GetConfig(key, orgID string) (any, error) {
	c, err := repo.GetConfig(key, orgID)
	return c.Value, err
}
