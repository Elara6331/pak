GOBUILD ?= go build

pak: main.go
	$(GOBUILD)

installbinonly: pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

aptinstall: install-config.sh pak.toml pak
	install -Dm644 pak.toml $(DESTDIR)/etc/pak.toml
	bash install-config.sh apt $(DESTDIR)

snapinstall: install-config.sh pak.toml pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak
	bash install-config.sh snap $(DESTDIR)

yayinstall: install-config.sh pak.toml pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak
	bash install-config.sh yay $(DESTDIR)

pacinstall: install-config.sh pak.toml pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak
	bash install-config.sh pacman $(DESTDIR)

aptitude: install-config.sh pak.toml pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak
	bash install-config.sh aptitude $(DESTDIR)

brewinstall: install-config.sh pak.toml pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak
	bash install-config.sh brew $(DESTDIR)

.PHONY: pak $(MAKECMDGOALS)
