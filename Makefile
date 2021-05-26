GO = $(shell which go)
OUTPUTS = $(shell pwd)/build/function

build-all:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(OUTPUTS)/user cmd/function/user/user.go

package-all:
	zip -D -j -r $(OUTPUTS)/user.zip $(OUTPUTS)/user

clean:
	rm -rf $(OUTPUTS)

update:
	go mod tidy
	go get -u github.com/gin-gonic/gin
	#chmod a+x update.sh
	./update.sh