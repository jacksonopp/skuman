.PHONY: run tw-watch run-server

run:
	make dbup && make sqlc && make -j tw-watch run-server

tw-watch:
	./tw -i ./static/styles/input.css -o ./static/styles/style.css --watch

run-server:
	go run cmd/htmx-app/main.go

sqlc:
	sqlc generate

dcup:
	docker-compose up -d

dcdown:
	docker-compose down

dcrm:
	docker-compose rm -s

dbup:
	migrate -source "file://db/migrations" -database "postgresql://postgres:postgres@localhost:5438/skuman?sslmode=disable" -verbose up

dbdown:
	migrate -source "file://db/migrations" -database "postgresql://postgres:postgres@localhost:5438/skuman?sslmode=disable" -verbose down

dbdown-one:
	migrate -source "file://db/migrations" -database "postgresql://postgres:postgres@localhost:5438/skuman?sslmode=disable" -verbose down 1
