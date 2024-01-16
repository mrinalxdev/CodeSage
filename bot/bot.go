package bot

import (
	"github.com/mrinalxdev/CodeSage/bot/analyzer"
	"github.com/mrinalxdev/CodeSage/models"
)

type CodeReviewBot struct {
	analyzers []analyzer.Analyzer
}

func NewCodeReviewBot() *CodeReviewBot {
	return &CodeReviewBot{
		analyzers: []analyzer.Analyzer{
			&analyzer.TODOAnalyzer{},
			&analyzer.VariableNameAnalyzer{},
		},
	}
}

func (bot *CodeReviewBot) ReviewCode(code string) []models.Comment {
	var comments []models.Comment
	for _, a := range bot.analyzers {
		comments = append(comments, a.Analyze(code)...)
	}

	return comments
}