module github.com/jyellick/mirbft-visualizer

go 1.13

replace github.com/IBM/mirbft => github.com/jyellick/mirbft v0.0.0-20200827044224-a07ea68efd60

replace github.com/vugu/vugu => /home/yellickj/oss/vugu

// require github.com/vugu/vugu v0.0.0-20200322002911-7d3a6c0162c2

require (
	github.com/IBM/mirbft v0.0.0-20190415184034-d1829758e6fc
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/golang/protobuf v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/vugu/vjson v0.0.0-20191111004939-722507e863cb
	github.com/vugu/vugu v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.14.0
)
