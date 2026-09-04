package resilience

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/MamangRust/monolith-graphql-pointofsale-pkg/logger"
)

// noopLogger implements logger.LoggerInterface without emitting anything,
// keeping unit tests side-effect free.
type noopLogger struct{}

func (noopLogger) Info(string, ...zap.Field)  {}
func (noopLogger) Fatal(string, ...zap.Field) {}
func (noopLogger) Debug(string, ...zap.Field) {}
func (noopLogger) Error(string, ...zap.Field) {}
func (noopLogger) Warn(string, ...zap.Field)  {}
func (noopLogger) Check(zapcore.Level, string) *zapcore.CheckedEntry {
	return nil
}
func (noopLogger) With(...zap.Field) logger.LoggerInterface { return noopLogger{} }
func (noopLogger) Sync() error                              { return nil }
