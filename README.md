# 目标
一件事情，如果自己不想重复去做，那就交给程序。

# 背景
在我们成长的过程中，脑海中会时刻产生灵感，

* 大部分人没有把握好灵感，
* 【少部分人】能做1次，然后束之高阁，
* 极少数人才坚持一直做，重复做，直到拿到结果。

我希望能帮助这【少部分人】拿到结果。

Auto-updating 项目，把你的灵感变成现实，并且在后台持续更新，自动化运行。

# 项目初始化及更新

**项目选型：**

* HTTP框架 Hertz https://github.com/cloudwego/hertz
* ORM框架 Gorm https://github.com/go-gorm/gorm
* 配置 Viper https://github.com/spf13/viper

开发者（一般用户可跳过）
```shell
# 安装 hz
go install github.com/cloudwego/hertz/cmd/hz@latest

# 新建
hz new  -mod github.com/jiangjilu/auto-updating

# 更新
go get github.com/cloudwego/thriftgo@latest
hz update --model_dir biz/hertz_gen -idl idl/api.thrift
hz update --model_dir biz/hertz_gen -idl idl/news.thrift
```

一般用户
```shell
# 下载
git clone git@github.com:jiangjilu/auto-updating.git

# 启动
go run .

# 访问
http://localhost:9090/
```

