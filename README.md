# 配置 goproxy 

$ go env -w GO111MODULE=on

$ go env -w GOPROXY=https://goproxy.cn,direct

# 克隆iris-blog


$  git clone https://github.com/songjiangfeng/iris-blog.git

## 安装依赖 

$ go get

## 安装 Iris CLI 
$ go get github.com/kataras/iris-cli


## 启动 redis server

$ redis-server 
## 创建数据库

$ mysqladmin create iris_blog
## 导入数据库

$ mysql -uroot -p iris_blog < iris_blog.sql

## 配置数据库文件

$ cp config.json.default config.json

## 初始化项目
$ iris-cli init

## 运行项目  (http://localhost:8080)
$ iris-cli run .

## 后台 http://localhost:8888/admin  （注意记得更改密码）
###  run local
`
 "extra": {
    "domain":"www.go365.tech go365.tech",
    "site_theme": "./app/theme/default",
    "env": "local"
  },
`
## 后台 https://domain/admin
###  run prod
`
 "extra": {
    "domain":"www.go365.tech go365.tech",
    "site_theme": "./app/theme/default",
    "env": "prod"
  },
`

用户名： admin
密码： admin

