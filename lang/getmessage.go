package lang

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed messages.json
var content embed.FS

// GetMessage
//
// will return messageType if it cannot find the specified provided argument.
func GetMessage(messageType, language string) (string, error) {

	file, err := content.Open("messages.json")
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var data map[string]map[string]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return "", fmt.Errorf("error decoding JSON: %w", err)
	}

	languageMessages, exists := data[messageType]
	if !exists {
		return messageType, nil
	}

	message, exists := languageMessages[language]
	if !exists {
		return "", fmt.Errorf("unsupported language: %s", language)
	}

	return message, nil
}
