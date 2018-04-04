SOURCES := $(shell find . -name "*.go")
PACKAGE_PATH:=github.com/solo-io/gloo-storage

# kubernetes custom clientsets
clientset:
	cd ${GOPATH}/src/k8s.io/code-generator && \
	./generate-groups.sh all \
		$(PACKAGE_PATH)/crd/client \
		$(PACKAGE_PATH)/crd \
		"solo.io:v1"

clean:
	rm -rf $(PACKAGE_PATH)/crd/client
