##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development
.DEFAULT_GOAL := image
build:	## Build binary
	go build -o dpll-cli main.go
test: ## Run unit tests
	go test -timeout 30s github.com/vitus133/go-dpll/pkg/dpll-ynl -v
clean:	## Remove build results
	rm dpll-cli
tidy:	## Clean dependencies
	go mod tidy && go mod vendor
run:	## Run (no build). This is using sudo and assuming go is installed in /usr/local/
	sudo /usr/local/go/bin/go run main.go
image: next ## build and push container image
	podman tag quay.io/vgrinber/tools:dpll-next quay.io/vgrinber/tools:dpll && podman push quay.io/vgrinber/tools:dpll
next: ## build and push container image for net-next
	podman build --secret id=rh_user,src=.rh_user --secret id=rh_pass,src=.rh_pass --arch=x86_64 -t quay.io/vgrinber/tools:dpll-next -f Containerfile . && podman push quay.io/vgrinber/tools:dpll-next
all: clean tidy build ## Clean, tidy, and build
