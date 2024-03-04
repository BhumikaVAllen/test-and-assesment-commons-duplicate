package event

import (
	"encoding/json"
	"github.com/Allen-Career-Institute/test-and-assessment-commons/pkg/constants"
)

type TestAndAssessmentEvent struct {
	BaseEvent
	Data map[string]interface{} `json:"data"`
}

func (m TestAndAssessmentEvent) GetEntityID() string {
	return m.EntityID
}

func (m TestAndAssessmentEvent) GetStringMessage() (string, error) {
	stringData, err := json.Marshal(m.Data)
	if err != nil {
		return "", err
	}
	if m.TenantID == "" {
		m.TenantID = constants.DefaultTenantID
	}
	message := struct {
		BaseEvent
		Data string `json:"data"`
	}{
		BaseEvent: m.BaseEvent,
		Data:      string(stringData),
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
