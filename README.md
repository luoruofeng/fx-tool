# fx-tool
该项目皆在使用命令行快速搭建go项目，搭建的项目将会使用Fx框架做为项目主体。

目前go的常用框架都会有过度封装的问题。我不认同这些框架的做法。比较喜欢uber出的fx框架，简单，规范。但是每次搭建都很麻烦。故写了该项目，来快熟搭建fx的项目。后续我会将常用的库都以模块的形式集成进来。    

基础版本包含功能如下：    
https://github.com/luoruofeng/fxdemo#basic%E5%88%86%E6%94%AF%E6%8F%90%E4%BE%9B%E7%9A%84%E5%8A%9F%E8%83%BD

常用三方库模块：   
https://github.com/luoruofeng/fx-component        
   
# 脚手架安装    

执行条件:安装了go    

```shell
# 下载项目
git clone https://github.com/luoruofeng/fx-tool.git

# 进入项目
cd  fx-tool

# 安装
go install .

# #删除源码
cd ..
rm -rf fx-tool
```

# 脚手架使用方法
创建项目
```shell
# 项目项目的URL格式如： github.com/org_name/project_name
fx-tool init -url="项目项目的URL"

# 例如：
fx-tool init -url="github.com/luoruofeng/xxxproj"
```
  
创建带三方模块的项目
```shell
# 具体有哪些常用三方库模块，以及使用方法，参考上方链接。
fx-tool init -url="项目项目的URL" 模块名称...

# 例如：grpc consul 可以这样写。
fx-tool init -url="项目项目的URL" grpc consul
```

对已有项目添加模块
```shell
# 如果需要对一个已有的项目添加模块：
cd 已经使用fx-tool创建好的项目
fx-tool add 模块名称

# 例如：
cd xxxproj
fx-tool add etcd kafka
```

# 项目使用说明


# 项目结构说明
https://github.com/luoruofeng/fxdemo#%E7%BB%93%E6%9E%84%E8%AF%B4%E6%98%8E
