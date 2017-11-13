# 说明

这是一个用来初始化一个go项目的模板。
主要完成的事情是设定了一个go项目的配置文件和日志系统。

配置文件和命令参数处理选择的库是
- multiconfig

日志选择的库是：
- go-logging
- lumberjack.v2

需要注意的是，整个系统只设定了一个日志处理器，如果需要更多不同
模块的不同日志，需要修改`log.go` 文件。

# 命令

> `gotemp --conf ./conf/conf.toml  --loglevel debug`

`conf` 是一定需要的
`loglevel` 是可选的，如果没有，则会选择配置文件里面的，如果配置文件里面也没有，
则默认是`debug`


# 日志输出
日志输出缺省使用了rotate模式的日志，在一定大小和日期会备份下来。
缺省的备份数量为3，可以在程序中修改。这样可以保证日志文件的大小是可控的。

日志可以输出到日志文件，同时也输出到stderr。