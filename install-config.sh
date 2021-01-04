sed "s/activeManager = \"\"/activeManager = \"$1\"/" pak.toml > pak-new.toml
install -Dm644 pak-new.toml "$2"/etc/pak.toml