package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/crush/internal/message"
	"github.com/charmbracelet/crush/internal/session"
)

const (
	ArchitectToolName = "architect"
)

type ArchitectParams struct {
	Requirement string `json:"requirement"`
	Context     string `json:"context,omitempty"`
	Existing    string `json:"existing,omitempty"`
}

type architectTool struct {
	sessions session.Service
	messages message.Service
}

func (b *architectTool) Name() string {
	return ArchitectToolName
}

func (b *architectTool) Info() ToolInfo {
	return ToolInfo{
		Name:        ArchitectToolName,
		Description: "Launch a specialized architecture agent to design software solutions and systems. Use this tool when you need to plan complex features, design system architectures, or create technical specifications.",
		Parameters: map[string]any{
			"requirement": map[string]any{
				"type":        "string",
				"description": "The architectural requirement or feature to design",
			},
			"context": map[string]any{
				"type":        "string",
				"description": "Additional context about the project or system",
			},
			"existing": map[string]any{
				"type":        "string",
				"description": "Description of existing systems or components that need to be integrated",
			},
		},
		Required: []string{"requirement"},
	}
}

func (b *architectTool) Run(ctx context.Context, call ToolCall) (ToolResponse, error) {
	var params ArchitectParams
	if err := json.Unmarshal([]byte(call.Input), &params); err != nil {
		return NewTextErrorResponse(fmt.Sprintf("error parsing parameters: %s", err)), nil
	}

	sessionID, messageID := GetContextValues(ctx)
	if sessionID == "" || messageID == "" {
		return ToolResponse{}, fmt.Errorf("session_id and message_id are required")
	}

	// Create a prompt for the architect agent
	prompt := fmt.Sprintf("Design an architecture for the following requirement:\n\nRequirement: %s\n\nContext: %s\n\nExisting: %s\n\nPlease create a detailed technical design including components, interactions, data flow, and implementation considerations.", 
		params.Requirement, params.Context, params.Existing)

	// For now, we'll return a placeholder response
	// In a full implementation, this would create a sub-agent session and run it
	return NewTextResponse(fmt.Sprintf("ARCHITECTING: %s\n\n[In a full implementation, this would launch a specialized architect agent to design the solution.]", prompt)), nil
}

func NewArchitectTool(
	sessions session.Service,
	messages message.Service,
) BaseTool {
	return &architectTool{
		sessions: sessions,
		messages: messages,
	}
}