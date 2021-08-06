.PHONY: all build db execDB

build:
	docker-compose build --force-rm --no-cache

stage:
	docker push vivaconagua/mail-backend:stage

prod:
	docker tag vivaconagua/mail-backend:stage vivaconagua/mail-backend:latest
	docker push vivaconagua/mail-backend:latest
