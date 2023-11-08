run:
	docker-compose up --build web-api

test:
	go test -v .\internal\repository\user\