### 1.Install RabbitMQ image as docker (linux or macos may require sudo)
```
docker run -d --name rabbitmq -p 8081:15672 -p 5672:5672 rabbitmq:3-management
```
### 2.Go to browser localhost:8081 to access rabbitmq management
