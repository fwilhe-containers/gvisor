// automatically generated by stateify.

package stack

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (t *tuple) StateTypeName() string {
	return "pkg/tcpip/stack.tuple"
}

func (t *tuple) StateFields() []string {
	return []string{
		"tupleEntry",
		"conn",
		"reply",
		"tupleID",
	}
}

func (t *tuple) beforeSave() {}

// +checklocksignore
func (t *tuple) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.tupleEntry)
	stateSinkObject.Save(1, &t.conn)
	stateSinkObject.Save(2, &t.reply)
	stateSinkObject.Save(3, &t.tupleID)
}

func (t *tuple) afterLoad() {}

// +checklocksignore
func (t *tuple) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.tupleEntry)
	stateSourceObject.Load(1, &t.conn)
	stateSourceObject.Load(2, &t.reply)
	stateSourceObject.Load(3, &t.tupleID)
}

func (ti *tupleID) StateTypeName() string {
	return "pkg/tcpip/stack.tupleID"
}

func (ti *tupleID) StateFields() []string {
	return []string{
		"srcAddr",
		"srcPort",
		"dstAddr",
		"dstPort",
		"transProto",
		"netProto",
	}
}

func (ti *tupleID) beforeSave() {}

// +checklocksignore
func (ti *tupleID) StateSave(stateSinkObject state.Sink) {
	ti.beforeSave()
	stateSinkObject.Save(0, &ti.srcAddr)
	stateSinkObject.Save(1, &ti.srcPort)
	stateSinkObject.Save(2, &ti.dstAddr)
	stateSinkObject.Save(3, &ti.dstPort)
	stateSinkObject.Save(4, &ti.transProto)
	stateSinkObject.Save(5, &ti.netProto)
}

func (ti *tupleID) afterLoad() {}

// +checklocksignore
func (ti *tupleID) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ti.srcAddr)
	stateSourceObject.Load(1, &ti.srcPort)
	stateSourceObject.Load(2, &ti.dstAddr)
	stateSourceObject.Load(3, &ti.dstPort)
	stateSourceObject.Load(4, &ti.transProto)
	stateSourceObject.Load(5, &ti.netProto)
}

func (cn *conn) StateTypeName() string {
	return "pkg/tcpip/stack.conn"
}

func (cn *conn) StateFields() []string {
	return []string{
		"ct",
		"original",
		"reply",
		"finalizeOnce",
		"finalizeResult",
		"sourceManip",
		"destinationManip",
		"tcb",
		"lastUsed",
	}
}

func (cn *conn) beforeSave() {}

// +checklocksignore
func (cn *conn) StateSave(stateSinkObject state.Sink) {
	cn.beforeSave()
	stateSinkObject.Save(0, &cn.ct)
	stateSinkObject.Save(1, &cn.original)
	stateSinkObject.Save(2, &cn.reply)
	stateSinkObject.Save(3, &cn.finalizeOnce)
	stateSinkObject.Save(4, &cn.finalizeResult)
	stateSinkObject.Save(5, &cn.sourceManip)
	stateSinkObject.Save(6, &cn.destinationManip)
	stateSinkObject.Save(7, &cn.tcb)
	stateSinkObject.Save(8, &cn.lastUsed)
}

func (cn *conn) afterLoad() {}

// +checklocksignore
func (cn *conn) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &cn.ct)
	stateSourceObject.Load(1, &cn.original)
	stateSourceObject.Load(2, &cn.reply)
	stateSourceObject.Load(3, &cn.finalizeOnce)
	stateSourceObject.Load(4, &cn.finalizeResult)
	stateSourceObject.Load(5, &cn.sourceManip)
	stateSourceObject.Load(6, &cn.destinationManip)
	stateSourceObject.Load(7, &cn.tcb)
	stateSourceObject.Load(8, &cn.lastUsed)
}

func (ct *ConnTrack) StateTypeName() string {
	return "pkg/tcpip/stack.ConnTrack"
}

func (ct *ConnTrack) StateFields() []string {
	return []string{
		"seed",
		"clock",
		"rand",
		"buckets",
	}
}

func (ct *ConnTrack) beforeSave() {}

// +checklocksignore
func (ct *ConnTrack) StateSave(stateSinkObject state.Sink) {
	ct.beforeSave()
	stateSinkObject.Save(0, &ct.seed)
	stateSinkObject.Save(1, &ct.clock)
	stateSinkObject.Save(2, &ct.rand)
	stateSinkObject.Save(3, &ct.buckets)
}

func (ct *ConnTrack) afterLoad() {}

// +checklocksignore
func (ct *ConnTrack) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ct.seed)
	stateSourceObject.Load(1, &ct.clock)
	stateSourceObject.Load(2, &ct.rand)
	stateSourceObject.Load(3, &ct.buckets)
}

