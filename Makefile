install-deps:
	go get github.com/rc1140/next45/cmd
	go get github.com/stretchr/testify/assert

install: cmd/main.go
	go build -o $(GOPATH)/bin/projectmars cmd/main.go 

uninstall:
	rm $(GOPATH)/bin/projectmars

test:
	go test -timeout 30s github.com/rc1140/next45/projectmars
