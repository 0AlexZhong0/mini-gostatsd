run-docker:
	docker-compose up --build

clean-docker:
	docker-compose rm -f
