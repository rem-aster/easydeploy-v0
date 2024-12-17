from pathlib import Path
from dataclasses import dataclass
import os
from enum import StrEnum

@dataclass
class AnsibleRequest:
    playbook_name: str
    host: str
    port: int
    username: str
    password: str
    extra_vars: dict

    @property
    def playbook_path(self) -> Path:
        """Returns the full path to the playbook."""
        return Path(os.getcwd(), "playbooks", self.playbook_name)

class TaskStatus(StrEnum):
    running = "running"
    success = "success"
    failed = "failed"

@dataclass
class AnsibleResponse:
    status: TaskStatus
    output: str  # Playbook output
    error: str   # Error details, if any