# Pak
#### Wrapper written in Go designed for package managers to unify software management commands between distros
[![Build status](https://ci.appveyor.com/api/projects/status/e4yacqd78gkte8a0?svg=true)](https://ci.appveyor.com/project/moussaelianarsen/pak)
[![Download Binary](https://img.shields.io/static/v1.svg?label=download&message=binary&color=blue)](https://minio.arsenm.dev/minio/pak/)

---

## Installation
###### APT Installation
- Install package "go" (You may remove after installation)
- Clone or download this repository
- Inside the repository, run:
```bash
make
sudo make aptinstall
```
###### Brew Installation

- Install package "go" (You may remove after installation)
- Clone or download this repository
- Inside the repository, run:
```bash
make
sudo make snapinstall
```

###### Snap Installation

- Install package "go" (You may remove after installation)
- Clone or download this repository
- Inside the repository, run:
```bash
make
sudo make snapinstall
```

###### Pacman Installation

- Install AUR package "pak" and choose pak-config-pacman when prompted

###### Yay Installation

- Install AUR package "pak" and choose pak-config-yay when prompted

###### Aptitude Installation

- Install package "go" (You may remove after installation)
- Clone or download this repository
- Inside the repository, run:
```bash
make
sudo make aptitude
```