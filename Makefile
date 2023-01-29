run:
	go run cmd/main.go

build:
	go build -o ad-api cmd/web/main.go

dbuild:
	docker image build -t api-image .

drun: 
	docker container run -p 9090:5432 -d --name api-container api-image

dstop:
	docker stop api-container

drm: 
	docker rm api-container

drim: 
	docker rmi api-image

dclear:
	docker system prune -a

dcomposebuild:
	docker compose build

dcomposeup:
	docker compose up

dcomposedown:
	docker compose down 

test:
	go test ./...