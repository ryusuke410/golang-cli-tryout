# golang cli tryout

## tryout

```sh
go run cmd/cobratryout/main.go help
```

```sh
go run cmd/urfavetryout/main.go help
```

## gen mock

```sh
find internal -name "mocks_test.go" -type f -delete && go tool mockery
```

## coding style

- (golang) Don't return interface from constructor

## go test

### All tests in parallel

```sh
go test ./... -parallel=$(getconf _NPROCESSORS_ONLN)
```

### Individual test example

```sh
go test -timeout 30s -run ^TestNewMathMultiplyInput$ ryusuke410/golang-cli-tryout/internal/usecase
```

Note: The `-parallel=$(getconf _NPROCESSORS_ONLN)` flag runs tests in parallel using all available CPU cores.

## windsurf

```text
Initialize your Memory Bank with the "SessionStart" workflow.
```

```text
You need to initialize your Memory Bank with the "Windsurf Memory System"
```

```text
Remember to follow the Windsurf Memory System

createTaskLog()
```
