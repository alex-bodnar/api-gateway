.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY: clean vendor start-docker-compose

clean:
	cd Authenticator; go clean ./...
	cd Proxy; go clean ./...
	cd User; go clean ./...
	cd Authenticator; rm -r vendor
	cd Proxy; rm -r vendor
	cd User; rm -r vendor

vendor:
	cd Authenticator; go mod vendor
	cd Proxy; go mod vendor
	cd User; go mod vendor

start-docker-compose: vendor
	docker-compose up
