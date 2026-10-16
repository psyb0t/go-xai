package xai_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psyb0t/go-xai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testAPIKey = "test-key"

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *xai.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := xai.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+testAPIKey)

			return nil
		},
	)

	client, err := xai.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func assertBearer(t *testing.T, r *http.Request) {
	t.Helper()

	assert.Equal(t, "Bearer "+testAPIKey, r.Header.Get("Authorization"))
}

func TestGetAPIKeyInfo(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/api-key", r.URL.Path)
		assertBearer(t, r)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"api_key_id":"key-1","name":"dev","acls":["api-key:model:*"],` +
				`"api_key_blocked":false,"team_id":"team-1"}`,
		))
	})

	res, err := client.HandleGetApiKeyInfoRequestWithResponse(
		context.Background(),
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	assert.Equal(t, "key-1", res.JSON200.ApiKeyId)
	assert.Equal(t, "dev", res.JSON200.Name)
	assert.Equal(t, []string{"api-key:model:*"}, res.JSON200.Acls)
	assert.False(t, res.JSON200.ApiKeyBlocked)
	assert.Equal(t, "team-1", res.JSON200.TeamId)
}

func TestListModels(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/models", r.URL.Path)
		assertBearer(t, r)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"object":"list","data":[{"id":"model-a","object":"model",` +
				`"created":1700000000,"owned_by":"xai"}]}`,
		))
	})

	res, err := client.HandleModelsListRequestWithResponse(context.Background())
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Data, 1)

	assert.Equal(t, "list", res.JSON200.Object)
	assert.Equal(t, "model-a", res.JSON200.Data[0].Id)
	assert.Equal(t, int64(1700000000), res.JSON200.Data[0].Created)
}

func TestChatCompletion(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/chat/completions", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assertBearer(t, r)

		var body struct {
			Model    string `json:"model"`
			Messages []struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"messages"`
		}

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "model-a", body.Model)

		if assert.Len(t, body.Messages, 1) {
			assert.Equal(t, "user", body.Messages[0].Role)
			assert.Equal(t, "hello", body.Messages[0].Content)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"id":"chat-1","object":"chat.completion","created":1,` +
				`"model":"model-a","service_tier":"default",` +
				`"choices":[{"index":0,"finish_reason":"stop",` +
				`"message":{"role":"assistant","content":"hi"}}]}`,
		))
	})

	var content xai.Content
	require.NoError(t, content.FromContent0("hello"))

	var msg xai.Message
	require.NoError(t, msg.FromMessage2(xai.Message2{
		Content: content,
		Role:    xai.Message2RoleUser,
	}))

	model := "model-a"
	messages := []xai.Message{msg}

	res, err := client.HandleGenericCompletionRequestWithResponse(
		context.Background(),
		xai.ChatRequest{Model: &model, Messages: &messages},
	)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Choices, 1)

	assert.Equal(t, "chat-1", res.JSON200.Id)
	assert.Equal(t, "assistant", res.JSON200.Choices[0].Message.Role)
	assert.Equal(t, "hi", *res.JSON200.Choices[0].Message.Content)
}
