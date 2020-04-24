package utils

import (
	"github.com/dustin/go-humanize"
	"reflect"
	"runtime"
	"strings"
	"time"
)

const (
	// Slash separator.
	SlashSeparator = "/"

	WindowsOSName     = "windows"
	NetBSDOSName      = "netbsd"
	DiskMinFreeSpace  = 400 * humanize.MiByte // Min 400MiB free space.
	DiskMinTotalSpace = DiskMinFreeSpace      // Min 400MiB total space.
)

// UTCNow - returns current UTC time.
func UTCNow() time.Time {
	return time.Now().UTC()
}

func Contains(slice interface{}, elem interface{}) bool {
	v := reflect.ValueOf(slice)
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

// Suffix matcher string matches suffix in a platform specific way.
// For example on windows since its case insensitive we are supposed
// to do case insensitive checks.
func HasSuffix(s string, suffix string) bool {
	if runtime.GOOS == "windows" {
		return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
	}
	return strings.HasSuffix(s, suffix)
}
