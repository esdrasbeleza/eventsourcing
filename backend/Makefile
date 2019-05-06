.PHONY: database

build-docker-database:
	docker build -t postgres-eventsourcing:latest ./database

run-docker-database:
	 docker run --rm --name pg-docker -d -p 5432:5432 postgres-eventsourcing
