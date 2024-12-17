from grpc import aio
from concurrent import futures
import logging
import signal


class GRPCServer:
    def __init__(self, address='[::]', port=50051, max_workers=10):
        self.__address = address
        self.__port = port
        self.__server = aio.server(futures.ThreadPoolExecutor(max_workers=10))
        self.__running = False

    @property
    def instance(self):
        """Expose the underlying gRPC server instance."""
        return self.__server

    def add_service(self, add_servicer_fn, servicer):
        add_servicer_fn(servicer, self.__server)

    async def serve(self):
        """Start the server and wait for termination."""
        endpoint = f'{self.__address}:{str(self.__port)}'
        self.__server.add_insecure_port(endpoint)
        self.__running = True

        logging.info(f"Starting gRPC Server at {endpoint}")
        await self.__server.start()

        # Handle graceful shutdown on signals
        self._setup_signal_handlers()

        try:
            logging.info("Server is running. Waiting for termination...")
            await self.__server.wait_for_termination()
        except KeyboardInterrupt:
            self.stop()

    async def stop(self, grace=3):
        """Stop the server gracefully."""
        if self.__running:
            logging.info("Stopping gRPC Server gracefully...")
            await self.__server.stop(grace)
            self.__running = False
            logging.info("Server stopped.")

    def _setup_signal_handlers(self):
        """Setup signal handlers for graceful shutdown."""
        def shutdown_signal_handler(signal_received, frame):
            self.stop()

        signal.signal(signal.SIGINT, shutdown_signal_handler)
        signal.signal(signal.SIGTERM, shutdown_signal_handler)
        