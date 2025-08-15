package agent

import (
	"github.com/charmbracelet/crush/internal/llm/tools"
	"github.com/charmbracelet/crush/internal/message"
	"github.com/charmbracelet/crush/internal/session"
)

// initializeSpecializedTools creates and returns the specialized tools (debugger, architect, etc.)
func initializeSpecializedTools(
	sessions session.Service,
	messages message.Service,
) []tools.BaseTool {
	return []tools.BaseTool{
		tools.NewDebuggerTool(sessions, messages),
		tools.NewArchitectTool(sessions, messages),
	}
}