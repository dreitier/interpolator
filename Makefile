IMAGE=dreitier/interpolator
EXECUTABLE=interpolator

all: build

build:
	go mod download
	go build -o $(EXECUTABLE)
	strip $(EXECUTABLE)

clean:
	go clean
	rm -rf vendor $(EXECUTABLE)

docker-build:
	docker build --no-cache -t ${IMAGE}:${TAG} .

docker-push: docker-build
	docker tag ${IMAGE}:${TAG} ${IMAGE}:latest
	docker push ${IMAGE}:${TAG}
	docker push ${IMAGE}:latest
