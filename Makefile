BINARY_NAME=go-crud
MAIN_FILE=./cmd/main.go
GOARCH=${_GOARCH}
# || "arm64 # amd64, 386, arm64, arm
GOOS=${_GOOS}
# || darwin # linux, windows

help:
	@echo "Select a command to run from the makefile"
	@echo "\t build -> build the binary"
	@echo "\t build-prod -> build the binary with tidy"
	@echo "\t run -> run the binary"
	@echo "\t watch -> run the binary with CompileDaemon"
	@echo "\t clean -> clean the binary"
	@echo "\t ready -> get all the dependencies"

tidy:
	go mod tidy

build:
	GOARCH=${GOARCH} GOOS=${GOOS} go build -o ./bin/${BINARY_NAME}-${GOOS} ${MAIN_FILE}

build-prod: tidy build

run: build
	./bin/${BINARY_NAME}-${GOOS}

watch:
	@CompileDaemon -polling -color="true" -directory="." -graceful-kill="true" -build="make build" -command="make run"|| echo "CompileDaemon not found"

clean:
	go clean
	@[[ -s ./bin/${BINARY_NAME}-${GOOS} ]] && rm ./bin/${BINARY_NAME}-${GOOS} || echo "Binary not found, moving on..."

# keeping this list, until I use a dependency manager like godep
ready:
	go get -u github.com/gin-gonic/gin
	go get -u github.com/gin-contrib/cors
	go get -u github.com/joho/godotenv
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite
	go get -u github.com/redis/go-redis/v9
	go get -u github.com/getsentry/sentry-go/gin
	go get -u github.com/rs/zerolog/log
	go get -u github.com/golang-jwt/jwt/v5
	go get- u github.com/gosimple/slug
	go get -u github.com/prometheus/client_golang/prometheus
	go get -u github.com/prometheus/client_golang/prometheus/promauto
	go get -u github.com/prometheus/client_golang/prometheus/promhttp

