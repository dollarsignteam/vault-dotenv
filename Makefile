build:
	docker build -t local/vault-dotenv .

shell:
	docker run --rm -it local/vault-dotenv sh
