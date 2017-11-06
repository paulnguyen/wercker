
all: clean

clean: 
	find . -name 'main' -type f -exec rm -f {} \; 
	go clean

format:
	go fmt main

install:
	go install main

build:
	go build main

test:
	go test main

start:
	./main 

test-ping:
	curl localhost:3000/ping

test-gumball:
	curl localhost:3000/gumball

docker-build: 
	docker build -t gumball .
	docker images

docker-run:
	docker run --name gumball -td -p 3000:3000 gumball
	docker ps

docker-network:
	docker network inspect host

docker-shell:
	docker exec -it gumball bash 

docker-clean:
	docker stop gumball
	docker rm gumball
	docker rmi gumball

docker-ip:
	docker-machine ip

heroku-tag:
	docker build -t gumball .
	docker tag gumball registry.heroku.com/pnguyen-goapi/web

heroku-push:
	heroku login 
	heroku container:login
	docker push registry.heroku.com/pnguyen-goapi/web
