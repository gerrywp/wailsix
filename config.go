package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config 应用程序配置结构体
type Config struct {
	// ESB REST API 配置
	ESBAPI struct {
		// API 基础URL
		BaseURL string `json:"baseURL"`
		// 请求超时时间（秒）
		Timeout int `json:"timeout"`
	} `json:"esbAPI"`
}

// 默认配置
func defaultConfig() *Config {
	return &Config{
		ESBAPI: struct {
			BaseURL string `json:"baseURL"`
			Timeout int    `json:"timeout"`
		}{
			BaseURL: "http://172.16.34.12:8080/esbrest/RestService/postrequest/",
			Timeout: 30,
		},
	}
}

// 配置文件路径
func configFilePath() string {
	// 获取应用程序配置目录
	// 在 Wails 中，我们可以使用 appdir 或直接使用当前目录
	exePath, err := os.Executable()
	if err != nil {
		return "config.json"
	}
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, "config.json")
}

// LoadConfig 加载配置文件
func LoadConfig() (*Config, error) {
	filePath := configFilePath()

	// 检查配置文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 配置文件不存在，创建默认配置
		config := defaultConfig()
		if err := config.Save(); err != nil {
			return nil, err
		}
		return config, nil
	}

	// 读取配置文件
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// 解析配置
	var config Config
	if err := json.Unmarshal(fileContent, &config); err != nil {
		return nil, err
	}

	// 确保所有必要字段都有默认值
	if config.ESBAPI.BaseURL == "" {
		config.ESBAPI.BaseURL = defaultConfig().ESBAPI.BaseURL
	}
	if config.ESBAPI.Timeout <= 0 {
		config.ESBAPI.Timeout = defaultConfig().ESBAPI.Timeout
	}

	return &config, nil
}

// Save 保存配置到文件
func (c *Config) Save() error {
	filePath := configFilePath()

	// 序列化配置为 JSON
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	// 写入配置文件
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return err
	}

	return nil
}

// GetESBAPIURL 获取ESB API URL
func (c *Config) GetESBAPIURL() string {
	return c.ESBAPI.BaseURL
}

// GetESBAPITimeout 获取ESB API超时时间
func (c *Config) GetESBAPITimeout() int {
	return c.ESBAPI.Timeout
}

// SetESBAPIURL 设置ESB API URL
func (c *Config) SetESBAPIURL(url string) {
	c.ESBAPI.BaseURL = url
}

// SetESBAPITimeout 设置ESB API超时时间
func (c *Config) SetESBAPITimeout(timeout int) {
	c.ESBAPI.Timeout = timeout
}

// Validate 验证配置的完整性和有效性
func (c *Config) Validate() error {
	if c.ESBAPI.BaseURL == "" {
		return fmt.Errorf("ESB API 基础URL不能为空")
	}
	if c.ESBAPI.Timeout <= 0 {
		return fmt.Errorf("ESB API 超时时间必须大于0")
	}
	return nil
}
