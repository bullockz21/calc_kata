# Имя бинарного файла
BINARY_NAME = calculator

# Каталог для сборки
BUILD_DIR = bin

# Цель по умолчанию
all: build

# Сборка проекта
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

# Запуск собранного бинарного файла
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Запуск тестов
test:
	go test ./...

# Очистка сгенерированных файлов
clean:
	rm -rf $(BUILD_DIR)

# Форматирование исходного кода
fmt:
	go fmt ./...

# Анализ кода (go vet)
vet:
	go vet ./...