func (bkt *bucket) StateTypeName() string {
	return "pkg/tcpip/stack.bucket"
}

func (bkt *bucket) StateFields() []string {
	return []string{
		"tuples",
	}
}

func (bkt *bucket) beforeSave() {}

// +checklocksignore
func (bkt *bucket) StateSave(stateSinkObject state.Sink) {
	bkt.beforeSave()
	stateSinkObject.Save(0, &bkt.tuples)
}

func (bkt *bucket) afterLoad() {}

// +checklocksignore
func (bkt *bucket) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &bkt.tuples)
}

func (it *IPTables) StateTypeName() string {
	return "pkg/tcpip/stack.IPTables"
}

func (it *IPTables) StateFields() []string {
	return []string{
		"connections",
		"reaperDone",
		"mu",
		"v4Tables",
		"v6Tables",
		"modified",
	}
}

// +checklocksignore
func (it *IPTables) StateSave(stateSinkObject state.Sink) {
	it.beforeSave()
	stateSinkObject.Save(0, &it.connections)
	stateSinkObject.Save(1, &it.reaperDone)
	stateSinkObject.Save(2, &it.mu)
	stateSinkObject.Save(3, &it.v4Tables)
	stateSinkObject.Save(4, &it.v6Tables)
	stateSinkObject.Save(5, &it.modified)
}

// +checklocksignore
func (it *IPTables) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &it.connections)
	stateSourceObject.Load(1, &it.reaperDone)
	stateSourceObject.Load(2, &it.mu)
	stateSourceObject.Load(3, &it.v4Tables)
	stateSourceObject.Load(4, &it.v6Tables)
	stateSourceObject.Load(5, &it.modified)
	stateSourceObject.AfterLoad(it.afterLoad)
}

func (table *Table) StateTypeName() string {
	return "pkg/tcpip/stack.Table"
}

func (table *Table) StateFields() []string {
	return []string{
		"Rules",
		"BuiltinChains",
		"Underflows",
	}
}

func (table *Table) beforeSave() {}

// +checklocksignore
func (table *Table) StateSave(stateSinkObject state.Sink) {
	table.beforeSave()
	stateSinkObject.Save(0, &table.Rules)
	stateSinkObject.Save(1, &table.BuiltinChains)
	stateSinkObject.Save(2, &table.Underflows)
}

func (table *Table) afterLoad() {}

// +checklocksignore
func (table *Table) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &table.Rules)
	stateSourceObject.Load(1, &table.BuiltinChains)
	stateSourceObject.Load(2, &table.Underflows)
}

func (r *Rule) StateTypeName() string {
	return "pkg/tcpip/stack.Rule"
}

func (r *Rule) StateFields() []string {
	return []string{
		"Filter",
		"Matchers",
		"Target",
	}
}

func (r *Rule) beforeSave() {}

// +checklocksignore
func (r *Rule) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.Filter)
	stateSinkObject.Save(1, &r.Matchers)
	stateSinkObject.Save(2, &r.Target)
}

func (r *Rule) afterLoad() {}

// +checklocksignore
func (r *Rule) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.Filter)
	stateSourceObject.Load(1, &r.Matchers)
	stateSourceObject.Load(2, &r.Target)
}

func (fl *IPHeaderFilter) StateTypeName() string {
	return "pkg/tcpip/stack.IPHeaderFilter"
}

func (fl *IPHeaderFilter) StateFields() []string {
	return []string{
		"Protocol",
		"CheckProtocol",
		"Dst",
		"DstMask",
		"DstInvert",
		"Src",
		"SrcMask",
		"SrcInvert",
		"InputInterface",
		"InputInterfaceMask",
		"InputInterfaceInvert",
		"OutputInterface",
		"OutputInterfaceMask",
		"OutputInterfaceInvert",
	}
}

func (fl *IPHeaderFilter) beforeSave() {}

// +checklocksignore
func (fl *IPHeaderFilter) StateSave(stateSinkObject state.Sink) {
	fl.beforeSave()
	stateSinkObject.Save(0, &fl.Protocol)
	stateSinkObject.Save(1, &fl.CheckProtocol)
	stateSinkObject.Save(2, &fl.Dst)
	stateSinkObject.Save(3, &fl.DstMask)
	stateSinkObject.Save(4, &fl.DstInvert)
	stateSinkObject.Save(5, &fl.Src)
	stateSinkObject.Save(6, &fl.SrcMask)
	stateSinkObject.Save(7, &fl.SrcInvert)
	stateSinkObject.Save(8, &fl.InputInterface)
	stateSinkObject.Save(9, &fl.InputInterfaceMask)
	stateSinkObject.Save(10, &fl.InputInterfaceInvert)
	stateSinkObject.Save(11, &fl.OutputInterface)
	stateSinkObject.Save(12, &fl.OutputInterfaceMask)
	stateSinkObject.Save(13, &fl.OutputInterfaceInvert)
}

