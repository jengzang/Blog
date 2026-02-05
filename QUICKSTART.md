# 快速開始指南 | Quick Start Guide

## 最小化設置步驟 | Minimal Setup Steps

### 1. 安裝依賴 | Install Dependencies

```bash
# 確保已安裝 Node.js 和 Go
node --version  # 應該 >= 16
go version      # 應該 >= 1.21

# 安裝 Hexo CLI
npm install -g hexo-cli

# 安裝項目依賴
cd blog
npm install

# 安裝 Butterfly 主題依賴
cd themes/butterfly
npm install
cd ../..
```

### 2. 配置後端 | Configure Backend

```bash
cd ../backend

# 創建 .env 文件
cp .env.example .env

# 編輯 .env，至少修改以下內容：
# JWT_SECRET=your-random-secret-key
# ACCESS_KEY_1=your-bcrypt-hashed-password
```

### 3. 啟動項目 | Start Project

#### 方式一：開發模式（推薦用於測試）| Method 1: Development Mode

```bash
# 終端 1: 啟動 Hexo
cd blog
hexo server

# 終端 2: 啟動 Go 後端
cd backend
go run main.go
```

訪問 http://localhost:4000 查看 Hexo 預覽
訪問 http://localhost:8080 查看完整應用（包含認證功能）

#### 方式二：生產模式 | Method 2: Production Mode

```bash
# 生成靜態文件
cd blog
hexo clean && hexo generate

# 啟動 Go 服務器（會自動服務靜態文件）
cd ../backend
go run main.go
```

訪問 http://localhost:8080

### 4. 測試簡歷頁面認證 | Test Resume Authentication

1. 訪問 http://localhost:8080/resume/
2. 輸入密鑰：`secret123`（默認密鑰）
3. 點擊驗證按鈕
4. 成功後應該看到簡歷內容

## 自定義配置 | Customization

### 修改網站信息 | Modify Site Information

編輯 `blog/_config.yml`:

```yaml
title: 你的博客名稱
subtitle: 副標題
description: 網站描述
author: 你的名字
language: zh-TW
```

### 修改導航菜單 | Modify Navigation Menu

編輯 `blog/_config.butterfly.yml`:

```yaml
menu:
  首頁: / || fas fa-home
  分類: /categories/ || fas fa-folder-open
  # 添加更多菜單項...
```

### 生成新的訪問密鑰 | Generate New Access Key

使用在線工具生成 bcrypt 哈希：
https://bcrypt-generator.com/

或使用 Go 代碼：

```go
package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    password := "your-new-password"
    hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    fmt.Println(string(hash))
}
```

將生成的哈希值更新到 `backend/.env` 的 `ACCESS_KEY_1`。

## 創建第一篇文章 | Create Your First Post

```bash
cd blog
hexo new post "我的第一篇文章"
```

編輯生成的文件 `source/_posts/我的第一篇文章.md`:

```markdown
---
title: 我的第一篇文章 | My First Post
date: 2026-02-05 10:00:00
categories:
  - 語言學
tags:
  - 測試
cover: https://source.unsplash.com/800x600/?blog
---

這是我的第一篇博客文章！

This is my first blog post!

<!-- more -->

## 正文

在這裡寫你的內容...
```

保存後刷新瀏覽器即可看到新文章。

## 部署到生產環境 | Deploy to Production

### 選項 1: 使用 GitHub Pages（僅前端）

```bash
cd blog
npm install hexo-deployer-git --save

# 編輯 _config.yml
deploy:
  type: git
  repo: https://github.com/yourusername/yourusername.github.io.git
  branch: main

# 部署
hexo clean && hexo deploy
```

### 選項 2: 使用 VPS（前端 + 後端）

```bash
# 1. 生成靜態文件
cd blog
hexo clean && hexo generate

# 2. 構建 Go 後端
cd ../backend
go build -o blog-server main.go

# 3. 上傳到服務器
scp -r blog/public user@your-server:/var/www/blog/
scp backend/blog-server user@your-server:/var/www/blog/
scp backend/.env user@your-server:/var/www/blog/

# 4. 在服務器上運行
ssh user@your-server
cd /var/www/blog
./blog-server
```

### 選項 3: 使用 Docker

創建 `Dockerfile`:

```dockerfile
FROM golang:1.21 AS builder
WORKDIR /app
COPY backend/ .
RUN go build -o blog-server main.go

FROM node:18 AS hexo-builder
WORKDIR /app
COPY blog/ .
RUN npm install && npm run build

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/blog-server .
COPY --from=hexo-builder /app/public ./public
COPY backend/.env .
EXPOSE 8080
CMD ["./blog-server"]
```

構建和運行：

```bash
docker build -t my-blog .
docker run -p 8080:8080 my-blog
```

## 故障排除 | Troubleshooting

### Hexo 無法啟動

```bash
# 清理緩存
hexo clean

# 重新安裝依賴
rm -rf node_modules package-lock.json
npm install
```

### Go 後端無法啟動

```bash
# 檢查 Go 版本
go version

# 重新下載依賴
go mod download

# 檢查端口是否被佔用
netstat -ano | findstr :8080  # Windows
lsof -i :8080                 # Linux/Mac
```

### 簡歷頁面認證失敗

1. 檢查 Go 後端是否正在運行
2. 檢查瀏覽器控制台是否有 CORS 錯誤
3. 確認密鑰哈希正確
4. 檢查 JWT_SECRET 是否設置

## 下一步 | Next Steps

- 📝 添加更多文章
- 🎨 自定義主題樣式
- 🔧 配置評論系統（Gitalk、Disqus 等）
- 📊 添加網站統計（Google Analytics）
- 🔍 配置 SEO 優化
- 🌐 設置自定義域名

## 獲取幫助 | Get Help

- Hexo 文檔: https://hexo.io/docs/
- Butterfly 主題文檔: https://butterfly.js.org/
- Gin 框架文檔: https://gin-gonic.com/docs/

祝你使用愉快！ | Enjoy your blogging!
