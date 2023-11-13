run:
	docker-compose up --build web-api

test:
	go test -v .\internal\repository\user\ .\internal\repository\revenue\ .\internal\middleware -coverprofile=all_test

view_test:
	go tool cover -html=all_test