func (fl *IPHeaderFilter) afterLoad() {}

// +checklocksignore
func (fl *IPHeaderFilter) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fl.Protocol)
	stateSourceObject.Load(1, &fl.CheckProtocol)
	stateSourceObject.Load(2, &fl.Dst)
	stateSourceObject.Load(3, &fl.DstMask)
	stateSourceObject.Load(4, &fl.DstInvert)
	stateSourceObject.Load(5, &fl.Src)
	stateSourceObject.Load(6, &fl.SrcMask)
	stateSourceObject.Load(7, &fl.SrcInvert)
	stateSourceObject.Load(8, &fl.InputInterface)
	stateSourceObject.Load(9, &fl.InputInterfaceMask)
	stateSourceObject.Load(10, &fl.InputInterfaceInvert)
	stateSourceObject.Load(11, &fl.OutputInterface)
	stateSourceObject.Load(12, &fl.OutputInterfaceMask)
	stateSourceObject.Load(13, &fl.OutputInterfaceInvert)
}

func (l *neighborEntryList) StateTypeName() string {
	return "pkg/tcpip/stack.neighborEntryList"
}

func (l *neighborEntryList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *neighborEntryList) beforeSave() {}

// +checklocksignore
func (l *neighborEntryList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *neighborEntryList) afterLoad() {}

// +checklocksignore
func (l *neighborEntryList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *neighborEntryEntry) StateTypeName() string {
	return "pkg/tcpip/stack.neighborEntryEntry"
}

func (e *neighborEntryEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *neighborEntryEntry) beforeSave() {}

// +checklocksignore
func (e *neighborEntryEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *neighborEntryEntry) afterLoad() {}

// +checklocksignore
func (e *neighborEntryEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (p *PacketBufferList) StateTypeName() string {
	return "pkg/tcpip/stack.PacketBufferList"
}

func (p *PacketBufferList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (p *PacketBufferList) beforeSave() {}

// +checklocksignore
func (p *PacketBufferList) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.head)
	stateSinkObject.Save(1, &p.tail)
}

func (p *PacketBufferList) afterLoad() {}

// +checklocksignore
func (p *PacketBufferList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.head)
	stateSourceObject.Load(1, &p.tail)
}

func (e *PacketBufferEntry) StateTypeName() string {
	return "pkg/tcpip/stack.PacketBufferEntry"
}

func (e *PacketBufferEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *PacketBufferEntry) beforeSave() {}

// +checklocksignore
func (e *PacketBufferEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *PacketBufferEntry) afterLoad() {}

// +checklocksignore
func (e *PacketBufferEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (r *packetBufferRefs) StateTypeName() string {
	return "pkg/tcpip/stack.packetBufferRefs"
}

func (r *packetBufferRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *packetBufferRefs) beforeSave() {}

// +checklocksignore
func (r *packetBufferRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *packetBufferRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (t *TransportEndpointID) StateTypeName() string {
	return "pkg/tcpip/stack.TransportEndpointID"
}

func (t *TransportEndpointID) StateFields() []string {
	return []string{
		"LocalPort",
		"LocalAddress",
		"RemotePort",
		"RemoteAddress",
	}
}

func (t *TransportEndpointID) beforeSave() {}

// +checklocksignore
func (t *TransportEndpointID) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.LocalPort)
	stateSinkObject.Save(1, &t.LocalAddress)
	stateSinkObject.Save(2, &t.RemotePort)
	stateSinkObject.Save(3, &t.RemoteAddress)
}

func (t *TransportEndpointID) afterLoad() {}

// +checklocksignore
func (t *TransportEndpointID) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.LocalPort)
	stateSourceObject.Load(1, &t.LocalAddress)
	stateSourceObject.Load(2, &t.RemotePort)
	stateSourceObject.Load(3, &t.RemoteAddress)
}

func (g *GSOType) StateTypeName() string {
	return "pkg/tcpip/stack.GSOType"
}

func (g *GSOType) StateFields() []string {
	return nil
}

func (g *GSO) StateTypeName() string {
	return "pkg/tcpip/stack.GSO"
}

func (g *GSO) StateFields() []string {
	return []string{
		"Type",
		"NeedsCsum",
		"CsumOffset",
		"MSS",
		"L3HdrLen",
		"MaxSize",
	}
}

func (g *GSO) beforeSave() {}

// +checklocksignore
func (g *GSO) StateSave(stateSinkObject state.Sink) {
	g.beforeSave()
	stateSinkObject.Save(0, &g.Type)
	stateSinkObject.Save(1, &g.NeedsCsum)
	stateSinkObject.Save(2, &g.CsumOffset)
	stateSinkObject.Save(3, &g.MSS)
	stateSinkObject.Save(4, &g.L3HdrLen)
	stateSinkObject.Save(5, &g.MaxSize)
}

