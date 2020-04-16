###  部署必读

#### 开箱即用
>   1.使用 `goland(>=2019.3版本)` 打开本项目，运行自动下载依赖.  
>   2.建议go语言使用>=1.14版本,才能更好的支持 `go module` 包管理.  
>![avatar](http://139.196.101.31:2080/GinSkeleton.jpg)  

#### 版本 
V 1.0.02   2020-04-16 
>   1.容器、事件注册器调整命名规范，增加模糊处理函数。        

V 1.0.01   2020-04-15 
>   1.增加容器，将一些比较繁琐的功能模块率先注册在容器，方便后续调用。  
>   2.表单参数验证器首先注册在容器，避免在路由模块不停地引入表单验证器造成该文件过于庞大。   
>   3.函数类事件精简代码，删除原有的一个键对应多个事件的逻辑，目前设置为键值一一对应关系。   
>   4.Mysql、Redis数据库连接的释放统一注册在函数事件管理器，由程序退出时统一释放。   
>   5.容器存储变量修改为sync.map，避免了并发情况下发生bug。     

V 1.0.00   2020-04-14 
>   1.基于gin框架的web项目骨架.  
>   2.开发单体应用基本的功能模块全部已经封装完毕.  
