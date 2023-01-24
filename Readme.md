# 基于go-zero框架的cloud-disk微服务网盘系统

*感谢[GetcharZp/cloud-disk: 基于go-zero实现的网盘系统](https://github.com/GetcharZp/cloud-disk)作者的贡献，本项目的业务逻辑思路基本来源于该项目。*

## 项目基本介绍和技术使用情况
> 这是一个轻量级云盘微服务系统，基于go-zero实现，官网：[go-zero帮助文档](https://go-zero.dev/cn/docs/introduction)
>
> 云盘存储存储引擎使用的是COS，官网：[腾讯云COS帮助文档](https://cloud.tencent.com/document/product/436/31215)
> 
> 使用JWT Token工具，生成了接口验证Auth的token，保证用户数据传输的安全
> 
> 使用Squirrel工具，在go-zero框架中简化了Sql语句的编写
> 
> 使用jordan-wright写的email工具，进行邮箱验证码的发送
> 
> 使用go-uuid工具，方便UUID的生成
> 
> 使用ApiFox工具，生成了在根目录下的接口测试导出文件cloud-disk.openapi.json，标准是openapi-3.0.1版本
> 
> 使用Goctl-Swagger插件，可以自行生成用于接口介绍的swagger网页

## 测试环境
* golang版本：`1.19.4 windows/amd64`
* redis版本：`6.2.6 linux`
* etcd版本：`3.4.23 windows` 
* Mysql版本：`5.7.17 windows`

## 项目目录树介绍
```text
cloud-disk:.
├─common //通用工具包
│  └─utils
└─service //服务层
   ├─repository //中心存储池服务
   │  ├─api 
   │  │  ├─etc
   │  │  └─internal
   │  │      ├─config
   │  │      ├─handler
   │  │      ├─logic
   │  │      ├─svc
   │  │      └─types
   │  ├─filePath
   │  ├─model
   │  └─rpc
   │      ├─etc
   │      ├─internal
   │      │  ├─config
   │      │  ├─logic
   │      │  ├─server
   │      │  └─svc
   │      ├─repository
   │      └─repositoryclient
   ├─share //文件分享服务
   │  ├─api
   │  │  ├─etc
   │  │  └─internal
   │  │      ├─config
   │  │      ├─handler
   │  │      ├─logic
   │  │      ├─svc
   │  │      └─types
   │  └─model
   ├─user //用户信息服务
   │  ├─api
   │  │  ├─etc
   │  │  └─internal
   │  │      ├─config
   │  │      ├─handler
   │  │      ├─logic
   │  │      ├─svc
   │  │      └─types
   │  ├─model
   │  └─rpc
   │      ├─etc
   │      ├─internal
   │      │  ├─config
   │      │  ├─logic
   │      │  ├─server
   │      │  └─svc
   │      ├─user
   │      └─userclient
   └─user_repository //个人存储池服务
       ├─api
       │  ├─etc
       │  └─internal
       │      ├─config
       │      ├─handler
       │      ├─logic
       │      ├─svc
       │      └─types
       ├─model
       └─rpc
           ├─etc
           ├─internal
           │  ├─config
           │  ├─logic
           │  ├─server
           │  └─svc
           ├─userRepository
           └─userrepositoryclient

```

## 微服务内容简述
该项目的业务逻辑思想来源于当前市场上的其他主流网盘。
### repository_pool 中心存储池资源管理服务
存储所有上传文件（注意不含文件夹）。 

根据网盘存储系统的思想，文件分享保存后，不需要在中心存储池中复制一份，只需要在个人存储池复制即可，即不同个人存储池的文件共享一个中心存储池的文件。

### share_basic 文件分享服务
将用户个人存储池中的文件分享等服务。

### user_basic 用户信息服务
用户的信息相关等服务。

### user_repository 个人存储池资源管理服务
存储用户上传文件、文件夹的简单信息。

存储结构就像一颗树一样，存储的每个值可能对应是文件夹或者文件，如果是文件夹则repository_identity值为空，ext是0； 
如果是文件，则repository_identity可以去关联中心存储池的identity，ext表示文件后缀，parent_id就表示该节点的父母节点的id。

根据网盘存储系统的思想，中心存储池的文件，映射到个人存储池中对应的userIdentity数目可能为0，也可能不止1，即拥有同一个文件的用户可能不止一个，也可能一个也没有。
但个人存储池中的文件一定在中心存储池中找得到，即一定拥有repositoryIdentity。


## 部分接口特别说明
* `fileListQuery`是指在指定的文件夹里查询文件。
* `shareBasicSave`是指把中心存储池的文件保存到自己的存储池下。
* `fileUpload`中的文件秒传思想就是根据文件的md5码判断该文件在中心存储池是否存在，若存在则秒传成功，不存在则继续传。

## 尚未完成的：
* [x] ~~model层的sql语句使用Squrrel工具帮助生成。~~
* [ ] go-zero框架数据库的查询语句，蜜汁不能返回多行数据，尝试使用gorm后就可以了，打算找一下bug
* [ ] 结果、错误模板修改。

## 如何运行该系统
### 先做好准备工作
* 填写工具类中的设置常量
  * 为系统的验证码发送邮箱申请授权码，并填写EmailAuthCode等值。注意邮箱要开启SMTP功能服务。
  * 申请COS存储桶，并填写BucketNameWithAPPID、SecretID、SecretKEY等值。
### 再运行基础服务：
* 运行ectd，并在每个服务的yaml文件中配置相应的地址和端口号
* 运行redis，并在每个服务的yaml文件中配置相应的地址和端口号
* 运行Mysql，并在每个服务的yaml文件中配置相应的地址、端口号、用户名和密码等信息
  记得每次都要改一下redis的host！
### 最后运行项目：
* Windows环境直接运行脚本start.bat文件即可。
* Linux环境，直接运行四个服务中的共7个yaml文件即可。（Windows环境也可以通过该方式运行项目）

## 最后
本项目是作者初学golang，用来练手的项目。在构造并完善该项目的过程中，还是学习到了很多内容的。
对于该项目，很明显不仅是在代码方面，而且业务逻辑方面也还是有许多问题的，包括但不限于：

若有其他问题的欢迎指出。
