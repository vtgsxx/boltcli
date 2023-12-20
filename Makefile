.PHONY: build compile upx clean all

GOCMD=go
GOBUILD=$(GOCMD) build
LDFLAGS=-ldflags="-w -s -buildid="
GCFLAGS=-gcflags=all="-l -B"
OUTFOLDER=bin
BUILDOPTIONS=-trimpath $(LDFLAGS) $(GCFLAGS)

GO_BINARY=boltcli

build:
	@$(GOBUILD) $(BUILDOPTIONS) -o $(OUTFOLDER)/$(GO_BINARY) .

compile:
	@GOOS=linux GOARCH=386 $(GOBUILD) $(BUILDOPTIONS) -o $(OUTFOLDER)/$(GO_BINARY)-linux-386 .
	@GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILDOPTIONS) -o $(OUTFOLDER)/$(GO_BINARY)-linux-amd64 .
	@GOOS=windows GOARCH=386 $(GOBUILD) $(BUILDOPTIONS) -o $(OUTFOLDER)/$(GO_BINARY)-windows-386.exe .
	@GOOS=windows GOARCH=amd64 $(GOBUILD) $(BUILDOPTIONS) -o $(OUTFOLDER)/$(GO_BINARY)-windows-amd64.exe .

upx:
	@upx -q9 $(OUTFOLDER)/$(GO_BINARY)*

clean:
	@rm -f $(OUTFOLDER)/$(GO_BINARY)*

all: build
