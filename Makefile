GOBUILD ?= go build

pak: main.go
	$(GOBUILD)
	
installbinonly: pak
	install -Dm755 pak $(DESTDIR)/usr/bin/pak
	mkdir -p $(DESTDIR)/etc/pak.d

aptinstall: plugins/apt/pak.cfg pak
	install -Dm644 plugins/apt/pak.cfg $(DESTDIR)/etc/pak.cfg
	mkdir -p $(DESTDIR)/etc/pak.d
	install -Dm644 plugins/snap/pak.cfg $(DESTDIR)/etc/pak.d/snap.cfg
	install -Dm644 plugins/yay/pak.cfg $(DESTDIR)/etc/pak.d/yay.cfg
	install -Dm644 plugins/zypper/pak.cfg $(DESTDIR)/etc/pak.d/zypper.cfg
	install -Dm644 plugins/pacman/pak.cfg $(DESTDIR)/etc/pak.d/pacman.cfg
	install -Dm644 plugins/brew/pak.cfg $(DESTDIR)/etc/pak.d/brew.cfg
	install -Dm644 plugins/aptitude/pak.cfg $(DESTDIR)/etc/pak.d/aptitude.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

snapinstall: plugins/snap/pak.cfg pak
	install -Dm644 plugins/snap/pak.cfg $(DESTDIR)/etc/pak.cfg
	mkdir -p $(DESTDIR)/etc/pak.d
	install -Dm644 plugins/apt/pak.cfg $(DESTDIR)/etc/pak.d/apt.cfg
	install -Dm644 plugins/yay/pak.cfg $(DESTDIR)/etc/pak.d/yay.cfg
	install -Dm644 plugins/zypper/pak.cfg $(DESTDIR)/etc/pak.d/zypper.cfg
	install -Dm644 plugins/pacman/pak.cfg $(DESTDIR)/etc/pak.d/pacman.cfg
	install -Dm644 plugins/brew/pak.cfg $(DESTDIR)/etc/pak.d/brew.cfg
	install -Dm644 plugins/aptitude/pak.cfg $(DESTDIR)/etc/pak.d/aptitude.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

yayinstall: plugins/yay/pak.cfg pak
	install -Dm644 plugins/yay/pak.cfg $(DESTDIR)/etc/pak.cfg
	mkdir -p $(DESTDIR)/etc/pak.d
	install -Dm644 plugins/snap/pak.cfg $(DESTDIR)/etc/pak.d/snap.cfg
	install -Dm644 plugins/apt/pak.cfg $(DESTDIR)/etc/pak.d/apt.cfg
	install -Dm644 plugins/zypper/pak.cfg $(DESTDIR)/etc/pak.d/zypper.cfg
	install -Dm644 plugins/pacman/pak.cfg $(DESTDIR)/etc/pak.d/pacman.cfg
	install -Dm644 plugins/brew/pak.cfg $(DESTDIR)/etc/pak.d/brew.cfg
	install -Dm644 plugins/aptitude/pak.cfg $(DESTDIR)/etc/pak.d/aptitude.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

pacinstall: plugins/pacman/pak.cfg pak
	install -Dm644 plugins/pacman/pak.cfg $(DESTDIR)/etc/pak.cfg
	mkdir -p $(DESTDIR)/etc/pak.d
	install -Dm644 plugins/snap/pak.cfg $(DESTDIR)/etc/pak.d/snap.cfg
	install -Dm644 plugins/yay/pak.cfg $(DESTDIR)/etc/pak.d/yay.cfg
	install -Dm644 plugins/zypper/pak.cfg $(DESTDIR)/etc/pak.d/zypper.cfg
	install -Dm644 plugins/apt/pak.cfg $(DESTDIR)/etc/pak.d/apt.cfg
	install -Dm644 plugins/brew/pak.cfg $(DESTDIR)/etc/pak.d/brew.cfg
	install -Dm644 plugins/aptitude/pak.cfg $(DESTDIR)/etc/pak.d/aptitude.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

aptitude: plugins/aptitude/pak.cfg pak
	install -Dm644 plugins/aptitude/pak.cfg $(DESTDIR)/etc/pak.cfg
	mkdir -p $(DESTDIR)/etc/pak.d
	install -Dm644 plugins/snap/pak.cfg $(DESTDIR)/etc/pak.d/snap.cfg
	install -Dm644 plugins/yay/pak.cfg $(DESTDIR)/etc/pak.d/yay.cfg
	install -Dm644 plugins/zypper/pak.cfg $(DESTDIR)/etc/pak.d/zypper.cfg
	install -Dm644 plugins/pacman/pak.cfg $(DESTDIR)/etc/pak.d/pacman.cfg
	install -Dm644 plugins/brew/pak.cfg $(DESTDIR)/etc/pak.d/brew.cfg
	install -Dm644 plugins/apt/pak.cfg $(DESTDIR)/etc/pak.d/apt.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

brewinstall: plugins/brew/pak.cfg pak
	install -Dm644 plugins/brew/pak.cfg $(DESTDIR)/etc/pak.cfg
	mkdir -p $(DESTDIR)/etc/pak.d
	install -Dm644 plugins/snap/pak.cfg $(DESTDIR)/etc/pak.d/snap.cfg
	install -Dm644 plugins/yay/pak.cfg $(DESTDIR)/etc/pak.d/yay.cfg
	install -Dm644 plugins/zypper/pak.cfg $(DESTDIR)/etc/pak.d/zypper.cfg
	install -Dm644 plugins/pacman/pak.cfg $(DESTDIR)/etc/pak.d/pacman.cfg
	install -Dm644 plugins/apt/pak.cfg $(DESTDIR)/etc/pak.d/apt.cfg
	install -Dm644 plugins/aptitude/pak.cfg $(DESTDIR)/etc/pak.d/aptitude.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

zyppinstall: plugins/zypper/pak.cfg pak
	install -Dm644 plugins/zypper/pak.cfg $(DESTDIR)/etc/pak.cfg
	mkdir -p $(DESTDIR)/etc/pak.d
	install -Dm644 plugins/snap/pak.cfg $(DESTDIR)/etc/pak.d/snap.cfg
	install -Dm644 plugins/yay/pak.cfg $(DESTDIR)/etc/pak.d/yay.cfg
	install -Dm644 plugins/apt/pak.cfg $(DESTDIR)/etc/pak.d/apt.cfg
	install -Dm644 plugins/pacman/pak.cfg $(DESTDIR)/etc/pak.d/pacman.cfg
	install -Dm644 plugins/brew/pak.cfg $(DESTDIR)/etc/pak.d/brew.cfg
	install -Dm644 plugins/aptitude/pak.cfg $(DESTDIR)/etc/pak.d/aptitude.cfg
	install -Dm755 pak $(DESTDIR)/usr/bin/pak

.PHONY: pak $(MAKECMDGOALS)
