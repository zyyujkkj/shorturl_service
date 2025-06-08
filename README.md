# ShortURL Service

一个基于 Go + go-zero 实现的高性能短链接系统，支持短链生成、跳转及访问统计功能。

## 🧩 技术栈

- Go 1.23+
- [go-zero](https://github.com/zeromicro/go-zero)
- MySQL
- Redis


## 🚀 快速启动

1. 修改配置文件（如数据库、Redis 等）
2. 启动 RPC 服务：

```bash
cd rpc/transform
go run transform.go
```
3. 启动 API 服务：
```bash
cd api
go run shorturl.go
```
##✅ 功能特性
✅ 短链接生成

✅ 短链接跳转

✅ 访问日志记录

✅ 防刷与缓存优化
