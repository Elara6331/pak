GOBUILD ?= go build

all: main.go
	$(GOBUILD)

install: PAK_CFG_MGR ?= apt
install: pak.toml pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak
	sed 's/activeManager = ""/activeManager = "$(PAK_CFG_MGR)"/' pak.toml > pak-new.toml
	install -Dm644 pak-new.toml $(DESTDIR)/etc/pak.toml

installbinonly: pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

.PHONY: all install installbinonly
