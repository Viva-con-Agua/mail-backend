.PHONY: db-dev

deploy = dev

up_db: $(object)
ifeq ($(deploy),prod) 
	docker-compose up -d db
else
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d db
endif

rm_db: $(object)
	docker-compose rm -f -s db

