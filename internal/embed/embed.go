package embed

import _ "embed"

//go:embed teams.kdl
var TeamsKDL []byte

//go:embed instruction_leader.md
var InstructionLeader []byte

//go:embed instruction_member.md
var InstructionMember []byte
