up:
	docker pull postgres
	docker run -d --name simbirGo -p 5432:5432 -e POSTGRES_PASSWORD=password postgres
down:
	docker stop simbirGo

run:
	docker start simbirGo
	go run cmd/api/main.go