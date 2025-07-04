package prompt

import (
	"strings"

	"github.com/bitrise-io/bitrise-plugins-ai-reviewer/common"
)

func GetSummarizePrompt(settings common.Settings) string {
	return `Provide your final response with the following content:
` + getSummary(settings) + `
` + getWalkthrough(settings) + `
## Line Feedback
A list of issues found in the diff hunks. Return the file ("file"), short title of the issue ("title"), category ("category"), brief description of the issue ("issue") and the exact line content ("content") you are commenting on.
Only include lines that appear in the diff hunk. Do not make up lines.
Quote the entire target line exactly as it appears in the diff.
Don't comment on lines that you already gave suggestion on.
Add a Prompt to AI Agents as "prompt" field. This should be a short, clear instruction for an AI agent to fix the issue. Don't include the file or line number in the prompt, just the prompt itself.
If you know how to fix the issue, you can include a "suggestion" field with a code snippet that fixes the issue. The suggestion should replace the flagged line(s) content. Suggestions should be valid code, always with "\t" indentation. Avoid adding just comments.

Focus on bugs, smells, security issues, and code quality improvements.
Categorize the issues as "issue", "refactor", improvement, "documentation", "nitpick", "test coverage"
` + getHaiku(settings) + `
---
Avoid additional commentary as the response will be added as a comment on the GitHub pull request.
` + getResponseFormat(settings)
}

func getSummary(settings common.Settings) string {
	if settings.Reviews.Summary {
		return `## Summary
A high-level, to-the-point, short summary of the overall change instead of specific files within 80 words.`
	}
	return ""
}

func getWalkthrough(settings common.Settings) string {
	if settings.Reviews.Walkthrough {
		return `## Walkthrough
A markdown table of file(s) (multiple files should be a string, separated with commas) and their summaries. Group files 
with similar changes together into a single row to save space. Return the file name(s) ("files") and a brief summary of the changes ("summary") in each row.`
	}
	return ""
}

func getHaiku(settings common.Settings) string {
	if settings.Reviews.Haiku {
		return `## Haiku
Write a whimsical, short haiku to celebrate the changes as "Bit Bot".
Format the haiku as a quote using the ">" symbol and feel free to use emojis where relevant.`
	}
	return ""
}

func getResponseFormat(settings common.Settings) string {
	headers := []string{}
	if settings.Reviews.Summary {
		headers = append(headers, "**summary**")
	}
	if settings.Reviews.Walkthrough {
		headers = append(headers, "**walkthrough**")
	}
	headers = append(headers, "**line-feedback**")
	if settings.Reviews.Haiku {
		headers = append(headers, "**haiku**")
	}
	return strings.Join(headers, ", ")
}
