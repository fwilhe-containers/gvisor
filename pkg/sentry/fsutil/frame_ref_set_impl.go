package fsutil

import (
	__generics_imported0 "gvisor.dev/gvisor/pkg/sentry/memmap"
)

import (
	"bytes"
	"context"
	"fmt"
)

// trackGaps is an optional parameter.
//
// If trackGaps is 1, the Set will track maximum gap size recursively,
// enabling the GapIterator.{Prev,Next}LargeEnoughGap functions. In this
// case, Key must be an unsigned integer.
//
// trackGaps must be 0 or 1.
const FrameReftrackGaps = 0

var _ = uint8(FrameReftrackGaps << 7) // Will fail if not zero or one.

// dynamicGap is a type that disappears if trackGaps is 0.
type FrameRefdynamicGap [FrameReftrackGaps]uint64

// Get returns the value of the gap.
//
// Precondition: trackGaps must be non-zero.
func (d *FrameRefdynamicGap) Get() uint64 {
	return d[:][0]
}

// Set sets the value of the gap.
//
// Precondition: trackGaps must be non-zero.
func (d *FrameRefdynamicGap) Set(v uint64) {
	d[:][0] = v
}

const (
	// minDegree is the minimum degree of an internal node in a Set B-tree.
	//
	//	- Any non-root node has at least minDegree-1 segments.
	//
	//	- Any non-root internal (non-leaf) node has at least minDegree children.
	//
	//	- The root node may have fewer than minDegree-1 segments, but it may
	// only have 0 segments if the tree is empty.
	//
	// Our implementation requires minDegree >= 3. Higher values of minDegree
	// usually improve performance, but increase memory usage for small sets.
	FrameRefminDegree = 3

	FrameRefmaxDegree = 2 * FrameRefminDegree
)

// A Set is a mapping of segments with non-overlapping Range keys. The zero
// value for a Set is an empty set. Set values are not safely movable nor
// copyable. Set is thread-compatible.
//
// +stateify savable
type FrameRefSet struct {
	root FrameRefnode `state:".([]FrameRefFlatSegment)"`
}

// IsEmpty returns true if the set contains no segments.
func (s *FrameRefSet) IsEmpty() bool {
	return s.root.nrSegments == 0
}

// IsEmptyRange returns true iff no segments in the set overlap the given
// range. This is semantically equivalent to s.SpanRange(r) == 0, but may be
// more efficient.
func (s *FrameRefSet) IsEmptyRange(r __generics_imported0.FileRange) bool {
	switch {
	case r.Length() < 0:
		panic(fmt.Sprintf("invalid range %v", r))
	case r.Length() == 0:
		return true
	}
	_, gap := s.Find(r.Start)
	if !gap.Ok() {
		return false
	}
	return r.End <= gap.End()
}

// Span returns the total size of all segments in the set.
func (s *FrameRefSet) Span() uint64 {
	var sz uint64
	for seg := s.FirstSegment(); seg.Ok(); seg = seg.NextSegment() {
		sz += seg.Range().Length()
	}
	return sz
}

// SpanRange returns the total size of the intersection of segments in the set
// with the given range.
func (s *FrameRefSet) SpanRange(r __generics_imported0.FileRange) uint64 {
	switch {
	case r.Length() < 0:
		panic(fmt.Sprintf("invalid range %v", r))
	case r.Length() == 0:
		return 0
	}
	var sz uint64
	for seg := s.LowerBoundSegment(r.Start); seg.Ok() && seg.Start() < r.End; seg = seg.NextSegment() {
		sz += seg.Range().Intersect(r).Length()
	}
	return sz
}

// FirstSegment returns the first segment in the set. If the set is empty,
// FirstSegment returns a terminal iterator.
func (s *FrameRefSet) FirstSegment() FrameRefIterator {
	if s.root.nrSegments == 0 {
		return FrameRefIterator{}
	}
	return s.root.firstSegment()
}

// LastSegment returns the last segment in the set. If the set is empty,
// LastSegment returns a terminal iterator.
func (s *FrameRefSet) LastSegment() FrameRefIterator {
	if s.root.nrSegments == 0 {
		return FrameRefIterator{}
	}
	return s.root.lastSegment()
}

// FirstGap returns the first gap in the set.
func (s *FrameRefSet) FirstGap() FrameRefGapIterator {
	n := &s.root
	for n.hasChildren {
		n = n.children[0]
	}
	return FrameRefGapIterator{n, 0}
}

// LastGap returns the last gap in the set.
func (s *FrameRefSet) LastGap() FrameRefGapIterator {
	n := &s.root
	for n.hasChildren {
		n = n.children[n.nrSegments]
	}
	return FrameRefGapIterator{n, n.nrSegments}
}

// Find returns the segment or gap whose range contains the given key. If a
// segment is found, the returned Iterator is non-terminal and the
// returned GapIterator is terminal. Otherwise, the returned Iterator is
// terminal and the returned GapIterator is non-terminal.
func (s *FrameRefSet) Find(key uint64) (FrameRefIterator, FrameRefGapIterator) {
	n := &s.root
	for {

		lower := 0
		upper := n.nrSegments
		for lower < upper {
			i := lower + (upper-lower)/2
			if r := n.keys[i]; key < r.End {
				if key >= r.Start {
					return FrameRefIterator{n, i}, FrameRefGapIterator{}
				}
				upper = i
			} else {
				lower = i + 1
			}
		}
		i := lower
		if !n.hasChildren {
			return FrameRefIterator{}, FrameRefGapIterator{n, i}
		}
		n = n.children[i]
	}
}

// FindSegment returns the segment whose range contains the given key. If no
// such segment exists, FindSegment returns a terminal iterator.
func (s *FrameRefSet) FindSegment(key uint64) FrameRefIterator {
	seg, _ := s.Find(key)
	return seg
}

// LowerBoundSegment returns the segment with the lowest range that contains a
// key greater than or equal to min. If no such segment exists,
// LowerBoundSegment returns a terminal iterator.
func (s *FrameRefSet) LowerBoundSegment(min uint64) FrameRefIterator {
	seg, gap := s.Find(min)
	if seg.Ok() {
		return seg
	}
	return gap.NextSegment()
}

// UpperBoundSegment returns the segment with the highest range that contains a
// key less than or equal to max. If no such segment exists, UpperBoundSegment
// returns a terminal iterator.
func (s *FrameRefSet) UpperBoundSegment(max uint64) FrameRefIterator {
	seg, gap := s.Find(max)
	if seg.Ok() {
		return seg
	}
	return gap.PrevSegment()
}

// FindGap returns the gap containing the given key. If no such gap exists
// (i.e. the set contains a segment containing that key), FindGap returns a
// terminal iterator.
func (s *FrameRefSet) FindGap(key uint64) FrameRefGapIterator {
	_, gap := s.Find(key)
	return gap
}

// LowerBoundGap returns the gap with the lowest range that is greater than or
// equal to min.
func (s *FrameRefSet) LowerBoundGap(min uint64) FrameRefGapIterator {
	seg, gap := s.Find(min)
	if gap.Ok() {
		return gap
	}
	return seg.NextGap()
}

// UpperBoundGap returns the gap with the highest range that is less than or
// equal to max.
func (s *FrameRefSet) UpperBoundGap(max uint64) FrameRefGapIterator {
	seg, gap := s.Find(max)
	if gap.Ok() {
		return gap
	}
	return seg.PrevGap()
}

// FirstLargeEnoughGap returns the first gap in the set with at least the given
// length. If no such gap exists, FirstLargeEnoughGap returns a terminal
// iterator.
//
// Precondition: trackGaps must be 1.
func (s *FrameRefSet) FirstLargeEnoughGap(minSize uint64) FrameRefGapIterator {
	if FrameReftrackGaps != 1 {
		panic("set is not tracking gaps")
	}
	gap := s.FirstGap()
	if gap.Range().Length() >= minSize {
		return gap
	}
	return gap.NextLargeEnoughGap(minSize)
}

// LastLargeEnoughGap returns the last gap in the set with at least the given
// length. If no such gap exists, LastLargeEnoughGap returns a terminal
// iterator.
//
// Precondition: trackGaps must be 1.
func (s *FrameRefSet) LastLargeEnoughGap(minSize uint64) FrameRefGapIterator {
	if FrameReftrackGaps != 1 {
		panic("set is not tracking gaps")
	}
	gap := s.LastGap()
	if gap.Range().Length() >= minSize {
		return gap
	}
	return gap.PrevLargeEnoughGap(minSize)
}

// LowerBoundLargeEnoughGap returns the first gap in the set with at least the
// given length and whose range contains a key greater than or equal to min. If
// no such gap exists, LowerBoundLargeEnoughGap returns a terminal iterator.
//
// Precondition: trackGaps must be 1.
func (s *FrameRefSet) LowerBoundLargeEnoughGap(min, minSize uint64) FrameRefGapIterator {
	if FrameReftrackGaps != 1 {
		panic("set is not tracking gaps")
	}
	gap := s.LowerBoundGap(min)
	if gap.Range().Length() >= minSize {
		return gap
	}
	return gap.NextLargeEnoughGap(minSize)
}

// UpperBoundLargeEnoughGap returns the last gap in the set with at least the
// given length and whose range contains a key less than or equal to max. If no
// such gap exists, UpperBoundLargeEnoughGap returns a terminal iterator.
//
// Precondition: trackGaps must be 1.
func (s *FrameRefSet) UpperBoundLargeEnoughGap(max, minSize uint64) FrameRefGapIterator {
	if FrameReftrackGaps != 1 {
		panic("set is not tracking gaps")
	}
	gap := s.UpperBoundGap(max)
	if gap.Range().Length() >= minSize {
		return gap
	}
	return gap.PrevLargeEnoughGap(minSize)
}

