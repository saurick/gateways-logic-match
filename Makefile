
sourcePath=./cmd/gateways-logic-match/main.go

build-gateways-logic-match:
		go build -trimpath -o bin/gateways-logic-match ${sourcePath}
build-gateways-logic-match-linux:
    	GOOS=linux GOARCH=amd64 go build -trimpath -o bin/gateways-logic-match-linux ${sourcePath}
build-gateways-logic-match-win:
    	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -trimpath -o bin/gateways-logic-match-win ${sourcePath}