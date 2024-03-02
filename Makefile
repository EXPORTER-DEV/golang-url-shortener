.PHONY=up test

publish: test git-push

git-push:
	git push

setup:
	cp ./.env.example ./.env

up:
	cp ./.env.example ./.env
	docker-compose up -d

dev:
	go run ./src

test:
	go test ./... -v