// automatically generated by stateify.

package pgalloc

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (s *aplUnloadedSet) StateTypeName() string {
	return "pkg/sentry/pgalloc.aplUnloadedSet"
}

func (s *aplUnloadedSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *aplUnloadedSet) beforeSave() {}

// +checklocksignore
func (s *aplUnloadedSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue []aplUnloadedFlatSegment
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *aplUnloadedSet) afterLoad(context.Context) {}

// +checklocksignore
func (s *aplUnloadedSet) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new([]aplUnloadedFlatSegment), func(y any) { s.loadRoot(ctx, y.([]aplUnloadedFlatSegment)) })
}

func (n *aplUnloadednode) StateTypeName() string {
	return "pkg/sentry/pgalloc.aplUnloadednode"
}

func (n *aplUnloadednode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *aplUnloadednode) beforeSave() {}

// +checklocksignore
func (n *aplUnloadednode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *aplUnloadednode) afterLoad(context.Context) {}

// +checklocksignore
func (n *aplUnloadednode) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (a *aplUnloadedFlatSegment) StateTypeName() string {
	return "pkg/sentry/pgalloc.aplUnloadedFlatSegment"
}

func (a *aplUnloadedFlatSegment) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Value",
	}
}

func (a *aplUnloadedFlatSegment) beforeSave() {}

// +checklocksignore
func (a *aplUnloadedFlatSegment) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.Start)
	stateSinkObject.Save(1, &a.End)
	stateSinkObject.Save(2, &a.Value)
}

func (a *aplUnloadedFlatSegment) afterLoad(context.Context) {}

// +checklocksignore
func (a *aplUnloadedFlatSegment) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.Start)
	stateSourceObject.Load(1, &a.End)
	stateSourceObject.Load(2, &a.Value)
}

func (r *EvictableRange) StateTypeName() string {
	return "pkg/sentry/pgalloc.EvictableRange"
}

func (r *EvictableRange) StateFields() []string {
	return []string{
		"Start",
		"End",
	}
}

func (r *EvictableRange) beforeSave() {}

// +checklocksignore
func (r *EvictableRange) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.Start)
	stateSinkObject.Save(1, &r.End)
}

func (r *EvictableRange) afterLoad(context.Context) {}

// +checklocksignore
func (r *EvictableRange) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.Start)
	stateSourceObject.Load(1, &r.End)
}

func (s *evictableRangeSet) StateTypeName() string {
	return "pkg/sentry/pgalloc.evictableRangeSet"
}

func (s *evictableRangeSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *evictableRangeSet) beforeSave() {}

// +checklocksignore
func (s *evictableRangeSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue []evictableRangeFlatSegment
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *evictableRangeSet) afterLoad(context.Context) {}

// +checklocksignore
func (s *evictableRangeSet) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new([]evictableRangeFlatSegment), func(y any) { s.loadRoot(ctx, y.([]evictableRangeFlatSegment)) })
}

func (n *evictableRangenode) StateTypeName() string {
	return "pkg/sentry/pgalloc.evictableRangenode"
}

func (n *evictableRangenode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *evictableRangenode) beforeSave() {}

// +checklocksignore
func (n *evictableRangenode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *evictableRangenode) afterLoad(context.Context) {}

// +checklocksignore
func (n *evictableRangenode) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (e *evictableRangeFlatSegment) StateTypeName() string {
	return "pkg/sentry/pgalloc.evictableRangeFlatSegment"
}

func (e *evictableRangeFlatSegment) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Value",
	}
}

func (e *evictableRangeFlatSegment) beforeSave() {}

// +checklocksignore
func (e *evictableRangeFlatSegment) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.Start)
	stateSinkObject.Save(1, &e.End)
	stateSinkObject.Save(2, &e.Value)
}

func (e *evictableRangeFlatSegment) afterLoad(context.Context) {}

// +checklocksignore
func (e *evictableRangeFlatSegment) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.Start)
	stateSourceObject.Load(1, &e.End)
	stateSourceObject.Load(2, &e.Value)
}

func (s *memAcctSet) StateTypeName() string {
	return "pkg/sentry/pgalloc.memAcctSet"
}

