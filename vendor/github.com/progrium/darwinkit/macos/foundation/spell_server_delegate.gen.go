// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The optional methods implemented by the delegate of a spell server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate?language=objc
type PSpellServerDelegate interface {
	// optional
	SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(sender SpellServer, range_ Range, string_ string, language string) []string
	HasSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage() bool

	// optional
	SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(sender SpellServer, stringToCheck string, language string, wordCount *int, countOnly bool) Range
	HasSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly() bool

	// optional
	SpellServerRecordResponseToCorrectionForWordLanguage(sender SpellServer, response uint, correction string, word string, language string)
	HasSpellServerRecordResponseToCorrectionForWordLanguage() bool

	// optional
	SpellServerSuggestGuessesForWordInLanguage(sender SpellServer, word string, language string) []string
	HasSpellServerSuggestGuessesForWordInLanguage() bool

	// optional
	SpellServerCheckGrammarInStringLanguageDetails(sender SpellServer, stringToCheck string, language string, details unsafe.Pointer) Range
	HasSpellServerCheckGrammarInStringLanguageDetails() bool

	// optional
	SpellServerDidLearnWordInLanguage(sender SpellServer, word string, language string)
	HasSpellServerDidLearnWordInLanguage() bool

	// optional
	SpellServerDidForgetWordInLanguage(sender SpellServer, word string, language string)
	HasSpellServerDidForgetWordInLanguage() bool

	// optional
	SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(sender SpellServer, stringToCheck string, offset uint, checkingTypes TextCheckingTypes, options map[string]objc.Object, orthography Orthography, wordCount *int) []TextCheckingResult
	HasSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount() bool
}

// A delegate implementation builder for the [PSpellServerDelegate] protocol.
type SpellServerDelegate struct {
	_SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage func(sender SpellServer, range_ Range, string_ string, language string) []string
	_SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly  func(sender SpellServer, stringToCheck string, language string, wordCount *int, countOnly bool) Range
	_SpellServerRecordResponseToCorrectionForWordLanguage             func(sender SpellServer, response uint, correction string, word string, language string)
	_SpellServerSuggestGuessesForWordInLanguage                       func(sender SpellServer, word string, language string) []string
	_SpellServerCheckGrammarInStringLanguageDetails                   func(sender SpellServer, stringToCheck string, language string, details unsafe.Pointer) Range
	_SpellServerDidLearnWordInLanguage                                func(sender SpellServer, word string, language string)
	_SpellServerDidForgetWordInLanguage                               func(sender SpellServer, word string, language string)
	_SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount     func(sender SpellServer, stringToCheck string, offset uint, checkingTypes TextCheckingTypes, options map[string]objc.Object, orthography Orthography, wordCount *int) []TextCheckingResult
}

func (di *SpellServerDelegate) HasSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage() bool {
	return di._SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage != nil
}

// This delegate method returns an array of possible word completions from the spell checker, based on a partially completed string and a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1414606-spellserver?language=objc
func (di *SpellServerDelegate) SetSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(f func(sender SpellServer, range_ Range, string_ string, language string) []string) {
	di._SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage = f
}

// This delegate method returns an array of possible word completions from the spell checker, based on a partially completed string and a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1414606-spellserver?language=objc
func (di *SpellServerDelegate) SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(sender SpellServer, range_ Range, string_ string, language string) []string {
	return di._SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(sender, range_, string_, language)
}
func (di *SpellServerDelegate) HasSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly() bool {
	return di._SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly != nil
}

// Asks the delegate to search for a misspelled word in a given string, using the specified language, and marking the first misspelled word found by returning its range within the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1413235-spellserver?language=objc
func (di *SpellServerDelegate) SetSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(f func(sender SpellServer, stringToCheck string, language string, wordCount *int, countOnly bool) Range) {
	di._SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly = f
}

// Asks the delegate to search for a misspelled word in a given string, using the specified language, and marking the first misspelled word found by returning its range within the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1413235-spellserver?language=objc
func (di *SpellServerDelegate) SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(sender SpellServer, stringToCheck string, language string, wordCount *int, countOnly bool) Range {
	return di._SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(sender, stringToCheck, language, wordCount, countOnly)
}
func (di *SpellServerDelegate) HasSpellServerRecordResponseToCorrectionForWordLanguage() bool {
	return di._SpellServerRecordResponseToCorrectionForWordLanguage != nil
}

// Notifies the spell checker of the users’s response to a correction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1412894-spellserver?language=objc
func (di *SpellServerDelegate) SetSpellServerRecordResponseToCorrectionForWordLanguage(f func(sender SpellServer, response uint, correction string, word string, language string)) {
	di._SpellServerRecordResponseToCorrectionForWordLanguage = f
}

// Notifies the spell checker of the users’s response to a correction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1412894-spellserver?language=objc
func (di *SpellServerDelegate) SpellServerRecordResponseToCorrectionForWordLanguage(sender SpellServer, response uint, correction string, word string, language string) {
	di._SpellServerRecordResponseToCorrectionForWordLanguage(sender, response, correction, word, language)
}
func (di *SpellServerDelegate) HasSpellServerSuggestGuessesForWordInLanguage() bool {
	return di._SpellServerSuggestGuessesForWordInLanguage != nil
}

