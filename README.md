- [ ] Gin

  - [ ] router

  - [ ] model

  - [ ] service

  - [ ] middleware

  - [ ] handler

  - [ ] Graceful restart or stop

- [ ] pkg

  - [ ] testify

  - [ ] errors

  - [ ] log

  - [ ] config

  - [ ] cache

  - [ ] validator

  - [ ] jwt

  - [ ] session

  - [ ] db

- [x] layout

遇到的问题：

1. 如何打印日志以定位错误位置
2. 如何优化代码结构
3. 数据库如何关闭连接

    https://github.com/EDDYCJY/go-gin-example/issues/18

    sql.Open() 实际上不建立与数据库的任何连接。也不验证驱动程序连接参数。相反，它只是准备数据库抽象以供以后使用。当第一次需要时，将懒惰地建立与底层数据库实例的第一个实际连接

    在本项目中，我们将 sql.DB 视为一个长期使用的包全局对象，它是一个连接池。不需要频繁 Open() 和 Close()，需要的是为不同的数据库请求返回不同的连接句柄

    而每次都 defer db.Close() 的行为更适合当该连接句柄已超出其生存函数时，应当关闭

    那么显然，用法上就不一样了
4. Fatal 和 Panic 有什么区别

    Fatal它会将错误信息写入日志，并且直接导致程序终止。
    
    Panic它会导致整个程序停止并立即引发一个运行时错误，并且Go语言会执行一系列的回滚操作，以释放已经分配的资源，并且将错误信息输出到错误日志