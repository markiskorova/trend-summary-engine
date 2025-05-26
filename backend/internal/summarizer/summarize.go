package summarizer

import (
	"fmt"
)

func Summarize(text string) (string, error) {
	// Call OpenAI API here. Placeholder response:
	return fmt.Sprintf("Summary of: %s", text[:100]), nil
}
