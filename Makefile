.PHONY: run-srv run-db stop-db deploy-dev

run-srv: 
	CompileDaemon -build "go build -o bin/automa8e_clone" -command "./bin/automa8e_clone"

run-db:
	docker-compose up db -d;

stop-db:
	docker-compose stop db; docker-compose rm -f db;

deploy-dev:
	docker-compose up erp-be -d;