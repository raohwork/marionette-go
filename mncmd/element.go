// This file is part of marionette-go
//
// marionette-go is distributed in two licenses: The Mozilla Public License,
// v. 2.0 and the GNU Lesser Public License.
//
// marionette-go is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.
//
// See License.txt for further information.

package mncmd

import marionette "github.com/raohwork/marionette-go/v3"

// ElementClear defines "WebDriver:ElementClear" command
//
// See GeckoDriver.prototype.clearElement
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2545
type ElementClear struct {
	Element *marionette.WebElement
}

func (c *ElementClear) Command() (ret string) {
	return "WebDriver:ElementClear"
}

func (c *ElementClear) Param() (ret interface{}) {
	return map[string]interface{}{
		"id": c.Element.UUID,
	}
}

func (c *ElementClear) Validate() (ok bool) {
	return c.Element != nil
}

// ElementClick defines "WebDriver:ElementClick" command
//
// See GeckoDriver.prototype.clickElement
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2094
type ElementClick struct {
	Element *marionette.WebElement
}

func (c *ElementClick) Command() (ret string) {
	return "WebDriver:ElementClick"
}

func (c *ElementClick) Param() (ret interface{}) {
	return map[string]interface{}{
		"id": c.Element.UUID,
	}
}

func (c *ElementClick) Validate() (ok bool) {
	return c.Element != nil
}

// ElementSendKeys defines "WebDriver:ElementSendKeys" command
//
// See GeckoDriver.prototype.sendKeysToElement
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2504
type ElementSendKeys struct {
	Element *marionette.WebElement `json:"id"`
	Text    string                 `json:"text"`
}

func (c *ElementSendKeys) Command() (ret string) {
	return "WebDriver:ElementSendKeys"
}

func (c *ElementSendKeys) Param() (ret interface{}) {
	return c
}

func (c *ElementSendKeys) Validate() (ok bool) {
	return c.Element != nil && c.Text != ""
}

// FindElement defines "WebDriver:FindElement" command
//
// See GeckoDriver.prototype.findElement
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L1974
type FindElement struct {
	Using       marionette.FindStrategy `json:"using"`
	Value       string                  `json:"value"`
	RootElement *marionette.WebElement  `json:"element,omitempty"`
	returnElem
}

func (c *FindElement) Command() (ret string) {
	return "WebDriver:FindElement"
}

func (c *FindElement) Param() (ret interface{}) {
	return c
}

func (c *FindElement) Validate() (ok bool) {
	return c.Using != "" && c.Value != ""
}

// FindElements defines "WebDriver:FindElements" command
//
// See GeckoDriver.prototype.findElements
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2024
type FindElements struct {
	Using       marionette.FindStrategy `json:"using"`
	Value       string                  `json:"value"`
	RootElement *marionette.WebElement  `json:"element,omitempty"`
	returnElems
}

func (c *FindElements) Command() (ret string) {
	return "WebDriver:FindElements"
}

func (c *FindElements) Param() (ret interface{}) {
	return c
}

func (c *FindElements) Validate() (ok bool) {
	return c.Using != "" && c.Value != ""
}

// GetActiveElement defines "WebDriver:GetActiveElement" command
//
// See GeckoDriver.prototype.getActiveElement
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2069
type GetActiveElement struct {
	noParam
	returnElem
}

func (c *GetActiveElement) Command() (ret string) {
	return "WebDriver:GetActiveElement"
}

// GetElementAttribute defines "WebDriver:GetElementAttribute" command
//
// See GeckoDriver.prototype.getElementAttribute
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2148
type GetElementAttribute struct {
	Element *marionette.WebElement `json:"id"`
	Name    string                 `json:"name"`
	returnStr
}

func (c *GetElementAttribute) Command() (ret string) {
	return "WebDriver:GetElementAttribute"
}

func (c *GetElementAttribute) Param() (ret interface{}) {
	return c
}

func (c *GetElementAttribute) Validate() (ok bool) {
	return c.Element != nil && c.Name != ""
}

