export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"
protoc -I calculator/ calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:calculator
