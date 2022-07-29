all: validate clean build

validate:
	swagger validate ./api/swagger.yml

build:
	swagger -q generate server -A hulhay-mall -f api/swagger.yml -s internal/apis -m internal/models
	go build -v -installsuffix cgo ./cmd/hulhay-mall-server

clean:
	rm -rf hulhay-mall-server
	go clean -i .

run: all
	./hulhay-mall-server --port=8080 --host=0.0.0.0