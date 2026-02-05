# 語言與地理研究博客 | Linguistics & Geography Research Blog

基於 Hexo 的多語言個人博客，展示語言學和地理學領域的研究文章，支持主題風格切換、日夜間模式，並使用 Go 後端實現簡歷頁面的密鑰驗證功能。

A multilingual personal blog based on Hexo, showcasing research articles in linguistics and geography, with theme style switching, day/night mode, and Go backend for resume page authentication.

## 功能特點 | Features

- ✅ **多語言支持** | Multi-language Support: 繁體中文、簡體中文、英文
- ✅ **主題風格切換** | Theme Style Switching: 技術感風格 / 簡約現代風格
- ✅ **日夜間模式** | Day/Night Mode: 自動保存用戶偏好
- ✅ **簡歷頁面認證** | Resume Page Authentication: 基於 JWT 的密鑰驗證
- ✅ **響應式設計** | Responsive Design: 適配桌面、平板、手機
- ✅ **文章分類** | Article Categories: 語言學、地理學、學習筆記
- ✅ **外部鏈接** | External Links: 鏈接到知乎原文

## 技術棧 | Tech Stack

### 前端 | Frontend
- **框架** | Framework: Hexo 8.1.1
- **主題** | Theme: Butterfly
- **語言** | Languages: HTML, CSS, JavaScript
- **樣式** | Styling: Stylus

### 後端 | Backend
- **語言** | Language: Go 1.21+
- **框架** | Framework: Gin
- **認證** | Authentication: JWT (golang-jwt/jwt)
- **加密** | Encryption: bcrypt

## 項目結構 | Project Structure

```
HomePage/
├── blog/                    # Hexo 博客目錄
│   ├── source/             # 源文件
│   │   ├── _posts/         # 文章
│   │   ├── about/          # 關於頁面
│   │   ├── resume/         # 簡歷頁面（需認證）
│   │   ├── categories/     # 分類頁面
│   │   ├── tags/           # 標籤頁面
│   │   ├── achievements/   # 成就展示
│   │   ├── projects/       # 項目展示
│   │   ├── links/          # 友鏈
│   │   └── guestbook/      # 留言板
│   ├── themes/
│   │   └── butterfly/      # Butterfly 主題
│   ├── _config.yml         # Hexo 配置
│   └── _config.butterfly.yml # 主題配置
│
└── backend/                # Go 後端
    ├── config/             # 配置管理
    ├── middleware/         # 中間件
    ├── handlers/           # 請求處理
    ├── utils/              # 工具函數
    ├── main.go             # 入口文件
    ├── go.mod              # Go 模塊
    └── .env.example        # 環境變量示例
```

## 安裝與配置 | Installation & Configuration

### 前置要求 | Prerequisites

- Node.js 16+ 和 npm
- Go 1.21+ (用於後端)
- Git

### 1. 克隆項目 | Clone Repository

```bash
git clone <your-repo-url>
cd HomePage
```

### 2. 安裝 Hexo 依賴 | Install Hexo Dependencies

```bash
cd blog
npm install
```

### 3. 配置 Hexo | Configure Hexo

編輯 `blog/_config.yml`，修改以下配置：

```yaml
title: 你的博客標題
author: 你的名字
url: http://your-domain.com
```

### 4. 配置 Butterfly 主題 | Configure Butterfly Theme

編輯 `blog/_config.butterfly.yml`，自定義：
- 導航菜單
- 社交鏈接
- 頭像和圖片
- 顏色主題

### 5. 設置 Go 後端 | Setup Go Backend

```bash
cd ../backend

# 複製環境變量配置
cp .env.example .env

# 編輯 .env 文件，設置你的配置
# 特別是 JWT_SECRET 和 ACCESS_KEY_1
```

生成新的訪問密鑰哈希：

```bash
# 使用 Go 生成 bcrypt 哈希
go run -c 'package main; import ("fmt"; "golang.org/x/crypto/bcrypt"); func main() { hash, _ := bcrypt.GenerateFromPassword([]byte("your-password"), bcrypt.DefaultCost); fmt.Println(string(hash)) }'

# 或使用在線工具: https://bcrypt-generator.com/
```

安裝 Go 依賴：

```bash
go mod download
```

## 使用方法 | Usage

### 開發模式 | Development Mode

#### 1. 啟動 Hexo 開發服務器 | Start Hexo Dev Server

```bash
cd blog
hexo server
# 訪問 http://localhost:4000
```

