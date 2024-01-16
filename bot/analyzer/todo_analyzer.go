package analyzer

import (
	"regexp"
	"strings"

	"github.com/mrinalxdev/CodeSage/models"
)

type TODOAnalyzer struct{}

func (a *TODOAnalyzer) Analyze(code string) []models.Comment{
	var comments []models.Comment
	

	lines := strings.Split(code, "\n")
	for i, line := range lines {
		if isTODOComment(line){
			comment := extractTODOComment(line)
			comments = append(comments, models.Comment{
				Message : comment,
				Line : i + 1,
			})
		}
	}

	return comments
}

func isTODOComment(line string) bool {
	return strings.Contains(strings.ToUpper(line), "TODO")
}

func extractTODOComment(line string) string {
	re := regexp.MustCompile(`(?i)\bTODO\b\s*:\s*(.*)$`)

	matches := re.FindStringSubmatch(line)
	if len(matches) == 2 {
		return strings.TrimSpace(matches[1])
	}

	return "TODO"
}