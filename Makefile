pak: pak.go
	go build pak.go

installbinonly: pak
	install -Dm755 pak $(DESTDIR)/usr/bin

aptinstall: pak.cfg pak
	install -Dm644 pak.cfg $(DESTDIR)/etc
	install -Dm755 pak $(DESTDIR)/usr/bin

snapinstall: plugins/snap/pak.cfg pak
	install -Dm644 plugins/snap/pak.cfg $(DESTDIR)/etc
	install -Dm755 pak $(DESTDIR)/usr/bin

yayinstall: plugins/yay/pak.cfg pak
	install -Dm644 plugins/yay/pak.cfg $(DESTDIR)/etc
	install -Dm755 pak $(DESTDIR)/usr/bin

pacinstall: plugins/pacman/pak.cfg pak
	install -Dm644 plugins/pacman/pak.cfg $(DESTDIR)/etc
	install -Dm755 pak $(DESTDIR)/usr/bin

aptitude: plugins/aptitude/pak.cfg pak
	install -Dm644 plugins/aptitude/pak.cfg $(DESTDIR)/etc
	install -Dm755 pak $(DESTDIR)/usr/bin