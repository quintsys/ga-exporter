app = gaexporter
directory = ./api
dir_target = $(directory)-$(wildcard $(directory))
dir_present = $(directory)-$(directory)
dir_absent = $(directory)-

all: | $(dir_target) swag goget
	@echo && echo "Done!"

$(dir_present):
	@echo "Folder $(directory) exists. Deleting generated files."
	@rm -vrf "$(directory)/cmd"|| true && \
		rm -vrf "$(directory)/restapi/operations" || true && \
		rm -vf "$(directory)/restapi/doc.go" || true && \
		rm -vf "$(directory)/restapi/embedded_spec.go" || true && \
		rm -vf "$(directory)/restapi/server.go" || true;

$(dir_absent):
	@echo "Folder $(directory) does not exist"
	mkdir -p $(directory)

swag:
	@swagger generate server \
		--target $(directory) \
		--name $(app) \
		--spec $(directory)/swagger.yml \
		--exclude-main

goget:
	@echo && echo "Following above recommendation:"
	@echo "Existing modules will be updated to newer minor or patch releases."
	@echo
	go get -v -u ./...

deploy: build push release

build:	
	docker build --tag gaexporter .

push:
	heroku container:push web

release:
	heroku container:release web

.PHONY: all swag goget deploy build push release