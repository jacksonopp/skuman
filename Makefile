.PHONY: run tw-watch run-server

run:
	make -j tw-watch run-server

tw-watch:
	./tw -i ./static/styles/input.css -o ./static/styles/style.css --watch

run-server:
	go run cmd/htmx-app/main.go