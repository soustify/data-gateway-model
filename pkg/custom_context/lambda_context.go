package custom_context

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"
)

type (
	Context struct {
		typeContext string
		fields      []field
	}

	field struct {
		Key   string
		Value string
	}
)

func Init(typeContext string) *Context {
	return &Context{
		typeContext: typeContext,
		fields:      make([]field, 0),
	}
}

func (c *Context) WithField(key, value string) *Context {
	c.fields = append(c.fields, field{
		Key:   key,
		Value: value,
	})
	return c
}

func (g *Context) Context(timeout time.Duration) (context.Context, context.CancelFunc) {
	if strings.ToUpper(g.typeContext) == "LAMBDA" {
		return g.lambdaContext(timeout)
	}

	if strings.ToUpper(g.typeContext) == "GRPC" {
		return g.grpcContext(timeout)
	}
	panic(fmt.Sprintf("invalid type context: %s", g.typeContext))
}

// Context para gRPC cria o contexto com metadata
func (g *Context) grpcContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	var pairs []string
	for _, val := range g.fields {
		pairs = append(pairs, val.Key, val.Value)
	}
	md := metadata.Pairs(pairs...)
	return context.WithTimeout(metadata.NewOutgoingContext(context.Background(), md), timeout)
}

// Context para Lambda cria o contexto com valores no contexto padrão
func (g *Context) lambdaContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx := context.Background()
	for _, val := range g.fields {
		ctx = context.WithValue(ctx, val.Key, val.Value) // Usando ContextKey como chave
	}
	return ctx, nil
}
