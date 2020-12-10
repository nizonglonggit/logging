package logging

import "go.uber.org/zap/zapcore"

var (
    // EncoderConfig Setting
    defaultMessageKey = "msg"
    defaultLevelKey   = "level"
    defaultTimeKey    = "time"
    defaultCallerKey  = "file"
)

var (
    // default logger Level
    logLevel = zapcore.DebugLevel
    // EncodeLevel Setting
    encodeLevel = zapcore.CapitalLevelEncoder
    // EncodeCaller Setting
    encodeCaller = zapcore.ShortCallerEncoder
    // EncodeTime Setting
    encodeTimeFormat = "2006-01-02 15:04:05.000000"
)

// lumberjack.Logger setting
var (
    MaxSize    = 0
    MaxBackups = 0
    MaxAge     = 0
    Compress   = false
)

// level const
const (
    DEBUGLevel = zapcore.DebugLevel
    INFOLevel  = zapcore.InfoLevel
    WARNLevel  = zapcore.WarnLevel
    ERRORLevel = zapcore.ErrorLevel
)

const (
    DEBUG = iota - 1
    INFO
    WARN
    ERROR
)

// EncodeLevel
const (
    LowercaseLevelEncoder int = iota
    LowercaseColorLevelEncoder
    CapitalLevelEncoder
    CapitalColorLevelEncoder
)

// EncodeCaller
const (
    ShortCallerEncoder int = iota
    FullCallerEncoder
)

// Lumberjacks
var (
    lumbDefault = map[string]interface{}{
        "max_size":    128,
        "max_backups": 100,
        "max_age":     31,
        "compress":    false,
    }
    lumbDebug = map[string]interface{}{
        "max_size":    128,
        "max_backups": 100,
        "max_age":     31,
        "compress":    false,
    }
    lumbInfo = map[string]interface{}{
        "max_size":    128,
        "max_backups": 100,
        "max_age":     31,
        "compress":    false,
    }
    lumbWarn = map[string]interface{}{
        "max_size":    128,
        "max_backups": 100,
        "max_age":     31,
        "compress":    false,
    }
    lumbError = map[string]interface{}{
        "max_size":    128,
        "max_backups": 100,
        "max_age":     31,
        "compress":    false,
    }
    lumbMap = map[string]map[string]interface{}{
        "debug": lumbDebug,
        "info":  lumbInfo,
        "warn":  lumbWarn,
        "error": lumbError,
    }
)
