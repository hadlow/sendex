# If the first argument is "run"
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

build:
	@go build -o bin/sendex

run: build
	@./bin/sendex $(RUN_ARGS)

test:
	@go test -v ./...
