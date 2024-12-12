package interceptor

import (
	"context"
	"time"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// LogInterceptor logs requests and responses.
func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	now := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(), zap.String("method", info.FullMethod), zap.Any("req", req))
	}

	logger.Info("request", zap.String("method", info.FullMethod), zap.Any("req", req), zap.Any("res", res), zap.Duration("duration", time.Since(now)))

	return res, err
}
