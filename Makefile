GO = $(shell which go)
OUTPUTS = $(shell pwd)/build/function

build-all:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(OUTPUTS)/users cmd/function/user.go

package-all:
	zip -D -j -r $(OUTPUTS)/user.zip $(OUTPUTS)/user

clean:
	rm -rf $(OUTPUTS)
