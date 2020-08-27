module github.com/vugu-examples/simple

go 1.14

replace github.com/IBM/mirbft => github.com/jyellick/mirbft v0.0.0-20200827044224-a07ea68efd60

require (
	github.com/IBM/mirbft v0.0.0-20200820193629-05a8c61dd0f9
	github.com/golang/protobuf v1.4.2
	github.com/vugu/vjson v0.0.0-20200505061711-f9cbed27d3d9
	github.com/vugu/vugu v0.3.2
	go.uber.org/zap v1.9.1
)
