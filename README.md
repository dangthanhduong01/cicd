# Go Gin Login API

API đơn giản để xử lý login với Go và Gin framework.

## 🚀 Features

- ✅ Login endpoint với validation
- ✅ Health check endpoint
- ✅ CORS middleware
- ✅ Profile endpoint với authentication check
- ✅ CI/CD với GitHub Actions
- ✅ Deploy tự động lên Render

## 📋 API Endpoints

### GET /
Health check cơ bản
```bash
curl https://your-app.onrender.com/
```

### GET /health
Kiểm tra trạng thái API
```bash
curl https://your-app.onrender.com/health
```

### POST /api/login
Đăng nhập
```bash
curl -X POST https://your-app.onrender.com/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "password123"
  }'
```

**Response thành công:**
```json
{
  "message": "Login successful",
  "token": "fake-jwt-token-admin",
  "user": "admin"
}
```

**Response thất bại:**
```json
{
  "message": "Invalid username or password"
}
```

### GET /api/profile
Lấy thông tin profile (cần token)
```bash
curl https://your-app.onrender.com/api/profile \
  -H "Authorization: Bearer your-token"
```

## 🛠️ Local Development

### Prerequisites
- Go 1.21 or higher
- Git

### Setup

1. Clone repository:
```bash
git clone https://github.com/dangthanhduong01/cicd.git
cd cicd
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run main.go
```

Server sẽ chạy tại `http://localhost:8080`

### Test API locally

```bash
# Health check
curl http://localhost:8080/

# Login
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password123"}'

# Profile
curl http://localhost:8080/api/profile \
  -H "Authorization: Bearer fake-jwt-token-admin"
```

## 🚢 Deploy to Render

### Manual Deploy

1. Push code lên GitHub
2. Tạo tài khoản tại [Render](https://render.com)
3. Tạo new Web Service
4. Connect với GitHub repository
5. Render sẽ tự động detect và build

### Automatic Deploy với CI/CD

Push code lên GitHub và CI/CD sẽ tự động:
- Build và test code
- Deploy lên Render
- Chạy health check

## 📝 Default Credentials

**Username:** admin  
**Password:** password123

⚠️ **Lưu ý:** Đây chỉ là demo. Trong production, hãy sử dụng database và JWT token thật.

## 🔧 Environment Variables

- `PORT`: Port để chạy server (default: 8080)

## 📦 Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- Go 1.21+

## 📄 License

MIT License
