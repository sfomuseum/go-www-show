// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A set of methods for managing interactions with an open or save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate?language=objc
type POpenSavePanelDelegate interface {
	// optional
	PanelUserEnteredFilenameConfirmed(sender objc.Object, filename string, okFlag bool) string
	HasPanelUserEnteredFilenameConfirmed() bool

	// optional
	PanelShouldEnableURL(sender objc.Object, url foundation.URL) bool
	HasPanelShouldEnableURL() bool

	// optional
	PanelValidateURLError(sender objc.Object, url foundation.URL, outError unsafe.Pointer) bool
	HasPanelValidateURLError() bool

	// optional
	PanelDidChangeToDirectoryURL(sender objc.Object, url foundation.URL)
	HasPanelDidChangeToDirectoryURL() bool

	// optional
	PanelSelectionDidChange(sender objc.Object)
	HasPanelSelectionDidChange() bool

	// optional
	PanelWillExpand(sender objc.Object, expanding bool)
	HasPanelWillExpand() bool
}

// A delegate implementation builder for the [POpenSavePanelDelegate] protocol.
type OpenSavePanelDelegate struct {
	_PanelUserEnteredFilenameConfirmed func(sender objc.Object, filename string, okFlag bool) string
	_PanelShouldEnableURL              func(sender objc.Object, url foundation.URL) bool
	_PanelValidateURLError             func(sender objc.Object, url foundation.URL, outError unsafe.Pointer) bool
	_PanelDidChangeToDirectoryURL      func(sender objc.Object, url foundation.URL)
	_PanelSelectionDidChange           func(sender objc.Object)
	_PanelWillExpand                   func(sender objc.Object, expanding bool)
}

func (di *OpenSavePanelDelegate) HasPanelUserEnteredFilenameConfirmed() bool {
	return di._PanelUserEnteredFilenameConfirmed != nil
}

// Tells the delegate that the user confirmed a filename choice by clicking Save in a Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1524630-panel?language=objc
func (di *OpenSavePanelDelegate) SetPanelUserEnteredFilenameConfirmed(f func(sender objc.Object, filename string, okFlag bool) string) {
	di._PanelUserEnteredFilenameConfirmed = f
}

// Tells the delegate that the user confirmed a filename choice by clicking Save in a Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1524630-panel?language=objc
func (di *OpenSavePanelDelegate) PanelUserEnteredFilenameConfirmed(sender objc.Object, filename string, okFlag bool) string {
	return di._PanelUserEnteredFilenameConfirmed(sender, filename, okFlag)
}
func (di *OpenSavePanelDelegate) HasPanelShouldEnableURL() bool {
	return di._PanelShouldEnableURL != nil
}

// Asks the delegate whether the specified URL should be enabled in the Open panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1535200-panel?language=objc
func (di *OpenSavePanelDelegate) SetPanelShouldEnableURL(f func(sender objc.Object, url foundation.URL) bool) {
	di._PanelShouldEnableURL = f
}

// Asks the delegate whether the specified URL should be enabled in the Open panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1535200-panel?language=objc
func (di *OpenSavePanelDelegate) PanelShouldEnableURL(sender objc.Object, url foundation.URL) bool {
	return di._PanelShouldEnableURL(sender, url)
}
func (di *OpenSavePanelDelegate) HasPanelValidateURLError() bool {
	return di._PanelValidateURLError != nil
}

// Asks the delegate to validate the URL for a file that the user selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1535141-panel?language=objc
func (di *OpenSavePanelDelegate) SetPanelValidateURLError(f func(sender objc.Object, url foundation.URL, outError unsafe.Pointer) bool) {
	di._PanelValidateURLError = f
}

// Asks the delegate to validate the URL for a file that the user selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1535141-panel?language=objc
func (di *OpenSavePanelDelegate) PanelValidateURLError(sender objc.Object, url foundation.URL, outError unsafe.Pointer) bool {
	return di._PanelValidateURLError(sender, url, outError)
}
func (di *OpenSavePanelDelegate) HasPanelDidChangeToDirectoryURL() bool {
	return di._PanelDidChangeToDirectoryURL != nil
}

// Tells the delegate that the user changed the selected directory to the directory located at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1527117-panel?language=objc
func (di *OpenSavePanelDelegate) SetPanelDidChangeToDirectoryURL(f func(sender objc.Object, url foundation.URL)) {
	di._PanelDidChangeToDirectoryURL = f
}

// Tells the delegate that the user changed the selected directory to the directory located at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1527117-panel?language=objc
func (di *OpenSavePanelDelegate) PanelDidChangeToDirectoryURL(sender objc.Object, url foundation.URL) {
	di._PanelDidChangeToDirectoryURL(sender, url)
}
func (di *OpenSavePanelDelegate) HasPanelSelectionDidChange() bool {
	return di._PanelSelectionDidChange != nil
}