// GetElementCSSValue defines "WebDriver:GetElementCSSValue" command
//
// See GeckoDriver.prototype.getElementCSSValue
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2348
type GetElementCSSValue struct {
	Element *marionette.WebElement `json:"id"`
	Prop    string                 `json:"propertyName"`
	returnStr
}

func (c *GetElementCSSValue) Command() (ret string) {
	return "WebDriver:GetElementCSSValue"
}

func (c *GetElementCSSValue) Param() (ret interface{}) {
	return c
}

func (c *GetElementCSSValue) Validate() (ok bool) {
	return c.Element != nil && c.Prop != ""
}

// GetElementProperty defines "WebDriver:GetElementProperty" command
//
// See GeckoDriver.prototype.getElementProperty
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2189
type GetElementProperty struct {
	Element *marionette.WebElement `json:"id"`
	Name    string                 `json:"name"`
	returnMixed
}

func (c *GetElementProperty) Command() (ret string) {
	return "WebDriver:GetElementProperty"
}

func (c *GetElementProperty) Param() (ret interface{}) {
	return c
}

func (c *GetElementProperty) Validate() (ok bool) {
	return c.Element != nil && c.Name != ""
}

// GetElementRect defines "WebDriver:GetElementRect" command
//
// See GeckoDriver.prototype.getElementRect
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2468
type GetElementRect struct {
	Element *marionette.WebElement `json:"id"`
	returnRect
}

func (c *GetElementRect) Command() (ret string) {
	return "WebDriver:GetElementRect"
}

func (c *GetElementRect) Param() (ret interface{}) {
	return c
}

func (c *GetElementRect) Validate() (ok bool) {
	return c.Element != nil
}

// GetElementTagName defines "WebDriver:GetElementTagName" command
//
// See GeckoDriver.prototype.getElementTagName
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2272
type GetElementTagName struct {
	Element *marionette.WebElement `json:"id"`
	returnStr
}

func (c *GetElementTagName) Command() (ret string) {
	return "WebDriver:GetElementTagName"
}

func (c *GetElementTagName) Param() (ret interface{}) {
	return c
}

func (c *GetElementTagName) Validate() (ok bool) {
	return c.Element != nil
}

// GetElementText defines "WebDriver:GetElementText" command
//
// See GeckoDriver.prototype.getElementText
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2230
type GetElementText struct {
	Element *marionette.WebElement `json:"id"`
	returnStr
}

func (c *GetElementText) Command() (ret string) {
	return "WebDriver:GetElementText"
}

func (c *GetElementText) Param() (ret interface{}) {
	return c
}

func (c *GetElementText) Validate() (ok bool) {
	return c.Element != nil
}

// IsElementDisplayed defines "WebDriver:IsElementDisplayed" command
//
// See GeckoDriver.prototype.isElementDisplayed
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2310
type IsElementDisplayed struct {
	Element *marionette.WebElement `json:"id"`
	returnBool
}

func (c *IsElementDisplayed) Command() (ret string) {
	return "WebDriver:IsElementDisplayed"
}

func (c *IsElementDisplayed) Param() (ret interface{}) {
	return c
}

func (c *IsElementDisplayed) Validate() (ok bool) {
	return c.Element != nil
}

// IsElementEnabled defines "WebDriver:IsElementEnabled" command
//
// See GeckoDriver.prototype.isElementEnabled
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2390
type IsElementEnabled struct {
	Element *marionette.WebElement `json:"id"`
	returnBool
}

func (c *IsElementEnabled) Command() (ret string) {
	return "WebDriver:IsElementEnabled"
}

func (c *IsElementEnabled) Param() (ret interface{}) {
	return c
}

func (c *IsElementEnabled) Validate() (ok bool) {
	return c.Element != nil
}

// IsElementSelected defines "WebDriver:IsElementSelected" command
//
// See GeckoDriver.prototype.isElementSelected
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js#L2429
type IsElementSelected struct {
	Element *marionette.WebElement `json:"id"`
	returnBool
}

func (c *IsElementSelected) Command() (ret string) {
	return "WebDriver:IsElementSelected"
}

func (c *IsElementSelected) Param() (ret interface{}) {
	return c
}

func (c *IsElementSelected) Validate() (ok bool) {
	return c.Element != nil
}
