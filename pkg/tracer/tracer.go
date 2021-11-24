package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	//Configuration 为jaeger client配置项，主要设置应用的基本信息
	//Sampler (固定采样，对所有数据都进行采样)，Report（是否启用LoggerReport，刷新缓冲区的频率，上报的Agent地址）等
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	//初始化Tracer对象
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	//设置全局的Tracer对象
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
