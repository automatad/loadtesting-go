VERSION:=$(shell date +%Y%m%d%H%M%S)

LOADTESTING_GO_STAGING_TAG=registry.digitalocean.com/automatad/bidder/loadingtesting-go:staging_$(VERSION)

.PHONY: all staging prod
all: staging
staging: staging_loadtesting_go
prod: 

staging_loadtesting_go:
	docker build -t $(LOADTESTING_GO_STAGING_TAG) -f ./Dockerfile .
	docker push $(LOADTESTING_GO_STAGING_TAG)
	sed s%IMAGE_TAG_PLACEHOLDER%$(LOADTESTING_GO_STAGING_TAG)% ./deployment/deployment.yml | kubectl apply -f - --record -n meru
