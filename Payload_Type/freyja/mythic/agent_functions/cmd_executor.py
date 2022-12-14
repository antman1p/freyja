from mythic_payloadtype_container.MythicCommandBase import *
import json
from mythic_payloadtype_container.MythicRPC import *


class CmdArguments(TaskArguments):
    def __init__(self, command_line, **kwargs):
        super().__init__(command_line, **kwargs)
        self.args = []

    async def parse_arguments(self):
        pass


class CmdCommand(CommandBase):
    cmd = "cmd_executor"
    needs_admin = False
    help_cmd = "cmd_executor {command}"
    description = "Execute a shell command using 'cmd /C'"
    version = 1
    author = "@antman1p"
    attackmapping = ["T1059", "T1059.003"]
    argument_class = CmdArguments
    attributes = CommandAttributes(
        suggested_command=True,
        supported_os=[SupportedOS.Windows]
    )

    async def create_tasking(self, task: MythicTask) -> MythicTask:
        resp = await MythicRPC().execute("create_artifact", task_id=task.id,
            artifact="cmd /C {}".format(task.args.command_line),
            artifact_type="Process Create",
        )
        resp = await MythicRPC().execute("create_artifact", task_id=task.id,
            artifact="{}".format(task.args.command_line),
            artifact_type="Process Create",
        )
        return task

    async def process_response(self, response: AgentResponse):
        pass