// Insert inserts the given segment into the given gap. If the new segment can
// be merged with adjacent segments, Insert will do so. Insert returns an
// iterator to the segment containing the inserted value (which may have been
// merged with other values). All existing iterators (including gap, but not
// including the returned iterator) are invalidated.
//
// If the gap cannot accommodate the segment, or if r is invalid, Insert panics.
//
// Insert is semantically equivalent to a InsertWithoutMerging followed by a
// Merge, but may be more efficient. Note that there is no unchecked variant of
// Insert since Insert must retrieve and inspect gap's predecessor and
// successor segments regardless.
func (s *FrameRefSet) Insert(gap FrameRefGapIterator, r __generics_imported0.FileRange, val FrameRefSegInfo) FrameRefIterator {
	if r.Length() <= 0 {
		panic(fmt.Sprintf("invalid segment range %v", r))
	}
	prev, next := gap.PrevSegment(), gap.NextSegment()
	if prev.Ok() && prev.End() > r.Start {
		panic(fmt.Sprintf("new segment %v overlaps predecessor %v", r, prev.Range()))
	}
	if next.Ok() && next.Start() < r.End {
		panic(fmt.Sprintf("new segment %v overlaps successor %v", r, next.Range()))
	}
	if prev.Ok() && prev.End() == r.Start {
		if mval, ok := (FrameRefSetFunctions{}).Merge(prev.Range(), prev.Value(), r, val); ok {
			shrinkMaxGap := FrameReftrackGaps != 0 && gap.Range().Length() == gap.node.maxGap.Get()
			prev.SetEndUnchecked(r.End)
			prev.SetValue(mval)
			if shrinkMaxGap {
				gap.node.updateMaxGapLeaf()
			}
			if next.Ok() && next.Start() == r.End {
				val = mval
				if mval, ok := (FrameRefSetFunctions{}).Merge(prev.Range(), val, next.Range(), next.Value()); ok {
					prev.SetEndUnchecked(next.End())
					prev.SetValue(mval)
					return s.Remove(next).PrevSegment()
				}
			}
			return prev
		}
	}
	if next.Ok() && next.Start() == r.End {
		if mval, ok := (FrameRefSetFunctions{}).Merge(r, val, next.Range(), next.Value()); ok {
			shrinkMaxGap := FrameReftrackGaps != 0 && gap.Range().Length() == gap.node.maxGap.Get()
			next.SetStartUnchecked(r.Start)
			next.SetValue(mval)
			if shrinkMaxGap {
				gap.node.updateMaxGapLeaf()
			}
			return next
		}
	}

	return s.InsertWithoutMergingUnchecked(gap, r, val)
}

// InsertWithoutMerging inserts the given segment into the given gap and
// returns an iterator to the inserted segment. All existing iterators
// (including gap, but not including the returned iterator) are invalidated.
//
// If the gap cannot accommodate the segment, or if r is invalid,
// InsertWithoutMerging panics.
func (s *FrameRefSet) InsertWithoutMerging(gap FrameRefGapIterator, r __generics_imported0.FileRange, val FrameRefSegInfo) FrameRefIterator {
	if r.Length() <= 0 {
		panic(fmt.Sprintf("invalid segment range %v", r))
	}
	if gr := gap.Range(); !gr.IsSupersetOf(r) {
		panic(fmt.Sprintf("cannot insert segment range %v into gap range %v", r, gr))
	}
	return s.InsertWithoutMergingUnchecked(gap, r, val)
}

// InsertWithoutMergingUnchecked inserts the given segment into the given gap
// and returns an iterator to the inserted segment. All existing iterators
// (including gap, but not including the returned iterator) are invalidated.
//
// Preconditions:
//   - r.Start >= gap.Start().
//   - r.End <= gap.End().
func (s *FrameRefSet) InsertWithoutMergingUnchecked(gap FrameRefGapIterator, r __generics_imported0.FileRange, val FrameRefSegInfo) FrameRefIterator {
	gap = gap.node.rebalanceBeforeInsert(gap)
	splitMaxGap := FrameReftrackGaps != 0 && (gap.node.nrSegments == 0 || gap.Range().Length() == gap.node.maxGap.Get())
	copy(gap.node.keys[gap.index+1:], gap.node.keys[gap.index:gap.node.nrSegments])
	copy(gap.node.values[gap.index+1:], gap.node.values[gap.index:gap.node.nrSegments])
	gap.node.keys[gap.index] = r
	gap.node.values[gap.index] = val
	gap.node.nrSegments++
	if splitMaxGap {
		gap.node.updateMaxGapLeaf()
	}
	return FrameRefIterator{gap.node, gap.index}
}

// InsertRange inserts the given segment into the set. If the new segment can
// be merged with adjacent segments, InsertRange will do so. InsertRange
// returns an iterator to the segment containing the inserted value (which may
// have been merged with other values). All existing iterators (excluding the
// returned iterator) are invalidated.
//
// If the new segment would overlap an existing segment, or if r is invalid,
// InsertRange panics.
//
// InsertRange searches the set to find the gap to insert into. If the caller
// already has the appropriate GapIterator, or if the caller needs to do
// additional work between finding the gap and insertion, use Insert instead.
func (s *FrameRefSet) InsertRange(r __generics_imported0.FileRange, val FrameRefSegInfo) FrameRefIterator {
	if r.Length() <= 0 {
		panic(fmt.Sprintf("invalid segment range %v", r))
	}
	seg, gap := s.Find(r.Start)
	if seg.Ok() {
		panic(fmt.Sprintf("new segment %v overlaps existing segment %v", r, seg.Range()))
	}
	if gap.End() < r.End {
		panic(fmt.Sprintf("new segment %v overlaps existing segment %v", r, gap.NextSegment().Range()))
	}
	return s.Insert(gap, r, val)
}

// InsertWithoutMergingRange inserts the given segment into the set and returns
// an iterator to the inserted segment. All existing iterators (excluding the
// returned iterator) are invalidated.
//
// If the new segment would overlap an existing segment, or if r is invalid,
// InsertWithoutMergingRange panics.
//
// InsertWithoutMergingRange searches the set to find the gap to insert into.
// If the caller already has the appropriate GapIterator, or if the caller
// needs to do additional work between finding the gap and insertion, use
// InsertWithoutMerging instead.
func (s *FrameRefSet) InsertWithoutMergingRange(r __generics_imported0.FileRange, val FrameRefSegInfo) FrameRefIterator {
	if r.Length() <= 0 {
		panic(fmt.Sprintf("invalid segment range %v", r))
	}
	seg, gap := s.Find(r.Start)
	if seg.Ok() {
		panic(fmt.Sprintf("new segment %v overlaps existing segment %v", r, seg.Range()))
	}
	if gap.End() < r.End {
		panic(fmt.Sprintf("new segment %v overlaps existing segment %v", r, gap.NextSegment().Range()))
	}
	return s.InsertWithoutMerging(gap, r, val)
}

// TryInsertRange attempts to insert the given segment into the set. If the new
// segment can be merged with adjacent segments, TryInsertRange will do so.
// TryInsertRange returns an iterator to the segment containing the inserted
// value (which may have been merged with other values). All existing iterators
// (excluding the returned iterator) are invalidated.
//
// If the new segment would overlap an existing segment, TryInsertRange does
// nothing and returns a terminal iterator.
//
// TryInsertRange searches the set to find the gap to insert into. If the
// caller already has the appropriate GapIterator, or if the caller needs to do
// additional work between finding the gap and insertion, use Insert instead.
func (s *FrameRefSet) TryInsertRange(r __generics_imported0.FileRange, val FrameRefSegInfo) FrameRefIterator {
	if r.Length() <= 0 {
		panic(fmt.Sprintf("invalid segment range %v", r))
	}
	seg, gap := s.Find(r.Start)
	if seg.Ok() {
		return FrameRefIterator{}
	}
	if gap.End() < r.End {
		return FrameRefIterator{}
	}
	return s.Insert(gap, r, val)
}

// TryInsertWithoutMergingRange attempts to insert the given segment into the
// set. If successful, it returns an iterator to the inserted segment; all
// existing iterators (excluding the returned iterator) are invalidated. If the
// new segment would overlap an existing segment, TryInsertWithoutMergingRange
// does nothing and returns a terminal iterator.
//
// TryInsertWithoutMergingRange searches the set to find the gap to insert
// into. If the caller already has the appropriate GapIterator, or if the
// caller needs to do additional work between finding the gap and insertion,
// use InsertWithoutMerging instead.
func (s *FrameRefSet) TryInsertWithoutMergingRange(r __generics_imported0.FileRange, val FrameRefSegInfo) FrameRefIterator {
	if r.Length() <= 0 {
		panic(fmt.Sprintf("invalid segment range %v", r))
	}
	seg, gap := s.Find(r.Start)
	if seg.Ok() {
		return FrameRefIterator{}
	}
	if gap.End() < r.End {
		return FrameRefIterator{}
	}
	return s.InsertWithoutMerging(gap, r, val)
}

