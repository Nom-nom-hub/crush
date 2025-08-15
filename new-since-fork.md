# New Since Fork

This document tracks new features and enhancements implemented since forking the project.

## Provider/Model Management

Added new CLI commands for viewing available providers and models:

### `providers`
- Lists all configured providers and their available models with context window information
- Usage: `crush providers`

## Configuration Viewing

Added new CLI commands for viewing the current configuration:

### `config view`
- Displays the current configuration settings in a formatted JSON structure
- Usage: `crush config view`

## Context Management

Added new CLI commands for viewing context paths:

### `context list`
- Lists the current context paths that are automatically included in prompts
- Usage: `crush context list`

## Tool Management

Added new CLI commands for viewing available tools:

### `tools list`
- Lists the tools that are available for use by Crush
- Usage: `crush tools list`

## Sessions Management

Added new CLI commands for managing chat sessions:

### `sessions list`
- Lists all saved chat sessions with ID, title, creation time, and last used time
- Usage: `crush sessions list`

### `sessions continue`
- Provides instructions on how to continue a conversation from a saved session
- Usage: `crush sessions continue [session-id]`
- Note: Currently directs users to use the TUI session switcher (Ctrl+S) to load the specified session