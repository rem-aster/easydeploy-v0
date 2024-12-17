from connector.connector import AnsibleConnector
from repo.repo import AnsibleRepo
from models.models import AnsibleRequest, TaskStatus
from uuid import UUID
import asyncio

class AnsibleService:
    def __init__(self, repo: AnsibleRepo) -> None:
        self.__connector = AnsibleConnector()
        self.__repo = repo
    async def run_playbook(self, id: UUID, request: AnsibleRequest) -> None:
        asyncio.create_task(self.__repo.add_task(id, TaskStatus.running))
        res = await self.__connector.run_playbook_async(id, request)
        asyncio.create_task(self.__repo.update_task(id, res.status, res.error))
        
    async def get_task_status(self, id: UUID) -> tuple[TaskStatus, str]:
        status, error = await self.__repo.get_task(id)
        return status, error