func (s *memAcctSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *memAcctSet) beforeSave() {}

// +checklocksignore
func (s *memAcctSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue []memAcctFlatSegment
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *memAcctSet) afterLoad(context.Context) {}

// +checklocksignore
func (s *memAcctSet) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new([]memAcctFlatSegment), func(y any) { s.loadRoot(ctx, y.([]memAcctFlatSegment)) })
}

func (n *memAcctnode) StateTypeName() string {
	return "pkg/sentry/pgalloc.memAcctnode"
}

func (n *memAcctnode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *memAcctnode) beforeSave() {}

// +checklocksignore
func (n *memAcctnode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *memAcctnode) afterLoad(context.Context) {}

// +checklocksignore
func (n *memAcctnode) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (m *memAcctFlatSegment) StateTypeName() string {
	return "pkg/sentry/pgalloc.memAcctFlatSegment"
}

func (m *memAcctFlatSegment) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Value",
	}
}

func (m *memAcctFlatSegment) beforeSave() {}

// +checklocksignore
func (m *memAcctFlatSegment) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.Start)
	stateSinkObject.Save(1, &m.End)
	stateSinkObject.Save(2, &m.Value)
}

func (m *memAcctFlatSegment) afterLoad(context.Context) {}

// +checklocksignore
func (m *memAcctFlatSegment) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.Start)
	stateSourceObject.Load(1, &m.End)
	stateSourceObject.Load(2, &m.Value)
}

func (c *chunkInfo) StateTypeName() string {
	return "pkg/sentry/pgalloc.chunkInfo"
}

func (c *chunkInfo) StateFields() []string {
	return []string{
		"huge",
	}
}

func (c *chunkInfo) beforeSave() {}

// +checklocksignore
func (c *chunkInfo) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.huge)
}

func (c *chunkInfo) afterLoad(context.Context) {}

// +checklocksignore
func (c *chunkInfo) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.huge)
}

func (u *unwasteInfo) StateTypeName() string {
	return "pkg/sentry/pgalloc.unwasteInfo"
}

func (u *unwasteInfo) StateFields() []string {
	return []string{}
}

func (u *unwasteInfo) beforeSave() {}

// +checklocksignore
func (u *unwasteInfo) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
}

func (u *unwasteInfo) afterLoad(context.Context) {}

// +checklocksignore
func (u *unwasteInfo) StateLoad(ctx context.Context, stateSourceObject state.Source) {
}

func (u *unfreeInfo) StateTypeName() string {
	return "pkg/sentry/pgalloc.unfreeInfo"
}

func (u *unfreeInfo) StateFields() []string {
	return []string{
		"refs",
	}
}

func (u *unfreeInfo) beforeSave() {}

// +checklocksignore
func (u *unfreeInfo) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.refs)
}

func (u *unfreeInfo) afterLoad(context.Context) {}

// +checklocksignore
func (u *unfreeInfo) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.refs)
}

func (m *memAcctInfo) StateTypeName() string {
	return "pkg/sentry/pgalloc.memAcctInfo"
}

func (m *memAcctInfo) StateFields() []string {
	return []string{
		"kind",
		"memCgID",
		"knownCommitted",
		"wasteOrReleasing",
		"commitSeq",
	}
}

func (m *memAcctInfo) beforeSave() {}

// +checklocksignore
func (m *memAcctInfo) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.kind)
	stateSinkObject.Save(1, &m.memCgID)
	stateSinkObject.Save(2, &m.knownCommitted)
	stateSinkObject.Save(3, &m.wasteOrReleasing)
	stateSinkObject.Save(4, &m.commitSeq)
}

func (m *memAcctInfo) afterLoad(context.Context) {}

// +checklocksignore
func (m *memAcctInfo) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.kind)
	stateSourceObject.Load(1, &m.memCgID)
	stateSourceObject.Load(2, &m.knownCommitted)
	stateSourceObject.Load(3, &m.wasteOrReleasing)
	stateSourceObject.Load(4, &m.commitSeq)
}