// Remove removes the given segment and returns an iterator to the vacated gap.
// All existing iterators (including seg, but not including the returned
// iterator) are invalidated.
func (s *FrameRefSet) Remove(seg FrameRefIterator) FrameRefGapIterator {

	if seg.node.hasChildren {

		victim := seg.PrevSegment()

		seg.SetRangeUnchecked(victim.Range())
		seg.SetValue(victim.Value())

		nextAdjacentNode := seg.NextSegment().node
		if FrameReftrackGaps != 0 {
			nextAdjacentNode.updateMaxGapLeaf()
		}
		return s.Remove(victim).NextGap()
	}
	copy(seg.node.keys[seg.index:], seg.node.keys[seg.index+1:seg.node.nrSegments])
	copy(seg.node.values[seg.index:], seg.node.values[seg.index+1:seg.node.nrSegments])
	FrameRefSetFunctions{}.ClearValue(&seg.node.values[seg.node.nrSegments-1])
	seg.node.nrSegments--
	if FrameReftrackGaps != 0 {
		seg.node.updateMaxGapLeaf()
	}
	return seg.node.rebalanceAfterRemove(FrameRefGapIterator{seg.node, seg.index})
}

// RemoveAll removes all segments from the set. All existing iterators are
// invalidated.
func (s *FrameRefSet) RemoveAll() {
	s.root = FrameRefnode{}
}

// RemoveRange removes all segments in the given range. An iterator to the
// newly formed gap is returned, and all existing iterators are invalidated.
//
// RemoveRange searches the set to find segments to remove. If the caller
// already has an iterator to either end of the range of segments to remove, or
// if the caller needs to do additional work before removing each segment,
// iterate segments and call Remove in a loop instead.
func (s *FrameRefSet) RemoveRange(r __generics_imported0.FileRange) FrameRefGapIterator {
	return s.RemoveRangeWith(r, nil)
}

// RemoveFullRange is equivalent to RemoveRange, except that if any key in the
// given range does not correspond to a segment, RemoveFullRange panics.
func (s *FrameRefSet) RemoveFullRange(r __generics_imported0.FileRange) FrameRefGapIterator {
	return s.RemoveFullRangeWith(r, nil)
}

// RemoveRangeWith removes all segments in the given range. An iterator to the
// newly formed gap is returned, and all existing iterators are invalidated.
//
// The function f is applied to each segment immediately before it is removed,
// in order of ascending keys. Segments that lie partially outside r are split
// before f is called, such that f only observes segments entirely within r.
// Non-empty gaps between segments are skipped.
//
// RemoveRangeWith searches the set to find segments to remove. If the caller
// already has an iterator to either end of the range of segments to remove, or
// if the caller needs to do additional work before removing each segment,
// iterate segments and call Remove in a loop instead.
//
// N.B. f must not invalidate iterators into s.
func (s *FrameRefSet) RemoveRangeWith(r __generics_imported0.FileRange, f func(seg FrameRefIterator)) FrameRefGapIterator {
	seg, gap := s.Find(r.Start)
	if seg.Ok() {
		seg = s.Isolate(seg, r)
		if f != nil {
			f(seg)
		}
		gap = s.Remove(seg)
	}
	for seg = gap.NextSegment(); seg.Ok() && seg.Start() < r.End; seg = gap.NextSegment() {
		seg = s.SplitAfter(seg, r.End)
		if f != nil {
			f(seg)
		}
		gap = s.Remove(seg)
	}
	return gap
}

// RemoveFullRangeWith is equivalent to RemoveRangeWith, except that if any key
// in the given range does not correspond to a segment, RemoveFullRangeWith
// panics.
func (s *FrameRefSet) RemoveFullRangeWith(r __generics_imported0.FileRange, f func(seg FrameRefIterator)) FrameRefGapIterator {
	seg := s.FindSegment(r.Start)
	if !seg.Ok() {
		panic(fmt.Sprintf("missing segment at %v", r.Start))
	}
	seg = s.SplitBefore(seg, r.Start)
	for {
		seg = s.SplitAfter(seg, r.End)
		if f != nil {
			f(seg)
		}
		end := seg.End()
		gap := s.Remove(seg)
		if r.End <= end {
			return gap
		}
		seg = gap.NextSegment()
		if !seg.Ok() || seg.Start() != end {
			panic(fmt.Sprintf("missing segment at %v", end))
		}
	}
}

// Merge attempts to merge two neighboring segments. If successful, Merge
// returns an iterator to the merged segment, and all existing iterators are
// invalidated. Otherwise, Merge returns a terminal iterator.
//
// If first is not the predecessor of second, Merge panics.
func (s *FrameRefSet) Merge(first, second FrameRefIterator) FrameRefIterator {
	if first.NextSegment() != second {
		panic(fmt.Sprintf("attempt to merge non-neighboring segments %v, %v", first.Range(), second.Range()))
	}
	return s.MergeUnchecked(first, second)
}

// MergeUnchecked attempts to merge two neighboring segments. If successful,
// MergeUnchecked returns an iterator to the merged segment, and all existing
// iterators are invalidated. Otherwise, MergeUnchecked returns a terminal
// iterator.
//
// Precondition: first is the predecessor of second: first.NextSegment() ==
// second, first == second.PrevSegment().
func (s *FrameRefSet) MergeUnchecked(first, second FrameRefIterator) FrameRefIterator {
	if first.End() == second.Start() {
		if mval, ok := (FrameRefSetFunctions{}).Merge(first.Range(), first.Value(), second.Range(), second.Value()); ok {

			first.SetEndUnchecked(second.End())
			first.SetValue(mval)

			return s.Remove(second).PrevSegment()
		}
	}
	return FrameRefIterator{}
}

// MergePrev attempts to merge the given segment with its predecessor if
// possible, and returns an updated iterator to the extended segment. All
// existing iterators (including seg, but not including the returned iterator)
// are invalidated.
//
// MergePrev is usually used when mutating segments while iterating them in
// order of increasing keys, to attempt merging of each mutated segment with
// its previously-mutated predecessor. In such cases, merging a mutated segment
// with its unmutated successor would incorrectly cause the latter to be
// skipped.
func (s *FrameRefSet) MergePrev(seg FrameRefIterator) FrameRefIterator {
	if prev := seg.PrevSegment(); prev.Ok() {
		if mseg := s.MergeUnchecked(prev, seg); mseg.Ok() {
			seg = mseg
		}
	}
	return seg
}

// MergeNext attempts to merge the given segment with its successor if
// possible, and returns an updated iterator to the extended segment. All
// existing iterators (including seg, but not including the returned iterator)
// are invalidated.
//
// MergeNext is usually used when mutating segments while iterating them in
// order of decreasing keys, to attempt merging of each mutated segment with
// its previously-mutated successor. In such cases, merging a mutated segment
// with its unmutated predecessor would incorrectly cause the latter to be
// skipped.
func (s *FrameRefSet) MergeNext(seg FrameRefIterator) FrameRefIterator {
	if next := seg.NextSegment(); next.Ok() {
		if mseg := s.MergeUnchecked(seg, next); mseg.Ok() {
			seg = mseg
		}
	}
	return seg
}

// Unisolate attempts to merge the given segment with its predecessor and
// successor if possible, and returns an updated iterator to the extended
// segment. All existing iterators (including seg, but not including the
// returned iterator) are invalidated.
//
// Unisolate is usually used in conjunction with Isolate when mutating part of
// a single segment in a way that may affect its mergeability. For the reasons
// described by MergePrev and MergeNext, it is usually incorrect to use the
// return value of Unisolate in a loop variable.
func (s *FrameRefSet) Unisolate(seg FrameRefIterator) FrameRefIterator {
	if prev := seg.PrevSegment(); prev.Ok() {
		if mseg := s.MergeUnchecked(prev, seg); mseg.Ok() {
			seg = mseg
		}
	}
	if next := seg.NextSegment(); next.Ok() {
		if mseg := s.MergeUnchecked(seg, next); mseg.Ok() {
			seg = mseg
		}
	}
	return seg
}

// MergeAll merges all mergeable adjacent segments in the set. All existing
// iterators are invalidated.
func (s *FrameRefSet) MergeAll() {
	seg := s.FirstSegment()
	if !seg.Ok() {
		return
	}
	next := seg.NextSegment()
	for next.Ok() {
		if mseg := s.MergeUnchecked(seg, next); mseg.Ok() {
			seg, next = mseg, mseg.NextSegment()
		} else {
			seg, next = next, next.NextSegment()
		}
	}
}

// MergeInsideRange attempts to merge all adjacent segments that contain a key
// in the specific range. All existing iterators are invalidated.
//
// MergeInsideRange only makes sense after mutating the set in a way that may
// change the mergeability of modified segments; callers should prefer to use
// MergePrev or MergeNext during the mutating loop instead (depending on the
// direction of iteration), in order to avoid a redundant search.
func (s *FrameRefSet) MergeInsideRange(r __generics_imported0.FileRange) {
	seg := s.LowerBoundSegment(r.Start)
	if !seg.Ok() {
		return
	}
	next := seg.NextSegment()
	for next.Ok() && next.Start() < r.End {
		if mseg := s.MergeUnchecked(seg, next); mseg.Ok() {
			seg, next = mseg, mseg.NextSegment()
		} else {
			seg, next = next, next.NextSegment()
		}
	}
}

// MergeOutsideRange attempts to merge the segment containing r.Start with its
// predecessor, and the segment containing r.End-1 with its successor.
//
// MergeOutsideRange only makes sense after mutating the set in a way that may
// change the mergeability of modified segments; callers should prefer to use
// MergePrev or MergeNext during the mutating loop instead (depending on the
// direction of iteration), in order to avoid two redundant searches.
func (s *FrameRefSet) MergeOutsideRange(r __generics_imported0.FileRange) {
	first := s.FindSegment(r.Start)
	if first.Ok() {
		if prev := first.PrevSegment(); prev.Ok() {
			s.Merge(prev, first)
		}
	}
	last := s.FindSegment(r.End - 1)
	if last.Ok() {
		if next := last.NextSegment(); next.Ok() {
			s.Merge(last, next)
		}
	}
}

