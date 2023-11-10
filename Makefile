run:
	docker-compose up --build web-api

test:
	go test -v .\internal\repository\user\ .\internal\repository\revenue\ -coverprofile=all_test

view_test:
	go tool cover -html=all_test