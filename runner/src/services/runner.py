from services.service import AnsibleService
from proto.runner_pb2_grpc import PlaybookServiceV1
from proto.runner_pb2 import RunPlaybookResponse, RunPlaybookRequest, PlaybookStatusRequest, PlaybookStatusResponse
from models.models import AnsibleRequest, TaskStatus
from google.protobuf.json_format import MessageToDict
from repo.repo import AnsibleRepo
from uuid import uuid4, UUID
import asyncio


class RunnerServiceV1(AnsibleService, PlaybookServiceV1):
    
    def __init__(self, repo: AnsibleRepo):
        super().__init__(repo)
        super(AnsibleService, self).__init__()
    async def RunPlaybook(self, request: RunPlaybookRequest, context) -> RunPlaybookResponse:
        data = MessageToDict(request)
        id = uuid4()
        pb_name = data.get('playbookName')
        if pb_name is None:
            return RunPlaybookResponse(id=str(id), status=TaskStatus.failed)
        try:
            ssh_addr = data.get('sshAddress').split("@")
            username = ssh_addr[0]
            host, port = ssh_addr[1].split(":")
            passw = data.get('sshPassword')
        except Exception:
            return RunPlaybookResponse(id=str(id), status=TaskStatus.failed)
        extra = data.get('extraVars')
        req = AnsibleRequest(pb_name, host, port, username, passw, extra)
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
