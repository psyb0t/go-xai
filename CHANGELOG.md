# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.1, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.0, 2026-10-16

- Added `xai`, a REST client generated with oapi-codegen from xAI's OpenAPI document (xAI's REST API, spec version 1.0.0, OpenAPI 3.1.0, from docs.x.ai). It covers all 39 paths: chat completions, the Responses API, messages, legacy completions, embeddings, tokenization, image and video generation, files, skills, document search, models and API key info.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec from docs.x.ai and regenerates the client. The spec is not committed.
- Added the `go-xai` agent skill: `claude plugin install go-xai@psyb0t` (Claude Code), `codex plugin add go-xai@psyb0t` (Codex), `openclaw skills install @psyb0t/go-xai` (ClawHub).
