upBuild:
	make clean
	docker-compose up --build

up:
	docker-compose up 

build:
	make clean
	docker-compose build
	make rm
	make prune

clean:
	make down
	make rm
	make prune

down:
	docker-compose down -v

prune:
	docker image prune -f

rm:
	docker rm $$(docker ps -aq)