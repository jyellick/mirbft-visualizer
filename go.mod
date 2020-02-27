module github.com/jyellick/mirbft-visualizer

go 1.13

replace github.com/IBM/mirbft => github.com/jyellick/mirbft v0.0.0-20200225211558-28f622ff11e5

replace github.com/vugu/vugu => github.com/jyellick/vugu v0.1.0-fork

require (
	github.com/IBM/mirbft v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/vugu/vugu v0.1.0
	go.uber.org/zap v1.14.0
)
