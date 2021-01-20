package tracer

import (
	"github.com/opentracing/opentracing-go"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

const (
	callbackTraceBefore = "opentracing:before"
	callbackTraceAfter  = "opentracing:after"
	gormSpanKey         = "__gorm_span"
)

func before(db *gorm.DB) {
	// 先从父级spans生成子span ---> 这里命名为gorm
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, "gorm")
	db.InstanceSet(gormSpanKey, span)
}

func after(db *gorm.DB) {
	// 从GORM的DB实例中取出span
	_span, isExist := db.InstanceGet(gormSpanKey)
	if !isExist {
		return
	}

	span, ok := _span.(opentracing.Span)
	if !ok {
		return
	}
	// <---- 一定一定一定要Finish掉！！！
	defer span.Finish()

	if db.Error != nil {
		span.LogFields(tracerLog.Error(db.Error))
	}

	// sql --> 写法来源GORM V2的日志
	span.LogFields(tracerLog.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
}

type OpenTracingPlugin struct {
}

func (op *OpenTracingPlugin) Name() string {
	return "opentracingPlugin"
}

func (op *OpenTracingPlugin) Initialize(db *gorm.DB) (err error) {
	//
	db.Callback().Create().Before("gorm:before_create").Register(callbackTraceBefore, before)
	db.Callback().Query().Before("gorm:query").Register(callbackTraceBefore, before)
	db.Callback().Delete().Before("gorm:before_delete").Register(callbackTraceBefore, before)
	db.Callback().Update().Before("gorm:setup_reflect_value").Register(callbackTraceBefore, before)
	db.Callback().Row().Before("gorm:row").Register(callbackTraceBefore, before)
	db.Callback().Raw().Before("gorm:raw").Register(callbackTraceBefore, before)

	// 结束后 - 并不是都用相同的方法，可以自己自定义
	db.Callback().Create().After("gorm:after_create").Register(callbackTraceAfter, after)
	db.Callback().Query().After("gorm:after_query").Register(callbackTraceAfter, after)
	db.Callback().Delete().After("gorm:after_delete").Register(callbackTraceAfter, after)
	db.Callback().Update().After("gorm:after_update").Register(callbackTraceAfter, after)
	db.Callback().Row().After("gorm:row").Register(callbackTraceAfter, after)
	db.Callback().Raw().After("gorm:raw").Register(callbackTraceAfter, after)
	return
}

var _ gorm.Plugin = &OpenTracingPlugin{}
