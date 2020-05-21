# Blockchain
> Simple blockchain!

### Implemented:
1. PoW
2. P2P

### Components are used:
1. Go: go-sqlite3

### Component installation:
```
$ go get github.com/mattn/go-sqlite3
```

### Build and run
```
$ go build main.go
$ ./main -a 127.0.0.1:8080
```

### Commands
1. [:help = print help info]
2. [:exit = exit from client]
3. [:init = create genesis block]
4. [:whoami = print hashname]
5. [:network = print connections]
6. [:connect = connect to node in p2p]
7. [:blocks = read blockchain]
8. [:balance = get own or node balance]
9. [:push = create transaction]
