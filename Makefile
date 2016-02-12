.PHONY: test build fmt install lint

CWD := $(shell pwd)
VERSION   := 0.0.1
RELEASE   := 1
TARDIR    := glusterrest-$(VERSION)
RPMBUILD  := $(HOME)/rpmbuild
# Set the Value if not set
PREFIX ?= /usr/local

devbuild:
	@echo "Doing $@"
	@GO15VENDOREXPERIMENT=1 go build -v -o ./bin/glusterrestd ./rest
	@GO15VENDOREXPERIMENT=1 go build -v -o ./bin/glustereventsd ./events

integration_test:
	py.test tests/

test:
	@echo "Doing $@"
	@GO15VENDOREXPERIMENT=1 go test $$(GO15VENDOREXPERIMENT=1 glide nv)

getdeps:
	@echo "Doing $@"
	@go get github.com/golang/lint/golint
	@go get github.com/Masterminds/glide

verifiers: getdeps vet fmt lint

vendor-update:
	@echo "Updating vendored packages"
	@GO15VENDOREXPERIMENT=1 glide -q up 2> /dev/null

docgen:
	@echo "Doing $@"
	python tools/docgen.py

build: verifiers vendor-update test docgen devbuild

hello:
	@echo A${PREFIX}A${DESTDIR}A

fmt:
	go fmt ./rest/...
	go fmt ./events/...
	go fmt ./glustercli/...
	go fmt ./grutil/...

install:
	@echo "Doing $@"
	install -d ${DESTDIR}${PREFIX}/sbin/
	install -d ${DESTDIR}/etc/glusterfs/
	install -d ${DESTDIR}${PREFIX}/lib/systemd/system/
	install -d ${DESTDIR}/var/log/glusterfs/
	install -m 755 ./bin/glusterrestd ${DESTDIR}${PREFIX}/sbin/glusterrestd
	install -m 755 ./bin/glustereventsd ${DESTDIR}${PREFIX}/sbin/glustereventsd
	install -m 0600 ./extra/glusterrest.json ${DESTDIR}/etc/glusterfs/glusterrest.json
	install -m 0644 ./extra/glusterrestd.service ${DESTDIR}${PREFIX}/lib/systemd/system/
	install -m 0644 ./extra/glustereventsd.service ${DESTDIR}${PREFIX}/lib/systemd/system/
	install -m 755 -d ${DESTDIR}/var/log/glusterfs/rest

lint:
	golint ./rest/...
	golint ./events/...
	golint ./glustercli/...
	golint ./grutil/...

vet:
	go vet ./rest/...
	go vet ./events/...
	go vet ./glustercli/...
	go vet ./grutil/...

dist:
	@echo "Doing $@"
	rm -fr ./dist
	mkdir -p ./dist/$(TARDIR)
	rsync -r --exclude .git/ --exclude dist/ --exclude bin/ $(CWD)/ ./dist/$(TARDIR)
	cd ./dist/; tar -zcf $(TARDIR).tar.gz $(TARDIR);

rpm: dist
	@echo "Doing $@"
	rm -rf $(RPMBUILD)/SOURCES/glusterrest*
	rm -rf $(RPMBUILD)/BUILD/glusterrest*
	mkdir -p $(RPMBUILD)/SOURCES
	cp ./dist/glusterrest-$(VERSION).tar.gz $(RPMBUILD)/SOURCES; \
	rpmbuild -ba glusterrest.spec
