.PHONY: build run test deploy

build:
	docker-compose build

run:
	docker-compose up

test:
	cd server && go test ./... -v
	cd client && yarn test

deploy:
	terraform -chdir=infra/terraform apply
	kubectl apply -f infra/kubernetes/

clean:
	docker-compose down
	rm -rf server/tmp client/.next

.PHONY: generate-server
generate-server:
	cd server && go generate ./...

.PHONY: lint
lint:
	cd server && golangci-lint run
	cd client && yarn lint

.PHONY: format
format:
	cd server && gofmt -s -w .
	cd client && yarn format