func (g *GSO) afterLoad() {}

// +checklocksignore
func (g *GSO) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &g.Type)
	stateSourceObject.Load(1, &g.NeedsCsum)
	stateSourceObject.Load(2, &g.CsumOffset)
	stateSourceObject.Load(3, &g.MSS)
	stateSourceObject.Load(4, &g.L3HdrLen)
	stateSourceObject.Load(5, &g.MaxSize)
}

func (t *TransportEndpointInfo) StateTypeName() string {
	return "pkg/tcpip/stack.TransportEndpointInfo"
}

func (t *TransportEndpointInfo) StateFields() []string {
	return []string{
		"NetProto",
		"TransProto",
		"ID",
		"BindNICID",
		"BindAddr",
		"RegisterNICID",
	}
}

func (t *TransportEndpointInfo) beforeSave() {}

// +checklocksignore
func (t *TransportEndpointInfo) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.NetProto)
	stateSinkObject.Save(1, &t.TransProto)
	stateSinkObject.Save(2, &t.ID)
	stateSinkObject.Save(3, &t.BindNICID)
	stateSinkObject.Save(4, &t.BindAddr)
	stateSinkObject.Save(5, &t.RegisterNICID)
}

func (t *TransportEndpointInfo) afterLoad() {}

// +checklocksignore
func (t *TransportEndpointInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.NetProto)
	stateSourceObject.Load(1, &t.TransProto)
	stateSourceObject.Load(2, &t.ID)
	stateSourceObject.Load(3, &t.BindNICID)
	stateSourceObject.Load(4, &t.BindAddr)
	stateSourceObject.Load(5, &t.RegisterNICID)
}

func (t *TCPCubicState) StateTypeName() string {
	return "pkg/tcpip/stack.TCPCubicState"
}

func (t *TCPCubicState) StateFields() []string {
	return []string{
		"WLastMax",
		"WMax",
		"T",
		"TimeSinceLastCongestion",
		"C",
		"K",
		"Beta",
		"WC",
		"WEst",
	}
}

func (t *TCPCubicState) beforeSave() {}

// +checklocksignore
func (t *TCPCubicState) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.WLastMax)
	stateSinkObject.Save(1, &t.WMax)
	stateSinkObject.Save(2, &t.T)
	stateSinkObject.Save(3, &t.TimeSinceLastCongestion)
	stateSinkObject.Save(4, &t.C)
	stateSinkObject.Save(5, &t.K)
	stateSinkObject.Save(6, &t.Beta)
	stateSinkObject.Save(7, &t.WC)
	stateSinkObject.Save(8, &t.WEst)
}

func (t *TCPCubicState) afterLoad() {}

// +checklocksignore
func (t *TCPCubicState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.WLastMax)
	stateSourceObject.Load(1, &t.WMax)
	stateSourceObject.Load(2, &t.T)
	stateSourceObject.Load(3, &t.TimeSinceLastCongestion)
	stateSourceObject.Load(4, &t.C)
	stateSourceObject.Load(5, &t.K)
	stateSourceObject.Load(6, &t.Beta)
	stateSourceObject.Load(7, &t.WC)
	stateSourceObject.Load(8, &t.WEst)
}

func (t *TCPRACKState) StateTypeName() string {
	return "pkg/tcpip/stack.TCPRACKState"
}

func (t *TCPRACKState) StateFields() []string {
	return []string{
		"XmitTime",
		"EndSequence",
		"FACK",
		"RTT",
		"Reord",
		"DSACKSeen",
		"ReoWnd",
		"ReoWndIncr",
		"ReoWndPersist",
		"RTTSeq",
	}
}

func (t *TCPRACKState) beforeSave() {}

// +checklocksignore
func (t *TCPRACKState) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.XmitTime)
	stateSinkObject.Save(1, &t.EndSequence)
	stateSinkObject.Save(2, &t.FACK)
	stateSinkObject.Save(3, &t.RTT)
	stateSinkObject.Save(4, &t.Reord)
	stateSinkObject.Save(5, &t.DSACKSeen)
	stateSinkObject.Save(6, &t.ReoWnd)
	stateSinkObject.Save(7, &t.ReoWndIncr)
	stateSinkObject.Save(8, &t.ReoWndPersist)
	stateSinkObject.Save(9, &t.RTTSeq)
}

func (t *TCPRACKState) afterLoad() {}

