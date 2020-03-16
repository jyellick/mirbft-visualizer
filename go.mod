module github.com/jyellick/mirbft-visualizer

go 1.13

// replace github.com/IBM/mirbft => /home/yellickj/oss/mirbft
replace github.com/IBM/mirbft => github.com/jyellick/mirbft v0.0.0-20200313203515-1e074c88c08f

// replace github.com/vugu/vugu => /home/yellickj/oss/vugu

require github.com/vugu/vugu v0.0.0-20191230174203-fe17e20eb6ce

require (
	github.com/IBM/mirbft v0.0.0-20190415184034-d1829758e6fc
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/golang/protobuf v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/vugu/vjson v0.0.0-20191111004939-722507e863cb
	go.uber.org/zap v1.14.0
)
