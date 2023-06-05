package respond

import (
	"sync"
)

type Messages struct {
	Lang      string
	Success   string
	Failed    string
	Errors    map[string]map[string]interface{}
	Languages map[string]map[string]interface{}
	sync.RWMutex
}

func NewMessages() *Messages {
	return &Messages{
		Lang: "en",
		Languages: map[string]map[string]interface{}{
			"fa": fa.Messages,
			"en": en.Messages,
		},
	}
}

func (m *Messages) AddLanguageTranslation(lang string, messages map[string]interface{}) {
	m.Lock()
	m.Languages[lang] = messages
	m.Unlock()
}

func (m *Messages) load() {
	m.RLock()
	translation := m.Languages[m.Lang]
	m.Errors = translation["errors"].(map[string]map[string]interface{})
	m.Success = translation["success"].(string)
	m.Failed = translation["failed"].(string)
	m.RUnlock()
}
