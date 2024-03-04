.PHONY=publish up git-push setup docker-up dev test

publish: test git-push
up: setup docker-up

git-push:
	git push

setup:
	cp ./.env.redis.example ./.env.redis
	cp ./.env.example ./.env
	mkdir -p ./redis

docker-up:
	docker-compose up -d

dev:
	cd ./src/; go run .

test:
	cd ./src/; go test ./... -v