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

### To create the sealed secrets
```bash
~/bin/kubeseal --fetch-cert --controller-name ofc-sealedsecrets-sealed-secrets  > pub-cert.pem
faas-cli cloud seal --name hbouvier-function-simple-kafka-consumer --literal kafka-response-topic=response --literal kafka-url=kafka.openfaas:9092
```
