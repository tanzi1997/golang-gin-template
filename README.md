## 三层架构

基于三层架构进行开发  
表示层 router  
业务逻辑: buss  
数据访问层: Dao  

## 目录解释
buss 业务  
common 通用  
conf 配置文件  
dao 数据访问层  
log 日志  
middlewear 中间件  
model 数据库模型  
router 路由  
main 为启动  
## 表示层 router
入口功能rpc，api  
生成api 注解  
校验传入参数  
输出数据  

## 业务逻辑: buss
负责业务处理  
所有的业务相关的逻辑都在buss  

## 数据访问层: Dao
数据库访问  
数据库相关操作都在这里  

### 日志记录
log处理  

### 部署上线
Mackefile  

### 中间件

### 错误处理


