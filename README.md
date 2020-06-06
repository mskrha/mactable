## mactable

### Description
Mactable is a tool used to show the MAC table of the specified Linux bridge, sorted by the MAC address and without the local ones. It can be used as an alternative to the ```brctl showmacs BRIDGE```.

### Build
```shell
git clone https://github.com/mskrha/mactable.git
cd mactable/source
make
```

### Build DEB package
```shell
git clone https://github.com/mskrha/mactable.git
cd mactable/source
make deb
```

### Usage
```shell
./mactable BRIDGE
```
