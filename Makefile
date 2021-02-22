.PHONY: all build db execDB

build:
	docker-compose build --force-rm --no-cache

stage:
	docker push vivaconagua/mail-backend:stage

prod:
	docker push vivaconagua/mail-backend:stage

db:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d db

execDB:
	docker-compose exec mail-db mongo mail
