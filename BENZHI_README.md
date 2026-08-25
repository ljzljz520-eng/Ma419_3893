# 小城非遗影像志

## 标准命令

```bash
CGO_ENABLED=0 go build ./...
go test -count=1 ./...
go run ./cmd/heritage -db heritage.db -addr :8080
```

服务提供 `/api/home` 首页聚合接口，支持 `category` 查询参数；`POST /api/articles` 创建文章。
