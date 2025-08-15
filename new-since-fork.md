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

### `config set`
- Sets a configuration value by specifying the key and value
- Usage: `crush config set [key] [value]`

### `config edit`
- Opens the configuration file in the default editor
- Usage: `crush config edit`

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

## Specialized Agents

Added new specialized agents and tools:

### Debugger Agent
- A specialized agent for debugging code and fixing errors
- Accessible through the `debugger` tool

### Architect Agent
- A specialized agent for software architecture and design
- Accessible through the `architect` tool

## Session Export/Import

Added new CLI commands for exporting and importing sessions:

### `export`
- Exports a session's conversation history to a file
- Supports JSON and markdown formats
- Usage: `crush export [session-id] [--format json|markdown]`

### `import`
- Imports a session from a file
- Usage: `crush import [file]`

## Session Management

Added new CLI commands for managing chat sessions:

### `sessions list`
- Lists all saved chat sessions with ID, title, creation time, and last used time
- Usage: `crush sessions list`

### `sessions continue`
- Provides instructions on how to continue a conversation from a saved session
- Usage: `crush sessions continue [session-id]`
- Note: Currently directs users to use the TUI session switcher (Ctrl+S) to load the specified session

### `sessions delete`
- Deletes a saved session by providing the session ID
- Usage: `crush sessions delete [session-id]`

### `sessions delete-all`
- Deletes all saved sessions with confirmation
- Usage: `crush sessions delete-all`

### `sessions export-all`
- Exports all sessions to JSON files
- Usage: `crush sessions export-all`

## Session Summarization

Added new CLI command for summarizing sessions:

### `summarize`
- Generates a summary of a session's conversation
- Usage: `crush summarize [session-id]`

## Session Search

Added new CLI command for searching sessions:

### `search`
- Searches through session titles or content
- Usage: `crush search [query]`

## Template Management

Added new CLI commands for managing prompt templates:

### `templates list`
- Lists all prompt templates
- Usage: `crush templates list`

### `templates create`
- Creates a new prompt template
- Usage: `crush templates create [name]`

### `templates use`
- Uses a prompt template to generate a prompt
- Usage: `crush templates use [name]`

## Usage Tracking

Added new CLI command for viewing usage and costs:

### `usage`
- Views token usage and associated costs for sessions or overall usage
- Usage: `crush usage`

## Model Testing

Added new CLI command for testing models:

### `test-model`
- Quickly tests a specific model with a prompt without starting a full session
- Usage: `crush test-model [provider] [model] [prompt]`