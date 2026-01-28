# 🚀 Hướng dẫn Deploy lên Render

## Bước 1: Chuẩn bị Repository

1. **Commit và push code lên GitHub:**
```bash
git add .
git commit -m "Add Go Gin login API with CI/CD"
git push origin main
```

## Bước 2: Tạo tài khoản Render

1. Truy cập [render.com](https://render.com)
2. Click **"Get Started"** hoặc **"Sign Up"**
3. Đăng nhập bằng GitHub account

## Bước 3: Deploy Web Service

### Cách 1: Auto-deploy từ render.yaml (Recommended)

1. Trong Render Dashboard, click **"New +"** → **"Blueprint"**
2. Chọn repository: **dangthanhduong01/cicd**
3. Render sẽ tự động detect file `render.yaml` và tạo service
4. Click **"Apply"**
5. Đợi deployment hoàn tất (2-5 phút)

### Cách 2: Manual setup

1. Trong Render Dashboard, click **"New +"** → **"Web Service"**
2. Connect với GitHub repository **dangthanhduong01/cicd**
3. Cấu hình service:
   - **Name**: `go-gin-api` (hoặc tên bạn muốn)
   - **Region**: `Singapore` (hoặc gần bạn nhất)
   - **Branch**: `main`
   - **Runtime**: `Go`
   - **Build Command**: `go build -o main .`
   - **Start Command**: `./main`
4. Click **"Create Web Service"**
5. Đợi deployment hoàn tất

## Bước 4: Kiểm tra Deployment

Sau khi deploy xong, bạn sẽ nhận được URL dạng:
```
https://go-gin-api-xxxx.onrender.com
```

### Test API:

**Health check:**
```bash
curl https://your-app.onrender.com/
```

**Login:**
```bash
curl -X POST https://your-app.onrender.com/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password123"}'
```

**Profile:**
```bash
curl https://your-app.onrender.com/api/profile \
  -H "Authorization: Bearer fake-jwt-token-admin"
```

## Bước 5: Setup Auto-Deploy (CI/CD)

Render đã tự động setup auto-deploy! Mỗi khi bạn push code lên branch `main`:

1. GitHub Actions sẽ chạy CI pipeline:
   - ✅ Build code
   - ✅ Run tests
   - ✅ Check formatting

2. Render tự động detect và deploy:
   - 🔄 Pull code mới
   - 📦 Build application
   - 🚀 Deploy lên production

## 🎯 CI/CD Workflow

```
Push to GitHub (main branch)
        ↓
GitHub Actions CI
├── Build & Test
├── Code Quality Check
└── Success ✅
        ↓
Render Auto-Deploy
├── Pull latest code
├── Build Go binary
├── Deploy to production
└── Health check ✅
        ↓
🎉 Live on Render!
```

## 📊 Monitor Deployment

1. Vào Render Dashboard
2. Click vào service của bạn
3. Xem logs trong tab **"Logs"**
4. Xem metrics trong tab **"Metrics"**

## 🔧 Environment Variables (Optional)

Nếu cần thêm environment variables:

1. Vào service trong Render Dashboard
2. Click tab **"Environment"**
3. Add variables:
   - `GIN_MODE=release` (cho production)
   - Các biến khác nếu cần

## ⚠️ Lưu ý

### Free Tier Limitations:
- Service sẽ sleep sau 15 phút không hoạt động
- Cold start có thể mất 30-60 giây
- 750 giờ miễn phí mỗi tháng

### Để tránh sleep:
- Upgrade lên paid plan ($7/tháng)
- Hoặc dùng cron job để ping API định kỳ

## 🐛 Troubleshooting

### Deployment failed?
1. Check logs trong Render Dashboard
2. Verify `go.mod` và `go.sum` đã được commit
3. Đảm bảo code build thành công locally: `go build -o main .`

### API không response?
1. Check logs trong Render
2. Verify environment variable `PORT` đang được set
3. Test locally trước: `PORT=10000 go run main.go`

### GitHub Actions failed?
1. Check Actions tab trong GitHub repo
2. Xem logs để tìm lỗi
3. Fix lỗi và push lại

## 🎉 Hoàn tất!

Giờ bạn đã có:
- ✅ Go Gin Login API
- ✅ CI/CD với GitHub Actions
- ✅ Auto-deploy lên Render
- ✅ Production-ready infrastructure

Mỗi lần push code, CI/CD sẽ tự động test và deploy! 🚀
