build:
	mkdir -p bin
	mkdir -p bin/linux
	mkdir -p bin/win
	mkdir -p bin/osx
	GOARCH=amd64 GOOS=linux go build -o bin/linux/loch disinformationindex.org/loch/cmd/loch
	GOARCH=amd64 GOOS=darwin go build -o bin/osx/loch disinformationindex.org/loch/cmd/loch
	GOARCH=amd64 GOOS=windows go build -o bin/win/loch.exe disinformationindex.org/loch/cmd/loch
