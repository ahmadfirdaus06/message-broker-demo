### 1. Install required dependencies
```
go get github.com/streadway/amqp github.com/ricochet2200/go-disk-usage/du
```
### 2. Pay attention to snipper ```usage := du.NewDiskUsage("/")``` change the path whenener necessary "/" for linux, "c:/" for windows

### 3. Run provider in directory
```
go run provider.go
```
