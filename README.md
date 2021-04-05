# go-jwt-example
### 記錄自己練習寫go-jwt

### 說明：使用jwt-go產生token，輸入後可以看到機密資料(只是一串文字)

## 使用方法
1. 分別cd到 client和server 資料夾中，分別在port 9002和9000啟動  
```bash
go run main.go
```
2. Get http://localhost:9002
得到一串token  
3. 複製這串token，使用postman如下  

GET  http://localhost:9000
|  KEY   | VALUE  |
|  ----  | ----  |
| token  | 這串token |
