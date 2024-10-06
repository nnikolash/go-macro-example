default:
	@echo "Choose a target"
	@echo "  build-tools: Build the tools"
	@echo "  generate: Generate code with macro replacements"
	@echo "  run: Generate code and run it"
	@echo "  run-no-logging: Run the project with logging disabled at compile time"
	@echo "  run-no-macro: Run the project without macro replacements"

build-tools:
	@echo "***** Building tools *****"
	@echo
	go build -o bin/tools/include tools/include/*.go

generate:
	@echo
	@echo "***** Generating the project *****"
	@echo
	@go generate ./...

run: generate
	@echo
	@echo "***** Running the project *****"
	@echo
	@go run -tags=use_macro generated/*.go 

run-no-logging:
	@echo
	@echo "***** Running the project with logging disabled in compile time *****"
	@echo
	@GO_INCLUDE_DEFINES="DISABLE_LOGGING" go generate ./...
	@go run -tags=use_macro generated/*.go

run-no-macro:
	@echo
	@echo "***** Running the project without macro *****"
	@echo
	@go run *.go