// Split splits the given segment at the given key and returns iterators to the
// two resulting segments. All existing iterators (including seg, but not
// including the returned iterators) are invalidated.
//
// If the segment cannot be split at split (because split is at the start or
// end of the segment's range, so splitting would produce a segment with zero
// length, or because split falls outside the segment's range altogether),
// Split panics.
func (s *FrameRefSet) Split(seg FrameRefIterator, split uint64) (FrameRefIterator, FrameRefIterator) {
	if !seg.Range().CanSplitAt(split) {
		panic(fmt.Sprintf("can't split %v at %v", seg.Range(), split))
	}
	return s.SplitUnchecked(seg, split)
}

// SplitUnchecked splits the given segment at the given key and returns
// iterators to the two resulting segments. All existing iterators (including
// seg, but not including the returned iterators) are invalidated.
//
// Preconditions: seg.Start() < key < seg.End().
func (s *FrameRefSet) SplitUnchecked(seg FrameRefIterator, split uint64) (FrameRefIterator, FrameRefIterator) {
	val1, val2 := (FrameRefSetFunctions{}).Split(seg.Range(), seg.Value(), split)
	end2 := seg.End()
	seg.SetEndUnchecked(split)
	seg.SetValue(val1)
	seg2 := s.InsertWithoutMergingUnchecked(seg.NextGap(), __generics_imported0.FileRange{split, end2}, val2)

	return seg2.PrevSegment(), seg2
}

// SplitBefore ensures that the given segment's start is at least start by
// splitting at start if necessary, and returns an updated iterator to the
// bounded segment. All existing iterators (including seg, but not including
// the returned iterator) are invalidated.
//
// SplitBefore is usually when mutating segments in a range. In such cases,
// when iterating segments in order of increasing keys, the first segment may
// extend beyond the start of the range to be mutated, and needs to be
// SplitBefore to ensure that only the part of the segment within the range is
// mutated. When iterating segments in order of decreasing keys, SplitBefore
// and SplitAfter; i.e. SplitBefore needs to be invoked on each segment, while
// SplitAfter only needs to be invoked on the first.
//
// Preconditions: start < seg.End().
func (s *FrameRefSet) SplitBefore(seg FrameRefIterator, start uint64) FrameRefIterator {
	if seg.Range().CanSplitAt(start) {
		_, seg = s.SplitUnchecked(seg, start)
	}
	return seg
}

// SplitAfter ensures that the given segment's end is at most end by splitting
// at end if necessary, and returns an updated iterator to the bounded segment.
// All existing iterators (including seg, but not including the returned
// iterator) are invalidated.
//
// SplitAfter is usually used when mutating segments in a range. In such cases,
// when iterating segments in order of increasing keys, each iterated segment
// may extend beyond the end of the range to be mutated, and needs to be
// SplitAfter to ensure that only the part of the segment within the range is
// mutated. When iterating segments in order of decreasing keys, SplitBefore
// and SplitAfter exchange roles; i.e. SplitBefore needs to be invoked on each
// segment, while SplitAfter only needs to be invoked on the first.
//
// Preconditions: seg.Start() < end.
func (s *FrameRefSet) SplitAfter(seg FrameRefIterator, end uint64) FrameRefIterator {
	if seg.Range().CanSplitAt(end) {
		seg, _ = s.SplitUnchecked(seg, end)
	}
	return seg
}

// Isolate ensures that the given segment's range is a subset of r by splitting
// at r.Start and r.End if necessary, and returns an updated iterator to the
// bounded segment. All existing iterators (including seg, but not including
// the returned iterators) are invalidated.
//
// Isolate is usually used when mutating part of a single segment, or when
// mutating segments in a range where the first segment is not necessarily
// split, making use of SplitBefore/SplitAfter complex.
//
// Preconditions: seg.Range().Overlaps(r).
func (s *FrameRefSet) Isolate(seg FrameRefIterator, r __generics_imported0.FileRange) FrameRefIterator {
	if seg.Range().CanSplitAt(r.Start) {
		_, seg = s.SplitUnchecked(seg, r.Start)
	}
	if seg.Range().CanSplitAt(r.End) {
		seg, _ = s.SplitUnchecked(seg, r.End)
	}
	return seg
}

// LowerBoundSegmentSplitBefore combines LowerBoundSegment and SplitBefore.
//
// LowerBoundSegmentSplitBefore is usually used when mutating segments in a
// range while iterating them in order of increasing keys. In such cases,
// LowerBoundSegmentSplitBefore provides an iterator to the first segment to be
// mutated, suitable as the initial value for a loop variable.
func (s *FrameRefSet) LowerBoundSegmentSplitBefore(min uint64) FrameRefIterator {
	seg, gap := s.Find(min)
	if seg.Ok() {
		return s.SplitBefore(seg, min)
	}
	return gap.NextSegment()
}

// UpperBoundSegmentSplitAfter combines UpperBoundSegment and SplitAfter.
//
// UpperBoundSegmentSplitAfter is usually used when mutating segments in a
// range while iterating them in order of decreasing keys. In such cases,
// UpperBoundSegmentSplitAfter provides an iterator to the first segment to be
// mutated, suitable as the initial value for a loop variable.
func (s *FrameRefSet) UpperBoundSegmentSplitAfter(max uint64) FrameRefIterator {
	seg, gap := s.Find(max)
	if seg.Ok() {
		return s.SplitAfter(seg, max)
	}
	return gap.PrevSegment()
}

// VisitRange applies the function f to all segments intersecting the range r,
// in order of ascending keys. Segments will not be split, so f may be called
// on segments lying partially outside r. Non-empty gaps between segments are
// skipped. If a call to f returns false, VisitRange stops iteration
// immediately.
//
// N.B. f must not invalidate iterators into s.
func (s *FrameRefSet) VisitRange(r __generics_imported0.FileRange, f func(seg FrameRefIterator) bool) {
	for seg := s.LowerBoundSegment(r.Start); seg.Ok() && seg.Start() < r.End; seg = seg.NextSegment() {
		if !f(seg) {
			return
		}
	}
}

// VisitFullRange is equivalent to VisitRange, except that if any key in r that
// is visited before f returns false does not correspond to a segment,
// VisitFullRange panics.
func (s *FrameRefSet) VisitFullRange(r __generics_imported0.FileRange, f func(seg FrameRefIterator) bool) {
	pos := r.Start
	seg := s.FindSegment(r.Start)
	for {
		if !seg.Ok() {
			panic(fmt.Sprintf("missing segment at %v", pos))
		}
		if !f(seg) {
			return
		}
		pos = seg.End()
		if r.End <= pos {
			return
		}
		seg, _ = seg.NextNonEmpty()
	}
}

// MutateRange applies the function f to all segments intersecting the range r,
// in order of ascending keys. Segments that lie partially outside r are split
// before f is called, such that f only observes segments entirely within r.
// Iterated segments are merged again after f is called. Non-empty gaps between
// segments are skipped. If a call to f returns false, MutateRange stops
// iteration immediately.
//
// MutateRange invalidates all existing iterators.
//
// N.B. f must not invalidate iterators into s.
func (s *FrameRefSet) MutateRange(r __generics_imported0.FileRange, f func(seg FrameRefIterator) bool) {
	seg := s.LowerBoundSegmentSplitBefore(r.Start)
	for seg.Ok() && seg.Start() < r.End {
		seg = s.SplitAfter(seg, r.End)
		cont := f(seg)
		seg = s.MergePrev(seg)
		if !cont {
			s.MergeNext(seg)
			return
		}
		seg = seg.NextSegment()
	}
	if seg.Ok() {
		s.MergePrev(seg)
	}
}

// MutateFullRange is equivalent to MutateRange, except that if any key in r
// that is visited before f returns false does not correspond to a segment,
// MutateFullRange panics.
func (s *FrameRefSet) MutateFullRange(r __generics_imported0.FileRange, f func(seg FrameRefIterator) bool) {
	seg := s.FindSegment(r.Start)
	if !seg.Ok() {
		panic(fmt.Sprintf("missing segment at %v", r.Start))
	}
	seg = s.SplitBefore(seg, r.Start)
	for {
		seg = s.SplitAfter(seg, r.End)
		cont := f(seg)
		end := seg.End()
		seg = s.MergePrev(seg)
		if !cont || r.End <= end {
			s.MergeNext(seg)
			return
		}
		seg = seg.NextSegment()
		if !seg.Ok() || seg.Start() != end {
			panic(fmt.Sprintf("missing segment at %v", end))
		}
	}
}

// +stateify savable
type FrameRefnode struct {
	// An internal binary tree node looks like:
	//
	//   K
	//  / \
	// Cl Cr
	//
	// where all keys in the subtree rooted by Cl (the left subtree) are less
	// than K (the key of the parent node), and all keys in the subtree rooted
	// by Cr (the right subtree) are greater than K.
	//
	// An internal B-tree node's indexes work out to look like:
	//
	//   K0 K1 K2  ...   Kn-1
	//  / \/ \/ \  ...  /  \
	// C0 C1 C2 C3 ... Cn-1 Cn
	//
	// where n is nrSegments.
	nrSegments int

	// parent is a pointer to this node's parent. If this node is root, parent
	// is nil.
	parent *FrameRefnode

	// parentIndex is the index of this node in parent.children.
	parentIndex int

	// Flag for internal nodes that is technically redundant with "children[0]
	// != nil", but is stored in the first cache line. "hasChildren" rather
	// than "isLeaf" because false must be the correct value for an empty root.
	hasChildren bool

	// The longest gap within this node. If the node is a leaf, it's simply the
	// maximum gap among all the (nrSegments+1) gaps formed by its nrSegments keys
	// including the 0th and nrSegments-th gap possibly shared with its upper-level
	// nodes; if it's a non-leaf node, it's the max of all children's maxGap.
	maxGap FrameRefdynamicGap

	// Nodes store keys and values in separate arrays to maximize locality in
	// the common case (scanning keys for lookup).
	keys     [FrameRefmaxDegree - 1]__generics_imported0.FileRange
	values   [FrameRefmaxDegree - 1]FrameRefSegInfo
	children [FrameRefmaxDegree]*FrameRefnode
}

