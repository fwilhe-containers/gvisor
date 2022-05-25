// Copyright 2022 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package procfs holds utilities for getting procfs information for sandboxed
// processes.
package procfs

import (
	"bytes"
	"fmt"
	"strings"

	"gvisor.dev/gvisor/pkg/context"
	"gvisor.dev/gvisor/pkg/log"
	"gvisor.dev/gvisor/pkg/sentry/fsimpl/proc"
	"gvisor.dev/gvisor/pkg/sentry/kernel"
	"gvisor.dev/gvisor/pkg/sentry/mm"
	"gvisor.dev/gvisor/pkg/sentry/vfs"
)

// FDInfo contains information about an application file descriptor.
type FDInfo struct {
	// Number is the FD number.
	Number int32 `json:"number,omitempty"`
	// Path is the path of the file that FD represents.
	Path string `json:"path,omitempty"`
}

// ProcessProcfsDump contains the procfs dump for one process.
type ProcessProcfsDump struct {
	// PID is the process ID.
	PID int32 `json:"pid,omitempty"`
	// Exe is the symlink target of /proc/[pid]/exe.
	Exe string `json:"exe,omitempty"`
	// Args is /proc/[pid]/cmdline split into an array.
	Args []string `json:"args,omitempty"`
	// Env is /proc/[pid]/environ split into an array.
	Env []string `json:"env,omitempty"`
	// CWD is the symlink target of /proc/[pid]/cwd.
	CWD string `json:"cwd,omitempty"`
	// FDs contains the directory entries of /proc/[pid]/fd and also contains the
	// symlink target for each FD.
	FDs []FDInfo `json:"fdlist,omitempty"`
	// StartTime is the process start time in nanoseconds since Unix epoch.
	StartTime int64 `json:"clone_ts,omitempty"`
}

// getMM returns t's MemoryManager. On success, the MemoryManager's users count
// is incremented, and must be decremented by the caller when it is no longer
// in use.
func getMM(t *kernel.Task) *mm.MemoryManager {
	var mm *mm.MemoryManager
	t.WithMuLocked(func(*kernel.Task) {
		mm = t.MemoryManager()
	})
	if mm == nil || !mm.IncUsers() {
		return nil
	}
	return mm
}

func getExecutablePath(ctx context.Context, pid kernel.ThreadID, mm *mm.MemoryManager) string {
	exec := mm.Executable()
	if exec == nil {
		log.Warningf("No executable found for PID %s", pid)
		return ""
	}
	defer exec.DecRef(ctx)

	return exec.PathnameWithDeleted(ctx)
}

func getMetadataArray(ctx context.Context, pid kernel.ThreadID, mm *mm.MemoryManager, metaType proc.MetadataType) []string {
	buf := bytes.Buffer{}
	if err := proc.GetMetadata(ctx, mm, &buf, metaType); err != nil {
		log.Warningf("failed to get %v metadata for PID %s: %v", metaType, pid, err)
		return nil
	}
	// As per proc(5), /proc/[pid]/cmdline may have "a further null byte after
	// the last string". Similarly, for /proc/[pid]/environ "there may be a null
	// byte at the end". So trim off the last null byte if it exists.
	return strings.Split(strings.TrimSuffix(buf.String(), "\000"), "\000")
}

func getCWD(ctx context.Context, t *kernel.Task, pid kernel.ThreadID) string {
	cwdDentry := t.FSContext().WorkingDirectoryVFS2()
	if !cwdDentry.Ok() {
		log.Warningf("No CWD dentry found for PID %s", pid)
		return ""
	}

	root := vfs.RootFromContext(ctx)
	if !root.Ok() {
		log.Warningf("no root could be found from context for PID %s", pid)
		return ""
	}
	defer root.DecRef(ctx)

	vfsObj := cwdDentry.Mount().Filesystem().VirtualFilesystem()
	name, err := vfsObj.PathnameWithDeleted(ctx, root, cwdDentry)
	if err != nil {
		log.Warningf("PathnameWithDeleted failed to find CWD: %v", err)
	}
	return name
}

func getFDs(ctx context.Context, t *kernel.Task, pid kernel.ThreadID) []FDInfo {
	type fdInfo struct {
		fd *vfs.FileDescription
		no int32
	}
	var fds []fdInfo
	defer func() {
		for _, fd := range fds {
			fd.fd.DecRef(ctx)
		}
	}()

	t.WithMuLocked(func(t *kernel.Task) {
		if fdTable := t.FDTable(); fdTable != nil {
			fdNos := fdTable.GetFDs(ctx)
			fds = make([]fdInfo, 0, len(fdNos))
			for _, fd := range fdNos {
				file, _ := fdTable.GetVFS2(fd)
				if file != nil {
					fds = append(fds, fdInfo{fd: file, no: fd})
				}
			}
		}
	})

	root := vfs.RootFromContext(ctx)
	defer root.DecRef(ctx)

	res := make([]FDInfo, 0, len(fds))
	for _, fd := range fds {
		path, err := t.Kernel().VFS().PathnameWithDeleted(ctx, root, fd.fd.VirtualDentry())
		if err != nil {
			log.Warningf("PathnameWithDeleted failed to find path for fd %d in PID %s: %v", fd.no, pid, err)
			path = ""
		}
		res = append(res, FDInfo{Number: fd.no, Path: path})
	}
	return res
}

// Dump returns a procfs dump for process pid. t must be a task in process pid.
func Dump(t *kernel.Task, pid kernel.ThreadID) (ProcessProcfsDump, error) {
	ctx := t.AsyncContext()

	mm := getMM(t)
	if mm == nil {
		return ProcessProcfsDump{}, fmt.Errorf("no MM found for PID %s", pid)
	}
	defer mm.DecUsers(ctx)

	return ProcessProcfsDump{
		PID:       int32(pid),
		Exe:       getExecutablePath(ctx, pid, mm),
		Args:      getMetadataArray(ctx, pid, mm, proc.Cmdline),
		Env:       getMetadataArray(ctx, pid, mm, proc.Environ),
		CWD:       getCWD(ctx, t, pid),
		FDs:       getFDs(ctx, t, pid),
		StartTime: t.StartTime().Nanoseconds(),
	}, nil
}