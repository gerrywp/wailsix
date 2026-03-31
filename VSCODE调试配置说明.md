# Wails 项目 VSCode 调试配置详解

## 概述

这个配置允许你在 VSCode 中使用 F9 设置断点，F5 启动调试，像调试普通 Go 程序一样调试 Wails 应用。

## 配置文件说明

### 1. `.vscode/launch.json` - 调试启动配置

这个文件告诉 VSCode 如何启动和附加到调试进程。

#### 配置项 1: "Debug Wails App" - 直接启动调试
```json
{
  "name": "Debug Wails App",
  "type": "go",           // 告诉 VSCode 这是 Go 语言调试
  "request": "launch",    // 启动新进程而不是附加
  "mode": "exec",         // 执行已编译的程序
  "program": "${workspaceFolder}/build/bin/wailsix.exe",  // 可执行文件路径
  "preLaunchTask": "build-wails",  // 调试前先运行构建任务
  "cwd": "${workspaceFolder}",      // 工作目录
  "showLog": true                    // 显示调试日志
}
```

**工作流程：**
1. 运行 `preLaunchTask` 指定的构建任务
2. 编译生成 `build/bin/wailsix.exe`
3. 启动这个程序并附加调试器
4. 你就可以在 Go 代码中设置断点了

#### 配置项 2: "Debug Wails (wails dev)" - 附加到现有进程
```json
{
  "name": "Debug Wails (wails dev)",
  "type": "go",
  "request": "attach",    // 附加到已运行的进程
  "mode": "local",
  "processId": "${command:pickProcess}",  // 让你选择要附加的进程
  "showLog": true
}
```

**工作流程：**
1. 先在终端运行 `wails dev`
2. 按 F5 选择这个配置
3. VSCode 会弹出进程列表，选择你的应用进程
4. 调试器附加成功后就可以设置断点了

### 2. `.vscode/tasks.json` - 任务配置

这个文件定义了 VSCode 可以执行的任务。

#### 任务 1: "build-wails" - 构建调试版本
```json
{
  "label": "build-wails",
  "type": "shell",
  "command": "wails",
  "args": ["build", "-debug"],  // 使用 -debug 参数构建调试版本
  "group": {
    "kind": "build",
    "isDefault": true            // 设为默认构建任务
  },
  "problemMatcher": ["$go"]      // Go 编译器错误匹配器
}
```

**关键点：**
- `-debug` 参数很重要，它会：
  - 保留调试符号
  - 不优化代码（方便调试）
  - 生成可调试的二进制文件

#### 任务 2: "wails-dev" - 运行开发服务器
```json
{
  "label": "wails-dev",
  "type": "shell",
  "command": "wails",
  "args": ["dev"],
  "isBackground": true,  // 作为后台任务运行
  "problemMatcher": {
    "background": {
      "beginsPattern": "Building application for development",
      "endsPattern": "Watching"
    }
  }
}
```

## 实际使用步骤

### 方法一：完整调试（推荐新手）

1. **在 Go 代码中设置断点**
   - 打开 `app.go` 或 `main.go`
   - 在你想调试的代码行号左侧点击，会出现红点
   - 或按 F9 在当前行设置断点

2. **启动调试**
   - 按 F5 键
   - 选择 "Debug Wails App"
   - 等待编译完成（第一次可能需要几分钟）

3. **调试**
   - 应用启动后，触发断点处的代码
   - 程序会在断点处暂停
   - 使用调试工具栏控制：
     - 继续 (F5)
     - 单步跳过 (F10)
     - 单步进入 (F11)
     - 单步跳出 (Shift+F11)
     - 重启 (Ctrl+Shift+F5)
     - 停止 (Shift+F5)

### 方法二：配合 wails dev（推荐开发时）

1. **启动 wails dev**
   ```bash
   wails dev
   ```

2. **在 Go 代码中设置断点**
   - 按 F9 在需要的地方设置断点

3. **附加调试器**
   - 按 F5
   - 选择 "Debug Wails (wails dev)"
   - 在弹出的进程列表中找到你的应用进程（通常叫 wailsix.exe）
   - 选择它

4. **开始调试**
   - 现在你可以正常使用应用
   - 当代码执行到断点时会暂停

## 调试技巧

### 1. 查看变量值
- 当程序暂停在断点时
- 鼠标悬停在变量上可以看到当前值
- 左侧 "变量" 面板显示所有局部变量
- "监视" 面板可以添加你想跟踪的表达式

### 2. 查看调用堆栈
- "调用堆栈" 面板显示函数调用链
- 点击可以跳转到对应的调用位置

### 3. 条件断点
- 右键点击断点 → "编辑断点"
- 输入条件表达式，例如 `username == "admin"`
- 只有条件满足时才会暂停

### 4. 日志点
- 右键点击断点 → "编辑断点"
- 选择 "日志点"
- 输入要输出的消息，不会暂停程序

## 前置要求

### 1. 安装 Go 扩展
在 VSCode 扩展市场搜索并安装：
- **Go** (作者: Go Team at Google)

### 2. 安装 delve 调试器
Go 扩展通常会自动安装，如果没有：
```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

### 3. 配置 Go 环境
确保 `GOPATH` 和 `GOROOT` 环境变量正确设置。

## 常见问题

### Q: 调试时提示找不到 dlv?
A: 确保 `$GOPATH/bin` 在你的 PATH 环境变量中。

### Q: 断点显示为"未验证"或不生效?
A: 
- 确保使用 `-debug` 参数构建
- 检查代码是否被编译器优化掉了
- 清理构建缓存：`wails build -clean`

### Q: 进程列表中找不到我的应用?
A: 
- 确保 `wails dev` 正在运行
- 尝试按名称搜索 "wails" 或你的应用名

### Q: 前端代码能调试吗?
A: 
- Go 代码用这个配置调试
- 前端代码可以用浏览器的开发者工具调试
- 或单独配置前端调试（使用 Chrome Debugger 扩展）

## 进阶配置

如果你想更精细地控制调试，可以修改 launch.json：

```json
{
  "env": {
    "MY_DEBUG_VAR": "value"  // 设置环境变量
  },
  "buildFlags": "-tags=debug",  // 构建标签
  "dlvFlags": "--check-go-version=false",  // delve 额外参数
  "showGlobalVariables": true  // 显示全局变量
}
```

希望这个详细的说明能帮助你理解和使用 Wails 项目的调试配置！🎉