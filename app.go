package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Request struct {
		Header struct {
			MessageName   string `json:"MESSAGENAME"`
			TransactionID string `json:"TRANSACTIONID"`
			OrgRRN        string `json:"ORGRRN"`
			OrgName       string `json:"ORGNAME"`
			UserName      string `json:"USERNAME"`
			Language      string `json:"LANGUAGE"`
			IP            string `json:"IP"`
		} `json:"Header"`
		Body struct {
			Check     bool   `json:"CHECK"`
			DirtyFlag bool   `json:"DIRTYFLAG"`
			UserName  string `json:"USERNAME"`
			Password  string `json:"PASSWORD"`
		} `json:"Body"`
	} `json:"Request"`
}

// LoginResponse 登录响应结构体
type LoginResponse struct {
	Response struct {
		Header struct {
			TransactionID string `json:"TRANSACTIONID"`
			Result        string `json:"RESULT"`
		} `json:"Header"`
		Body struct {
			LoginResult bool `json:"LOGINRESULT"`
			User        struct {
				UserName        string `json:"USERNAME"`
				Description     string `json:"DESCRIPTION"`
				DefaultLanguage string `json:"DEFAULTLANGUAGE"`
				DefaultOrgRRN   int    `json:"DEFAULTORGRRN"`
				IsClientEncry   string `json:"ISCLIENTENCRY"`
				Department      string `json:"DEPARTMENT"`
				AdUserGroupList []struct {
					UserGroupID string `json:"USERGROUPID"`
				} `json:"ADUSERGROUPLIST"`
			} `json:"USER"`
		} `json:"Body"`
	} `json:"Response"`
}

// App struct
type App struct {
	ctx    context.Context
	config *Config
}

// NewApp creates a new App application struct
func NewApp() *App {
	// 加载配置
	config, err := LoadConfig()
	if err != nil {
		fmt.Printf("配置加载错误: %v\n", err)
		config = defaultConfig() // 使用默认配置
	}

	return &App{
		config: config,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowMaximise(ctx)
}

// navigateTo 运行时方法，导航到指定页面
func (a *App) navigateTo(page string) {
	// 发送事件到前端
	runtime.EventsEmit(a.ctx, "navigate", page)
}

// SetConfig 设置配置
func (a *App) SetConfig(config *Config) {
	a.config = config
}

// GetConfig 获取配置
func (a *App) GetConfig() *Config {
	return a.config
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

// Login 验证用户登录
func (a *App) Login(username, password string) bool {
	// 构造登录请求
	reqData := LoginRequest{}
	reqData.Request.Header.MessageName = "PCB.MES.USER_LOGIN"
	reqData.Request.Header.TransactionID = fmt.Sprintf("%d", time.Now().UnixNano()/1000000) // 生成时间戳作为交易ID
	reqData.Request.Header.OrgRRN = "378341"
	reqData.Request.Header.OrgName = "新创元"
	reqData.Request.Header.UserName = username
	reqData.Request.Header.Language = "ZH"
	reqData.Request.Header.IP = "10.1.12.54"

	reqData.Request.Body.Check = false
	reqData.Request.Body.DirtyFlag = false
	reqData.Request.Body.UserName = username
	reqData.Request.Body.Password = password

	// 转换为JSON字符串
	reqJSON, err := json.Marshal(reqData)
	if err != nil {
		fmt.Printf("JSON序列化错误: %v\n", err)
		return false
	}

	// 构造表单数据
	formData := url.Values{}
	formData.Set("senderId", "PCBMESSender")
	formData.Set("requestMessage", string(reqJSON))

	// 发送POST请求
	client := &http.Client{
		Timeout: time.Duration(a.config.GetESBAPITimeout()) * time.Second,
	}

	resp, err := client.PostForm(a.config.GetESBAPIURL(), formData)
	if err != nil {
		fmt.Printf("请求发送错误: %v\n", err)
		return false
	}
	defer resp.Body.Close()

	// 读取响应
	var respData LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		fmt.Printf("响应解析错误: %v\n", err)
		return false
	}

	// 检查登录结果，根据实际响应格式
	if respData.Response.Header.Result == "SUCCESS" && respData.Response.Body.LoginResult {
		return true
	}
	fmt.Printf("登录失败: 结果=%s, 登录结果=%v\n", respData.Response.Header.Result, respData.Response.Body.LoginResult)
	return false
}
