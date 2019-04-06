# go_server

#### Create JarvisMessage
```shell
protoc --go_out=$GOPATH/src JarvisMessage/jarvis_message.proto #create go package
protoc --cpp_out=$GOPATH/src JarvisMessage/jarvis_message.proto #create c++ package
```


#### msg_rules
[goal_] : msg