// +checklocksignore
func (t *TCPRACKState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.XmitTime)
	stateSourceObject.Load(1, &t.EndSequence)
	stateSourceObject.Load(2, &t.FACK)
	stateSourceObject.Load(3, &t.RTT)
	stateSourceObject.Load(4, &t.Reord)
	stateSourceObject.Load(5, &t.DSACKSeen)
	stateSourceObject.Load(6, &t.ReoWnd)
	stateSourceObject.Load(7, &t.ReoWndIncr)
	stateSourceObject.Load(8, &t.ReoWndPersist)
	stateSourceObject.Load(9, &t.RTTSeq)
}

func (t *TCPEndpointID) StateTypeName() string {
	return "pkg/tcpip/stack.TCPEndpointID"
}

func (t *TCPEndpointID) StateFields() []string {
	return []string{
		"LocalPort",
		"LocalAddress",
		"RemotePort",
		"RemoteAddress",
	}
}

func (t *TCPEndpointID) beforeSave() {}

// +checklocksignore
func (t *TCPEndpointID) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.LocalPort)
	stateSinkObject.Save(1, &t.LocalAddress)
	stateSinkObject.Save(2, &t.RemotePort)
	stateSinkObject.Save(3, &t.RemoteAddress)
}

func (t *TCPEndpointID) afterLoad() {}

// +checklocksignore
func (t *TCPEndpointID) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.LocalPort)
	stateSourceObject.Load(1, &t.LocalAddress)
	stateSourceObject.Load(2, &t.RemotePort)
	stateSourceObject.Load(3, &t.RemoteAddress)
}

func (t *TCPFastRecoveryState) StateTypeName() string {
	return "pkg/tcpip/stack.TCPFastRecoveryState"
}

func (t *TCPFastRecoveryState) StateFields() []string {
	return []string{
		"Active",
		"First",
		"Last",
		"MaxCwnd",
		"HighRxt",
		"RescueRxt",
	}
}

func (t *TCPFastRecoveryState) beforeSave() {}

// +checklocksignore
func (t *TCPFastRecoveryState) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.Active)
	stateSinkObject.Save(1, &t.First)
	stateSinkObject.Save(2, &t.Last)
	stateSinkObject.Save(3, &t.MaxCwnd)
	stateSinkObject.Save(4, &t.HighRxt)
	stateSinkObject.Save(5, &t.RescueRxt)
}

func (t *TCPFastRecoveryState) afterLoad() {}

// +checklocksignore
func (t *TCPFastRecoveryState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.Active)
	stateSourceObject.Load(1, &t.First)
	stateSourceObject.Load(2, &t.Last)
	stateSourceObject.Load(3, &t.MaxCwnd)
	stateSourceObject.Load(4, &t.HighRxt)
	stateSourceObject.Load(5, &t.RescueRxt)
}

func (t *TCPReceiverState) StateTypeName() string {
	return "pkg/tcpip/stack.TCPReceiverState"
}

func (t *TCPReceiverState) StateFields() []string {
	return []string{
		"RcvNxt",
		"RcvAcc",
		"RcvWndScale",
		"PendingBufUsed",
	}
}

func (t *TCPReceiverState) beforeSave() {}

// +checklocksignore
func (t *TCPReceiverState) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.RcvNxt)
	stateSinkObject.Save(1, &t.RcvAcc)
	stateSinkObject.Save(2, &t.RcvWndScale)
	stateSinkObject.Save(3, &t.PendingBufUsed)
}

func (t *TCPReceiverState) afterLoad() {}

// +checklocksignore
func (t *TCPReceiverState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.RcvNxt)
	stateSourceObject.Load(1, &t.RcvAcc)
	stateSourceObject.Load(2, &t.RcvWndScale)
	stateSourceObject.Load(3, &t.PendingBufUsed)
}

func (t *TCPRTTState) StateTypeName() string {
	return "pkg/tcpip/stack.TCPRTTState"
}

func (t *TCPRTTState) StateFields() []string {
	return []string{
		"SRTT",
		"RTTVar",
		"SRTTInited",
	}
}

func (t *TCPRTTState) beforeSave() {}

// +checklocksignore
func (t *TCPRTTState) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.SRTT)
	stateSinkObject.Save(1, &t.RTTVar)
	stateSinkObject.Save(2, &t.SRTTInited)
}

func (t *TCPRTTState) afterLoad() {}

// +checklocksignore
func (t *TCPRTTState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.SRTT)
	stateSourceObject.Load(1, &t.RTTVar)
	stateSourceObject.Load(2, &t.SRTTInited)
}

func (t *TCPSenderState) StateTypeName() string {
	return "pkg/tcpip/stack.TCPSenderState"
}