// firstSegment returns the first segment in the subtree rooted by n.
//
// Preconditions: n.nrSegments != 0.
func (n *FrameRefnode) firstSegment() FrameRefIterator {
	for n.hasChildren {
		n = n.children[0]
	}
	return FrameRefIterator{n, 0}
}

// lastSegment returns the last segment in the subtree rooted by n.
//
// Preconditions: n.nrSegments != 0.
func (n *FrameRefnode) lastSegment() FrameRefIterator {
	for n.hasChildren {
		n = n.children[n.nrSegments]
	}
	return FrameRefIterator{n, n.nrSegments - 1}
}

func (n *FrameRefnode) prevSibling() *FrameRefnode {
	if n.parent == nil || n.parentIndex == 0 {
		return nil
	}
	return n.parent.children[n.parentIndex-1]
}

func (n *FrameRefnode) nextSibling() *FrameRefnode {
	if n.parent == nil || n.parentIndex == n.parent.nrSegments {
		return nil
	}
	return n.parent.children[n.parentIndex+1]
}

// rebalanceBeforeInsert splits n and its ancestors if they are full, as
// required for insertion, and returns an updated iterator to the position
// represented by gap.
func (n *FrameRefnode) rebalanceBeforeInsert(gap FrameRefGapIterator) FrameRefGapIterator {
	if n.nrSegments < FrameRefmaxDegree-1 {
		return gap
	}
	if n.parent != nil {
		gap = n.parent.rebalanceBeforeInsert(gap)
	}
	if n.parent == nil {

		left := &FrameRefnode{
			nrSegments:  FrameRefminDegree - 1,
			parent:      n,
			parentIndex: 0,
			hasChildren: n.hasChildren,
		}
		right := &FrameRefnode{
			nrSegments:  FrameRefminDegree - 1,
			parent:      n,
			parentIndex: 1,
			hasChildren: n.hasChildren,
		}
		copy(left.keys[:FrameRefminDegree-1], n.keys[:FrameRefminDegree-1])
		copy(left.values[:FrameRefminDegree-1], n.values[:FrameRefminDegree-1])
		copy(right.keys[:FrameRefminDegree-1], n.keys[FrameRefminDegree:])
		copy(right.values[:FrameRefminDegree-1], n.values[FrameRefminDegree:])
		n.keys[0], n.values[0] = n.keys[FrameRefminDegree-1], n.values[FrameRefminDegree-1]
		FrameRefzeroValueSlice(n.values[1:])
		if n.hasChildren {
			copy(left.children[:FrameRefminDegree], n.children[:FrameRefminDegree])
			copy(right.children[:FrameRefminDegree], n.children[FrameRefminDegree:])
			FrameRefzeroNodeSlice(n.children[2:])
			for i := 0; i < FrameRefminDegree; i++ {
				left.children[i].parent = left
				left.children[i].parentIndex = i
				right.children[i].parent = right
				right.children[i].parentIndex = i
			}
		}
		n.nrSegments = 1
		n.hasChildren = true
		n.children[0] = left
		n.children[1] = right

		if FrameReftrackGaps != 0 {
			left.updateMaxGapLocal()
			right.updateMaxGapLocal()
		}
		if gap.node != n {
			return gap
		}
		if gap.index < FrameRefminDegree {
			return FrameRefGapIterator{left, gap.index}
		}
		return FrameRefGapIterator{right, gap.index - FrameRefminDegree}
	}

	copy(n.parent.keys[n.parentIndex+1:], n.parent.keys[n.parentIndex:n.parent.nrSegments])
	copy(n.parent.values[n.parentIndex+1:], n.parent.values[n.parentIndex:n.parent.nrSegments])
	n.parent.keys[n.parentIndex], n.parent.values[n.parentIndex] = n.keys[FrameRefminDegree-1], n.values[FrameRefminDegree-1]
	copy(n.parent.children[n.parentIndex+2:], n.parent.children[n.parentIndex+1:n.parent.nrSegments+1])
	for i := n.parentIndex + 2; i < n.parent.nrSegments+2; i++ {
		n.parent.children[i].parentIndex = i
	}
	sibling := &FrameRefnode{
		nrSegments:  FrameRefminDegree - 1,
		parent:      n.parent,
		parentIndex: n.parentIndex + 1,
		hasChildren: n.hasChildren,
	}
	n.parent.children[n.parentIndex+1] = sibling
	n.parent.nrSegments++
	copy(sibling.keys[:FrameRefminDegree-1], n.keys[FrameRefminDegree:])
	copy(sibling.values[:FrameRefminDegree-1], n.values[FrameRefminDegree:])
	FrameRefzeroValueSlice(n.values[FrameRefminDegree-1:])
	if n.hasChildren {
		copy(sibling.children[:FrameRefminDegree], n.children[FrameRefminDegree:])
		FrameRefzeroNodeSlice(n.children[FrameRefminDegree:])
		for i := 0; i < FrameRefminDegree; i++ {
			sibling.children[i].parent = sibling
			sibling.children[i].parentIndex = i
		}
	}
	n.nrSegments = FrameRefminDegree - 1

	if FrameReftrackGaps != 0 {
		n.updateMaxGapLocal()
		sibling.updateMaxGapLocal()
	}

	if gap.node != n {
		return gap
	}
	if gap.index < FrameRefminDegree {
		return gap
	}
	return FrameRefGapIterator{sibling, gap.index - FrameRefminDegree}
}

// rebalanceAfterRemove "unsplits" n and its ancestors if they are deficient
// (contain fewer segments than required by B-tree invariants), as required for
// removal, and returns an updated iterator to the position represented by gap.
//
// Precondition: n is the only node in the tree that may currently violate a
// B-tree invariant.
func (n *FrameRefnode) rebalanceAfterRemove(gap FrameRefGapIterator) FrameRefGapIterator {
	for {
		if n.nrSegments >= FrameRefminDegree-1 {
			return gap
		}
		if n.parent == nil {

			return gap
		}

		if sibling := n.prevSibling(); sibling != nil && sibling.nrSegments >= FrameRefminDegree {
			copy(n.keys[1:], n.keys[:n.nrSegments])
			copy(n.values[1:], n.values[:n.nrSegments])
			n.keys[0] = n.parent.keys[n.parentIndex-1]
			n.values[0] = n.parent.values[n.parentIndex-1]
			n.parent.keys[n.parentIndex-1] = sibling.keys[sibling.nrSegments-1]
			n.parent.values[n.parentIndex-1] = sibling.values[sibling.nrSegments-1]
			FrameRefSetFunctions{}.ClearValue(&sibling.values[sibling.nrSegments-1])
			if n.hasChildren {
				copy(n.children[1:], n.children[:n.nrSegments+1])
				n.children[0] = sibling.children[sibling.nrSegments]
				sibling.children[sibling.nrSegments] = nil
				n.children[0].parent = n
				n.children[0].parentIndex = 0
				for i := 1; i < n.nrSegments+2; i++ {
					n.children[i].parentIndex = i
				}
			}
			n.nrSegments++
			sibling.nrSegments--

			if FrameReftrackGaps != 0 {
				n.updateMaxGapLocal()
				sibling.updateMaxGapLocal()
			}
			if gap.node == sibling && gap.index == sibling.nrSegments {
				return FrameRefGapIterator{n, 0}
			}
			if gap.node == n {
				return FrameRefGapIterator{n, gap.index + 1}
			}
			return gap
		}
		if sibling := n.nextSibling(); sibling != nil && sibling.nrSegments >= FrameRefminDegree {
			n.keys[n.nrSegments] = n.parent.keys[n.parentIndex]
			n.values[n.nrSegments] = n.parent.values[n.parentIndex]
			n.parent.keys[n.parentIndex] = sibling.keys[0]
			n.parent.values[n.parentIndex] = sibling.values[0]
			copy(sibling.keys[:sibling.nrSegments-1], sibling.keys[1:])
			copy(sibling.values[:sibling.nrSegments-1], sibling.values[1:])
			FrameRefSetFunctions{}.ClearValue(&sibling.values[sibling.nrSegments-1])
			if n.hasChildren {
				n.children[n.nrSegments+1] = sibling.children[0]
				copy(sibling.children[:sibling.nrSegments], sibling.children[1:])
				sibling.children[sibling.nrSegments] = nil
				n.children[n.nrSegments+1].parent = n
				n.children[n.nrSegments+1].parentIndex = n.nrSegments + 1
				for i := 0; i < sibling.nrSegments; i++ {
					sibling.children[i].parentIndex = i
				}
			}
			n.nrSegments++
			sibling.nrSegments--

			if FrameReftrackGaps != 0 {
				n.updateMaxGapLocal()
				sibling.updateMaxGapLocal()
			}
			if gap.node == sibling {
				if gap.index == 0 {
					return FrameRefGapIterator{n, n.nrSegments}
				}
				return FrameRefGapIterator{sibling, gap.index - 1}
			}
			return gap
		}

		p := n.parent
		if p.nrSegments == 1 {

			left, right := p.children[0], p.children[1]
			p.nrSegments = left.nrSegments + right.nrSegments + 1
			p.hasChildren = left.hasChildren
			p.keys[left.nrSegments] = p.keys[0]
			p.values[left.nrSegments] = p.values[0]
			copy(p.keys[:left.nrSegments], left.keys[:left.nrSegments])
			copy(p.values[:left.nrSegments], left.values[:left.nrSegments])
			copy(p.keys[left.nrSegments+1:], right.keys[:right.nrSegments])
			copy(p.values[left.nrSegments+1:], right.values[:right.nrSegments])
			if left.hasChildren {
				copy(p.children[:left.nrSegments+1], left.children[:left.nrSegments+1])
				copy(p.children[left.nrSegments+1:], right.children[:right.nrSegments+1])
				for i := 0; i < p.nrSegments+1; i++ {
					p.children[i].parent = p
					p.children[i].parentIndex = i
				}
			} else {
				p.children[0] = nil
				p.children[1] = nil
			}

			if gap.node == left {
				return FrameRefGapIterator{p, gap.index}
			}
			if gap.node == right {
				return FrameRefGapIterator{p, gap.index + left.nrSegments + 1}
			}
			return gap
		}
		// Merge n and either sibling, along with the segment separating the
		// two, into whichever of the two nodes comes first. This is the
		// reverse of the non-root splitting case in
		// node.rebalanceBeforeInsert.
		var left, right *FrameRefnode
		if n.parentIndex > 0 {
			left = n.prevSibling()
			right = n
		} else {
			left = n
			right = n.nextSibling()
		}

		if gap.node == right {
			gap = FrameRefGapIterator{left, gap.index + left.nrSegments + 1}
		}
		left.keys[left.nrSegments] = p.keys[left.parentIndex]
		left.values[left.nrSegments] = p.values[left.parentIndex]
		copy(left.keys[left.nrSegments+1:], right.keys[:right.nrSegments])
		copy(left.values[left.nrSegments+1:], right.values[:right.nrSegments])
		if left.hasChildren {
			copy(left.children[left.nrSegments+1:], right.children[:right.nrSegments+1])
			for i := left.nrSegments + 1; i < left.nrSegments+right.nrSegments+2; i++ {
				left.children[i].parent = left
				left.children[i].parentIndex = i
			}
		}
		left.nrSegments += right.nrSegments + 1
		copy(p.keys[left.parentIndex:], p.keys[left.parentIndex+1:p.nrSegments])
		copy(p.values[left.parentIndex:], p.values[left.parentIndex+1:p.nrSegments])
		FrameRefSetFunctions{}.ClearValue(&p.values[p.nrSegments-1])
		copy(p.children[left.parentIndex+1:], p.children[left.parentIndex+2:p.nrSegments+1])
		for i := 0; i < p.nrSegments; i++ {
			p.children[i].parentIndex = i
		}
		p.children[p.nrSegments] = nil
		p.nrSegments--

		if FrameReftrackGaps != 0 {
			left.updateMaxGapLocal()
		}

		n = p
	}
}

