.PHONY: build
build/linux/%: ## Build server for linux
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/$*-linux ./cmd/$*/
	# env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./bin/$*-linux ./cmd/$*/

build/mac/%: ## Show build.sh help for building binnary package under cmd
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/$*/
