

gen:
	kx -t ignore_initialisms=false -m github.com/daidai21/thrift-go-map-gen   ./thrift_go_map_gen.thrift

run:
	go run main.go

