// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"github.com/progrium/darwinkit/objc"
)

// An interface that enables customizing the behavior of the default notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationcenterdelegate?language=objc
type PUserNotificationCenterDelegate interface {
}

// A delegate implementation builder for the [PUserNotificationCenterDelegate] protocol.
type UserNotificationCenterDelegate struct {
}

// ensure impl type implements protocol interface
var _ PUserNotificationCenterDelegate = (*UserNotificationCenterDelegateObject)(nil)

// A concrete type for the [PUserNotificationCenterDelegate] protocol.
type UserNotificationCenterDelegateObject struct {
	objc.Object
}