// updateMaxGapLeaf updates maxGap bottom-up from the calling leaf until no
// necessary update.
//
// Preconditions: n must be a leaf node, trackGaps must be 1.
func (n *FrameRefnode) updateMaxGapLeaf() {
	if n.hasChildren {
		panic(fmt.Sprintf("updateMaxGapLeaf should always be called on leaf node: %v", n))
	}
	max := n.calculateMaxGapLeaf()
	if max == n.maxGap.Get() {

		return
	}
	oldMax := n.maxGap.Get()
	n.maxGap.Set(max)
	if max > oldMax {

		for p := n.parent; p != nil; p = p.parent {
			if p.maxGap.Get() >= max {

				break
			}

			p.maxGap.Set(max)
		}
		return
	}

	for p := n.parent; p != nil; p = p.parent {
		if p.maxGap.Get() > oldMax {

			break
		}

		parentNewMax := p.calculateMaxGapInternal()
		if p.maxGap.Get() == parentNewMax {

			break
		}

		p.maxGap.Set(parentNewMax)
	}
}

// updateMaxGapLocal updates maxGap of the calling node solely with no
// propagation to ancestor nodes.
//
// Precondition: trackGaps must be 1.
func (n *FrameRefnode) updateMaxGapLocal() {
	if !n.hasChildren {

		n.maxGap.Set(n.calculateMaxGapLeaf())
	} else {

		n.maxGap.Set(n.calculateMaxGapInternal())
	}
}

// calculateMaxGapLeaf iterates the gaps within a leaf node and calculate the
// max.
//
// Preconditions: n must be a leaf node.
func (n *FrameRefnode) calculateMaxGapLeaf() uint64 {
	max := FrameRefGapIterator{n, 0}.Range().Length()
	for i := 1; i <= n.nrSegments; i++ {
		if current := (FrameRefGapIterator{n, i}).Range().Length(); current > max {
			max = current
		}
	}
	return max
}

// calculateMaxGapInternal iterates children's maxGap within an internal node n
// and calculate the max.
//
// Preconditions: n must be a non-leaf node.
func (n *FrameRefnode) calculateMaxGapInternal() uint64 {
	max := n.children[0].maxGap.Get()
	for i := 1; i <= n.nrSegments; i++ {
		if current := n.children[i].maxGap.Get(); current > max {
			max = current
		}
	}
	return max
}

// searchFirstLargeEnoughGap returns the first gap having at least minSize length
// in the subtree rooted by n. If not found, return a terminal gap iterator.
func (n *FrameRefnode) searchFirstLargeEnoughGap(minSize uint64) FrameRefGapIterator {
	if n.maxGap.Get() < minSize {
		return FrameRefGapIterator{}
	}
	if n.hasChildren {
		for i := 0; i <= n.nrSegments; i++ {
			if largeEnoughGap := n.children[i].searchFirstLargeEnoughGap(minSize); largeEnoughGap.Ok() {
				return largeEnoughGap
			}
		}
	} else {
		for i := 0; i <= n.nrSegments; i++ {
			currentGap := FrameRefGapIterator{n, i}
			if currentGap.Range().Length() >= minSize {
				return currentGap
			}
		}
	}
	panic(fmt.Sprintf("invalid maxGap in %v", n))
}

// searchLastLargeEnoughGap returns the last gap having at least minSize length
// in the subtree rooted by n. If not found, return a terminal gap iterator.
func (n *FrameRefnode) searchLastLargeEnoughGap(minSize uint64) FrameRefGapIterator {
	if n.maxGap.Get() < minSize {
		return FrameRefGapIterator{}
	}
	if n.hasChildren {
		for i := n.nrSegments; i >= 0; i-- {
			if largeEnoughGap := n.children[i].searchLastLargeEnoughGap(minSize); largeEnoughGap.Ok() {
				return largeEnoughGap
			}
		}
	} else {
		for i := n.nrSegments; i >= 0; i-- {
			currentGap := FrameRefGapIterator{n, i}
			if currentGap.Range().Length() >= minSize {
				return currentGap
			}
		}
	}
	panic(fmt.Sprintf("invalid maxGap in %v", n))
}

// A Iterator is conceptually one of:
//
//   - A pointer to a segment in a set; or
//
//   - A terminal iterator, which is a sentinel indicating that the end of
//     iteration has been reached.
//
// Iterators are copyable values and are meaningfully equality-comparable. The
// zero value of Iterator is a terminal iterator.
//
// Unless otherwise specified, any mutation of a set invalidates all existing
// iterators into the set.
type FrameRefIterator struct {
	// node is the node containing the iterated segment. If the iterator is
	// terminal, node is nil.
	node *FrameRefnode

	// index is the index of the segment in node.keys/values.
	index int
}

// Ok returns true if the iterator is not terminal. All other methods are only
// valid for non-terminal iterators.
func (seg FrameRefIterator) Ok() bool {
	return seg.node != nil
}

// Range returns the iterated segment's range key.
func (seg FrameRefIterator) Range() __generics_imported0.FileRange {
	return seg.node.keys[seg.index]
}

// Start is equivalent to Range().Start, but should be preferred if only the
// start of the range is needed.
func (seg FrameRefIterator) Start() uint64 {
	return seg.node.keys[seg.index].Start
}

// End is equivalent to Range().End, but should be preferred if only the end of
// the range is needed.
func (seg FrameRefIterator) End() uint64 {
	return seg.node.keys[seg.index].End
}

// SetRangeUnchecked mutates the iterated segment's range key. This operation
// does not invalidate any iterators.
//
// Preconditions:
// - r.Length() > 0.
// - The new range must not overlap an existing one:
//   - If seg.NextSegment().Ok(), then r.end <= seg.NextSegment().Start().
//   - If seg.PrevSegment().Ok(), then r.start >= seg.PrevSegment().End().
func (seg FrameRefIterator) SetRangeUnchecked(r __generics_imported0.FileRange) {
	seg.node.keys[seg.index] = r
}

// SetRange mutates the iterated segment's range key. If the new range would
// cause the iterated segment to overlap another segment, or if the new range
// is invalid, SetRange panics. This operation does not invalidate any
// iterators.
func (seg FrameRefIterator) SetRange(r __generics_imported0.FileRange) {
	if r.Length() <= 0 {
		panic(fmt.Sprintf("invalid segment range %v", r))
	}
	if prev := seg.PrevSegment(); prev.Ok() && r.Start < prev.End() {
		panic(fmt.Sprintf("new segment range %v overlaps segment range %v", r, prev.Range()))
	}
	if next := seg.NextSegment(); next.Ok() && r.End > next.Start() {
		panic(fmt.Sprintf("new segment range %v overlaps segment range %v", r, next.Range()))
	}
	seg.SetRangeUnchecked(r)
}

// SetStartUnchecked mutates the iterated segment's start. This operation does
// not invalidate any iterators.
//
// Preconditions: The new start must be valid:
//   - start < seg.End()
//   - If seg.PrevSegment().Ok(), then start >= seg.PrevSegment().End().
func (seg FrameRefIterator) SetStartUnchecked(start uint64) {
	seg.node.keys[seg.index].Start = start
}

