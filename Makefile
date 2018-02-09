all: build

build:
	@go build -v ./cmd/dotenv

clean:
	@rm dotenv
