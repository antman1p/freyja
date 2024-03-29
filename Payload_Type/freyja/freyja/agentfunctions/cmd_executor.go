package agentfunctions

import (
	agentstructs "github.com/MythicMeta/MythicContainer/agent_structs"
	"github.com/MythicMeta/MythicContainer/logging"
	"github.com/MythicMeta/MythicContainer/mythicrpc"
)

var cmd_executor = agentstructs.Command{
	Name:                "cmd_executor",
	Author:              "@antman1p",
	Description:         "Execute a shell command using 'cmd /C'",
	MitreAttackMappings: []string{"T1059"},
	CommandAttributes: agentstructs.CommandAttribute{
		SupportedOS: []string{agentstructs.SUPPORTED_OS_WINDOWS},
	},
	TaskFunctionCreateTasking: cmd_executorCreateTasking,
	Version:                   1,
}

func init() {
	agentstructs.AllPayloadData.Get("freyja").AddCommand(cmd_executor)
}

func cmd_executorCreateTasking(taskData *agentstructs.PTTaskMessageAllData) agentstructs.PTTaskCreateTaskingMessageResponse {
	response := agentstructs.PTTaskCreateTaskingMessageResponse{
		Success: true,
		TaskID:  taskData.Task.ID,
	}
	if _, err := mythicrpc.SendMythicRPCArtifactCreate(mythicrpc.MythicRPCArtifactCreateMessage{
		BaseArtifactType: "ProcessCreate",
		ArtifactMessage:  "cmd /C " + taskData.Args.GetCommandLine(),
		TaskID:           taskData.Task.ID,
	}); err != nil {
		logging.LogError(err, "Failed to send mythicrpc artifact create")
	}
	return response
}
