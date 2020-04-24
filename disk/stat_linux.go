// +build linux

/*
 * MinIO Cloud Storage, (C) 2015, 2016, 2017 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package disk

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// GetInfo returns total and free bytes available in a directory, e.g. `/`.
func GetInfo(path string) (info Info, err error) {
	s := syscall.Statfs_t{}
	err = syscall.Statfs(path, &s)
	if err != nil {
		return Info{}, err
	}
	reservedBlocks := uint64(s.Bfree) - uint64(s.Bavail)
	info = Info{
		Total:  uint64(s.Frsize) * (uint64(s.Blocks) - reservedBlocks),
		Free:   uint64(s.Frsize) * uint64(s.Bavail),
		Files:  uint64(s.Files),
		Ffree:  uint64(s.Ffree),
		Reserved: uint64(s.Frsize) * reservedBlocks,
		FSType: getFSType(int64(s.Type)),
	}
	return info, nil
}

func GetDirUsage(path string) (uint64, error) {
	cmd := exec.Command("du", "-sk", path)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return 0, err
	}

	usageStr := strings.Split(out.String(), "\t")[0]
	usage, err := strconv.ParseUint(usageStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return usage * 1024, nil
}

func GetDiskDev(path string) (interface{}, error) {
	st := syscall.Stat_t{}

	err := syscall.Stat(path, &st)
	if err != nil {
		return nil, fmt.Errorf("sys stat failed, %s", err)
	}

	return st, nil
}
