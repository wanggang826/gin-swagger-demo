## gin-swagger-demo

### 介绍
gin 结合 swagger简单封装，方便开箱即用，API服务及接口文档

### 目录
1. admin          -- 管理后台api
2. app            -- 客户端api
3. conf           -- 配置文件
4. docs           -- 文档
5. middleware     -- 中间件
6. models         -- 数据层
7. pkg            -- 一些工具包
8. routers        --路由（多模块 如:admin、app）
9. service        -- 方便api与models交互

### swagger 
1. go get -u github.com/swaggo/swag/cmd/swag
2. swag init  // 注意，一定要和main.go处于同一级目录
3. 初始化命令，在根目录生成一个docs文件夹

### gin
go get -u github.com/gin-gonic/gin

### 项目安装
1. git clone https://gitee.com/wanggang826/gin-swagger-demo.git
2. go mod tidy
4. 修改配置: 复制配置文件config/app.ini.example 为config/app.ini，并修改其中配置
5. 非热更新启动: go run main.go,
6. 热更新启动 ：避免每次改代码手工重启服务
```
go get github.com/pilu/fresh
使用 fresh 启动服务： 项目根目录下执行： fresh

```
7. 测试接口,浏览器访问： localhost:8000/ping
8. 本地需要有 mysql 和 redis

### 调试
1. 日志文件保存在 ：runtime/logs 目录下
2. 可以单步调试，golang 或者 vscode 都可以，vscode 单步时，记得在 vscode 中当前打开文件是 main.go，否没办法启动调试

### xorm
1. xorm 文档地址： https://gobook.io/read/gitea.com/xorm/manual-zh-CN/
2. xorm坑  where = 0 
3. created 、updated、deleted xorm auto 可控



