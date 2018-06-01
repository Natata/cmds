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

## Client side
### send the code and param
```
conn, _ := grpc.Dial(addr, grpc.WithInsecure())
client := server.NewCommandServiceClient(conn)
r, _ := client.Send(context.Background(), &server.Request{
    Code:  1,
    Param: "btc",
})
fmt.Println(r)

r, _ = client.Send(context.Background(), &server.Request{
    Code:  1,
    Param: "eth",
})
fmt.Println(r)
```