// Gives the delegate the opportunity to suggest guesses to the sender for the correct spelling of the given misspelled word in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1410726-spellserver?language=objc
func (di *SpellServerDelegate) SetSpellServerSuggestGuessesForWordInLanguage(f func(sender SpellServer, word string, language string) []string) {
	di._SpellServerSuggestGuessesForWordInLanguage = f
}

// Gives the delegate the opportunity to suggest guesses to the sender for the correct spelling of the given misspelled word in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1410726-spellserver?language=objc
func (di *SpellServerDelegate) SpellServerSuggestGuessesForWordInLanguage(sender SpellServer, word string, language string) []string {
	return di._SpellServerSuggestGuessesForWordInLanguage(sender, word, language)
}
func (di *SpellServerDelegate) HasSpellServerCheckGrammarInStringLanguageDetails() bool {
	return di._SpellServerCheckGrammarInStringLanguageDetails != nil
}

// Gives the delegate the opportunity to customize the grammatical analysis of a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1409242-spellserver?language=objc
func (di *SpellServerDelegate) SetSpellServerCheckGrammarInStringLanguageDetails(f func(sender SpellServer, stringToCheck string, language string, details unsafe.Pointer) Range) {
	di._SpellServerCheckGrammarInStringLanguageDetails = f
}

// Gives the delegate the opportunity to customize the grammatical analysis of a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1409242-spellserver?language=objc
func (di *SpellServerDelegate) SpellServerCheckGrammarInStringLanguageDetails(sender SpellServer, stringToCheck string, language string, details unsafe.Pointer) Range {
	return di._SpellServerCheckGrammarInStringLanguageDetails(sender, stringToCheck, language, details)
}
func (di *SpellServerDelegate) HasSpellServerDidLearnWordInLanguage() bool {
	return di._SpellServerDidLearnWordInLanguage != nil
}

// Notifies the delegate that the sender has added the specified word to the user’s list of acceptable words in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1407851-spellserver?language=objc
func (di *SpellServerDelegate) SetSpellServerDidLearnWordInLanguage(f func(sender SpellServer, word string, language string)) {
	di._SpellServerDidLearnWordInLanguage = f
}

// Notifies the delegate that the sender has added the specified word to the user’s list of acceptable words in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1407851-spellserver?language=objc
func (di *SpellServerDelegate) SpellServerDidLearnWordInLanguage(sender SpellServer, word string, language string) {
	di._SpellServerDidLearnWordInLanguage(sender, word, language)
}
func (di *SpellServerDelegate) HasSpellServerDidForgetWordInLanguage() bool {
	return di._SpellServerDidForgetWordInLanguage != nil
}

// Notifies the delegate that the sender has removed the specified word from the user’s list of acceptable words in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1417315-spellserver?language=objc
func (di *SpellServerDelegate) SetSpellServerDidForgetWordInLanguage(f func(sender SpellServer, word string, language string)) {
	di._SpellServerDidForgetWordInLanguage = f
}

// Notifies the delegate that the sender has removed the specified word from the user’s list of acceptable words in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1417315-spellserver?language=objc
func (di *SpellServerDelegate) SpellServerDidForgetWordInLanguage(sender SpellServer, word string, language string) {
	di._SpellServerDidForgetWordInLanguage(sender, word, language)
}
func (di *SpellServerDelegate) HasSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount() bool {
	return di._SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount != nil
}

// Gives the delegate the opportunity to analyze both the spelling and grammar simultaneously, which is more efficient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1409733-spellserver?language=objc
func (di *SpellServerDelegate) SetSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(f func(sender SpellServer, stringToCheck string, offset uint, checkingTypes TextCheckingTypes, options map[string]objc.Object, orthography Orthography, wordCount *int) []TextCheckingResult) {
	di._SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount = f
}

// Gives the delegate the opportunity to analyze both the spelling and grammar simultaneously, which is more efficient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1409733-spellserver?language=objc
func (di *SpellServerDelegate) SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(sender SpellServer, stringToCheck string, offset uint, checkingTypes TextCheckingTypes, options map[string]objc.Object, orthography Orthography, wordCount *int) []TextCheckingResult {
	return di._SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(sender, stringToCheck, offset, checkingTypes, options, orthography, wordCount)
}

// ensure impl type implements protocol interface
var _ PSpellServerDelegate = (*SpellServerDelegateObject)(nil)

// A concrete type for the [PSpellServerDelegate] protocol.
type SpellServerDelegateObject struct {
	objc.Object
}

func (s_ SpellServerDelegateObject) HasSpellServerSuggestCompletionsForPartialWordRangeInStringLanguage() bool {
	return s_.RespondsToSelector(objc.Sel("spellServer:suggestCompletionsForPartialWordRange:inString:language:"))
}

