from services.service import AnsibleService
from proto.runner_pb2_grpc import PlaybookService
from proto.runner_pb2 import RunPlaybookResponse, RunPlaybookRequest, PlaybookStatusRequest, PlaybookStatusResponse
from models.models import AnsibleRequest, TaskStatus
from google.protobuf.json_format import MessageToDict
from repo.repo import AnsibleRepo
from uuid import uuid4, UUID
import asyncio


class Playbook(AnsibleService, PlaybookService):
    
    def __init__(self, repo: AnsibleRepo):
        super().__init__(repo)
        super(AnsibleService, self).__init__()
    async def RunPlaybook(self, request: RunPlaybookRequest, context) -> RunPlaybookResponse:
        data = MessageToDict(request)
        pb_name = data.get('playbookName')
        if pb_name is None:
            return RunPlaybookResponse(id="", status=TaskStatus.failed)
        try:
            ssh_addr = data.get('sshAddress').split("@")
            username = ssh_addr[0]
            host, port = ssh_addr[1].split(":")
            passw = data.get('sshPassword')
        except Exception as e:
            return RunPlaybookResponse(id="", status=TaskStatus.failed)
        extra = data.get('extraVars')
        req = AnsibleRequest(pb_name, host, port, username, passw, extra)
        id = uuid4()
        asyncio.create_task(self.run_playbook(id, req))
        return RunPlaybookResponse(id=str(id), status=TaskStatus.running)
    
    async def GetTaskStatus(self, request: PlaybookStatusRequest, context) -> PlaybookStatusResponse:
        data = MessageToDict(request)
        id = data.get('id')
        if id is None:
            return PlaybookStatusResponse(status=TaskStatus.failed, error="ID is required")
        try:
            id = UUID(id)
            status, error = await self.get_task_status(id)
            return PlaybookStatusResponse(status=status, error=error)
        except Exception as e:
            return PlaybookStatusResponse(status=TaskStatus.failed, error=str(e))
