
```
- node版本 > v20.14.0
- golang版本 1.20
```

### 2.1 server项目

```bash
# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o server.exe main.go

# 运行
./server.exe
```

### 2.2 web项目

```bash
# 进入web文件夹
cd web

# 安装依赖
npm install

# 启动web项目
npm run serve
```
