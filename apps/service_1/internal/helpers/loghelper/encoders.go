package loghelper

import "go.uber.org/zap/zapcore"

func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	s := "[" + caller.String() + "]"
	enc.AppendString(s)
}
