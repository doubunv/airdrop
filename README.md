# 目录简介

cmd: 项目核心执行文件
cmd/etc 配置文件位置 

# 部署步骤


#### 1. 修改为线上配置和数据库信息
    a. 修改 cmd/etc/mm-test.yaml中内容，其中#号开头的都需要修改，其他可以保持为默认值，修改完成后可以另存为mm-prod.yaml 

#### 2. go编译 编译为二进制文件
    进入cmd目录编译go程序
    cd .\cmd\
    go build -o main

#### 3. 部署文件
    将编译后的二进制文件和配置文件拷贝到服务器的某个位置
    例如拷贝到 /home/xyz下
    然后运行 /home/air-drop/main -f /home/air-drop/mm-prod.yaml 指定配置文件 
