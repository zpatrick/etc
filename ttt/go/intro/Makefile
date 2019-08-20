ADDR=localhost:8000

deps:
	go get golang.org/x/tools/cmd/present

present:
	present -http $(ADDR) | open "http://$(ADDR)"
