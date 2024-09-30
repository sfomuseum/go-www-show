// Code generated by DarwinKit. DO NOT EDIT.

package webkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The methods for presenting native user interface elements on behalf of a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate?language=objc
type PUIDelegate interface {
	// optional
	WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func())
	HasWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler() bool

	// optional
	WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) WebView
	HasWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures() bool

	// optional
	WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func(result bool))
	HasWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler() bool

	// optional
	WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler(webView WebView, parameters OpenPanelParameters, frame FrameInfo, completionHandler func(URLs []foundation.URL))
	HasWebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler() bool

	// optional
	WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision))
	HasWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler() bool

	// optional
	WebViewDidClose(webView WebView)
	HasWebViewDidClose() bool

	// optional
	WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string))
	HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler() bool
}

// A delegate implementation builder for the [PUIDelegate] protocol.
type UIDelegate struct {
	_WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler               func(webView WebView, message string, frame FrameInfo, completionHandler func())
	_WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures                   func(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) WebView
	_WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler             func(webView WebView, message string, frame FrameInfo, completionHandler func(result bool))
	_WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler                       func(webView WebView, parameters OpenPanelParameters, frame FrameInfo, completionHandler func(URLs []foundation.URL))
	_WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler         func(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision))
	_WebViewDidClose                                                                          func(webView WebView)
	_WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler func(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string))
}

func (di *UIDelegate) HasWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler() bool {
	return di._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler != nil
}

// Displays a JavaScript alert panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1537406-webview?language=objc
func (di *UIDelegate) SetWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(f func(webView WebView, message string, frame FrameInfo, completionHandler func())) {
	di._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler = f
}

// Displays a JavaScript alert panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1537406-webview?language=objc
func (di *UIDelegate) WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func()) {
	di._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView, message, frame, completionHandler)
}
func (di *UIDelegate) HasWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures() bool {
	return di._WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures != nil
}

// Creates a new web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1536907-webview?language=objc
func (di *UIDelegate) SetWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(f func(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) WebView) {
	di._WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures = f
}

// Creates a new web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1536907-webview?language=objc
func (di *UIDelegate) WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) WebView {
	return di._WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView, configuration, navigationAction, windowFeatures)
}
func (di *UIDelegate) HasWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler() bool {
	return di._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler != nil
}

// Displays a JavaScript confirm panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1536489-webview?language=objc
func (di *UIDelegate) SetWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(f func(webView WebView, message string, frame FrameInfo, completionHandler func(result bool))) {
	di._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler = f
}

// Displays a JavaScript confirm panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1536489-webview?language=objc
func (di *UIDelegate) WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func(result bool)) {
	di._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView, message, frame, completionHandler)
}
func (di *UIDelegate) HasWebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler() bool {
	return di._WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler != nil
}

// Displays a file upload panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1641952-webview?language=objc
func (di *UIDelegate) SetWebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler(f func(webView WebView, parameters OpenPanelParameters, frame FrameInfo, completionHandler func(URLs []foundation.URL))) {
	di._WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler = f
}

// Displays a file upload panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1641952-webview?language=objc
func (di *UIDelegate) WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler(webView WebView, parameters OpenPanelParameters, frame FrameInfo, completionHandler func(URLs []foundation.URL)) {
	di._WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler(webView, parameters, frame, completionHandler)
}
func (di *UIDelegate) HasWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler() bool {
	return di._WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler != nil
}

// Determines whether a web resource, which the security origin object describes, can access to the device’s microphone audio and camera video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/3763087-webview?language=objc
func (di *UIDelegate) SetWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(f func(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision))) {
	di._WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler = f
}

// Determines whether a web resource, which the security origin object describes, can access to the device’s microphone audio and camera video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/3763087-webview?language=objc
func (di *UIDelegate) WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision)) {
	di._WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView, origin, frame, type_, decisionHandler)
}
func (di *UIDelegate) HasWebViewDidClose() bool {
	return di._WebViewDidClose != nil
}

// Notifies your app that the DOM window closed successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1537390-webviewdidclose?language=objc
func (di *UIDelegate) SetWebViewDidClose(f func(webView WebView)) {
	di._WebViewDidClose = f
}

// Notifies your app that the DOM window closed successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1537390-webviewdidclose?language=objc
func (di *UIDelegate) WebViewDidClose(webView WebView) {
	di._WebViewDidClose(webView)
}
func (di *UIDelegate) HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler() bool {
	return di._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler != nil
}

