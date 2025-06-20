it: clean build

clean:
	rm -rf ./binaries || true
build:
	docker buildx bake binaries
