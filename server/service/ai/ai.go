package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/global"
	"sync"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

var (
	lastRequestTime                                                                                                                      time.Time
	mu                                                                                                                                   sync.Mutex
	TbaiURL, TbaiModel, TbaiApiKey, UnlimitURL, UnlimitModel, UnlimitApiKey, OpenrouterURL, OpenrouterModel, DeepSeekURL, DeepseekApiKey string
)

func GetResponse(originText string) string {
	TbaiURL = global.GVA_VP.GetString("aiconfig.tbai_url")
	TbaiModel = global.GVA_VP.GetString("aiconfig.tbai_model")
	TbaiApiKey = global.GVA_VP.GetString("aiconfig.tbai_api_key")
	UnlimitURL = global.GVA_VP.GetString("aiconfig.unlimit_url")
	UnlimitModel = global.GVA_VP.GetString("aiconfig.unlimit_model")
	UnlimitApiKey = global.GVA_VP.GetString("aiconfig.unlimit_api_key")
	OpenrouterURL = global.GVA_VP.GetString("aiconfig.openrouter_url")
	OpenrouterModel = global.GVA_VP.GetString("aiconfig.openrouter_model")
	DeepSeekURL = global.GVA_VP.GetString("aiconfig.deepseek_url")
	DeepseekApiKey = global.GVA_VP.GetString("aiconfig.deepseek_api_key")
	mu.Lock()
	now := time.Now()
	if !lastRequestTime.IsZero() {
		elapsed := now.Sub(lastRequestTime)
		if elapsed < 20*time.Second {
			waitTime := 20*time.Second - elapsed
			fmt.Printf("等待 %v 以避免速率限制\n", waitTime)
			time.Sleep(waitTime)
		}
	}
	lastRequestTime = time.Now()
	mu.Unlock()
	unlimitModel := UnlimitModel
	if unlimitModel == "" {
		unlimitModel = "gpt-4.1-nano"
	}
	requestData := ChatRequest{
		Model: unlimitModel,
		Messages: []Message{
			{
				Role:    "user",
				Content: originText,
			},
		},
	}
	maxRetries := 2
	baseDelay := 35 * time.Second
	// 尝试 Unlimit API
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			delay := time.Duration(1<<uint(attempt-1)) * baseDelay
			if delay > 2*time.Minute {
				delay = 2 * time.Minute
			}
			fmt.Printf("Unlimit API 第 %d 次重试，等待 %v\n", attempt+1, delay)
			time.Sleep(delay)
		}
		result, shouldRetry, err := makeUnlimitRequest(requestData, UnlimitApiKey)
		if err != nil {
			if !shouldRetry || attempt == maxRetries-1 {
				fmt.Printf("Unlimit API 失败: %v\n", err)
				break
			}
			fmt.Printf("Unlimit API 请求失败，准备重试: %v\n", err)
			continue
		}
		return result
	}
	fmt.Println("Unlimit API 不可用，切换到 Tbai API")
	tbaiModel := TbaiModel
	if tbaiModel == "" {
		tbaiModel = "gpt-4.1-nano"
	}
	requestData.Model = tbaiModel
	// 尝试 Tbai API
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			delay := time.Duration(1<<uint(attempt-1)) * baseDelay
			if delay > 2*time.Minute {
				delay = 2 * time.Minute
			}
			fmt.Printf("Tbai API 第 %d 次重试，等待 %v\n", attempt+1, delay)
			time.Sleep(delay)
		}
		result, shouldRetry, err := makeTbaiRequest(requestData, TbaiApiKey)
		if err != nil {
			if !shouldRetry || attempt == maxRetries-1 {
				fmt.Printf("Tbai API 失败: %v\n", err)
				break
			}
			fmt.Printf("Tbai API 请求失败，准备重试: %v\n", err)
			continue
		}
		return result
	}
	fmt.Println("Tbai API 不可用，切换到 DeepSeek API")
	requestData.Model = "deepseek-chat"
	// 尝试 DeepSeek API
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			delay := time.Duration(1<<uint(attempt-1)) * baseDelay
			if delay > 3*time.Minute {
				delay = 3 * time.Minute
			}
			fmt.Printf("DeepSeek API 第 %d 次重试，等待 %v\n", attempt+1, delay)
			time.Sleep(delay)
		}
		result, shouldRetry, err := makeDeepSeekRequest(requestData, DeepseekApiKey)
		if err != nil {
			if !shouldRetry || attempt == maxRetries-1 {
				fmt.Printf("DeepSeek API 最终错误: %v\n", err)
				return ""
			}
			fmt.Printf("DeepSeek API 请求失败，准备重试: %v\n", err)
			continue
		}
		return result
	}
	return ""
}

