package embed

import _ "embed"

//go:embed squad.kdl
var SquadKDL []byte

//go:embed instruction_leader.md
var InstructionLeader []byte

//go:embed instruction_member.md
var InstructionMember []byte
