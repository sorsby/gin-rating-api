download:
	go mod download

lint:
	$(MAKE) revive

revive:
	revive -config defaults.toml ./...

run-dynamodb:
	docker kill dynamodb-gin-rating || true
	docker run --rm -d --name dynamodb-gin-rating -p 4569:4569 -p 4568:4568 -e PORT_WEB_UI=8888 -e SERVICES=dynamodb localstack/localstack