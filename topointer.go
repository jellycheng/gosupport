package gosupport

import "time"

func TimePtr(t time.Time) *time.Time {
	return &t
}

func StringPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func Float32Ptr(f float32) *float32 {
	return &f
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func Int32Ptr(i int32) *int32 {
	return &i
}

func IntPtr(i int) *int {
	return &i
}

func Uint32Ptr(i uint32) *uint32 {
	return &i
}

func Uint64Ptr(i uint64) *uint64 {
	return &i
}
