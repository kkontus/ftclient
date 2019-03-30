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
$ protoc --proto_path=proto --go_out=plugins=grpc:proto questions.proto
$ go install

Run server (only one of the commands below, network, file or db):
$ ftserver questions --mode network -u https://gist.githubusercontent.com/kkontus/1cb8bc5aa35911df48c03e3854f82c16/raw/134d2cbd21cd25c01ac83d4d713a2bb4f7ec0c27/quiz

$ ftserver questions --mode file -f /full/path/to/project/src/quiz/ftserver/questions.json 
$ ftserver questions --mode file -f /Users/<user>/GolangProjects/src/ftserver/questions.json 

// for this to work mysql has to be installed and questions.sql imported
// this was tested on Docker container with following config
// https://github.com/kkontus/GoBittrexDockerConfig
$ ftserver questions --mode db

Run client:
$ ftclient questions

or

$ ftclient ranking
```

## TODO

Tests for util functions
