package main

import "fmt"

type Resp struct {
	Success bool
	Error   error
}

var resp map[string]Resp

func test(a int) map[string]Resp {
	if a > 5 {
		resp = make(map[string]Resp)
		resp["yes"] = Resp.Success
		resp["no"] = Resp.Error
		Resp.Success = true
		Resp.Error = nil
		return resp
	}
	resp = make(map[string]Resp)
	resp["yes"] = Resp.Success
	resp["no"] = Resp.Error
	Resp.Success = false
	Resp.Error = fmt.Errorf("nah it's small dude: %v", a)

	return resp
}

func main() {

	b := 10

	rsp, err := test(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("resp:", rsp)
}
