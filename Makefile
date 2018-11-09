all: build

build:
	@go build -v ./cmd/dotenv

install:
	@go install -v ./cmd/dotenv

clean:
	@rm dotenv
