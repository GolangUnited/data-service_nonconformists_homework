ALL:
	@echo "You can use protogen, run, start, stop or shell parameters"
	
protogen: 	
	protoc -I api/proto --go_out=. --go-grpc_out=. api/proto/homework.proto

run:
	go run ./cmd/main.go

start: stop
	docker build -t golang-united-homework .
	docker run --name golang-united-homework --network my-network -p 8080:8080 -d -e HOMEWORK_DB_HOST=postgres -e HOMEWORK_DB_PORT -e HOMEWORK_DB_USER -e HOMEWORK_DB_PASSWORD -e HOMEWORK_DB_DATABASE golang-united-homework

stop:
	docker stop golang-united-homework || true
	docker rm golang-united-homework || true
	docker rmi golang-united-homework || true	

shell:
	@docker exec -it golang-united-homework /bin/sh