download:
	go mod download

lint:
	$(MAKE) revive

revive:
	go get github.com/mgechev/revive
	revive -config defaults.toml ./...