package marionette

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

type Timeouts struct {
	Implicit int `json:"implicit,omitempty"`
	PageLoad int `json:"pageLoad,omitempty"`
	Script   int `json:"script,omitempty"`
}

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
