package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

const (
	/*
		Length (in bytes) of one element in the bridge table
	*/
	ENTRY_LENGTH = 16

	/*
		Root path of dir under the /sys filesystem
	*/
	SYS_ROOT = "/sys/class/net/"

	/*
		Directory name containing the bridge ports
	*/
	FILE_PORTS = "/brif"

	/*
		File name containing the bridge table
	*/
	FILE_MACS = "/brforward"
)

/*
	Entry in the MAC table
*/
type entry struct {
	/*
		Port name
	*/
	Port string

	/*
		MAC address
	*/
	Mac string

	/*
		Ageing time in seconds
	*/
	Age float64
}

var (
	/*
		Port index to port name mapping
	*/
	ports map[int]string

	/*
		MAC table entries
	*/
	macs []entry

	/*
		Length of the longest port name
	*/
	maxLength int

	version string
)

func main() {
	/*
		Initialise the global variables
	*/
	ports = make(map[int]string)

	fmt.Printf("Bridge MAC table, version %s\n\n", version)

	/*
		Validate the command-line arguments
	*/
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [ bridge name ]\n", os.Args[0])
		return
	}
	bridge := os.Args[1]

	/*
		Parse the ports of the bridge and prepare index to name map
	*/
	if err := parsePorts(os.Args[1]); err != nil {
		fmt.Printf("Failed to parse the port indexes for bridge %s!\n", bridge)
		fmt.Println(err)
		return
	}

	/*
		Parse the MACs of the bridge
	*/
	if err := parseTable(bridge); err != nil {
		fmt.Printf("Failed to parse the MAC table for bridge %s!\n", bridge)
		fmt.Println(err)
		return
	}

	/*
		Print the bridge MAC address table
	*/
	printTable()
}

/*
	Parse the ports of the bridge and prepare index to name map
*/
func parsePorts(n string) error {
	/*
		Read the port names
	*/
	f, err := ioutil.ReadDir(SYS_ROOT + n + FILE_PORTS)
	if err != nil {
		return err
	}

	/*
		Iterate over port names and update the maximum port name length if needed
	*/
	for k, v := range f {
		if len(v.Name()) > maxLength {
			maxLength = len(v.Name())
		}
		ports[k+1] = v.Name()
	}

	return nil
}

/*
	Parse the bridge table and sort by MACs
*/
func parseTable(n string) error {
	/*
		Read the bridge table
	*/
	data, err := ioutil.ReadFile(SYS_ROOT + n + FILE_MACS)
	if err != nil {
		return err
	}

	/*
		Iterate over the bridge table and try to parse the entries
	*/
	for i := 0; i < len(data)/ENTRY_LENGTH; i++ {
		parseEntry(data[ENTRY_LENGTH*i : ENTRY_LENGTH*(i+1)])
	}

	/*
		Sort the parsed table by MAC address
	*/
	sort.SliceStable(macs, func(a, b int) bool {
		return macs[a].Mac < macs[b].Mac
	})

	return nil
}

/*
	Parse one entry from the bridge table and
	add it to the MACs table if it is not a local
*/
func parseEntry(in []byte) {
	if l := int(in[7]); l == 0 {
		var m entry
		m.Port = ports[int(in[6])]
		m.Mac = parseMac(in[:6])
		m.Age = parseAge(in[8:10])
		macs = append(macs, m)
	}
}

/*
	Parse the MAC address from byte array
*/
func parseMac(m []byte) string {
	m1 := hex.EncodeToString(m[0:1])
	m2 := hex.EncodeToString(m[1:2])
	m3 := hex.EncodeToString(m[2:3])
	m4 := hex.EncodeToString(m[3:4])
	m5 := hex.EncodeToString(m[4:5])
	m6 := hex.EncodeToString(m[5:6])
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s", m1, m2, m3, m4, m5, m6)
}

/*
	Parse the MAC address age from byte array
*/
func parseAge(a []byte) float64 {
	return float64(int(a[1])*256+int(a[0])) / 100
}

/*
	Print the MAC address table
*/
func printTable() {
	/*
		Prepare the table format depending on the length of the longest port name
	*/
	format := fmt.Sprintf("%%%ds\t%%s\t%%.2f\n", maxLength)

	/*
		Print the table
	*/
	for _, v := range macs {
		fmt.Printf(format, v.Port, v.Mac, v.Age)
	}
}