func makeUnlimitRequest(requestData ChatRequest, apiKey string) (string, bool, error) {
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", false, fmt.Errorf("JSON序列化错误: %v", err)
	}
	req, err := http.NewRequest("POST", UnlimitURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", false, fmt.Errorf("创建请求错误: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	return sendRequest(req)
}

func makeTbaiRequest(requestData ChatRequest, apiKey string) (string, bool, error) {
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", false, fmt.Errorf("JSON序列化错误: %v", err)
	}
	req, err := http.NewRequest("POST", TbaiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", false, fmt.Errorf("创建请求错误: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	return sendRequest(req)
}

func makeDeepSeekRequest(requestData ChatRequest, apiKey string) (string, bool, error) {
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", false, fmt.Errorf("JSON序列化错误: %v", err)
	}
	req, err := http.NewRequest("POST", DeepSeekURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", false, fmt.Errorf("创建请求错误: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	return sendRequest(req)
}

func makeOpenRouterRequest(requestData ChatRequest, apiKey string) (string, bool, error) {
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", false, fmt.Errorf("JSON序列化错误: %v", err)
	}
	req, err := http.NewRequest("POST", OpenrouterURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", false, fmt.Errorf("创建请求错误: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	return sendRequest(req)
}

func sendRequest(req *http.Request) (string, bool, error) {
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", true, fmt.Errorf("发送请求错误: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", true, fmt.Errorf("读取响应错误: %v", err)
	}
	if resp.StatusCode == 429 {
		return "", true, fmt.Errorf("速率限制错误: HTTP %d, 响应: %s", resp.StatusCode, string(body))
	}
	if resp.StatusCode != http.StatusOK {
		shouldRetry := resp.StatusCode >= 500
		return "", shouldRetry, fmt.Errorf("HTTP错误: %d, 响应: %s", resp.StatusCode, string(body))
	}
	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		return "", false, fmt.Errorf("JSON解析错误: %v, 原始响应: %s", err, string(body))
	}
	if len(chatResponse.Choices) > 0 {
		fmt.Printf("使用的token数: %d\n", chatResponse.Usage.TotalTokens)
		return chatResponse.Choices[0].Message.Content, false, nil
	} else {
		return "", false, fmt.Errorf("没有收到有效回复")
	}
}

//func DeepSeekResponse(apiKey, originText string) string {
//	mu.Lock()
//	now := time.Now()
//	if !lastRequestTime.IsZero() {
//		elapsed := now.Sub(lastRequestTime)
//		if elapsed < 20*time.Second {
//			waitTime := 20*time.Second - elapsed
//			fmt.Printf("等待 %v 以避免速率限制\n", waitTime)
//			time.Sleep(waitTime)
//		}
//	}
//	lastRequestTime = time.Now()
//	mu.Unlock()
//	requestData := ChatRequest{
//		Model: "deepseek-chat",
//		Messages: []Message{
//			{
//				Role:    "user",
//				Content: originText,
//			},
//		},
//	}
//	maxRetries := 5
//	baseDelay := 35 * time.Second
//	for attempt := 0; attempt < maxRetries; attempt++ {
//		if attempt > 0 {
//			delay := time.Duration(1<<uint(attempt-1)) * baseDelay
//			if delay > 3*time.Minute {
//				delay = 3 * time.Minute
//			}
//			fmt.Printf("第 %d 次重试，等待 %v\n", attempt+1, delay)
//			time.Sleep(delay)
//		}
//		result, shouldRetry, err := makeDeepSeekRequest(requestData, apiKey)
//		if err != nil {
//			if !shouldRetry || attempt == maxRetries-1 {
//				fmt.Printf("最终错误: %v\n", err)
//				return ""
//			}
//			fmt.Printf("请求失败，准备重试: %v\n", err)
//			continue
//		}
//		return result
//	}
//	return ""
//}
