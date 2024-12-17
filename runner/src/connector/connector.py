import asyncio
from models.models import AnsibleRequest, AnsibleResponse, TaskStatus
import ansible_runner
import re
import tempfile
from pathlib import Path
from uuid import UUID
import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Function to extract error lines
def extract_errors(stdout: str, stderr: str):
    error_lines = []

    # Check stdout for lines containing 'fatal', 'UNREACHABLE!', or 'FAILED!'
    if stdout:
        for line in stdout.splitlines():
            if re.search(r"(fatal:|UNREACHABLE!|FAILED!)", line):
                error_lines.append(line.strip())

    # Add stderr output if it exists
    if stderr:
        error_lines.append(stderr.strip())

    return "\n".join(error_lines)


class AnsibleConnector:
    async def run_playbook_async(self, id: UUID, model: AnsibleRequest) -> AnsibleResponse:

        playbook_path = Path(model.playbook_path)
        if not playbook_path.is_file():
            return AnsibleResponse(
                status=TaskStatus.failed,
                output="",
                error=f"Playbook path does not exist: {playbook_path}"
            )

        inventory = {
            "all": {
                "hosts": {
                    model.host: {
                        "ansible_user": model.username,
                        "ansible_ssh_pass": model.password,
                        "ansible_connection": "ssh",
                        "ansible_port": model.port
                    }
                }
            }
        }
    
        with tempfile.TemporaryDirectory() as temp_dir:
            try:
                result = await asyncio.to_thread(
                    ansible_runner.run,
                    private_data_dir=temp_dir,
                    playbook=str(playbook_path),
                    inventory=inventory,
                    extravars=model.extra_vars,
                )
                if result.rc == 0:
                    return AnsibleResponse(
                            status=TaskStatus.success,
                            output=result.stdout or "",
                            error=""
                            )
                else:
                    errors_only = extract_errors(result.stdout.read() or "", result.stderr.read() or "")
                    return AnsibleResponse(
                            status=TaskStatus.failed,
                            output=result.stdout.read() or "",
                            error=errors_only
                    )
            except Exception as e:
                logger.error(f"Error running task {str(id)}: {e}")
                return AnsibleResponse(
                    status=TaskStatus.failed,
                    output="",
                    error=str(e)
                )