// Tells the delegate that the user changed the selection in the specified Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1533556-panelselectiondidchange?language=objc
func (di *OpenSavePanelDelegate) SetPanelSelectionDidChange(f func(sender objc.Object)) {
	di._PanelSelectionDidChange = f
}

// Tells the delegate that the user changed the selection in the specified Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1533556-panelselectiondidchange?language=objc
func (di *OpenSavePanelDelegate) PanelSelectionDidChange(sender objc.Object) {
	di._PanelSelectionDidChange(sender)
}
func (di *OpenSavePanelDelegate) HasPanelWillExpand() bool {
	return di._PanelWillExpand != nil
}

// Tells the delegate that the Save panel is about to expand or collapse because the user clicked the disclosure triangle that displays or hides the file browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1532953-panel?language=objc
func (di *OpenSavePanelDelegate) SetPanelWillExpand(f func(sender objc.Object, expanding bool)) {
	di._PanelWillExpand = f
}

// Tells the delegate that the Save panel is about to expand or collapse because the user clicked the disclosure triangle that displays or hides the file browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1532953-panel?language=objc
func (di *OpenSavePanelDelegate) PanelWillExpand(sender objc.Object, expanding bool) {
	di._PanelWillExpand(sender, expanding)
}

// ensure impl type implements protocol interface
var _ POpenSavePanelDelegate = (*OpenSavePanelDelegateObject)(nil)

// A concrete type for the [POpenSavePanelDelegate] protocol.
type OpenSavePanelDelegateObject struct {
	objc.Object
}

func (o_ OpenSavePanelDelegateObject) HasPanelUserEnteredFilenameConfirmed() bool {
	return o_.RespondsToSelector(objc.Sel("panel:userEnteredFilename:confirmed:"))
}

// Tells the delegate that the user confirmed a filename choice by clicking Save in a Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1524630-panel?language=objc
func (o_ OpenSavePanelDelegateObject) PanelUserEnteredFilenameConfirmed(sender objc.Object, filename string, okFlag bool) string {
	rv := objc.Call[string](o_, objc.Sel("panel:userEnteredFilename:confirmed:"), sender, filename, okFlag)
	return rv
}

func (o_ OpenSavePanelDelegateObject) HasPanelShouldEnableURL() bool {
	return o_.RespondsToSelector(objc.Sel("panel:shouldEnableURL:"))
}

// Asks the delegate whether the specified URL should be enabled in the Open panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1535200-panel?language=objc
func (o_ OpenSavePanelDelegateObject) PanelShouldEnableURL(sender objc.Object, url foundation.URL) bool {
	rv := objc.Call[bool](o_, objc.Sel("panel:shouldEnableURL:"), sender, url)
	return rv
}

func (o_ OpenSavePanelDelegateObject) HasPanelValidateURLError() bool {
	return o_.RespondsToSelector(objc.Sel("panel:validateURL:error:"))
}

// Asks the delegate to validate the URL for a file that the user selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1535141-panel?language=objc
func (o_ OpenSavePanelDelegateObject) PanelValidateURLError(sender objc.Object, url foundation.URL, outError unsafe.Pointer) bool {
	rv := objc.Call[bool](o_, objc.Sel("panel:validateURL:error:"), sender, url, outError)
	return rv
}

func (o_ OpenSavePanelDelegateObject) HasPanelDidChangeToDirectoryURL() bool {
	return o_.RespondsToSelector(objc.Sel("panel:didChangeToDirectoryURL:"))
}

// Tells the delegate that the user changed the selected directory to the directory located at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1527117-panel?language=objc
func (o_ OpenSavePanelDelegateObject) PanelDidChangeToDirectoryURL(sender objc.Object, url foundation.URL) {
	objc.Call[objc.Void](o_, objc.Sel("panel:didChangeToDirectoryURL:"), sender, url)
}

func (o_ OpenSavePanelDelegateObject) HasPanelSelectionDidChange() bool {
	return o_.RespondsToSelector(objc.Sel("panelSelectionDidChange:"))
}

// Tells the delegate that the user changed the selection in the specified Save panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1533556-panelselectiondidchange?language=objc
func (o_ OpenSavePanelDelegateObject) PanelSelectionDidChange(sender objc.Object) {
	objc.Call[objc.Void](o_, objc.Sel("panelSelectionDidChange:"), sender)
}

func (o_ OpenSavePanelDelegateObject) HasPanelWillExpand() bool {
	return o_.RespondsToSelector(objc.Sel("panel:willExpand:"))
}

// Tells the delegate that the Save panel is about to expand or collapse because the user clicked the disclosure triangle that displays or hides the file browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopensavepaneldelegate/1532953-panel?language=objc
func (o_ OpenSavePanelDelegateObject) PanelWillExpand(sender objc.Object, expanding bool) {
	objc.Call[objc.Void](o_, objc.Sel("panel:willExpand:"), sender, expanding)
}
