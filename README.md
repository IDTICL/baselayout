# BaseLayout

1. Edit line 14 in .air.toml file to add PG connection info.

```bigquery
	## binary will be $(go env GOPATH)/bin/air
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```