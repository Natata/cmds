# cmds
A simple command service implement with grpc.

# How to use
## Server side
### Register actions for codes
```
set := cmds.Set{
    1: func(s string) error {
        if s == "btc" {
            fmt.Println("mining")
            return nil 
        }   
        return fmt.Errorf("wrong")
    },  
    2: func(s string) error {
        fmt.Println(s)
        return nil 
    },  
}   
s := cmds.InitCMDS(set)
```

### Run the server
```
addr := ":8080"
s.Run(addr)
```

### Run the server with TLS
```
credOpt, _ := cmds.CreateCred(*certFile, *keyFile)
addr := ":8080"
s.Run(addr, credOpt)
```

## Client side

### create client
```
cli, _ := client.InitClient(addr)
```

### create client with TLS
```
cred, _ := client.CreateCred(*certFile, *serverName)
cli, _ := client.InitClient(addr, cred)
```

### send the code and param
```
err = cli.Send(1, "btc")
if err != nil {
	panic(err)
}
fmt.Println("success")

err = cli.Send(1, "eth")
fmt.Println(err)
```
