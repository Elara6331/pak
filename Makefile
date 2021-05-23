GOBUILD ?= go build

all: main.go
	$(GOBUILD)

install: PAK_CFG_MGR ?= apt
install: pak.toml pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak
	sed -i 's/activeManager = "\$PKGMANAGER"/activeManager = "$(PAK_CFG_MGR)"/' pak.toml
	install -Dm644 pak.toml $(DESTDIR)/etc/pak.toml

installbinonly: pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

.PHONY: all install installbinonly
