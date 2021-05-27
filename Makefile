GO = $(shell which go)
OUTPUTS = $(shell pwd)/build/function
SOURCE = $(shell pwd)/cmd/function

build-all:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(OUTPUTS)/user $(SOURCE)/user/main.go

build-all-j:
	GOOS=linux GOARCH=amd64 $(GO) build -tags=go_json -o $(OUTPUTS)/user $(SOURCE)/user/main.go

package-all:
	zip -D -j -r $(OUTPUTS)/user.zip $(OUTPUTS)/user

clean:
	rm -rf $(OUTPUTS)

update:
	go mod tidy
	go get -u github.com/gin-gonic/gin
	#chmod a+x update.sh
	./update.sh