// Displays a JavaScript text input panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1538086-webview?language=objc
func (di *UIDelegate) SetWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(f func(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string))) {
	di._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler = f
}

// Displays a JavaScript text input panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1538086-webview?language=objc
func (di *UIDelegate) WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string)) {
	di._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView, prompt, defaultText, frame, completionHandler)
}

// ensure impl type implements protocol interface
var _ PUIDelegate = (*UIDelegateObject)(nil)

// A concrete type for the [PUIDelegate] protocol.
type UIDelegateObject struct {
	objc.Object
}

func (u_ UIDelegateObject) HasWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler() bool {
	return u_.RespondsToSelector(objc.Sel("webView:runJavaScriptAlertPanelWithMessage:initiatedByFrame:completionHandler:"))
}

// Displays a JavaScript alert panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1537406-webview?language=objc
func (u_ UIDelegateObject) WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func()) {
	objc.Call[objc.Void](u_, objc.Sel("webView:runJavaScriptAlertPanelWithMessage:initiatedByFrame:completionHandler:"), webView, message, frame, completionHandler)
}

func (u_ UIDelegateObject) HasWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures() bool {
	return u_.RespondsToSelector(objc.Sel("webView:createWebViewWithConfiguration:forNavigationAction:windowFeatures:"))
}

// Creates a new web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1536907-webview?language=objc
func (u_ UIDelegateObject) WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) WebView {
	rv := objc.Call[WebView](u_, objc.Sel("webView:createWebViewWithConfiguration:forNavigationAction:windowFeatures:"), webView, configuration, navigationAction, windowFeatures)
	return rv
}

func (u_ UIDelegateObject) HasWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler() bool {
	return u_.RespondsToSelector(objc.Sel("webView:runJavaScriptConfirmPanelWithMessage:initiatedByFrame:completionHandler:"))
}

// Displays a JavaScript confirm panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1536489-webview?language=objc
func (u_ UIDelegateObject) WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func(result bool)) {
	objc.Call[objc.Void](u_, objc.Sel("webView:runJavaScriptConfirmPanelWithMessage:initiatedByFrame:completionHandler:"), webView, message, frame, completionHandler)
}

func (u_ UIDelegateObject) HasWebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler() bool {
	return u_.RespondsToSelector(objc.Sel("webView:runOpenPanelWithParameters:initiatedByFrame:completionHandler:"))
}

// Displays a file upload panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1641952-webview?language=objc
func (u_ UIDelegateObject) WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler(webView WebView, parameters OpenPanelParameters, frame FrameInfo, completionHandler func(URLs []foundation.URL)) {
	objc.Call[objc.Void](u_, objc.Sel("webView:runOpenPanelWithParameters:initiatedByFrame:completionHandler:"), webView, parameters, frame, completionHandler)
}

func (u_ UIDelegateObject) HasWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler() bool {
	return u_.RespondsToSelector(objc.Sel("webView:requestMediaCapturePermissionForOrigin:initiatedByFrame:type:decisionHandler:"))
}

// Determines whether a web resource, which the security origin object describes, can access to the device’s microphone audio and camera video. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/3763087-webview?language=objc
func (u_ UIDelegateObject) WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision)) {
	objc.Call[objc.Void](u_, objc.Sel("webView:requestMediaCapturePermissionForOrigin:initiatedByFrame:type:decisionHandler:"), webView, origin, frame, type_, decisionHandler)
}

func (u_ UIDelegateObject) HasWebViewDidClose() bool {
	return u_.RespondsToSelector(objc.Sel("webViewDidClose:"))
}

// Notifies your app that the DOM window closed successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1537390-webviewdidclose?language=objc
func (u_ UIDelegateObject) WebViewDidClose(webView WebView) {
	objc.Call[objc.Void](u_, objc.Sel("webViewDidClose:"), webView)
}

func (u_ UIDelegateObject) HasWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler() bool {
	return u_.RespondsToSelector(objc.Sel("webView:runJavaScriptTextInputPanelWithPrompt:defaultText:initiatedByFrame:completionHandler:"))
}

// Displays a JavaScript text input panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1538086-webview?language=objc
func (u_ UIDelegateObject) WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string)) {
	objc.Call[objc.Void](u_, objc.Sel("webView:runJavaScriptTextInputPanelWithPrompt:defaultText:initiatedByFrame:completionHandler:"), webView, prompt, defaultText, frame, completionHandler)
}