func (t *TCPSenderState) StateFields() []string {
	return []string{
		"LastSendTime",
		"DupAckCount",
		"SndCwnd",
		"Ssthresh",
		"SndCAAckCount",
		"Outstanding",
		"SackedOut",
		"SndWnd",
		"SndUna",
		"SndNxt",
		"RTTMeasureSeqNum",
		"RTTMeasureTime",
		"Closed",
		"RTO",
		"RTTState",
		"MaxPayloadSize",
		"SndWndScale",
		"MaxSentAck",
		"FastRecovery",
		"Cubic",
		"RACKState",
		"RetransmitTS",
		"SpuriousRecovery",
	}
}

func (t *TCPSenderState) beforeSave() {}

// +checklocksignore
func (t *TCPSenderState) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.LastSendTime)
	stateSinkObject.Save(1, &t.DupAckCount)
	stateSinkObject.Save(2, &t.SndCwnd)
	stateSinkObject.Save(3, &t.Ssthresh)
	stateSinkObject.Save(4, &t.SndCAAckCount)
	stateSinkObject.Save(5, &t.Outstanding)
	stateSinkObject.Save(6, &t.SackedOut)
	stateSinkObject.Save(7, &t.SndWnd)
	stateSinkObject.Save(8, &t.SndUna)
	stateSinkObject.Save(9, &t.SndNxt)
	stateSinkObject.Save(10, &t.RTTMeasureSeqNum)
	stateSinkObject.Save(11, &t.RTTMeasureTime)
	stateSinkObject.Save(12, &t.Closed)
	stateSinkObject.Save(13, &t.RTO)
	stateSinkObject.Save(14, &t.RTTState)
	stateSinkObject.Save(15, &t.MaxPayloadSize)
	stateSinkObject.Save(16, &t.SndWndScale)
	stateSinkObject.Save(17, &t.MaxSentAck)
	stateSinkObject.Save(18, &t.FastRecovery)
	stateSinkObject.Save(19, &t.Cubic)
	stateSinkObject.Save(20, &t.RACKState)
	stateSinkObject.Save(21, &t.RetransmitTS)
	stateSinkObject.Save(22, &t.SpuriousRecovery)
}

func (t *TCPSenderState) afterLoad() {}

// +checklocksignore
func (t *TCPSenderState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.LastSendTime)
	stateSourceObject.Load(1, &t.DupAckCount)
	stateSourceObject.Load(2, &t.SndCwnd)
	stateSourceObject.Load(3, &t.Ssthresh)
	stateSourceObject.Load(4, &t.SndCAAckCount)
	stateSourceObject.Load(5, &t.Outstanding)
	stateSourceObject.Load(6, &t.SackedOut)
	stateSourceObject.Load(7, &t.SndWnd)
	stateSourceObject.Load(8, &t.SndUna)
	stateSourceObject.Load(9, &t.SndNxt)
	stateSourceObject.Load(10, &t.RTTMeasureSeqNum)
	stateSourceObject.Load(11, &t.RTTMeasureTime)
	stateSourceObject.Load(12, &t.Closed)
	stateSourceObject.Load(13, &t.RTO)
	stateSourceObject.Load(14, &t.RTTState)
	stateSourceObject.Load(15, &t.MaxPayloadSize)
	stateSourceObject.Load(16, &t.SndWndScale)
	stateSourceObject.Load(17, &t.MaxSentAck)
	stateSourceObject.Load(18, &t.FastRecovery)
	stateSourceObject.Load(19, &t.Cubic)
	stateSourceObject.Load(20, &t.RACKState)
	stateSourceObject.Load(21, &t.RetransmitTS)
	stateSourceObject.Load(22, &t.SpuriousRecovery)
}

func (t *TCPSACKInfo) StateTypeName() string {
	return "pkg/tcpip/stack.TCPSACKInfo"
}

func (t *TCPSACKInfo) StateFields() []string {
	return []string{
		"Blocks",
		"ReceivedBlocks",
		"MaxSACKED",
	}
}

func (t *TCPSACKInfo) beforeSave() {}

// +checklocksignore
func (t *TCPSACKInfo) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.Blocks)
	stateSinkObject.Save(1, &t.ReceivedBlocks)
	stateSinkObject.Save(2, &t.MaxSACKED)
}

func (t *TCPSACKInfo) afterLoad() {}

// +checklocksignore
func (t *TCPSACKInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.Blocks)
	stateSourceObject.Load(1, &t.ReceivedBlocks)
	stateSourceObject.Load(2, &t.MaxSACKED)
}

func (r *RcvBufAutoTuneParams) StateTypeName() string {
	return "pkg/tcpip/stack.RcvBufAutoTuneParams"
}

func (r *RcvBufAutoTuneParams) StateFields() []string {
	return []string{
		"MeasureTime",
		"CopiedBytes",
		"PrevCopiedBytes",
		"RcvBufSize",
		"RTT",
		"RTTVar",
		"RTTMeasureSeqNumber",
		"RTTMeasureTime",
		"Disabled",
	}
}