func (s *unfreeSet) StateTypeName() string {
	return "pkg/sentry/pgalloc.unfreeSet"
}

func (s *unfreeSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *unfreeSet) beforeSave() {}

// +checklocksignore
func (s *unfreeSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue []unfreeFlatSegment
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *unfreeSet) afterLoad(context.Context) {}

// +checklocksignore
func (s *unfreeSet) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new([]unfreeFlatSegment), func(y any) { s.loadRoot(ctx, y.([]unfreeFlatSegment)) })
}

func (n *unfreenode) StateTypeName() string {
	return "pkg/sentry/pgalloc.unfreenode"
}

func (n *unfreenode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *unfreenode) beforeSave() {}

// +checklocksignore
func (n *unfreenode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *unfreenode) afterLoad(context.Context) {}

// +checklocksignore
func (n *unfreenode) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (u *unfreeFlatSegment) StateTypeName() string {
	return "pkg/sentry/pgalloc.unfreeFlatSegment"
}

func (u *unfreeFlatSegment) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Value",
	}
}

func (u *unfreeFlatSegment) beforeSave() {}

// +checklocksignore
func (u *unfreeFlatSegment) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.Start)
	stateSinkObject.Save(1, &u.End)
	stateSinkObject.Save(2, &u.Value)
}

func (u *unfreeFlatSegment) afterLoad(context.Context) {}

// +checklocksignore
func (u *unfreeFlatSegment) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.Start)
	stateSourceObject.Load(1, &u.End)
	stateSourceObject.Load(2, &u.Value)
}

func (s *unwasteSet) StateTypeName() string {
	return "pkg/sentry/pgalloc.unwasteSet"
}

func (s *unwasteSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *unwasteSet) beforeSave() {}

// +checklocksignore
func (s *unwasteSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue []unwasteFlatSegment
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *unwasteSet) afterLoad(context.Context) {}

// +checklocksignore
func (s *unwasteSet) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new([]unwasteFlatSegment), func(y any) { s.loadRoot(ctx, y.([]unwasteFlatSegment)) })
}

func (n *unwastenode) StateTypeName() string {
	return "pkg/sentry/pgalloc.unwastenode"
}

func (n *unwastenode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *unwastenode) beforeSave() {}

// +checklocksignore
func (n *unwastenode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *unwastenode) afterLoad(context.Context) {}

// +checklocksignore
func (n *unwastenode) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (u *unwasteFlatSegment) StateTypeName() string {
	return "pkg/sentry/pgalloc.unwasteFlatSegment"
}

func (u *unwasteFlatSegment) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Value",
	}
}

func (u *unwasteFlatSegment) beforeSave() {}

// +checklocksignore
func (u *unwasteFlatSegment) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.Start)
	stateSinkObject.Save(1, &u.End)
	stateSinkObject.Save(2, &u.Value)
}

func (u *unwasteFlatSegment) afterLoad(context.Context) {}

// +checklocksignore
func (u *unwasteFlatSegment) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.Start)
	stateSourceObject.Load(1, &u.End)
	stateSourceObject.Load(2, &u.Value)
}

func init() {
	state.Register((*aplUnloadedSet)(nil))
	state.Register((*aplUnloadednode)(nil))
	state.Register((*aplUnloadedFlatSegment)(nil))
	state.Register((*EvictableRange)(nil))
	state.Register((*evictableRangeSet)(nil))
	state.Register((*evictableRangenode)(nil))
	state.Register((*evictableRangeFlatSegment)(nil))
	state.Register((*memAcctSet)(nil))
	state.Register((*memAcctnode)(nil))
	state.Register((*memAcctFlatSegment)(nil))
	state.Register((*chunkInfo)(nil))
	state.Register((*unwasteInfo)(nil))
	state.Register((*unfreeInfo)(nil))
	state.Register((*memAcctInfo)(nil))
	state.Register((*unfreeSet)(nil))
	state.Register((*unfreenode)(nil))
	state.Register((*unfreeFlatSegment)(nil))
	state.Register((*unwasteSet)(nil))
	state.Register((*unwastenode)(nil))
	state.Register((*unwasteFlatSegment)(nil))
}
