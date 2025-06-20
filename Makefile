it: clean build

clean:
	rm -rf ./out || true
build:
	docker buildx bake binaries