func (r *RcvBufAutoTuneParams) beforeSave() {}

// +checklocksignore
func (r *RcvBufAutoTuneParams) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.MeasureTime)
	stateSinkObject.Save(1, &r.CopiedBytes)
	stateSinkObject.Save(2, &r.PrevCopiedBytes)
	stateSinkObject.Save(3, &r.RcvBufSize)
	stateSinkObject.Save(4, &r.RTT)
	stateSinkObject.Save(5, &r.RTTVar)
	stateSinkObject.Save(6, &r.RTTMeasureSeqNumber)
	stateSinkObject.Save(7, &r.RTTMeasureTime)
	stateSinkObject.Save(8, &r.Disabled)
}

func (r *RcvBufAutoTuneParams) afterLoad() {}

// +checklocksignore
func (r *RcvBufAutoTuneParams) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.MeasureTime)
	stateSourceObject.Load(1, &r.CopiedBytes)
	stateSourceObject.Load(2, &r.PrevCopiedBytes)
	stateSourceObject.Load(3, &r.RcvBufSize)
	stateSourceObject.Load(4, &r.RTT)
	stateSourceObject.Load(5, &r.RTTVar)
	stateSourceObject.Load(6, &r.RTTMeasureSeqNumber)
	stateSourceObject.Load(7, &r.RTTMeasureTime)
	stateSourceObject.Load(8, &r.Disabled)
}

func (t *TCPRcvBufState) StateTypeName() string {
	return "pkg/tcpip/stack.TCPRcvBufState"
}

func (t *TCPRcvBufState) StateFields() []string {
	return []string{
		"RcvBufUsed",
		"RcvAutoParams",
		"RcvClosed",
	}
}

func (t *TCPRcvBufState) beforeSave() {}

// +checklocksignore
func (t *TCPRcvBufState) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.RcvBufUsed)
	stateSinkObject.Save(1, &t.RcvAutoParams)
	stateSinkObject.Save(2, &t.RcvClosed)
}

func (t *TCPRcvBufState) afterLoad() {}

// +checklocksignore
func (t *TCPRcvBufState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.RcvBufUsed)
	stateSourceObject.Load(1, &t.RcvAutoParams)
	stateSourceObject.Load(2, &t.RcvClosed)
}

func (t *TCPSndBufState) StateTypeName() string {
	return "pkg/tcpip/stack.TCPSndBufState"
}

func (t *TCPSndBufState) StateFields() []string {
	return []string{
		"SndBufSize",
		"SndBufUsed",
		"SndClosed",
		"PacketTooBigCount",
		"SndMTU",
		"AutoTuneSndBufDisabled",
	}
}

func (t *TCPSndBufState) beforeSave() {}

// +checklocksignore
func (t *TCPSndBufState) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.SndBufSize)
	stateSinkObject.Save(1, &t.SndBufUsed)
	stateSinkObject.Save(2, &t.SndClosed)
	stateSinkObject.Save(3, &t.PacketTooBigCount)
	stateSinkObject.Save(4, &t.SndMTU)
	stateSinkObject.Save(5, &t.AutoTuneSndBufDisabled)
}

func (t *TCPSndBufState) afterLoad() {}

// +checklocksignore
func (t *TCPSndBufState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.SndBufSize)
	stateSourceObject.Load(1, &t.SndBufUsed)
	stateSourceObject.Load(2, &t.SndClosed)
	stateSourceObject.Load(3, &t.PacketTooBigCount)
	stateSourceObject.Load(4, &t.SndMTU)
	stateSourceObject.Load(5, &t.AutoTuneSndBufDisabled)
}

func (t *TCPEndpointStateInner) StateTypeName() string {
	return "pkg/tcpip/stack.TCPEndpointStateInner"
}

func (t *TCPEndpointStateInner) StateFields() []string {
	return []string{
		"TSOffset",
		"SACKPermitted",
		"SendTSOk",
		"RecentTS",
	}
}

func (t *TCPEndpointStateInner) beforeSave() {}

// +checklocksignore
func (t *TCPEndpointStateInner) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.TSOffset)
	stateSinkObject.Save(1, &t.SACKPermitted)
	stateSinkObject.Save(2, &t.SendTSOk)
	stateSinkObject.Save(3, &t.RecentTS)
}

func (t *TCPEndpointStateInner) afterLoad() {}

// +checklocksignore
func (t *TCPEndpointStateInner) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.TSOffset)
	stateSourceObject.Load(1, &t.SACKPermitted)
	stateSourceObject.Load(2, &t.SendTSOk)
	stateSourceObject.Load(3, &t.RecentTS)
}

func (t *TCPEndpointState) StateTypeName() string {
	return "pkg/tcpip/stack.TCPEndpointState"
}

