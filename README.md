# ftclient
Simple Client CLI Quiz using Cobra and gRPC

Command to compile proto files
```
protoc --proto_path=proto --go_out=plugins=grpc:proto questions.proto
```

Gist with questions
```
https://gist.githubusercontent.com/kkontus/323b05ed729e53c7dd5307bf6231693a/raw/2ca073e5dbfd10a7ded4883a565584db71aff85c/questions
https://gist.githubusercontent.com/kkontus/1cb8bc5aa35911df48c03e3854f82c16/raw/134d2cbd21cd25c01ac83d4d713a2bb4f7ec0c27/quiz
```

Installation:

```
$ mkdir quiz && cd quiz
$ git clone https://github.com/kkontus/ftserver.git
$ git clone https://github.com/kkontus/ftclient.git

$ cd ftserver
$ go install

$ cd ftclient
$ go install

Run server:
$ ftserver questions -u https://gist.githubusercontent.com/kkontus/1cb8bc5aa35911df48c03e3854f82c16/raw/134d2cbd21cd25c01ac83d4d713a2bb4f7ec0c27/quiz

Run client:
$ ftclient questions
or

$ ftclient ranking
```
