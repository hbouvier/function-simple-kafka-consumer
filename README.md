# Simple kafka function



### To use libraries in the function

```bash
go get -u github.com/golang/dep/cmd/dep
cd ~/gopath/src/github.com/hbouvier/function-simple-kafka-consume
cd kafka-message
dep init .
# dep ensure -add github.com/google/uuid
# dep ensure -add github.com/segmentio/kafka-go
cd ..
faas-cli build
```