// SetStart mutates the iterated segment's start. If the new start value would
// cause the iterated segment to overlap another segment, or would result in an
// invalid range, SetStart panics. This operation does not invalidate any
// iterators.
func (seg FrameRefIterator) SetStart(start uint64) {
	if start >= seg.End() {
		panic(fmt.Sprintf("new start %v would invalidate segment range %v", start, seg.Range()))
	}
	if prev := seg.PrevSegment(); prev.Ok() && start < prev.End() {
		panic(fmt.Sprintf("new start %v would cause segment range %v to overlap segment range %v", start, seg.Range(), prev.Range()))
	}
	seg.SetStartUnchecked(start)
}

// SetEndUnchecked mutates the iterated segment's end. This operation does not
// invalidate any iterators.
//
// Preconditions: The new end must be valid:
//   - end > seg.Start().
//   - If seg.NextSegment().Ok(), then end <= seg.NextSegment().Start().
func (seg FrameRefIterator) SetEndUnchecked(end uint64) {
	seg.node.keys[seg.index].End = end
}

// SetEnd mutates the iterated segment's end. If the new end value would cause
// the iterated segment to overlap another segment, or would result in an
// invalid range, SetEnd panics. This operation does not invalidate any
// iterators.
func (seg FrameRefIterator) SetEnd(end uint64) {
	if end <= seg.Start() {
		panic(fmt.Sprintf("new end %v would invalidate segment range %v", end, seg.Range()))
	}
	if next := seg.NextSegment(); next.Ok() && end > next.Start() {
		panic(fmt.Sprintf("new end %v would cause segment range %v to overlap segment range %v", end, seg.Range(), next.Range()))
	}
	seg.SetEndUnchecked(end)
}

// Value returns a copy of the iterated segment's value.
func (seg FrameRefIterator) Value() FrameRefSegInfo {
	return seg.node.values[seg.index]
}

// ValuePtr returns a pointer to the iterated segment's value. The pointer is
// invalidated if the iterator is invalidated. This operation does not
// invalidate any iterators.
func (seg FrameRefIterator) ValuePtr() *FrameRefSegInfo {
	return &seg.node.values[seg.index]
}

// SetValue mutates the iterated segment's value. This operation does not
// invalidate any iterators.
func (seg FrameRefIterator) SetValue(val FrameRefSegInfo) {
	seg.node.values[seg.index] = val
}

// PrevSegment returns the iterated segment's predecessor. If there is no
// preceding segment, PrevSegment returns a terminal iterator.
func (seg FrameRefIterator) PrevSegment() FrameRefIterator {
	if seg.node.hasChildren {
		return seg.node.children[seg.index].lastSegment()
	}
	if seg.index > 0 {
		return FrameRefIterator{seg.node, seg.index - 1}
	}
	if seg.node.parent == nil {
		return FrameRefIterator{}
	}
	return FrameRefsegmentBeforePosition(seg.node.parent, seg.node.parentIndex)
}

// NextSegment returns the iterated segment's successor. If there is no
// succeeding segment, NextSegment returns a terminal iterator.
func (seg FrameRefIterator) NextSegment() FrameRefIterator {
	if seg.node.hasChildren {
		return seg.node.children[seg.index+1].firstSegment()
	}
	if seg.index < seg.node.nrSegments-1 {
		return FrameRefIterator{seg.node, seg.index + 1}
	}
	if seg.node.parent == nil {
		return FrameRefIterator{}
	}
	return FrameRefsegmentAfterPosition(seg.node.parent, seg.node.parentIndex)
}

// PrevGap returns the gap immediately before the iterated segment.
func (seg FrameRefIterator) PrevGap() FrameRefGapIterator {
	if seg.node.hasChildren {

		return seg.node.children[seg.index].lastSegment().NextGap()
	}
	return FrameRefGapIterator{seg.node, seg.index}
}

// NextGap returns the gap immediately after the iterated segment.
func (seg FrameRefIterator) NextGap() FrameRefGapIterator {
	if seg.node.hasChildren {
		return seg.node.children[seg.index+1].firstSegment().PrevGap()
	}
	return FrameRefGapIterator{seg.node, seg.index + 1}
}

// PrevNonEmpty returns the iterated segment's predecessor if it is adjacent,
// or the gap before the iterated segment otherwise. If seg.Start() ==
// Functions.MinKey(), PrevNonEmpty will return two terminal iterators.
// Otherwise, exactly one of the iterators returned by PrevNonEmpty will be
// non-terminal.
func (seg FrameRefIterator) PrevNonEmpty() (FrameRefIterator, FrameRefGapIterator) {
	if prev := seg.PrevSegment(); prev.Ok() && prev.End() == seg.Start() {
		return prev, FrameRefGapIterator{}
	}
	return FrameRefIterator{}, seg.PrevGap()
}

// NextNonEmpty returns the iterated segment's successor if it is adjacent, or
// the gap after the iterated segment otherwise. If seg.End() ==
// Functions.MaxKey(), NextNonEmpty will return two terminal iterators.
// Otherwise, exactly one of the iterators returned by NextNonEmpty will be
// non-terminal.
func (seg FrameRefIterator) NextNonEmpty() (FrameRefIterator, FrameRefGapIterator) {
	if next := seg.NextSegment(); next.Ok() && next.Start() == seg.End() {
		return next, FrameRefGapIterator{}
	}
	return FrameRefIterator{}, seg.NextGap()
}

// A GapIterator is conceptually one of:
//
//   - A pointer to a position between two segments, before the first segment, or
//     after the last segment in a set, called a *gap*; or
//
//   - A terminal iterator, which is a sentinel indicating that the end of
//     iteration has been reached.
//
// Note that the gap between two adjacent segments exists (iterators to it are
// non-terminal), but has a length of zero. GapIterator.IsEmpty returns true
// for such gaps. An empty set contains a single gap, spanning the entire range
// of the set's keys.
//
// GapIterators are copyable values and are meaningfully equality-comparable.
// The zero value of GapIterator is a terminal iterator.
//
// Unless otherwise specified, any mutation of a set invalidates all existing
// iterators into the set.
type FrameRefGapIterator struct {
	// The representation of a GapIterator is identical to that of an Iterator,
	// except that index corresponds to positions between segments in the same
	// way as for node.children (see comment for node.nrSegments).
	node  *FrameRefnode
	index int
}

// Ok returns true if the iterator is not terminal. All other methods are only
// valid for non-terminal iterators.
func (gap FrameRefGapIterator) Ok() bool {
	return gap.node != nil
}

// Range returns the range spanned by the iterated gap.
func (gap FrameRefGapIterator) Range() __generics_imported0.FileRange {
	return __generics_imported0.FileRange{gap.Start(), gap.End()}
}

// Start is equivalent to Range().Start, but should be preferred if only the
// start of the range is needed.
func (gap FrameRefGapIterator) Start() uint64 {
	if ps := gap.PrevSegment(); ps.Ok() {
		return ps.End()
	}
	return FrameRefSetFunctions{}.MinKey()
}

// End is equivalent to Range().End, but should be preferred if only the end of
// the range is needed.
func (gap FrameRefGapIterator) End() uint64 {
	if ns := gap.NextSegment(); ns.Ok() {
		return ns.Start()
	}
	return FrameRefSetFunctions{}.MaxKey()
}

// IsEmpty returns true if the iterated gap is empty (that is, the "gap" is
// between two adjacent segments.)
func (gap FrameRefGapIterator) IsEmpty() bool {
	return gap.Range().Length() == 0
}

// PrevSegment returns the segment immediately before the iterated gap. If no
// such segment exists, PrevSegment returns a terminal iterator.
func (gap FrameRefGapIterator) PrevSegment() FrameRefIterator {
	return FrameRefsegmentBeforePosition(gap.node, gap.index)
}

// NextSegment returns the segment immediately after the iterated gap. If no
// such segment exists, NextSegment returns a terminal iterator.
func (gap FrameRefGapIterator) NextSegment() FrameRefIterator {
	return FrameRefsegmentAfterPosition(gap.node, gap.index)
}

// PrevGap returns the iterated gap's predecessor. If no such gap exists,
// PrevGap returns a terminal iterator.
func (gap FrameRefGapIterator) PrevGap() FrameRefGapIterator {
	seg := gap.PrevSegment()
	if !seg.Ok() {
		return FrameRefGapIterator{}
	}
	return seg.PrevGap()
}

// NextGap returns the iterated gap's successor. If no such gap exists, NextGap
// returns a terminal iterator.
func (gap FrameRefGapIterator) NextGap() FrameRefGapIterator {
	seg := gap.NextSegment()
	if !seg.Ok() {
		return FrameRefGapIterator{}
	}
	return seg.NextGap()
}

// NextLargeEnoughGap returns the iterated gap's first next gap with larger
// length than minSize.  If not found, return a terminal gap iterator (does NOT
// include this gap itself).
//
// Precondition: trackGaps must be 1.
func (gap FrameRefGapIterator) NextLargeEnoughGap(minSize uint64) FrameRefGapIterator {
	if FrameReftrackGaps != 1 {
		panic("set is not tracking gaps")
	}
	if gap.node != nil && gap.node.hasChildren && gap.index == gap.node.nrSegments {

		gap.node = gap.NextSegment().node
		gap.index = 0
		return gap.nextLargeEnoughGapHelper(minSize)
	}
	return gap.nextLargeEnoughGapHelper(minSize)
}

