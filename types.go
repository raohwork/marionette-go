package marionette

import "encoding/json"

// Proxy represents proxy info
type Proxy struct {
	Type          string   `json:"proxyType,omitempty"`
	AutoconfigUrl string   `json:"proxyAutoconfigUrl,omitempty"`
	FTP           string   `json:"ftpProxy,omitempty"`
	FTPPort       int      `json:"ftpProxyPort,omitempty"`
	HTTP          string   `json:"httpProxy,omitempty"`
	HTTPPort      int      `json:"httpProxyPort,omitempty"`
	SSL           string   `json:"sslProxy,omitempty"`
	SSLPort       int      `json:"sslProxyPort,omitempty"`
	Socks         string   `json:"socksProxy,omitempty"`
	SocksPort     int      `json:"socksProxyPort,omitempty"`
	SocksVersion  string   `json:"socksVersion,omitempty"`
	NoProxy       []string `json:"noProxy,omitempty"`
}

// Timeouts represnets timeout values
//
// Refer driver.js for further info
// https://github.com/mozilla/gecko-dev/blob/master/testing/marionette/driver.js
type Timeouts struct {
	Implicit int `json:"implicit,omitempty"`
	PageLoad int `json:"pageLoad,omitempty"`
	Script   int `json:"script,omitempty"`
}

// Capabilities represents marionette server capabilities
type Capabilities struct {
	// web driver
	BrowserName               string    `json:"browserName,omitempty"`
	BrowserVersion            string    `json:"browserVersion,omitempty"`
	PlatformName              string    `json:"platformName,omitempty"`
	PlatformVersion           string    `json:"platformversion,omitempty"`
	AcceptInsecureCerts       bool      `json:"acceptInsecureCerts,omitempty"`
	PageLoadStrategy          string    `json:"pageLoadStrategy,omitempty"`
	Proxy                     *Proxy    `json:"proxy,omitempty"`
	SetWindowRect             bool      `json:"setWindowRect,omitempty"`
	Timeouts                  *Timeouts `json:"timeouts,omitempty"`
	StrictFileInteractability bool      `json:"strictFileInteractability,omitempty"`
	UnhandledPromptBehavior   string    `json:"unhandledPromptBehavior,omitempty"`

	// features
	Rotatable bool `json:"rotatable,omitempty"`

	// proprietary
	AccessibilityChecks  bool   `json:"moz:accessibilityChecks,omitempty"`
	BuildID              string `json:"buildID,omitempty"`
	Headless             bool   `json:"headless,omitempty"`
	ProcessID            int    `json:"processID,omitempty"`
	Profile              string `json:"profile,omitempty"`
	ShutdownTimeout      int    `json:"shutdownTimeout,omitempty"`
	SpecialPointerOrigin bool   `json:"moz:useNonSpecCompliantPointerOrigin,omitempty"`
	WebdriverClick       bool   `json:"moz:webdriverClick,omitempty"`
}

// Rect represents size/placement info about a window/element/...
type Rect struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"width"`
	H float64 `json:"height"`
}

// FindStrategy denotes how you find element
type FindStrategy string

const (
	ClassName       FindStrategy = "class name"
	Selector        FindStrategy = "css selector"
	ID              FindStrategy = "id"
	Name            FindStrategy = "name"
	LinkText        FindStrategy = "link text"
	PartialLinkText FindStrategy = "partial link text"
	TagName         FindStrategy = "tag name"
	XPath           FindStrategy = "xpath"
	Anon            FindStrategy = "anon"
	AnonAttribute   FindStrategy = "anon attribute"
)

const (
	// Type of WebElements
	ElementType       = "element-6066-11e4-a52e-4f735466cecf"
	WindowType        = "window-fcc6-11e5-b4f8-330a88ab9d7f"
	FrameType         = "frame-075b-4da1-b6ba-e579c2d3230a"
	ChromeElementType = "chromeelement-9fc5-4b51-a3c8-01716eedeb04"
)

// WebElement is an element (window/frame/html element) referenced by an UUID
type WebElement struct {
	Type string
	UUID string
}

func (el *WebElement) MarshalJSON() (data []byte, err error) {
	return json.Marshal(el.UUID)
}

type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
	HttpOnly bool   `json:"httpOnly,omitempty"`
	Expiry   int64  `json:"expiry,omitempty"`
}

const (
	// screen orientations
	PORTRAIT            = "portrait"
	LANDSCAPE           = "landscape"
	PORTRAIT_PRIMARY    = "portrait-primary"
	LANDSCAPE_PRIMARY   = "landscape-primary"
	PORTRAIT_SECONDARY  = "portrait-secondary"
	LANDSCAPE_SECONDARY = "landscape-secondary"
)

const (
	// window types
	FirefoxWindow     = "navigator:browser"
	GeckoViewWindow   = "navigator:geckoview"
	ThunderbirdWindow = "mail:3pane"
)
