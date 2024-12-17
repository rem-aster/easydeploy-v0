from adapters.redis.server import RedisServer
from uuid import UUID
from models.models import TaskStatus
class AnsibleRepo:
    def __init__(self, redis: RedisServer):
        self.__server = redis
    
    async def add_task(self, id: UUID, status: TaskStatus) -> bool:
        n = self.__server.client.hset("task", str(id), str(status))
        n2 = self.__server.client.hset("task_error", str(id), "")
        if n == 1 and n2 == 1:
            return True
        return False

    async def get_task(self, id: UUID) -> tuple[TaskStatus, str]:
        status = self.__server.client.hget("task", str(id))
        error = self.__server.client.hget("task_error", str(id))
        if status:
            return TaskStatus(status.decode()), error
        return TaskStatus.failed, error
    
    async def update_task(self, id: UUID, status: TaskStatus, error: str="") -> bool:
        n = self.__server.client.hset("task", str(id), str(status))
        n2 = self.__server.client.hset("task_error", str(id), error)
        if n == 1 and n2 == 1:
            return True
        return False