// nextLargeEnoughGapHelper is the helper function used by NextLargeEnoughGap
// to do the real recursions.
//
// Preconditions: gap is NOT the trailing gap of a non-leaf node.
func (gap FrameRefGapIterator) nextLargeEnoughGapHelper(minSize uint64) FrameRefGapIterator {
	for {

		for gap.node != nil &&
			(gap.node.maxGap.Get() < minSize || (!gap.node.hasChildren && gap.index == gap.node.nrSegments)) {
			gap.node, gap.index = gap.node.parent, gap.node.parentIndex
		}

		if gap.node == nil {
			return FrameRefGapIterator{}
		}

		gap.index++
		for gap.index <= gap.node.nrSegments {
			if gap.node.hasChildren {
				if largeEnoughGap := gap.node.children[gap.index].searchFirstLargeEnoughGap(minSize); largeEnoughGap.Ok() {
					return largeEnoughGap
				}
			} else {
				if gap.Range().Length() >= minSize {
					return gap
				}
			}
			gap.index++
		}
		gap.node, gap.index = gap.node.parent, gap.node.parentIndex
		if gap.node != nil && gap.index == gap.node.nrSegments {

			gap.node, gap.index = gap.node.parent, gap.node.parentIndex
		}
	}
}

// PrevLargeEnoughGap returns the iterated gap's first prev gap with larger or
// equal length than minSize.  If not found, return a terminal gap iterator
// (does NOT include this gap itself).
//
// Precondition: trackGaps must be 1.
func (gap FrameRefGapIterator) PrevLargeEnoughGap(minSize uint64) FrameRefGapIterator {
	if FrameReftrackGaps != 1 {
		panic("set is not tracking gaps")
	}
	if gap.node != nil && gap.node.hasChildren && gap.index == 0 {

		gap.node = gap.PrevSegment().node
		gap.index = gap.node.nrSegments
		return gap.prevLargeEnoughGapHelper(minSize)
	}
	return gap.prevLargeEnoughGapHelper(minSize)
}

// prevLargeEnoughGapHelper is the helper function used by PrevLargeEnoughGap
// to do the real recursions.
//
// Preconditions: gap is NOT the first gap of a non-leaf node.
func (gap FrameRefGapIterator) prevLargeEnoughGapHelper(minSize uint64) FrameRefGapIterator {
	for {

		for gap.node != nil &&
			(gap.node.maxGap.Get() < minSize || (!gap.node.hasChildren && gap.index == 0)) {
			gap.node, gap.index = gap.node.parent, gap.node.parentIndex
		}

		if gap.node == nil {
			return FrameRefGapIterator{}
		}

		gap.index--
		for gap.index >= 0 {
			if gap.node.hasChildren {
				if largeEnoughGap := gap.node.children[gap.index].searchLastLargeEnoughGap(minSize); largeEnoughGap.Ok() {
					return largeEnoughGap
				}
			} else {
				if gap.Range().Length() >= minSize {
					return gap
				}
			}
			gap.index--
		}
		gap.node, gap.index = gap.node.parent, gap.node.parentIndex
		if gap.node != nil && gap.index == 0 {

			gap.node, gap.index = gap.node.parent, gap.node.parentIndex
		}
	}
}

// segmentBeforePosition returns the predecessor segment of the position given
// by n.children[i], which may or may not contain a child. If no such segment
// exists, segmentBeforePosition returns a terminal iterator.
func FrameRefsegmentBeforePosition(n *FrameRefnode, i int) FrameRefIterator {
	for i == 0 {
		if n.parent == nil {
			return FrameRefIterator{}
		}
		n, i = n.parent, n.parentIndex
	}
	return FrameRefIterator{n, i - 1}
}

// segmentAfterPosition returns the successor segment of the position given by
// n.children[i], which may or may not contain a child. If no such segment
// exists, segmentAfterPosition returns a terminal iterator.
func FrameRefsegmentAfterPosition(n *FrameRefnode, i int) FrameRefIterator {
	for i == n.nrSegments {
		if n.parent == nil {
			return FrameRefIterator{}
		}
		n, i = n.parent, n.parentIndex
	}
	return FrameRefIterator{n, i}
}

func FrameRefzeroValueSlice(slice []FrameRefSegInfo) {

	for i := range slice {
		FrameRefSetFunctions{}.ClearValue(&slice[i])
	}
}

func FrameRefzeroNodeSlice(slice []*FrameRefnode) {
	for i := range slice {
		slice[i] = nil
	}
}

// String stringifies a Set for debugging.
func (s *FrameRefSet) String() string {
	return s.root.String()
}

// String stringifies a node (and all of its children) for debugging.
func (n *FrameRefnode) String() string {
	var buf bytes.Buffer
	n.writeDebugString(&buf, "")
	return buf.String()
}

func (n *FrameRefnode) writeDebugString(buf *bytes.Buffer, prefix string) {
	if n.hasChildren != (n.nrSegments > 0 && n.children[0] != nil) {
		buf.WriteString(prefix)
		buf.WriteString(fmt.Sprintf("WARNING: inconsistent value of hasChildren: got %v, want %v\n", n.hasChildren, !n.hasChildren))
	}
	for i := 0; i < n.nrSegments; i++ {
		if child := n.children[i]; child != nil {
			cprefix := fmt.Sprintf("%s- % 3d ", prefix, i)
			if child.parent != n || child.parentIndex != i {
				buf.WriteString(cprefix)
				buf.WriteString(fmt.Sprintf("WARNING: inconsistent linkage to parent: got (%p, %d), want (%p, %d)\n", child.parent, child.parentIndex, n, i))
			}
			child.writeDebugString(buf, fmt.Sprintf("%s- % 3d ", prefix, i))
		}
		buf.WriteString(prefix)
		if n.hasChildren {
			if FrameReftrackGaps != 0 {
				buf.WriteString(fmt.Sprintf("- % 3d: %v => %v, maxGap: %d\n", i, n.keys[i], n.values[i], n.maxGap.Get()))
			} else {
				buf.WriteString(fmt.Sprintf("- % 3d: %v => %v\n", i, n.keys[i], n.values[i]))
			}
		} else {
			buf.WriteString(fmt.Sprintf("- % 3d: %v => %v\n", i, n.keys[i], n.values[i]))
		}
	}
	if child := n.children[n.nrSegments]; child != nil {
		child.writeDebugString(buf, fmt.Sprintf("%s- % 3d ", prefix, n.nrSegments))
	}
}

// FlatSegment represents a segment as a single object. FlatSegment is used as
// an intermediate representation for save/restore and tests.
//
// +stateify savable
type FrameRefFlatSegment struct {
	Start uint64
	End   uint64
	Value FrameRefSegInfo
}

// ExportSlice returns a copy of all segments in the given set, in ascending
// key order.
func (s *FrameRefSet) ExportSlice() []FrameRefFlatSegment {
	var fs []FrameRefFlatSegment
	for seg := s.FirstSegment(); seg.Ok(); seg = seg.NextSegment() {
		fs = append(fs, FrameRefFlatSegment{
			Start: seg.Start(),
			End:   seg.End(),
			Value: seg.Value(),
		})
	}
	return fs
}

// ImportSlice initializes the given set from the given slice.
//
// Preconditions:
//   - s must be empty.
//   - fs must represent a valid set (the segments in fs must have valid
//     lengths that do not overlap).
//   - The segments in fs must be sorted in ascending key order.
func (s *FrameRefSet) ImportSlice(fs []FrameRefFlatSegment) error {
	if !s.IsEmpty() {
		return fmt.Errorf("cannot import into non-empty set %v", s)
	}
	gap := s.FirstGap()
	for i := range fs {
		f := &fs[i]
		r := __generics_imported0.FileRange{f.Start, f.End}
		if !gap.Range().IsSupersetOf(r) {
			return fmt.Errorf("segment overlaps a preceding segment or is incorrectly sorted: %v => %v", r, f.Value)
		}
		gap = s.InsertWithoutMerging(gap, r, f.Value).NextGap()
	}
	return nil
}

// segmentTestCheck returns an error if s is incorrectly sorted, does not
// contain exactly expectedSegments segments, or contains a segment which
// fails the passed check.
//
// This should be used only for testing, and has been added to this package for
// templating convenience.
func (s *FrameRefSet) segmentTestCheck(expectedSegments int, segFunc func(int, __generics_imported0.FileRange, FrameRefSegInfo) error) error {
	havePrev := false
	prev := uint64(0)
	nrSegments := 0
	for seg := s.FirstSegment(); seg.Ok(); seg = seg.NextSegment() {
		next := seg.Start()
		if havePrev && prev >= next {
			return fmt.Errorf("incorrect order: key %d (segment %d) >= key %d (segment %d)", prev, nrSegments-1, next, nrSegments)
		}
		if segFunc != nil {
			if err := segFunc(nrSegments, seg.Range(), seg.Value()); err != nil {
				return err
			}
		}
		prev = next
		havePrev = true
		nrSegments++
	}
	if nrSegments != expectedSegments {
		return fmt.Errorf("incorrect number of segments: got %d, wanted %d", nrSegments, expectedSegments)
	}
	return nil
}

// countSegments counts the number of segments in the set.
//
// Similar to Check, this should only be used for testing.
func (s *FrameRefSet) countSegments() (segments int) {
	for seg := s.FirstSegment(); seg.Ok(); seg = seg.NextSegment() {
		segments++
	}
	return segments
}
func (s *FrameRefSet) saveRoot() []FrameRefFlatSegment {
	fs := s.ExportSlice()

	fs = fs[:len(fs):len(fs)]
	return fs
}

func (s *FrameRefSet) loadRoot(_ context.Context, fs []FrameRefFlatSegment) {
	if err := s.ImportSlice(fs); err != nil {
		panic(err)
	}
}
