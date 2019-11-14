download:
	go mod download

lint:
	$(MAKE) revive

revive:
	revive -config defaults.toml ./...