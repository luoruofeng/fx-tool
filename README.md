# fx-tool

![Image text](https://github.com/luoruofeng/fx-tool/blob/master/logo.jpg?raw=true)

### 主旨
* 项目皆在使用命令行`快速搭建go项目`，搭建的项目将会使用`Fx框架`做为项目主体，并且做到了`模块化，超轻量，少封装`。 

### 痛点
* 该项目将解决这些痛点：`go常用整体框架问题框架中的三方模块过度封装。`*学习成本巨高。使用原生技术不香吗？  但凡go做较大的项目都会继承一堆三方模块，例如：GRPC ETCD CONCEL等等，使用过度封装的框架不仅要学如何使用三方模块本身，还需要学习框架所封装的三方模块的使用方法。* 该项目将会让这些问题不在困扰你。

### 特色
1. uber的fx框架非常出色，简单，轻量，规范。使用该脚手架可以来快速搭建fx的项目。  
2. 以最原生的方法使用集成到项目中的三方库，并且自带三方库的配置文件，docker启动文件。免去配置三方库的烦恼。欢迎大家一起来添加[常用三方模块](https://github.com/luoruofeng/fx-component)。     
   

<br>

---
    
## 脚手架安装        
```shell
go install github.com/luoruofeng/fx-tool@latest

```

<br>

## 脚手架使用方法
* 创建项目
```shell
# 项目项目的URL格式如： github.com/org_name/project_name
fx-tool initial -url="项目项目的URL"

# 例如：
fx-tool initial -url="github.com/luoruofeng/xxxproj"
```
  
* 创建带三方模块的项目
```shell
# 具体有哪些常用三方库模块，以及使用方法，参考上方链接。
fx-tool initial -url="项目项目的URL" 模块名称...

# 例如：grpc consul 可以这样写。
fx-tool initial -url="项目项目的URL" grpc consul
```

* 对已有项目添加模块
```shell
# 如果需要对一个已有的项目添加模块：
cd 已经使用fx-tool创建好的项目
fx-tool add 模块名称

# 例如：
cd xxxproj
fx-tool add etcd kafka
```

<br>
<br>

---


## 项目使用说明   
### [运行项目](https://github.com/luoruofeng/fxdemo#%E8%BF%90%E8%A1%8C)  

<br>

### [项目教程](https://github.com/luoruofeng/fxdemo#%E6%95%99%E7%A8%8B)  

<br>

### [基础版的项目结构](https://github.com/luoruofeng/fxdemo#%E9%A1%B9%E7%9B%AE%E7%BB%93%E6%9E%84)  

<br>

### [导入三方模块](https://github.com/luoruofeng/fx-component#fx-tool%E5%AF%BC%E5%85%A5%E5%BA%93)

<br>
