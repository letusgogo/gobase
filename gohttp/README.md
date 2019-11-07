# 启动 http server 

启动多个 httpserver 监听不同端口

```

func TestNewMultiHttpServer( t *testing.T ){
	hand1 := HandlerMap{
		"/a":handA,
		"/b":handB,
	}
	server1 := NewHttpServer(":8081", hand1)
	go server1.StartHttpServer()
	go server1.WaitForTerminal()



	hand2 := HandlerMap{
		"/c":handC,
		"/d":handD,
	}

	server2 := NewHttpServer(":8082", hand2)
	go server2.StartHttpServer()
	server2.WaitForTerminal()
}
```