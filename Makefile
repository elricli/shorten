IMAGE = shorten
BINARY = cmd/shorten
TARGET = -o $(BINARY)
MAIN_FOLDER = ./cmd
BUILD_FLAGS = $(TARGET)

.PHONY: all
all: pull update restart

.PHONY: update
update: pull build

.PHONY: run
run:
	./$(BINARY)

.PHONY: pull
pull:
	git pull

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build $(BUILD_FLAGS) $(MAIN_FOLDER)
	docker build -t $(IMAGE):latest .

.PHONY: restart
restart:
	docker-compose restart

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: clean
clean:
	rm -rf $(BINARY)
	docker rmi -f $(shell docker images -f "dangling=true" -q) 2> /dev/null; true
	docker rmi -f $(IMAGE):latest 2> /dev/null; true