run:
	go run api/main.go

clean:
	@echo "Cleaning the dist directory \n"
	rm -rf dist/
	@echo "## Cleaned 'dist/' Successfully ##"


# Build a Binary for all major Operating Systems
build:
	@echo "Compiling for every OS and Platform"
	# 386 variants
	GOOS=freebsd GOARCH=386 go build -o dist/assignment-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o dist/assignment-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o dist/assignment-windows-386 main.go
	# amd64 variants
	GOOS=freebsd GOARCH=amd64 go build -o dist/assignment-freebsd-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o dist/assignment-darwin-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/assignment-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o dist/assignment-windows-amd64 main.go
	# arm64 variants
	GOOS=darwin GOARCH=amd64 go build -o dist/assignment-darwin-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/assignment-linux-arm64 main.go
	@echo "## Build completed successfully ##"