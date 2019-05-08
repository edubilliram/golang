package main

import "fmt"

type dispatcher interface {
	prints()
	set(string, string)
}
type ping struct {
	ip, div string
}

func (p *ping) prints() {
	fmt.Printf("ping:{ip :%v, div: %v}\n", p.ip, p.div)
}

/*func (p *ping) set(ip, div string) {
	p.ip = ip
	p.div = div
}*/
func entity(d dispatcher, hk string) {
	ip, div := hk+"newIp", hk+"newDiv"
	d.set(ip, div)
}

func Print(d dispatcher) {
	d.prints()
}

func main() {
	var d dispatcher
	var p = ping{"123423564", "ny"}
	d = &p

	hk := "1353456547"
	entity(d, hk)
	Print(d)

}
