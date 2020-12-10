package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

var (
	//只能输出结构化日志，但是性能要高于 SugaredLogger
	Logger *zap.Logger
	//可以输出 结构化日志、非结构化日志。性能低于 zap.Logger
	SugarLogger *zap.SugaredLogger
	// 通用配置
	LoggerEncoderConf zapcore.EncoderConfig
	// 核心配置
	Core zapcore.Core
)

func init() {
	LoggerEncoderConf = zapcore.EncoderConfig{
		MessageKey:   defaultMessageKey, // 结构化（json）输出：msg的key
		LevelKey:     defaultLevelKey,   // 结构化（json）输出：日志级别的key（INFO，WARN，ERROR等）
		TimeKey:      defaultTimeKey,    // 结构化（json）输出：时间的key（INFO，WARN，ERROR等）
		CallerKey:    defaultCallerKey,  // 结构化（json）输出：打印日志的文件对应的Key
		LineEnding:   zapcore.DefaultLineEnding,
		EncodeLevel:  encodeLevel,  // 将日志级别转换成大写（INFO，WARN，ERROR等）
		EncodeCaller: encodeCaller, // 采用短文件路径编码输出（test/main.go:14 ）
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(encodeTimeFormat))
		}, // 输出的时间格式
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}

	// 实现多个输出
	Core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(LoggerEncoderConf), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel),
	)
	Logger = zap.New(Core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	SugarLogger = Logger.Sugar()
}

func SetEncodeLevel(encoder uint8) {
	switch encoder {
	case 0:
		encodeLevel = zapcore.LowercaseLevelEncoder
	case 1:
		encodeLevel = zapcore.LowercaseColorLevelEncoder
	case 2:
		encodeLevel = zapcore.CapitalLevelEncoder
	case 3:
		encodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		encodeLevel = zapcore.CapitalLevelEncoder
	}
	LoggerEncoderConf.EncodeLevel = encodeLevel
}

func SetEncodeCaller(caller uint8) {
	switch caller {
	case 0:
		encodeCaller = zapcore.ShortCallerEncoder
	case 1:
		encodeCaller = zapcore.FullCallerEncoder
	default:
		encodeCaller = zapcore.ShortCallerEncoder
	}
	LoggerEncoderConf.EncodeCaller = encodeCaller
}

func SetEncodeTime(format string) {
	LoggerEncoderConf.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(format))
	}
}

func SetLogLevel(level zapcore.Level, logPath ...string) {
	logLevel = level
	if len(logPath) > 0 {
		writer := getWriter(level, logPath[0])
		Core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewConsoleEncoder(LoggerEncoderConf), zapcore.AddSync(writer), level),
			zapcore.NewCore(zapcore.NewJSONEncoder(LoggerEncoderConf), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), level),
		)
	} else {
		Core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewJSONEncoder(LoggerEncoderConf), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), level),
		)
	}

	Logger = zap.New(Core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	SugarLogger = Logger.Sugar()
}

func GetLogLevel() string {
	return logLevel.String()
}

func SetMultiLog(logs map[int]string) {
	var cores []zapcore.Core
	for logLevel, logFile := range logs {
		writer := getWriter(zapcore.Level(logLevel), logFile)
		// 若不写文件，则直接 os.Stdout 输出
		if logFile == "" && len(logFile) == 0 {
			writer = os.Stdout
		}
		core := zapcore.NewCore(zapcore.NewConsoleEncoder(LoggerEncoderConf), zapcore.AddSync(writer), getLevelEnablerFunc(zapcore.Level(logLevel)))
		cores = append(cores, core)
	}
	cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(LoggerEncoderConf), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel))
	// 实现多个输出
	Core = zapcore.NewTee(cores...)
	Logger = zap.New(Core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	SugarLogger = Logger.Sugar()
}

func getLevelEnablerFunc(targetLv zapcore.Level) zapcore.LevelEnabler {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == targetLv
	})
}

func getWriter(level zapcore.Level, filename string) io.Writer {
	var lumber *lumberjack.Logger

	lumber = &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    lumbMap[level.String()]["max_size"].(int),    // 最大M数，超过则切割
		MaxBackups: lumbMap[level.String()]["max_backups"].(int), // 最大文件保留数，超过就删除最老的日志文件
		MaxAge:     lumbMap[level.String()]["max_age"].(int),     // 保存30天
		Compress:   lumbMap[level.String()]["compress"].(bool),   // 是否压缩
	}

	if MaxSize > 0 {
		lumber.MaxSize = MaxSize
	}

	if MaxBackups > 0 {
		lumber.MaxBackups = MaxBackups
	}

	if MaxAge > 0 {
		lumber.MaxAge = MaxAge
	}

	if Compress {
		lumber.Compress = Compress
	}

	return lumber
}
