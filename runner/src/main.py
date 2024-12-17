import logging
import signal
import os
import asyncio
from adapters.grpc.server import GRPCServer
from adapters.redis.server import RedisServer
from repo.repo import AnsibleRepo
from proto import runner_pb2_grpc
from services.playbook import Playbook


def signalHandler(sig, loop, server):
    logging.info("Shutting down server...")
    asyncio.create_task(shutdown_server(loop, server))

async def shutdown_server(loop, server):
    await server.stop()
    loop.stop()  # This stops the asyncio event loop

async def main():
    logging.basicConfig(level=logging.INFO)

    # Initialize the server
    server_gRPC = GRPCServer(address='0.0.0.0', port=50051)

    server_redis = RedisServer(os.environ.get("REDIS_ADDRESS"), int(os.environ.get("REDIS_PORT")), int(os.environ.get("REDIS_DB")), os.environ.get("REDIS_PASSWORD"))
    ansible_repo = AnsibleRepo(server_redis)
    server_gRPC.add_service(runner_pb2_grpc.add_PlaybookServiceServicer_to_server, Playbook(ansible_repo))
    
    # Handle graceful shutdown
    loop = asyncio.get_event_loop()
    loop.add_signal_handler(signal.SIGINT, lambda: signalHandler(signal.SIGINT, loop, server_gRPC))

    logging.info("Started")
    await server_gRPC.serve()

if __name__ == '__main__':
    asyncio.run(main())