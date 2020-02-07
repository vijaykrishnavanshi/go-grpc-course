export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"
protoc -I greet/ greet/greetpb/greet.proto --go_out=plugins=grpc:greet
