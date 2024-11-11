package callstack

import (
	"runtime"
)

func Get() []runtime.Frame {
	ret := make([]runtime.Frame, 0)
	pc := make([]uintptr, 10) // at least 1 entry needed
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])

	for {
		frame, more := frames.Next()
		ret = append(ret, frame)
		if !more {
			break
		}
	}
	return ret
}
