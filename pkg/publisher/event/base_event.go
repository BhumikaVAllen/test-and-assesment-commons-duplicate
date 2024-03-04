package event

type BaseEvent struct {
	Event      string `json:"event"`
	EntityID   string `json:"entity_id"`
	EntityName string `json:"entity_name"`
	Timestamp  int64  `json:"timestamp"`
	TenantID   string `json:"tenant_id"`
}
