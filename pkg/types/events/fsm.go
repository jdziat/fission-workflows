package events

import (
	"github.com/fission/fission-workflow/pkg/types"
	"github.com/fission/fission-workflow/pkg/util/fsm"
)

var WorkflowInvocationFsm *fsm.Fsm = fsm.New(
	types.WorkflowInvocationStatus_UNKNOWN,
	[]fsm.Transition{
		{
			Event: Invocation_INVOCATION_CREATED,
			Src:   types.WorkflowInvocationStatus_UNKNOWN,
			Dst:   types.WorkflowInvocationStatus_SCHEDULED,
		},
		{
			Event: Invocation_INVOCATION_CANCELED,
			Src:   types.WorkflowInvocationStatus_SCHEDULED,
			Dst:   types.WorkflowInvocationStatus_ABORTED,
		},
		{
			Event: Invocation_INVOCATION_CANCELED,
			Src:   types.WorkflowInvocationStatus_IN_PROGRESS,
			Dst:   types.WorkflowInvocationStatus_ABORTED,
		},
		{
			Event: Invocation_INVOCATION_COMPLETED,
			Src:   types.WorkflowInvocationStatus_IN_PROGRESS,
			Dst:   types.WorkflowInvocationStatus_SUCCEEDED,
		},
		{
			Event: Invocation_INVOCATION_COMPLETED,
			Src:   types.WorkflowInvocationStatus_IN_PROGRESS,
			Dst:   types.WorkflowStatus_PARSING,
		},
	},
)