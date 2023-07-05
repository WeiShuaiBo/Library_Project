# 项目结构
## api
    所有的路由都会转发到这里，相当于controller层
    负责解析前端请求的数据，并返回数据响应
## service
    业务逻辑处理层，被api调用
    负责对api层数据的校验与逻辑处理
## dao
    调用数据库，查询数据，并将数据响应给service层
## doc
    swagger 生成的文档
## models
    请求数据的模型和数据库的模型以及创建redis和mysql的连接
## routes
    路由
## config
    配置文件
## common
    封装的工具包：
       1 config:解析配置文件 2 jwt 3 md5加密：数据库敏感数据加密 4 Response 统一响应格式
       5 SendEmail: 发送qq邮箱 6 snowflake雪花算法: 生成唯一id  
