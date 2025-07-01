package model

type ConfigEntry struct {
	Key   string `bson:"key" json:"key"`
	OrgID string `bson:"org_id,omitempty" json:"org_id,omitempty"`
	Value any    `bson:"value" json:"value"`
}
