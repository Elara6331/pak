pak: pak.go
	go build pak.go

installbinonly: pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

aptinstall: pak.cfg pak
	install -Dm644 pak.cfg $(DESTDIR)/etc/pak.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

snapinstall: plugins/snap/pak.cfg pak
	install -Dm644 plugins/snap/pak.cfg $(DESTDIR)/etc/pak.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

yayinstall: plugins/yay/pak.cfg pak
	install -Dm644 plugins/yay/pak.cfg $(DESTDIR)/etc/pak.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

pacinstall: plugins/pacman/pak.cfg pak
	install -Dm644 plugins/pacman/pak.cfg $(DESTDIR)/etc/pak.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

aptitude: plugins/aptitude/pak.cfg pak
	install -Dm644 plugins/aptitude/pak.cfg $(DESTDIR)/etc/pak.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak
