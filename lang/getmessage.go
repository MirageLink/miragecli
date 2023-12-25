package lang

import (
	"embed"
	"encoding/json"
)

//go:embed messages.json
var content embed.FS

// GetMessage
//
// will return messageType if it cannot find the specified provided argument.
func GetMessage(messageType, language string) string {

	file, err := content.Open("messages.json")
	if err != nil {
		return ""
	}
	defer file.Close()

	var data map[string]map[string]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return ""
	}

	languageMessages, exists := data[messageType]
	if !exists {
		return messageType
	}

	message, exists := languageMessages[language]
	if !exists {
		return ""
	}

	return message
}
