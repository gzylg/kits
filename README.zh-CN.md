# Y-Server

## 介绍

Y-Server 微服务

## 软件架构

软件架构说明

## 目录说明

```go
├── mp-ticket-recycle-server
    ├── api         接口（原 controllers）
    │   ├──── client    客户端接口
    │   │     ├──── v1    v1版本接口
    │   │     ├──── v2    v2版本接口
    │   │     └──── ...   ...
    │   ├──── mgr       管理端接口
    │   │     ├──── v1    v1版本接口
    │   │     ├──── v2    v2版本接口
    │   │     └──── ...   ...
    │   └──── apimp     mpSrv接口
    ├── configs     配置文件
    ├── docs        文档
    ├── examples    实例
    ├── global      全局对象
    ├── init        初始化
    ├── logs        日志
    ├── models      模型？？？？？？
    ├── modules     应用/功能模块
    │   ├──── cfg       配置文件读写模块
    │   └────
    ├── resource    资源目录
    ├── routers     路由
    │   ├──── client    用户端路由组
    │   ├──── filters   路由中间件
    │   ├──── mgr       管理端路由组
    │   └──── rtmp      mrSrv路由组
    ├── main.go
    └── go.mod
```

## git

git 操作记录

#### --- 基础操作

`git config --global gui.encoding utf-8` //git GUI 显示中文乱码的话，执行这个命令更换编码

`git remote add origin https://gitee.com/YourGiteeName/YourProjName.git` //把一个本地仓库与一个云端 Gitee 仓库关联

`git remote -v` //查看仓库.

#### --- 提交操作

`git add -A` //将所有文件的修改，文件的删除，文件的新建，都添加到暂存区。

`git commit -m "第一次提交"` //提交到本地库，并附加注释。

`git push origin xxx` //提交代码到指定分支（xxx 为分支名）

#### --- 分支操作

`git branch xxx` //新建分支（xxx 为分支名）

`git branch -d xxx` //删除本地分支（xxx 为分支名）

`git push origin --delete xxx` //删除远程分支（xxx 为分支名）

`git checkout xxx` //切换分支（xxx 为分支名）

`git checkout -b xxx` //新建并切换到新分支（xxx 为分支名）

`git branch -a` //查看分支

`git merge OtherXXX` //合并某分支到当前分支（OtherXXX 为其他分支名）

## go mod

`go mod init` //生成 go.mod 文件

`go mod download` //下载 go.mod 文件中指明的所有依赖

`go mod tidy` //整理现有的依赖

`go mod graph` //查看现有的依赖结构

`go mod edit` //编辑 go.mod 文件

`go mod vendor` //导出项目所有的依赖到 vendor 目录

`go mod verify` //校验一个模块是否被篡改过

`go mod why` //查看为什么需要依赖某模块

`go build -ldflags "-w -s"` //编译的时候删除 DWARF 信息

## 管理员权限

#### 1、打开命令终端输入：

go get -v -u github.com/akavel/rsrc

#### 2、在项目根目录创建名为 nac.manifest 的文件，并填充以下内容：

```go
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
    <trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
        <security>
            <requestedPrivileges>
                <requestedExecutionLevel level="requireAdministrator"/>
            </requestedPrivileges>
        </security>
    </trustInfo>
</assembly>
```

#### 3、项目根目录下打开命令行输入：`rsrc -manifest nac.manifest -o nac.syso`
