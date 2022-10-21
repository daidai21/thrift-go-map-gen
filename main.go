package main

import (
	"fmt"
	mdl "github.com/daidai21/thrift-go-map-gen/kitex_gen/thrift_go_map_gen"
)

func main() {
	req := mdl.NewReq()
	k := mdl.UserKey{ID: 123, Gender: false}
	req.UsersScore = map[*mdl.UserKey]int32{}
	req.UsersScore[&k] = 123
	fmt.Println(req)
	fmt.Println(req.UsersScore[&k])

	k2 := mdl.UserKey{ID: 13123214, Gender: false}
	v, ok := req.UsersScore[&k2]
	fmt.Println(v, ok)
}
