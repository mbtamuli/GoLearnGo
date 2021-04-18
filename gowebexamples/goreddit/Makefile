.PHONY: postgres adminer migrate

postgres:
	docker run --rm -it -p 5432:5432 --env POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -it -p 8080:8080 adminer

migrate:
	docker run --rm -it \
		-v ${PWD}/migrations:/migrations \
		migrate/migrate \
		-source file:///migrations \
		-database postgres://postgres:secret@host.docker.internal/postgres?sslmode=disable \
		up

migrate-down:
	docker run --rm -it \
		-v ${PWD}/migrations:/migrations \
		migrate/migrate \
		-source file:///migrations \
		-database postgres://postgres:secret@host.docker.internal/postgres?sslmode=disable \
		down