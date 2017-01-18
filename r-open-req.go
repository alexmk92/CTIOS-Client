package main

type OpenReq struct {
	Body

	InvokeID uint32
	VersionNumber uint32
	IdleTimeout uint32
	PeripheralID uint32
	ServicesRequested uint32
	CallMsgMask uint32
	AgentStateMask uint32
	ConfigMsgMask uint32
	Reserved1 uint32
	Reserved2 uint32
	Reserved3 uint32

	ClientID_Tag NamedVar
	ClientPassword_Tag NamedVar
}
