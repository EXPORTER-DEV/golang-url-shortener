.PHONY=up test

publish: test git-push

git-push:
	git push

up:
	cp ./.env.example ./.env
	docker-compose up -d

dev:
	cp ./.env.example ./.env
	go run ./src

test:
	go test ./... -v