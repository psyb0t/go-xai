# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.3, 2026-10-17

- `ChatRequest.Tools` and `MessageRequest.Tools` doc comments now say the API accepts up to 1000 functions, up from 350.
- `ModelRequest.Tools` and `ModelResponse.Tools` doc comments now say the API accepts up to 1000 tools, up from 350.
- No field, method or type changed. This is a comment-only update to the generated client.

## v0.1.2, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.1, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.0, 2026-10-16

- Added `xai`, a REST client generated with oapi-codegen from xAI's OpenAPI document (xAI's REST API, spec version 1.0.0, OpenAPI 3.1.0, from docs.x.ai). It covers all 39 paths: chat completions, the Responses API, messages, legacy completions, embeddings, tokenization, image and video generation, files, skills, document search, models and API key info.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec from docs.x.ai and regenerates the client. The spec is not committed.
- Added the `go-xai` agent skill: `claude plugin install go-xai@psyb0t` (Claude Code), `codex plugin add go-xai@psyb0t` (Codex), `openclaw skills install @psyb0t/go-xai` (ClawHub).
