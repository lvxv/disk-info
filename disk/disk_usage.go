/*
 * MinIO Cloud Storage, (C) 2018 MinIO, Inc.
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
	"context"
	"github.com/lvxv/disk-info/utils"
	"io/ioutil"
	"path"
)

// getDiskUsage walks the file tree rooted at root, calling usageFn
// for each file or directory in the tree, including root.
func GetDiskUsage(ctx context.Context, root string, usageFn usageFunc) error {
	return walk(ctx, root+utils.SlashSeparator, usageFn)
}

// pathJoin - like path.Join() but retains trailing "/" of the last element
func PathJoin(elem ...string) string {
	trailingSlash := ""
	if len(elem) > 0 {
		if utils.HasSuffix(elem[len(elem)-1], utils.SlashSeparator) {
			trailingSlash = "/"
		}
	}
	return path.Join(elem...) + trailingSlash
}

type usageFunc func(ctx context.Context, entry string) error

// walk recursively descends path, calling walkFn.
func walk(ctx context.Context, path string, usageFn usageFunc) error {
	if err := usageFn(ctx, path); err != nil {
		return err
	}

	if !utils.HasSuffix(path, utils.SlashSeparator) {
		return nil
	}

	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return usageFn(ctx, path)
	}

	for _, entry := range entries {
		fname := PathJoin(path, entry.Name())
		if err = walk(ctx, fname, usageFn); err != nil {
			return err
		}
	}

	return nil
}
