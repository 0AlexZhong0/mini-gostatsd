run-docker:
	docker-compose up --build

clean-docker:
	docker-compose stop
	docker-compose rm
