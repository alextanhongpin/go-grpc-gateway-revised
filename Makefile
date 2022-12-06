.PHONY: gen

install:
	brew install bufbuild/buf/buf


gen: # Targets the buf.gen.yaml
	buf generate


server:
	go run server.go


client:
	go run client.go


call-api: # Make sure to run the `make server`, then `make client` before calling this.
	curl -XPOST localhost:8081/your.service.v1.YourService/Echo -d '{"value": "John"}'
