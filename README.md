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

### Example output
```shell
# brctl showmacs br0
port no	mac addr		is local?	ageing timer
  1	00:0c:42:xx:xx:xx	no		   0.10
  2	52:54:00:e0:bd:82	no		   0.10
  1	e0:d5:5e:xx:xx:xx	yes		   0.00
  1	e0:d5:5e:xx:xx:xx	yes		   0.00
  2	fe:54:00:e0:bd:82	yes		   0.00
  2	fe:54:00:e0:bd:82	yes		   0.00
```
```shell
# mactable br0
Bridge MAC table, version 0.2

enp3s0.41	00:0c:42:xx:xx:xx	0.13
   vnet16	52:54:00:e0:bd:82	0.14
```
