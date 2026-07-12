# Agent Instructions

## Coding Standards
- **No repeated string literals:** Do not repeat long string literals or character blocks. Use constants or package-level variables instead to improve maintainability.
- **No standard library for core functions:** Always use `tinywasm` equivalents instead of the Go standard library where available:
    - `github.com/tinywasm/fmt` instead of `fmt`
    - `github.com/tinywasm/json` instead of `encoding/json`
    - `github.com/tinywasm/model`
    - `github.com/tinywasm/orm`
    - `github.com/tinywasm/mcp`
- **JSON in tests:** For JSON operations in tests, use **only** `tinywasm/json`.

## MCP and sqlmcp specifics
- **JSON Schema:** Generating JSON Schema is NOT the responsibility of `sqlmcp`. This is handled by `tinywasm/mcp` via `Tool.Args` (which must implement `model.Fielder`).
- **No JSON Schema logic:** Do not include `encodeSchema` functions or manual `InputSchema` JSON strings in tool definitions.
- **SQL Validation:** When defining models for SQL input, use `model.Text()` and provide an explicit `Permitted` whitelist that includes necessary SQL symbols (quotes, operators, etc.), as the default `Text` kind blocks them for XSS protection.
