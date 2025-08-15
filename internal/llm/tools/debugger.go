package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/crush/internal/message"
	"github.com/charmbracelet/crush/internal/session"
)

const (
	DebuggerToolName = "debugger"
)

type DebuggerParams struct {
	Error   string `json:"error"`
	Context string `json:"context,omitempty"`
	File    string `json:"file,omitempty"`
}

type debuggerTool struct {
	sessions session.Service
	messages message.Service
}

func (b *debuggerTool) Name() string {
	return DebuggerToolName
}

func (b *debuggerTool) Info() ToolInfo {
	return ToolInfo{
		Name:        DebuggerToolName,
		Description: "Launch a specialized debugging agent to analyze and fix code errors. Use this tool when you encounter errors in code execution or need to debug specific issues.",
		Parameters: map[string]any{
			"error": map[string]any{
				"type":        "string",
				"description": "The error message or issue to debug",
			},
			"context": map[string]any{
				"type":        "string",
				"description": "Additional context about the error, such as what operation was being performed",
			},
			"file": map[string]any{
				"type":        "string",
				"description": "The file where the error occurred, if known",
			},
		},
		Required: []string{"error"},
	}
}

func (b *debuggerTool) Run(ctx context.Context, call ToolCall) (ToolResponse, error) {
	var params DebuggerParams
	if err := json.Unmarshal([]byte(call.Input), &params); err != nil {
		return NewTextErrorResponse(fmt.Sprintf("error parsing parameters: %s", err)), nil
	}

	sessionID, messageID := GetContextValues(ctx)
	if sessionID == "" || messageID == "" {
		return ToolResponse{}, fmt.Errorf("session_id and message_id are required")
	}

	// Create a prompt for the debugger agent
	prompt := fmt.Sprintf("Debug the following error:\n\nError: %s\n\nContext: %s\n\nFile: %s\n\nPlease analyze the error, identify the root cause, and provide a solution to fix it.", 
		params.Error, params.Context, params.File)

	// For now, we'll return a placeholder response
	// In a full implementation, this would create a sub-agent session and run it
	return NewTextResponse(fmt.Sprintf("DEBUGGING: %s\n\n[In a full implementation, this would launch a specialized debugger agent to analyze and fix the issue.]", prompt)), nil
}

func NewDebuggerTool(
	sessions session.Service,
	messages message.Service,
) BaseTool {
	return &debuggerTool{
		sessions: sessions,
		messages: messages,
	}
}