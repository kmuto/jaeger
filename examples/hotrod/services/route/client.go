// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package route

import (
	"context"
	"net/url"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/log"
	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/tracing"
	"github.com/jaegertracing/jaeger/examples/hotrod/services/config"
)

// Client is a remote client that implements route.Interface
type Client struct {
	logger   log.Factory
	client   *tracing.HTTPClient
	hostPort string
}

// NewClient creates a new route.Client
func NewClient(tracer trace.TracerProvider, logger log.Factory, hostPort string) *Client {
	return &Client{
		logger:   logger,
		client:   tracing.NewHTTPClient(tracer),
		hostPort: hostPort,
	}
}

// FindRoute implements route.Interface#FindRoute as an RPC
func (c *Client) FindRoute(ctx context.Context, pickup, dropoff string, vip bool) (*Route, error) {
	if config.MaxRouteLength < 9999 {
		return c.FindRoute_span(ctx, pickup, dropoff, vip)
	}

	c.logger.For(ctx).Info("Finding route", zap.String("pickup", pickup), zap.String("dropoff", dropoff))

	v := url.Values{}
	v.Set("pickup", pickup)
	v.Set("dropoff", dropoff)
	if vip {
		v.Set("vip", "yes")
	}
	url := "http://" + c.hostPort + "/route?" + v.Encode()
	var route Route
	if err := c.client.GetJSON(ctx, "/route", url, &route); err != nil {
		c.logger.For(ctx).Error("Error getting route", zap.Error(err))
		return nil, err
	}
	return &route, nil
}

func (c *Client) FindRoute_span(ctx context.Context, pickup, dropoff string, vip bool) (*Route, error) {
	tracer := c.client.TracerProvider.Tracer("route-client")
	ctx, span := tracer.Start(ctx, "FindRoute")
	defer span.End()

	c.logger.For(ctx).Info("Finding route", zap.String("pickup", pickup), zap.String("dropoff", dropoff))

	v := url.Values{}
	v.Set("pickup", pickup)
	v.Set("dropoff", dropoff)
	if vip {
		v.Set("vip", "yes")
	}
	url := "http://" + c.hostPort + "/route?" + v.Encode()
	var route Route
	if err := c.client.GetJSON(ctx, "/route", url, &route); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		c.logger.For(ctx).Error("Error getting route", zap.Error(err))
		return nil, err
	}
	return &route, nil
}
