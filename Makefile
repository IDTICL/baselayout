GO = $(shell which go)
OUTPUTS = $(shell pwd)/build/function

build-all:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(OUTPUTS)/users cmd/function/user/user.go

package-all:
	zip -D -j -r $(OUTPUTS)/user.zip $(OUTPUTS)/user

clean:
	rm -rf $(OUTPUTS)

go_mod_update:
	go mod tidy
	## Update gin
	go get -u github.com/gin-gonic/gin
	## binary will be $(go env GOPATH)/bin/air
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin