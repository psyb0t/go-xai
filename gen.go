// Package xai is a Go client for the xAI (Grok) REST API, generated with
// oapi-codegen from the OpenAPI document xAI publishes.
package xai

//go:generate bash scripts/fetch-spec.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.json