func (t *TCPEndpointState) StateFields() []string {
	return []string{
		"TCPEndpointStateInner",
		"ID",
		"SegTime",
		"RcvBufState",
		"SndBufState",
		"SACK",
		"Receiver",
		"Sender",
	}
}

func (t *TCPEndpointState) beforeSave() {}

// +checklocksignore
func (t *TCPEndpointState) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.TCPEndpointStateInner)
	stateSinkObject.Save(1, &t.ID)
	stateSinkObject.Save(2, &t.SegTime)
	stateSinkObject.Save(3, &t.RcvBufState)
	stateSinkObject.Save(4, &t.SndBufState)
	stateSinkObject.Save(5, &t.SACK)
	stateSinkObject.Save(6, &t.Receiver)
	stateSinkObject.Save(7, &t.Sender)
}

func (t *TCPEndpointState) afterLoad() {}

// +checklocksignore
func (t *TCPEndpointState) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.TCPEndpointStateInner)
	stateSourceObject.Load(1, &t.ID)
	stateSourceObject.Load(2, &t.SegTime)
	stateSourceObject.Load(3, &t.RcvBufState)
	stateSourceObject.Load(4, &t.SndBufState)
	stateSourceObject.Load(5, &t.SACK)
	stateSourceObject.Load(6, &t.Receiver)
	stateSourceObject.Load(7, &t.Sender)
}

func (ep *multiPortEndpoint) StateTypeName() string {
	return "pkg/tcpip/stack.multiPortEndpoint"
}

func (ep *multiPortEndpoint) StateFields() []string {
	return []string{
		"demux",
		"netProto",
		"transProto",
		"flags",
		"endpoints",
	}
}

func (ep *multiPortEndpoint) beforeSave() {}

// +checklocksignore
func (ep *multiPortEndpoint) StateSave(stateSinkObject state.Sink) {
	ep.beforeSave()
	stateSinkObject.Save(0, &ep.demux)
	stateSinkObject.Save(1, &ep.netProto)
	stateSinkObject.Save(2, &ep.transProto)
	stateSinkObject.Save(3, &ep.flags)
	stateSinkObject.Save(4, &ep.endpoints)
}

func (ep *multiPortEndpoint) afterLoad() {}

// +checklocksignore
func (ep *multiPortEndpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ep.demux)
	stateSourceObject.Load(1, &ep.netProto)
	stateSourceObject.Load(2, &ep.transProto)
	stateSourceObject.Load(3, &ep.flags)
	stateSourceObject.Load(4, &ep.endpoints)
}

func (l *tupleList) StateTypeName() string {
	return "pkg/tcpip/stack.tupleList"
}

func (l *tupleList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *tupleList) beforeSave() {}

// +checklocksignore
func (l *tupleList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *tupleList) afterLoad() {}

// +checklocksignore
func (l *tupleList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *tupleEntry) StateTypeName() string {
	return "pkg/tcpip/stack.tupleEntry"
}

func (e *tupleEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *tupleEntry) beforeSave() {}

// +checklocksignore
func (e *tupleEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *tupleEntry) afterLoad() {}

// +checklocksignore
func (e *tupleEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*tuple)(nil))
	state.Register((*tupleID)(nil))
	state.Register((*conn)(nil))
	state.Register((*ConnTrack)(nil))
	state.Register((*bucket)(nil))
	state.Register((*IPTables)(nil))
	state.Register((*Table)(nil))
	state.Register((*Rule)(nil))
	state.Register((*IPHeaderFilter)(nil))
	state.Register((*neighborEntryList)(nil))
	state.Register((*neighborEntryEntry)(nil))
	state.Register((*PacketBufferList)(nil))
	state.Register((*PacketBufferEntry)(nil))
	state.Register((*packetBufferRefs)(nil))
	state.Register((*TransportEndpointID)(nil))
	state.Register((*GSOType)(nil))
	state.Register((*GSO)(nil))
	state.Register((*TransportEndpointInfo)(nil))
	state.Register((*TCPCubicState)(nil))
	state.Register((*TCPRACKState)(nil))
	state.Register((*TCPEndpointID)(nil))
	state.Register((*TCPFastRecoveryState)(nil))
	state.Register((*TCPReceiverState)(nil))
	state.Register((*TCPRTTState)(nil))
	state.Register((*TCPSenderState)(nil))
	state.Register((*TCPSACKInfo)(nil))
	state.Register((*RcvBufAutoTuneParams)(nil))
	state.Register((*TCPRcvBufState)(nil))
	state.Register((*TCPSndBufState)(nil))
	state.Register((*TCPEndpointStateInner)(nil))
	state.Register((*TCPEndpointState)(nil))
	state.Register((*multiPortEndpoint)(nil))
	state.Register((*tupleList)(nil))
	state.Register((*tupleEntry)(nil))
}