// This delegate method returns an array of possible word completions from the spell checker, based on a partially completed string and a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1414606-spellserver?language=objc
func (s_ SpellServerDelegateObject) SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(sender SpellServer, range_ Range, string_ string, language string) []string {
	rv := objc.Call[[]string](s_, objc.Sel("spellServer:suggestCompletionsForPartialWordRange:inString:language:"), sender, range_, string_, language)
	return rv
}

func (s_ SpellServerDelegateObject) HasSpellServerFindMisspelledWordInStringLanguageWordCountCountOnly() bool {
	return s_.RespondsToSelector(objc.Sel("spellServer:findMisspelledWordInString:language:wordCount:countOnly:"))
}

// Asks the delegate to search for a misspelled word in a given string, using the specified language, and marking the first misspelled word found by returning its range within the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1413235-spellserver?language=objc
func (s_ SpellServerDelegateObject) SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(sender SpellServer, stringToCheck string, language string, wordCount *int, countOnly bool) Range {
	rv := objc.Call[Range](s_, objc.Sel("spellServer:findMisspelledWordInString:language:wordCount:countOnly:"), sender, stringToCheck, language, wordCount, countOnly)
	return rv
}

func (s_ SpellServerDelegateObject) HasSpellServerRecordResponseToCorrectionForWordLanguage() bool {
	return s_.RespondsToSelector(objc.Sel("spellServer:recordResponse:toCorrection:forWord:language:"))
}

// Notifies the spell checker of the users’s response to a correction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1412894-spellserver?language=objc
func (s_ SpellServerDelegateObject) SpellServerRecordResponseToCorrectionForWordLanguage(sender SpellServer, response uint, correction string, word string, language string) {
	objc.Call[objc.Void](s_, objc.Sel("spellServer:recordResponse:toCorrection:forWord:language:"), sender, response, correction, word, language)
}

func (s_ SpellServerDelegateObject) HasSpellServerSuggestGuessesForWordInLanguage() bool {
	return s_.RespondsToSelector(objc.Sel("spellServer:suggestGuessesForWord:inLanguage:"))
}

// Gives the delegate the opportunity to suggest guesses to the sender for the correct spelling of the given misspelled word in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1410726-spellserver?language=objc
func (s_ SpellServerDelegateObject) SpellServerSuggestGuessesForWordInLanguage(sender SpellServer, word string, language string) []string {
	rv := objc.Call[[]string](s_, objc.Sel("spellServer:suggestGuessesForWord:inLanguage:"), sender, word, language)
	return rv
}

func (s_ SpellServerDelegateObject) HasSpellServerCheckGrammarInStringLanguageDetails() bool {
	return s_.RespondsToSelector(objc.Sel("spellServer:checkGrammarInString:language:details:"))
}

// Gives the delegate the opportunity to customize the grammatical analysis of a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1409242-spellserver?language=objc
func (s_ SpellServerDelegateObject) SpellServerCheckGrammarInStringLanguageDetails(sender SpellServer, stringToCheck string, language string, details unsafe.Pointer) Range {
	rv := objc.Call[Range](s_, objc.Sel("spellServer:checkGrammarInString:language:details:"), sender, stringToCheck, language, details)
	return rv
}

func (s_ SpellServerDelegateObject) HasSpellServerDidLearnWordInLanguage() bool {
	return s_.RespondsToSelector(objc.Sel("spellServer:didLearnWord:inLanguage:"))
}

// Notifies the delegate that the sender has added the specified word to the user’s list of acceptable words in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1407851-spellserver?language=objc
func (s_ SpellServerDelegateObject) SpellServerDidLearnWordInLanguage(sender SpellServer, word string, language string) {
	objc.Call[objc.Void](s_, objc.Sel("spellServer:didLearnWord:inLanguage:"), sender, word, language)
}

func (s_ SpellServerDelegateObject) HasSpellServerDidForgetWordInLanguage() bool {
	return s_.RespondsToSelector(objc.Sel("spellServer:didForgetWord:inLanguage:"))
}

// Notifies the delegate that the sender has removed the specified word from the user’s list of acceptable words in the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1417315-spellserver?language=objc
func (s_ SpellServerDelegateObject) SpellServerDidForgetWordInLanguage(sender SpellServer, word string, language string) {
	objc.Call[objc.Void](s_, objc.Sel("spellServer:didForgetWord:inLanguage:"), sender, word, language)
}

func (s_ SpellServerDelegateObject) HasSpellServerCheckStringOffsetTypesOptionsOrthographyWordCount() bool {
	return s_.RespondsToSelector(objc.Sel("spellServer:checkString:offset:types:options:orthography:wordCount:"))
}

// Gives the delegate the opportunity to analyze both the spelling and grammar simultaneously, which is more efficient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsspellserverdelegate/1409733-spellserver?language=objc
func (s_ SpellServerDelegateObject) SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(sender SpellServer, stringToCheck string, offset uint, checkingTypes TextCheckingTypes, options map[string]objc.Object, orthography Orthography, wordCount *int) []TextCheckingResult {
	rv := objc.Call[[]TextCheckingResult](s_, objc.Sel("spellServer:checkString:offset:types:options:orthography:wordCount:"), sender, stringToCheck, offset, checkingTypes, options, orthography, wordCount)
	return rv
}
