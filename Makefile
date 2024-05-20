
build:
	docker build -t containers -f Dockerfile . 
run:
	docker run -it containers

start: build run

build_tests:
	docker build -t containers_tests -f Dockerfile.test .

run_tests:
	docker run containers_tests

tests: build_tests run_tests
