.PHONY: all test clean

NAME_IMAGE="rinha"

generate: 
	@echo rinha: Construindo a aplicacao...
	go build -o dist/rinha/main cmd/main.go 
	@echo rinha: Construindo a imagem
	@docker build -t betamedina/rinha .

publish:
	@echo rinha: Publicando a imagem
	@docker push betamedina/rinha

run:
	docker-compose -f scripts/docker-compose.yml up -d
	go run cmd/main.go

stop:
	docker-compose -f scripts/docker-compose.yml down

start:
	docker-compose -f scripts/docker-compose.yml up -d