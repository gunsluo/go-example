package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set the OPENAI_API_KEY environment variable")
		os.Exit(1)
	}
	openaiClient := NewClient(apiKey)

	ctx := context.Background()
	// models, err := openaiClient.ListModels(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, m := range models {
	// 	fmt.Println("->", m)
	// }

	// openaiClient.UseModel("gpt-4")
	questions := []string{
		`I want you to act as a Records Management Specialist. I will give you meeting talks in Engligh, and you will give me the agenda from that. Do not include any explanations or additional information in your response, simply provide the generated agenda.`,
		`Learning a new language can be a challenging but rewarding experience.  It can help you connect with new people and cultures.  And it can also boost your cognitive abilities and enhance your career prospects.  Whether you're learning a language for fun, travel, or work,  there are many resources available to help you get started,  such as language apps, classes, and language exchange programs.  With dedication and practice, you can make significant progress  and achieve your language learning goals.`,
	}
	// `create a agenda from follow talk: `,

	answers, err := openaiClient.Chat(ctx, questions)
	if err != nil {
		panic(err)
	}

	for i, answer := range answers {
		fmt.Printf("question %d and answer is: \n%s\n", i, answer)
	}
}

type Client struct {
	apiKey string
	model  string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		model:  "gpt-3.5-turbo",
	}
}

func (c *Client) UseModel(model string) {
	c.model = model
}

type Model struct {
	Id          string       `json:"id"`
	Object      string       `json:"object"`
	Created     int          `json:"created"`
	OwnedBy     string       `json:"owned_by"`
	Root        string       `json:"root"`
	Permissions []Permission `json:"permission"`
}

type Permission struct {
	Id                 string `json:"id"`
	Object             string `json:"object"`
	Created            int    `json:"created"`
	AllowCreateEngine  bool   `json:"allow_create_engine"`
	AllowSampling      bool   `json:"allow_sampling"`
	AllowLogprobs      bool   `json:"allow_logprobs"`
	AllowSearchIndices bool   `json:"allow_search_indices"`
	AllowView          bool   `json:"allow_view"`
	AllowFineTuning    bool   `json:"allow_fine_tuning"`
	Is_blocking        bool   `json:"is_blocking"`
	Organization       string `json:"organization"`
}

type listModelsResponse struct {
	Data []Model `json:"data"`
}

func (c *Client) ListModels(ctx context.Context) ([]Model, error) {
	req := &openaiRequest{
		Method: http.MethodGet,
		ApiUrl: "https://api.openai.com/v1/models",
	}
	resp := &listModelsResponse{}

	if err := c.callApi(ctx, req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

type chatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func (c *Client) Chat(ctx context.Context, messages []string) ([]string, error) {
	data := &chatRequest{
		Model: c.model,
	}
	for _, message := range messages {
		data.Messages = append(data.Messages, Message{
			Role:    "user",
			Content: message,
		})
	}

	req := &openaiRequest{
		Method: http.MethodPost,
		ApiUrl: "https://api.openai.com/v1/chat/completions",
		Data:   data,
	}
	resp := &chatResponse{}

	if err := c.callApi(ctx, req, resp); err != nil {
		return nil, err
	}

	var contents []string
	for _, c := range resp.Choices {
		contents = append(contents, c.Message.Content)
	}

	if len(contents) == 0 {
		return nil, fmt.Errorf("No response from API")
	}

	return contents, nil
}

type openaiRequest struct {
	Method string
	ApiUrl string
	Data   interface{}
}

func (c *Client) callApi(ctx context.Context, req *openaiRequest, output interface{}) error {
	var body io.Reader
	if req.Data != nil {
		buffer, err := json.Marshal(req.Data)
		if err != nil {
			return err
		}

		body = bytes.NewBuffer(buffer)
	}

	httpReq, err := http.NewRequest(req.Method, req.ApiUrl, body)
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("code: %d, body: %v", httpResp.StatusCode, string(respBody))
	}

	if output != nil {
		err = json.Unmarshal(respBody, output)
		if err != nil {
			return err
		}
	}

	return nil
}
