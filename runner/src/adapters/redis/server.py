import redis
import json

class RedisServer:
    def __init__(self, host='localhost', port=6379, db=0, password=None):
        """
        Initializes the Redis client wrapper.

        Args:
            host (str): Redis server hostname.
            port (int): Redis server port.
            db (int): Redis database number.
            password (str, optional): Password for Redis authentication.
        """
        self.__pool = redis.ConnectionPool(host=host, port=port, db=db, password=password)
        self.client = redis.Redis(connection_pool=self.__pool)
