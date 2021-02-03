.PHONY: all build db

build:
	docker-compose build --force-rm --no-cache

db:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d db