#### 2. 生成靜態文件 | Generate Static Files

```bash
cd blog
hexo clean
hexo generate
```

#### 3. 啟動 Go 後端 | Start Go Backend

```bash
cd backend
go run main.go
# 服務器運行在 http://localhost:8080
```

### 生產部署 | Production Deployment

#### 1. 生成 Hexo 靜態文件 | Generate Hexo Static Files

```bash
cd blog
hexo clean
hexo generate
```

#### 2. 構建 Go 後端 | Build Go Backend

```bash
cd backend
go build -o blog-server main.go
```

#### 3. 運行服務器 | Run Server

```bash
./blog-server
```

服務器將在端口 8080 上運行，並自動服務 Hexo 生成的靜態文件。

## 創建新文章 | Creating New Posts

### 使用 Hexo 命令 | Using Hexo Command

```bash
cd blog
hexo new post "文章標題"
```

### 文章模板 | Post Template

```markdown
---
title: 中文標題 | English Title | 简体中文标题
date: 2026-02-05 10:00:00
categories:
  - 語言學  # 或 地理學、學習筆記
  - 子分類
tags:
  - 標籤1
  - 標籤2
cover: https://source.unsplash.com/800x600/?keyword
external_link: https://zhuanlan.zhihu.com/p/your-article-id
---

文章摘要（中文）

Article summary (English)

<!-- more -->

## 正文內容

...

---

**閱讀全文請訪問知乎 | Read full article on Zhihu:** [點擊這裡 | Click here](https://zhuanlan.zhihu.com/p/your-article-id)
```

## 簡歷頁面認證 | Resume Page Authentication

### 默認密鑰 | Default Key

- 密鑰 | Key: `secret123`
- 哈希 | Hash: `$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy`

### 修改密鑰 | Change Key

1. 生成新的 bcrypt 哈希
2. 更新 `backend/.env` 中的 `ACCESS_KEY_1`
3. 重啟 Go 服務器

### 認證流程 | Authentication Flow

1. 用戶訪問 `/resume`
2. 如果未認證，顯示密鑰輸入表單
3. 用戶輸入密鑰，前端調用 `/api/auth`
4. 後端驗證密鑰，生成 JWT token
5. Token 存儲在 HttpOnly Cookie 中
6. 認證成功後顯示簡歷內容

## API 端點 | API Endpoints

- `POST /api/auth` - 驗證密鑰並生成 token
- `GET /api/verify` - 驗證當前認證狀態
- `POST /api/logout` - 登出並清除 token

## 主題風格切換 | Theme Style Switching

博客支持兩種主題風格：

1. **技術感風格** | Tech Style
   - 深色背景
   - 代碼風格字體
   - 科技感動畫

2. **簡約現代風格** | Modern Minimalist Style
   - 淺色背景
   - 無襯線字體
   - 簡潔動畫

切換按鈕位於導航欄右側，用戶偏好保存在 LocalStorage 中。

## 多語言支持 | Multi-language Support

博客支持三種語言：
- 繁體中文 (zh-TW) - 默認
- 簡體中文 (zh-CN) - 通過轉換按鈕切換
- 英文 (en) - 文章內容包含英文翻譯

## 常見問題 | FAQ

### Q: 如何修改網站標題和描述？
A: 編輯 `blog/_config.yml` 文件中的 `title`、`subtitle` 和 `description` 字段。

### Q: 如何添加社交媒體鏈接？
A: 編輯 `blog/_config.butterfly.yml` 文件中的 `social` 部分。

### Q: 如何更改主題顏色？
A: 編輯 Butterfly 主題的 CSS 文件或在 `_config.butterfly.yml` 中配置顏色。

### Q: 簡歷頁面無法訪問？
A: 確保 Go 後端正在運行，並檢查瀏覽器控制台是否有錯誤信息。

### Q: 如何部署到生產環境？
A: 生成靜態文件後，可以部署到任何支持靜態網站的平台（GitHub Pages、Netlify、Vercel 等）。Go 後端需要部署到支持 Go 的服務器。

## 貢獻 | Contributing

歡迎提交 Issue 和 Pull Request！

## 許可證 | License

MIT License

## 聯繫方式 | Contact

- 博客 | Blog: [Your Blog URL]
- 知乎 | Zhihu: @YourZhihuID
- Email: your.email@example.com

---

**注意** | Note: 這是一個模板項目，請根據實際需求進行定制和修改。

This is a template project, please customize and modify according to your actual needs.

