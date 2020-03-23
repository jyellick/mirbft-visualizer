package main

import (
	"bufio"
	"bytes"
	"fmt"
	"sync"

	"github.com/vugu/vugu"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type lockingWriterSync struct {
	eventEnv vugu.EventEnv
	buffer   *bytes.Buffer
	mutex    sync.Mutex
	cond     *sync.Cond
}

func newLockingWriterSync(e vugu.EventEnv) *lockingWriterSync {
	lws := &lockingWriterSync{
		eventEnv: e,
		buffer:   &bytes.Buffer{},
	}

	lws.cond = sync.NewCond(&lws.mutex)
	return lws
}

func (lws *lockingWriterSync) Write(p []byte) (int, error) {
	lws.mutex.Lock()
	defer lws.mutex.Unlock()
	defer lws.cond.Signal()
	return lws.buffer.Write(p)
}

func (lws *lockingWriterSync) drain() {
	for {
		lws.mutex.Lock()
		for lws.buffer.Len() == 0 {
			lws.cond.Wait()
		}
		pending := lws.buffer
		lws.buffer = &bytes.Buffer{}
		lws.mutex.Unlock()

		scanner := bufio.NewScanner(pending)

		lws.eventEnv.Lock()
		for scanner.Scan() {
			fmt.Printf("MirBFT Logger: %s\n", scanner.Bytes())
		}
		lws.eventEnv.UnlockOnly()
	}
}

func (lws *lockingWriterSync) Sync() {}

func wasmZap(eventEnv vugu.EventEnv) *zap.Logger {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	lockingWriteSync := newLockingWriterSync(eventEnv)
	go lockingWriteSync.drain()

	lockingWriteSyncer := zapcore.AddSync(lockingWriteSync)

	allPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	core := zapcore.NewCore(consoleEncoder, lockingWriteSyncer, allPriority)

	return zap.New(core)
}
