package main

import (
	"github.com/IBM/mirbft/testengine"
	"github.com/vugu/vugu"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type lockingWriterSync struct {
	EventEnv vugu.EventEnv
}

func (lws *lockingWriterSync) Write(p []byte) (int, error) {
	// XXX if we try to acquire the lock here, but another
	// go routine is waiting on the state machine, then
	// we will deadlock.
	// lws.EventEnv.Lock()
	// fmt.Printf("MirBFT Logger: %s\n", string(p))
	// lws.EventEnv.UnlockOnly()
	return len(p), nil
}

func (lws *lockingWriterSync) Sync() {}

func wasmZap(eventEnv vugu.EventEnv) *zap.Logger {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	lockingWriteSyncer := zapcore.AddSync(&lockingWriterSync{
		EventEnv: eventEnv,
	})

	allPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	core := zapcore.NewCore(consoleEncoder, lockingWriteSyncer, allPriority)

	return zap.New(core)
}

type Network struct {
	Parameters *BootstrapParameters
	Recording  *testengine.Recording
}

func (n *Network) BeforeBuild() {
	if n.Recording != nil {
		return
	}

	logger := wasmZap(n.Parameters.EventEnv)

	recorder := testengine.BasicRecorder(n.Parameters.NodeCount, 0, 100)
	recorder.NetworkConfig.NumberOfBuckets = int32(n.Parameters.BucketCount)
	recorder.Logger = logger
	for _, clientConfig := range recorder.ClientConfigs {
		clientConfig.MaxInFlight = 1
	}

	recording, err := recorder.Recording()
	if err != nil {
		panic(err)
	}

	n.Recording = recording
}
