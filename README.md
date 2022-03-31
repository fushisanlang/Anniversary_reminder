# Anniversary_reminder

### 说明
用于记录一些纪念日，重要活动，比如开会日期，约会日期等。
通过提前添加相关日期，并设置提前通知天数。
服务运行时，会根据配置生成一个定时任务，根据配置的频率去扫描数据库，如果查询到响应记录，就会通过企业微信发送消息进行通知。
使用时，需要有一个企业微信的自定义应用来做信息的收发。
也可以自己修改模块，更改为邮件通知的模式。

### 开发环境说明
在ubuntu20.04LTS上，基于go1.15开发。数据库使用mysql5.5。

### 项目文件说明
src 源码目录 
    * Anniversary_reminder.go 项目主文件 
    * until 工具包目录 
        * Alert.go 定时任务功能功能 
        * DataBase.go 数据库相关功能 
        * DateSwitch.go 日期转换功能 
        * ReadConf.go 读取配置文件功能 
        * SendWx.go 微信推送功能 
bin 编译产物路径 
    * linux_64 64位linux环境 
        * Anniversary_reminder 编译产物 
    * macos_64 64位mac环境 
    *注：理论上mac可用，但是未进行实际测试* 
        * Anniversary_reminder 编译产物 
sql 初始化sql
     sql文件
conf.ini 配置文件

### 使用方法
*注：下文所有命令均基于linux，其他操作系统可能会有不同*
```shell
mkdir /usr/local/Anniversary_reminder -p

#首先需要初始化数据库
#使用sql/Anniversary_reminder.sql文件

#编译，需要有golang环境已经响应的包
go build Anniversary_reminder.go #会在当前路径生成Anniversary_reminder二进制文件
cp Anniversary_reminder /usr/local/Anniversary_reminder

#也可以使用预先编译好的二进制文件
cp bin/linux/Anniversary_reminder /usr/local/Anniversary_reminder

#修改配置文件
cp conf.ini /usr/local/Anniversary_reminder


cd /usr/local/Anniversary_reminder
chmod +x Anniversary_reminder

./Anniversary_reminder -cli on #打开cli命令行，进行数据的增删改查
./Anniversary_reminder -service start #开始定时任务
#可以使用以下命令将服务改为后台运行，适合服务器上使用
nohup ./Anniversary_reminder -service start &
#也可以结合supervisor或者systemctl做进程守护，本处不做赘述。
```

### 配置文件说明
```shell
[Anniversary_reminder]   #配置头
dbhost=0.0.0.0 #数据库地址
dbport=3306 #数据库端口
dbuser=root #数据库用户
dbpass=xxxx #数据库密码
dbname=Anniversary_reminder #数据库名
dbsuffix=charset=utf8 #数据库连接配置，可以使用此处的配置，即通过utf8编码连接
 
wxcorpid=xxx #企业微信id，通过网页版企业微信-管理后台-我的企业获得
wxcorpsecret=xxx #企业微信密码，网页版企业微信-应用管理-想要使用的应用
wxagentid=xxx #企业微信服务id，网页版企业微信-应用管理-想要使用的应用

cronvalue=* * * * * * #定时任务频率。分别为 秒 分 时 日 月 周，与linux的crontab类似
```
