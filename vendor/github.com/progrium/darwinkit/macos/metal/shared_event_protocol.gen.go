// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"github.com/progrium/darwinkit/objc"
)

// An object you use to synchronize access to Metal resources across multiple CPUs, GPUs, and processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedevent?language=objc
type PSharedEvent interface {
	// optional
	NotifyListenerAtValueBlock(listener SharedEventListener, value uint64, block SharedEventNotificationBlock)
	HasNotifyListenerAtValueBlock() bool

	// optional
	NewSharedEventHandle() SharedEventHandle
	HasNewSharedEventHandle() bool

	// optional
	SetSignaledValue(value uint64)
	HasSetSignaledValue() bool

	// optional
	SignaledValue() uint64
	HasSignaledValue() bool
}

// ensure impl type implements protocol interface
var _ PSharedEvent = (*SharedEventObject)(nil)

// A concrete type for the [PSharedEvent] protocol.
type SharedEventObject struct {
	objc.Object
}

func (s_ SharedEventObject) HasNotifyListenerAtValueBlock() bool {
	return s_.RespondsToSelector(objc.Sel("notifyListener:atValue:block:"))
}

// Schedules a notification handler to be called after the shareable event’s signal value equals or exceeds a given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedevent/2966574-notifylistener?language=objc
func (s_ SharedEventObject) NotifyListenerAtValueBlock(listener SharedEventListener, value uint64, block SharedEventNotificationBlock) {
	objc.Call[objc.Void](s_, objc.Sel("notifyListener:atValue:block:"), listener, value, block)
}

func (s_ SharedEventObject) HasNewSharedEventHandle() bool {
	return s_.RespondsToSelector(objc.Sel("newSharedEventHandle"))
}

// Creates a new shareable event handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedevent/2981025-newsharedeventhandle?language=objc
func (s_ SharedEventObject) NewSharedEventHandle() SharedEventHandle {
	rv := objc.Call[SharedEventHandle](s_, objc.Sel("newSharedEventHandle"))
	return rv
}

func (s_ SharedEventObject) HasSetSignaledValue() bool {
	return s_.RespondsToSelector(objc.Sel("setSignaledValue:"))
}

// The current signal value for the shareable event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedevent/2966575-signaledvalue?language=objc
func (s_ SharedEventObject) SetSignaledValue(value uint64) {
	objc.Call[objc.Void](s_, objc.Sel("setSignaledValue:"), value)
}

func (s_ SharedEventObject) HasSignaledValue() bool {
	return s_.RespondsToSelector(objc.Sel("signaledValue"))
}

// The current signal value for the shareable event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedevent/2966575-signaledvalue?language=objc
func (s_ SharedEventObject) SignaledValue() uint64 {
	rv := objc.Call[uint64](s_, objc.Sel("signaledValue"))
	return rv
}
