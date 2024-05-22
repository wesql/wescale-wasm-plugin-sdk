package common

import (
	"errors"
	"strconv"
)

type Status uint32

const (
	StatusOK              Status = 0
	StatusNotFound        Status = 1
	StatusBadArgument     Status = 2
	StatusEmpty           Status = 7
	StatusCasMismatch     Status = 8
	StatusInternalFailure Status = 10
	StatusUnimplemented   Status = 12
)

var (
	// ErrorStatusNotFound means not found for various hostcalls.
	ErrorStatusNotFound = errors.New("error status returned by host: not found")
	// ErrorStatusBadArgument means the arguments for a hostcall are invalid.
	ErrorStatusBadArgument = errors.New("error status returned by host: bad argument")
	// ErrorStatusEmpty means the target queue of DequeueSharedQueue call is empty.
	ErrorStatusEmpty = errors.New("error status returned by host: empty")
	// ErrorStatusCasMismatch means the CAS value provided to the SetSharedData
	// does not match the current value. It indicates that other Wasm VMs
	// have already set a value for the same key, and the current CAS
	// for the key gets incremented.
	// Having retry logic in the face of this error is recommended.
	ErrorStatusCasMismatch = errors.New("error status returned by host: cas mismatch")
	// ErrorInternalFailure indicates an internal failure in hosts.
	// When this error occurs, there's nothing we could do in the Wasm VM.
	// Abort or panic after this error is recommended.
	ErrorInternalFailure = errors.New("error status returned by host: internal failure")
	// ErrorUnimplemented indicates the API is not implemented in the host yet.
	ErrorUnimplemented = errors.New("error status returned by host: unimplemented")
)

func StatusToError(status Status) error {
	switch Status(status) {
	case StatusOK:
		return nil
	case StatusNotFound:
		return ErrorStatusNotFound
	case StatusBadArgument:
		return ErrorStatusBadArgument
	case StatusEmpty:
		return ErrorStatusEmpty
	case StatusCasMismatch:
		return ErrorStatusCasMismatch
	case StatusInternalFailure:
		return ErrorInternalFailure
	case StatusUnimplemented:
		return ErrorUnimplemented
	}
	return errors.New("unknown status code: " + strconv.Itoa(int(status)))
}
