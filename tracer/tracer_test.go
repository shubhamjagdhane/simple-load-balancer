package tracer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/trace"

	"github.com/shubhamjagdhane/simple-load-balancer/logger"
)

func TestNew(t *testing.T) {
	var tLogger logger.Logger

	tracerImpl := New(false, "test", "test", tLogger)
	assert.NotNil(t, tracerImpl)

}

func TestSetSpanAttributes(t *testing.T) {
	type args struct {
		span  trace.Span
		input map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "OK",
			args: args{
				span:  trace.SpanFromContext(context.TODO()),
				input: map[string]string{"test": "test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				SetSpanAttributes(tt.args.span, tt.args.input)
			},
		)
	}
}

func TestShutdown(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				Shutdown()
			},
		)
	}
}

func TestStartSpan(t *testing.T) {
	tCtx := context.Background()
	ctx, span := StartSpan(tCtx, "test")
	assert.NotNil(t, ctx)
	assert.NotNil(t, span)
}
