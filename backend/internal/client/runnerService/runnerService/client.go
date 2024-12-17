package runnerService

import (
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/client/runnerService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client структура клиента, реализующая интерфейсa runnerService.RunnerService
type Client struct {
	conn *grpc.ClientConn
}

// NewClient создает нового клиента, подключаясь к gRPC серверу
func NewClient(address string) (runnerService.RunnerService, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

// Close закрывает соединение с gRPC сервером
func (c *Client) Close() {
	_ = c.conn.Close()
}
