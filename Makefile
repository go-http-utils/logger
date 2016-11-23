test:
	go test -v

cover:
	rm -rf *.coverprofile
	go test -coverprofile=logger.coverprofile
	gover
	go tool cover -html=logger.coverprofile