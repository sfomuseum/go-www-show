// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a stepper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystepper?language=objc
type PAccessibilityStepper interface {
	// optional
	AccessibilityPerformIncrement() bool
	HasAccessibilityPerformIncrement() bool

	// optional
	AccessibilityLabel() string
	HasAccessibilityLabel() bool

	// optional
	AccessibilityValue() objc.Object
	HasAccessibilityValue() bool

	// optional
	AccessibilityPerformDecrement() bool
	HasAccessibilityPerformDecrement() bool
}

// ensure impl type implements protocol interface
var _ PAccessibilityStepper = (*AccessibilityStepperObject)(nil)

// A concrete type for the [PAccessibilityStepper] protocol.
type AccessibilityStepperObject struct {
	objc.Object
}

func (a_ AccessibilityStepperObject) HasAccessibilityPerformIncrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformIncrement"))
}

// Increments the stepper’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystepper/1533764-accessibilityperformincrement?language=objc
func (a_ AccessibilityStepperObject) AccessibilityPerformIncrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

func (a_ AccessibilityStepperObject) HasAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabel"))
}

// Returns a short description of the stepper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystepper/1528702-accessibilitylabel?language=objc
func (a_ AccessibilityStepperObject) AccessibilityLabel() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityLabel"))
	return rv
}

func (a_ AccessibilityStepperObject) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// Returns the stepper’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystepper/1528167-accessibilityvalue?language=objc
func (a_ AccessibilityStepperObject) AccessibilityValue() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityValue"))
	return rv
}

func (a_ AccessibilityStepperObject) HasAccessibilityPerformDecrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformDecrement"))
}

// Decrements the stepper’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitystepper/1525327-accessibilityperformdecrement?language=objc
func (a_ AccessibilityStepperObject) AccessibilityPerformDecrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformDecrement"))
	return rv
}
