IMAGE = shorten
BINARY = cmd/shorten
TARGET = -o $(BINARY)
MAIN_FOLDER = ./cmd
BUILD_FLAGS = $(TARGET)

all:
	go build $(BUILD_FLAGS) $(MAIN_FOLDER)
run:
	./$(BINARY)
build:
	CGO_ENABLED=0 GOOS=linux go build $(BUILD_FLAGS) $(MAIN_FOLDER)
	docker build -t $(IMAGE):latest .
up:
	docker-compose up -d
down:
	docker-compose down
clean:
	rm -rf $(BINARY)
	docker rmi -f $(shell docker images -f "dangling=true" -q) 2> /dev/null; true
	docker rmi -f $(IMAGE):latest 2> /dev/null; true