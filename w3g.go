package w3g

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// Accept request HTTP header advertises which content types, expressed as MIME types, the client is able to understand.
const Accept string = "Accept"

// AcceptCH HTTP header is set by the server to specify which Client Hints headers client should include in subsequent requests.
const AcceptCH string = "Accept-CH"

// AcceptCHLifetime HTTP header is set by the server to specify the persistence of Accept-CH header value
// that specifies for which Client Hints headers client should include in subsequent requests.
const AcceptCHLifetime string = "Accept-CH-Lifetime"

// AcceptCharset request HTTP header advertises which character encodings the client understands.
const AcceptCharset string = "Accept-Charset"

// AcceptEncoding request HTTP header advertises which content encoding,
// usually a compression algorithm, the client is able to understand.
const AcceptEncoding string = "Accept-Encoding"

// AcceptLanguage request HTTP header advertises which languages the client is able to understand,
// and which locale variant is preferred.
const AcceptLanguage string = "Accept-Language"

// AcceptPatch response HTTP header advertises which media-type the server is able to understand.
const AcceptPatch string = "Accept-Patch"

// AcceptRanges response HTTP header is a marker used by the server to advertise its support of partial requests.
const AcceptRanges string = "Accept-Ranges"

// AccessControlAllowCredentials response HTTP header tells browsers whether to expose the response to frontend JavaScript code
// when the request's credentials mode is include.
const AccessControlAllowCredentials string = "Access-Control-Allow-Credentials"

// AccessControlAllowHeaders response HTTP header is used in response to a preflight request which includes the
// Access-Control-Request-Headers to indicate which HTTP headers can be used during the actual request.
const AccessControlAllowHeaders string = "Access-Control-Allow-Headers"

// AccessControlAllowMethod response HTTP header specifies the method or methods allowed when accessing
// the resource in response to a preflight request.
const AccessControlAllowMethod string = "Access-Control-Allow-Method"

// AccessControlAllowOrigin response HTTP header indicates whether the response can be shared with requesting code from the given origin.
const AccessControlAllowOrigin string = "Access-Control-Allow-Origin"

// AccessControlExposeHeaders response HTTP header indicates which headers can be exposed as part of the response by listing their names.
const AccessControlExposeHeaders string = "Access-Control-Expose-Headers"

// AccessControlMaxAge response HTTP header indicates how long the results of a preflight request
// (that is the information contained in the Access-Control-Allow-Methods and Access-Control-Allow-Headers headers) can be cached.
const AccessControlMaxAge string = "Access-Control-Max-Age"

// AccessControlRequestHeaders request HTTP header is used by browsers when issuing a preflight request, to let the server know which
// HTTP headers the client might send when the actual request is made.
const AccessControlRequestHeaders string = "Access-Control-Request-Headers"

// AccessControlRequestMethod request HTTP header is used by browsers when issuing a preflight request, to let the server know which
// HTTP method will be used when the actual request is made.
const AccessControlRequestMethod string = "Access-Control-Request-Method"

// Age HTTP header contains the time in seconds the object has been in a proxy cache.
const Age string = "Age"

// Allow HTTP header lists the set of methods supported by a resource.
const Allow string = "Allow"

// AltSvc HTTP response header is used to advertise alternative services through which the same resource can be reached.
const AltSvc string = "Alt-Svc"

// Authorization HTTP request header contains the credentials to authenticate a user agent with a server.
const Authorization string = "Authorization"

// CacheControl HTTP header holds directives (instructions) for caching in both requests and responses.
const CacheControl string = "Cache-Control"

// ClearSiteData HTTP header clears browsing data (cookies, storage, cache) associated with the requesting website
const ClearSiteData string = "Clear-Site-Data"

// Connection HTTP header controls whether or not the network connection stays open after the current transaction finishes.
const Connection string = "Connection"

// ContentDisposition HTTP response header is a header indicating if the content is expected to be displayed inline in the browser,
// that is, as a Web page or as part of a Web page, or as an attachment, that is downloaded and saved locally.
const ContentDisposition string = "Content-Disposition"

// ContentDPR HTTP header is a header that indicates the ratio between physical pixels over CSS pixels of the selected image.
const ContentDPR string = "Content-DPR"

// ContentEncoding HTTP header is used to compress the media-type.
const ContentEncoding string = "Content-Encoding"

// ContentLanguage HTTP header is used to describe the language(s) intended for the audience, so that it allows a user to
// differentiate according to the users' own preferred language.
const ContentLanguage string = "Content-Language"

// ContentLength HTTP header indicates the size of the entity-body, in bytes, sent to the recipient.
const ContentLength string = "Content-Length"

// ContentLocation header indicates an alternate location for the returned data.
const ContentLocation string = "Content-Location"

// ContentRange response HTTP header indicates where in a full body message a partial message belongs.
const ContentRange string = "Content-Range"

// ContentSecurityPolicy response header allows web site administrators to control resources the user agent
// is allowed to load for a given page.
const ContentSecurityPolicy string = "Content-Security-Policy"

// ContentSecurityPolicyReportOnly response header allows web developers to experiment with policies by monitoring (but not enforcing) their effects.
const ContentSecurityPolicyReportOnly string = "Content-Security-Policy-Report-Only"

// ContentType HTTP header is used to indicate the media type of the resource.
const ContentType string = "Content-Type"

// Cookie HTTP request header contains stored HTTP cookies previously sent by the server with the Set-Cookie header.
const Cookie string = "Cookie"

// Cookie2 HTTP request header used to advise the server that the user agent understands "new-style" cookies.
const Cookie2 string = "Cookie2"

// CrossOriginResourcePolicy HTTP response header conveys a desire that the browser blocks no-cors cross-origin/cross-site requests to the given resource.
const CrossOriginResourcePolicy string = "Cross-Origin-Resource-Policy"

// DNT HTTP request header indicates the user's tracking preference.
const DNT string = "DNT"

// DPR HTTP header is a Client Hints headers which represents the client device pixel ratio (DPR), which is the the number of
// physical device pixels corresponding to every CSS pixel.
const DPR string = "DPR"

// Date HTTP header contains the date and time at which the message was originated.
const Date string = "Date"

// DeviceMemory HTTP header is a Device Memory API header that works like Client Hints header which represents the approximate amount of RAM client device has.
const DeviceMemory string = "Device-Memory"

// Digest response HTTP header provides a digest of the requested resource.
const Digest string = "Digest"

// ETag HTTP response header is an identifier for a specific version of a resource.
const ETag string = "ETag"

// EarlyData HTTP header is set by an intermediate to indicate that the request has been conveyed in TLS early data.
const EarlyData string = "Early-Data"

// Expect HTTP request header indicates expectations that need to be fulfilled by the server in order to properly handle the request.
const Expect string = "Expect"

// ExpectCT HTTP header allows sites to opt in to reporting and/or enforcement of Certificate Transparency requirements.
const ExpectCT string = "Expect-CT"

// Expires HTTP header contains the date/time after which the response is considered stale.
const Expires string = "Expires"

// FeaturePolicy HTTP header provides a mechanism to allow and deny the use of browser features in its own frame.
const FeaturePolicy string = "Feature-Policy"

// Forwarded HTTP header contains information from the client-facing side of proxy servers that is altered or lost when a proxy is involved in the path of the request.
const Forwarded string = "Forwarded"

// From HTTP request header contains an Internet email address for a human user who controls the requesting user agent.
const From string = "From"

// Host HTTP request header specifies the domain name of the server (for virtual hosting), and (optionally) the TCP port number on which the server is listening.
const Host string = "Host"

// IfMatch HTTP request header makes the request conditional.
const IfMatch string = "If-Match"

// IfModifiedSince request HTTP header makes the request conditional.
const IfModifiedSince string = "If-Modified-Since"

// IfNoneMatch HTTP request header makes the request conditional.
const IfNoneMatch string = "If-None-Match"

// IfRange HTTP request header makes a range request conditional.
const IfRange string = "If-Range"

// IfUnmodifiedSince request HTTP header makes the request conditional.
const IfUnmodifiedSince string = "If-Unmodified-Since"

// KeepAlive HTTP header allows the sender to hint about how the connection may be used to set a timeout and a maximum amount of requests.
const KeepAlive string = "Keep-Alive"

// LargeAllocation response HTTP header tells the browser that the page being loaded is going to want to perform a large allocation.
const LargeAllocation string = "Large-Allocation"

// LastModified response HTTP header contains the date and time at which the origin server believes the resource was last modified.
const LastModified string = "Last-Modified"

// Link HTTP header field provides a means for serialising one or more links in HTTP headers.
const Link string = "Link"

// Location response HTTP header indicates the URL to redirect a page to.
const Location string = "Location"

// Origin request HTTP header indicates where a fetch originates from.
const Origin string = "Origin"

// Pragma HTTP header is an implementation-specific header that is used for backwards compatibility with HTTP/1.0.
const Pragma string = "Pragma"

// ProxyAuthenticate response HTTP header defines the authentication method that should be used to gain access to a resource behind a proxy server.
const ProxyAuthenticate string = "Proxy-Authenticate"

// ProxyAuthorization request HTTP header contains the credentials to authenticate a user agent to a proxy server.
const ProxyAuthorization string = "Proxy-Authorization"

// PublicKeyPins response HTTP header associates a specific cryptographic public key with a certain web server.
const PublicKeyPins string = "Public-Key-Pins"

// PublicKeyPinsReportOnly response HTTP header sends reports of pinning violation to the report-uri specified in the header.
const PublicKeyPinsReportOnly string = "Public-Key-Pins-Report-Only"

// Range HTTP request header indicates the part of a document that the server should return.
const Range string = "Range"

// Referer request HTTP header contains the address of the previous web page from which a link to the currently requested page was followed.
const Referer string = "Referer"

// ReferrerPolicy HTTP header controls how much referrer information should be included with requests.
const ReferrerPolicy string = "Referrer-Policy"

// RetryAfter response HTTP header indicates how long the user agent should wait before making a follow-up request.
const RetryAfter string = "Retry-After"

// SaveData HTTP header field is a boolean which, in requests, indicates the client's preference for reduced data usage.
const SaveData string = "Save-Data"

// SecFetchDest HTTP header indicates the request's destination.
const SecFetchDest string = "Sec-Fetch-Dest"

// SecFetchMode HTTP header indicates the request's mode.
const SecFetchMode string = "Sec-Fetch-Mode"

// SecFetchSite HTTP header indicates the relationship between a request initiator's origin and the origin of the resource.
const SecFetchSite string = "Sec-Fetch-Site"

// SecFetchUser HTTP header indicates whether or not a navigation request was triggered by a user activation.
const SecFetchUser string = "Sec-Fetch-User"

// SecWebSocketAccept WS header is used in the websocket opening handshake.
const SecWebSocketAccept string = "Sec-WebSocket-Accept"

// Server response HTTP header contains information about the software used by the origin server to handle the request.
const Server string = "Server"

// ServerTiming response HTTP header communicates one or more metrics and descriptions for a given request-response cycle.
const ServerTiming string = "Server-Timing"

// SetCookie HTTP response header is used to send cookies from the server to the user agent, so the user agent can send them back to the server later.
const SetCookie string = "Set-Cookie"

// SetCookie2 HTTP response header used to send cookies from the server to the user agent.
const SetCookie2 string = "Set-Cookie2"

// SourceMap HTTP response header links generated code to a source map, enabling the browser to reconstruct the original source and present
// the reconstructed original in the debugger.
const SourceMap string = "SourceMap"

// StrictTransportSecurity response HTTP header lets a web site tell browsers that it should only be accessed using HTTPS, instead of using HTTP.
const StrictTransportSecurity string = "Strict-Transport-Security"

// TE request HTTP header specifies the transfer encodings the user agent is willing to accept.
const TE string = "TE"

// TimingAllowOrigin response HTTP header specifies origins that are allowed to see values of attributes retrieved via features of the Resource Timing API.
const TimingAllowOrigin string = "Timing-Allow-Origin"

// Tk response HTTP header indicates the tracking status that applied to the corresponding request.
const Tk string = "Tk"

// Trailer response HTTP header allows the sender to include additional fields at the end of chunked messages.
const Trailer string = "Trailer"

// TransferEncoding response HTTP header specifies the form of encoding used to safely transfer the payload body to the user.
const TransferEncoding string = "Transfer-Encoding"

// UpgradeInsecureRequests request HTTP header sends a signal to the server expressing the client’s preference for an encrypted and authenticated response.
const UpgradeInsecureRequests string = "Upgrade-Insecure-Requests"

// UserAgent request HTTP header is a characteristic string that lets servers and network peers identify the application, operating system, vendor, and/or version of the requesting user agent.
const UserAgent string = "User-Agent"

// Vary response HTTP header determines how to match future request headers to decide whether a cached response can be used rather than requesting a fresh one from the origin server.
const Vary string = "Vary"

// ViewportWidth response HTTP header is a HTTP header that indicates the layout viewport width in CSS pixels.
const ViewportWidth string = "Viewport-Width"

// Via HTTP header is added by proxies, both forward and reverse proxies.
const Via string = "Via"

// WWWAuthenticate response HTTP header defines the authentication method that should be used to gain access to a resource.
const WWWAuthenticate string = "WWW-Authenticate"

// WantDigest HTTP header is used to ask the responder to provide a digest of the requested resource.
const WantDigest string = "Want-Digest"

// Warning HTTP header contains information about possible problems with the status of the message.
const Warning string = "Warning"

// Width request HTTP header is a number that indicates the desired resource width in physical pixels.
const Width string = "Width"

// XContentTypeOptions response HTTP header is a marker used by the server to indicate that the MIME types advertised in the Content-Type headers should not be changed.
const XContentTypeOptions string = " X-Content-Type-Options"

// XDNSPrefetchControl response HTTP header controls DNS prefetching.
const XDNSPrefetchControl string = "X-DNS-Prefetch-Control"

// XForwardedFor HTTP header is a standard header for identifying the originating IP address of a client connecting to a web server through an HTTP proxy or a load balancer.
const XForwardedFor string = "X-Forwarded-For"

// XForwardedHost HTTP header is a standard header for identifying the original host requested by the client.
const XForwardedHost string = "X-Forwarded-Host"

// XForwardedProto HTTP header is a standard header for identifying the protocol that a client used to connect to your proxy or load balancer.
const XForwardedProto string = "X-Forwarded-Proto"

// XFrameOptions response HTTP header can be used to indicate whether or not a browser should be allowed to render a page in a <frame>, <iframe>, <embed> or <object>.
const XFrameOptions string = "X-Frame-Options"

// XXSSProtection response HTTP header is a feature of Internet Explorer, Chrome and Safari that stops pages from loading when they detect reflected cross-site scripting.
const XXSSProtection string = "X-XSS-Protection"

// AcceptHeader is a struct to prepare a Accept HTTP request header.
type AcceptHeader struct {
	MIMESubType string  `json:"mime_subtype"`
	MIMEType    string  `json:"mime_type"`
	Q           float32 `json:"q"`
}

// String returns a string representation of a Accept HTTP header value.
func (a AcceptHeader) String() string {
	var mimeSubTypeLengthOK, mimeTypeLengthOK, qOK bool = (len(a.MIMEType) != 0), (len(a.MIMESubType) != 0), (a.Q != 0)
	var mimeSubType, mimeType string = "*", "*"
	var substrings ([]string) = (make([]string, 2))
	var s string
	if mimeTypeLengthOK {
		mimeType = a.MIMEType
	}
	substrings[0] = (mimeType)
	if mimeSubTypeLengthOK {
		mimeSubType = a.MIMESubType
	}
	substrings[1] = (mimeSubType)
	s = (strings.Join(substrings, "/"))
	if qOK {
		s = (fmt.Sprintf("%s;q=%1.1f", s, a.Q))
	}
	return s
}

// AcceptCHHeader is a struct to prepare a Accept-CH HTTP response header.
type AcceptCHHeader struct {
	AcceptCH         bool `json:"accept_ch"`
	AcceptCHLifetime bool `json:"accept_ch_lifetime"`
	ContentDPR       bool `json:"content_dpr"`
	DeviceMemory     bool `json:"device_memory"`
	DPR              bool `json:"dpr"`
	EarlyData        bool `json:"early_data"`
	SaveData         bool `json:"save_data"`
	ViewportWidth    bool `json:"viewport_width"`
	Width            bool `json:"width"`
}

// String returns a string representation of a Accept-CH HTTP header value.
func (a AcceptCHHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var substring string
	var s string
	var r reflect.Value = reflect.ValueOf(a)
	for _, substring = range []string{AcceptCH, AcceptCHLifetime, ContentDPR, DPR, EarlyData, SaveData, ViewportWidth, Width} {
		var f reflect.Value = (r.FieldByName(strings.ReplaceAll(substring, "-", "")))
		if f.IsValid() && f.Bool() {
			substrings = (append(substrings, substring))
		}
	}
	s = (strings.Join(substrings, ", "))
	return s
}

// AcceptCHLifetimeHeader is a struct to prepare a Accept-CH-Lifetime HTTP response header.
type AcceptCHLifetimeHeader struct {
	Age int64 `json:"age"`
}

// String returns a string representation of a Accept-CH-Lifetime HTTP header value.
func (a AcceptCHLifetimeHeader) String() string {
	return (fmt.Sprintf("%d", a.Age))
}

// AcceptCharsetHeader is a struct to prepare a Accept-Charset HTTP request header.
type AcceptCharsetHeader struct {
	Charset string  `json:"charset"`
	Q       float32 `json:"q"`
}

// String returns a string representation of a Accept-Charset HTTP header value.
func (a AcceptCharsetHeader) String() string {
	var charsetOK, qOK bool = (len(a.Charset) != 0), (a.Q != 0)
	var charset string = "*"
	var substrings ([]string) = (make([]string, 1))
	var s string
	if charsetOK {
		charset = a.Charset
	}
	(substrings[0]) = charset
	if qOK {
		(substrings) = (append(substrings, fmt.Sprintf("q=%1.1f", a.Q)))
	}
	s = (strings.Join(substrings, "/"))
	return s
}

// AcceptEncodingHeader is a struct to prepare a Accept-Encoding HTTP request header.
type AcceptEncodingHeader struct {
	Encoding string  `json:"encoding"`
	Q        float32 `json:"q"`
}

// String returns a string representation of a Accept-Charset HTTP header value.
func (a AcceptEncodingHeader) String() string {
	var encodingOK, qOK bool = (len(a.Encoding) != 0), (a.Q != 0)
	var encoding string = "*"
	var substrings ([]string) = (make([]string, 1))
	var s string
	if encodingOK {
		encoding = a.Encoding
	}
	(substrings[0]) = encoding
	if qOK {
		(substrings) = (append(substrings, fmt.Sprintf("q=%1.1f", a.Q)))
	}
	s = (strings.Join(substrings, "/"))
	return s
}

// AcceptLanguageHeader is a struct to prepare a Accept-Language HTTP request header.
type AcceptLanguageHeader struct {
	Language string  `json:"language"`
	Q        float32 `json:"q"`
}

// String returns a string representation of a Accept-Language HTTP header value.
func (a AcceptLanguageHeader) String() string {
	var languageOK, qOK bool = (len(a.Language) != 0), (a.Q != 0)
	var language string = "*"
	var substrings ([]string) = (make([]string, 1))
	var s string
	if languageOK {
		language = a.Language
	}
	(substrings[0]) = language
	if qOK {
		(substrings) = (append(substrings, fmt.Sprintf("q=%1.1f", a.Q)))
	}
	s = (strings.Join(substrings, "/"))
	return s
}

// AcceptPatchHeader is a struct to prepare a Accept-Patch HTTP header value.
type AcceptPatchHeader struct {
	MIMESubType string `json:"mime_subtype"`
	MIMEType    string `json:"mime_type"`
	Charset     string `json:"charset"`
}

// String returns a string representation of a Accept-Patch HTTP header value.
func (a AcceptPatchHeader) String() string {
	var charsetOK, mimeSubTypeLengthOK, mimeTypeLengthOK bool = (len(a.Charset) != 0), (len(a.MIMEType) != 0), (len(a.MIMESubType) != 0)
	var mimeSubType, mimeType string = "*", "*"
	var substrings ([]string) = (make([]string, 0))
	var s string
	if mimeTypeLengthOK {
		mimeType = a.MIMEType
	}
	(substrings) = (append(substrings, (mimeType)))
	if mimeSubTypeLengthOK {
		mimeSubType = a.MIMESubType
	}
	(substrings) = (append(substrings, (mimeSubType)))
	s = (strings.Join(substrings, "/"))
	(substrings) = ([]string{s})
	if charsetOK {
		(substrings) = (append(substrings, fmt.Sprintf("charset=%s", a.Charset)))
	}
	s = (strings.Join(substrings, ";"))
	return s
}

// AcceptRangesHeader is a struct to prepare a Accept-Ranges HTTP header value.
type AcceptRangesHeader struct {
	Bytes bool `json:"bytes"`
}

// String returns a string representation of a Accept-Ranges HTTP header value.
func (a AcceptRangesHeader) String() string {
	if a.Bytes {
		return "bytes"
	}
	return "none"
}

// AccessControlAllowCredentialsHeader is a struct to prepare a Access-Control-Allow-Credentials HTTP header value.
type AccessControlAllowCredentialsHeader struct {
	Allow bool `json:"allow"`
}

// String returns a string representation of a Access-Control-Allow-Credentials HTTP header value.
func (a AccessControlAllowCredentialsHeader) String() string {
	return (fmt.Sprintf("%t", a.Allow))
}

// AccessControlAllowHeadersHeader is a struct to prepare a Access-Control-Allow-Headers HTTP header value.
type AccessControlAllowHeadersHeader struct {
	Headers []string `json:"headers"`
}

// String returns a string representation of a Access-Control-Allow-Headers HTTP header value.
func (a AccessControlAllowHeadersHeader) String() string {
	var headersOK bool = (len(a.Headers) != 0)
	var s string = "*"
	if headersOK {
		s = (strings.Join(a.Headers, ", "))
	}
	return s
}

// AcceptControlAllowOriginHeader is a struct to prepare a Accept-Control-Allow-Origin HTTP header value.
type AcceptControlAllowOriginHeader struct {
	Origin string `json:"origin"`
}

// String returns a string representation of a Access-Control-Allow-Origin HTTP header value.
func (a AcceptControlAllowOriginHeader) String() string {
	var originOK bool = (len(a.Origin) != 0)
	var s string = "*"
	if originOK {
		return (a.Origin)
	}
	return s
}

// AcceptControlExposeHeadersHeader is a struct to prepare a Access-Control-Expose-Headers HTTP header value.
type AcceptControlExposeHeadersHeader struct {
	Headers []string `json:"headers"`
}

// String returns a string representation of a Access-Control-Expose-Headers HTTP header value.
func (a AcceptControlExposeHeadersHeader) String() string {
	var headersOK bool = (len(a.Headers) != 0)
	var s string = "*"
	if headersOK {
		s = (strings.Join(a.Headers, ", "))
	}
	return s
}

// AccessControlMaxAgeHeader is a struct to prepare a Access-Control-Max-Age HTTP response header.
type AccessControlMaxAgeHeader struct {
	Age int64 `json:"age"`
}

// String returns a string representation of a Access-Control-Max-Age HTTP header value.
func (a AccessControlMaxAgeHeader) String() string {
	return (fmt.Sprintf("%d", a.Age))
}

// AcceptControlRequestHeadersHeader is a struct to prepare a Access-Control-Request-Headers HTTP header value.
type AcceptControlRequestHeadersHeader struct {
	Headers []string `json:"headers"`
}

// String returns a string representation of a Access-Control-Request-Headers HTTP header value.
func (a AcceptControlRequestHeadersHeader) String() string {
	var headersOK bool = (len(a.Headers) != 0)
	var s string = "*"
	if headersOK {
		s = (strings.Join(a.Headers, ", "))
	}
	return s
}

// AcceptControlRequestMethodHeader is a struct to prepare a Access-Control-Request-Method HTTP header value.
type AcceptControlRequestMethodHeader struct {
	Method string `json:"method"`
}

// String returns a string representation of a Access-Control-Request-Method HTTP header value.
func (a AcceptControlRequestMethodHeader) String() string {
	var methodsOK bool = (len(a.Method) != 0)
	if methodsOK {
		return a.Method
	}
	return http.MethodGet
}

// AgeHeader is a struct to prepare a Age HTTP response header.
type AgeHeader struct {
	Age int64 `json:"age"`
}

// String returns a string representation of a Age HTTP header value.
func (a AgeHeader) String() string {
	return (fmt.Sprintf("%d", a.Age))
}

// AllowHeader is a struct to prepare a Allow HTTP header value.
type AllowHeader struct {
	Methods []string `json:"methods"`
}

// String returns a string representation of a Allow HTTP header value.
func (a AllowHeader) String() string {
	var methods []string = (make([]string, 0))
	var methodsOK bool = (len(a.Methods) != 0)
	if methodsOK {
		methods = (a.Methods)
	}
	return strings.Join(methods, ", ")
}

// AltSvcHeader is a struct to prepare a Alt-Svc HTTP header value.
type AltSvcHeader struct {
	AltAuthority string `json:"alt_authority"`
	Clear        bool   `json:"clear"`
	MaxAge       int64  `json:"max_age"`
	Persist      bool   `json:"persist"`
	ProtocolID   string `json:"protocol_id"`
}

// String returns a string representation of a Alt-Svc HTTP header value.
func (a AltSvcHeader) String() string {
	var s string = "clear"
	if a.Clear {
		return s
	}
	s = (fmt.Sprintf(("%s=\"%s\""), a.ProtocolID, a.AltAuthority))
	if a.MaxAge != 0 {
		s = (fmt.Sprintf(("%s;%d"), s, a.MaxAge))
	}
	if a.Persist {
		s = (fmt.Sprintf(("%s;persist=1"), s))
	}
	return s
}

// AuthorizationHeader is a struct to prepare a Authorization HTTP header value.
type AuthorizationHeader struct {
	Credentials string `json:"credentials"`
	Type        string `json:"type"`
}

// String returns a string representation of a Authorization HTTP header value.
func (a AuthorizationHeader) String() string {
	return (fmt.Sprintf("%s %s", a.Type, a.Credentials))
}

// CacheControlHeader is a struct to prepare a Cache-Control HTTP header value.
type CacheControlHeader struct {
	Immutable            bool  `json:"immutable"`
	MaxAge               int64 `json:"max_age"`
	MaxStale             int64 `json:"max_stale"`
	MinFresh             int64 `json:"min_stale"`
	MustRevalidate       bool  `json:"must_revalidate"`
	NoCache              bool  `json:"no_cache"`
	NoStore              bool  `json:"no_store"`
	NoTransform          bool  `json:"no_transform"`
	OnlyIfCached         bool  `json:"only_if_cached"`
	Private              bool  `json:"private"`
	ProxyRevalidate      bool  `json:"proxy_revalidate"`
	Public               bool  `json:"public"`
	SMaxAge              int64 `json:"s_max_age"`
	StaleIfError         int64 `json:"stale_if_error"`
	StaleWhileRevalidate int64 `json:"stale_while_revalidate"`
}

// String returns a string representation of a Cache-Control HTTP header value.
func (c CacheControlHeader) String() string {
	var regex *regexp.Regexp = regexp.MustCompile(`[A-Z][^A-Z]*`)
	var substrings ([]string) = (make([]string, 0))
	var s string
	var r reflect.Value = reflect.ValueOf(c)
	var v reflect.Type = (reflect.Indirect(r).Type())
	for i, n := 0, r.NumField(); i < n; i++ {
		var f reflect.Value = r.Field(i)
		if f.IsValid() && !f.IsZero() {
			var name string = strings.Join(regex.FindAllString(v.Field(i).Name, -1), "-")
			name = strings.ToLower(name)
			switch f.Kind() {
			case reflect.Bool:
				(substrings) = (append(substrings, (name)))
			case reflect.Int64:
				(substrings) = (append(substrings, fmt.Sprintf("%s=%d", name, r.Int())))
			}
		}
	}
	(s) = (strings.Join(substrings, ", "))
	return s
}

// ClearSiteDataHeader is a struct to prepare a Clear-Site-Data HTTP header value.
type ClearSiteDataHeader struct {
	All               bool `json:"all"`
	ExecutionContexts bool `json:"execution_contexts"`
	Cache             bool `json:"cache"`
	Cookies           bool `json:"cookies"`
	Storage           bool `json:"storage"`
}

// String returns a string representation of a Clear-Site-Data HTTP header value.
func (c ClearSiteDataHeader) String() string {
	if c.All {
		return "*"
	}
	var substrings ([]string) = (make([]string, 0))
	var s string
	if c.ExecutionContexts {
		(substrings) = (append(substrings, ("executionContexts")))
	}
	if c.Cache {
		(substrings) = (append(substrings, ("cache")))
	}
	if c.Cookies {
		(substrings) = (append(substrings, ("cookies")))
	}
	if c.Storage {
		(substrings) = (append(substrings, ("storage")))
	}
	s = (strings.Join(substrings, ", "))
	return s
}

// ConnectionHeader is a struct to prepare a Connection HTTP header.
type ConnectionHeader struct {
	Close bool `json:"close"`
}

// String returns a string representation of a Connection HTTP header value.
func (c ConnectionHeader) String() string {
	if c.Close {
		return "close"
	}
	return "keep-alive"
}

// ContentDispositionHeader is a struct to prepare a Content-Disposition HTTP header.
type ContentDispositionHeader struct {
	Attachment bool   `json:"attachment"`
	FileName   string `json:"file_name"`
	FormData   bool   `json:"form_data"`
	Name       string `json:"name"`
}

// String returns a string representation of a Content-Disposition HTTP header value.
func (c ContentDispositionHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	if c.Attachment {
		(substrings) = (append(substrings, ("attachment")))
	}
	if !c.Attachment && c.FormData {
		(substrings) = (append(substrings, ("form-data")))
	}
	if len(c.FileName) != 0 {
		(substrings) = (append(substrings, fmt.Sprintf("filename=\"%s\"", c.FileName)))
	}
	if len(c.Name) != 0 {
		(substrings) = (append(substrings, fmt.Sprintf("name=\"%s\"", c.Name)))
	}
	s = (strings.Join(substrings, "; "))
	return s
}

// ContentEncodingHeader is a struct to prepare a Content-Encoding HTTP header.
type ContentEncodingHeader struct {
	Br       bool `json:"br"`
	Compress bool `json:"compress"`
	Deflate  bool `json:"deflate"`
	Identity bool `json:"identity"`
	GZip     bool `json:"gizp"`
}

// String returns a string representation of a Content-Encoding HTTP header value.
func (c ContentEncodingHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	if c.Br {
		(substrings) = (append(substrings, "br"))
	}
	if c.Compress {
		(substrings) = (append(substrings, "compress"))
	}
	if c.Deflate {
		(substrings) = (append(substrings, "deflate"))
	}
	if c.Identity {
		(substrings) = (append(substrings, "identity"))
	}
	if c.GZip {
		(substrings) = (append(substrings, "gzip"))
	}
	s = (strings.Join(substrings, ", "))
	return s
}

// ContentLanguageHeader is a struct to prepare a Content-Language HTTP header.
type ContentLanguageHeader struct {
	LanguageTags []string `json:"language_tags"`
}

// String returns a string representation of a Content-Language HTTP header value.
func (c ContentLanguageHeader) String() string {
	return (strings.Join(c.LanguageTags, ", "))
}

// ContentLengthHeader is a struct to prepare a Content-Length HTTP header.
type ContentLengthHeader struct {
	Length int64 `json:"length"`
}

// String returns a string representation of a Content-Length HTTP header value.
func (c ContentLengthHeader) String() string {
	return (fmt.Sprintf("%d", c.Length))
}

// ContentLocationHeader is a struct to prepare a Content-Location HTTP header.
type ContentLocationHeader struct {
	URL string `json:"url"`
}

// String returns a string representation of a Content-Location HTTP header value.
func (c ContentLocationHeader) String() string {
	return c.URL
}

// ContentRangeHeader is a struct to prepare a Content-Range HTTP header.
type ContentRangeHeader struct {
	RangeEnd   int64  `json:"range_end"`
	RangeStart int64  `json:"range_start"`
	Size       int64  `json:"size"`
	Units      string `json:"units"`
}

// String returns a string representation of a Content-Range HTTP header value.
func (c ContentRangeHeader) String() string {
	var units string = "*"
	if !reflect.ValueOf(c.Units).IsZero() {
		units = (c.Units)
	}
	return (fmt.Sprintf("%s %d-%d/%d", units, c.RangeStart, c.RangeEnd, c.Size))
}

// ContentSecurityPolicyHeader is a struct to prepare a Content-Security-Policy HTTP header.
type ContentSecurityPolicyHeader struct {
	ChildSrc      []string `json:"child_src"`
	ConnectSrc    []string `json:"connect_src"`
	DefaultSrc    []string `json:"default_src"`
	FontSrc       []string `json:"font_src"`
	FrameSrc      []string `json:"frame_src"`
	ImgSrc        []string `json:"img_src"`
	ManifestSrc   []string `json:"manifest_src"`
	MediaSrc      []string `json:"media_src"`
	ObjectSrc     []string `json:"object_src"`
	PrefetchSrc   []string `json:"prefetch_src"`
	ScriptSrc     []string `json:"script_src"`
	ScriptSrcElem []string `json:"script_src_elem"`
	ScriptSrcAttr []string `json:"script_src_attr"`
	StyleSrc      []string `json:"style_src"`
	StyleSrcElem  []string `json:"style_src_elem"`
	StyleSrcAttr  []string `json:"style_src_attr"`
	WorkerSrc     []string `json:"worker_src"`
}

// ContentTypeHeader is a struct to prepare a Content-Type HTTP header.
type ContentTypeHeader struct {
	Boundary    string `json:"boundary"`
	Charset     string `json:"charset"`
	MIMESubType string `json:"mime_subtype"`
	MIMEType    string `json:"mime_type"`
}

// String returns a string representation of a Content-Type HTTP header value.
func (c ContentTypeHeader) String() string {
	var boundaryOK, charsetOK bool = (len(c.Boundary) != 0), (len(c.Charset) != 0)
	var mimeSubTypeLengthOK, mimeTypeLengthOK bool = (len(c.MIMEType) != 0), (len(c.MIMESubType) != 0)
	var mimeSubType, mimeType string = "*", "*"
	var substrings ([]string) = (make([]string, 2))
	var s string
	if mimeTypeLengthOK {
		mimeType = c.MIMEType
	}
	substrings[0] = (mimeType)
	if mimeSubTypeLengthOK {
		mimeSubType = c.MIMESubType
	}
	substrings[1] = (mimeSubType)
	s = (strings.Join(substrings, "/"))
	(substrings) = ([]string{s})
	if boundaryOK {
		(substrings) = (append(substrings, (fmt.Sprintf("boundary=%s", c.Boundary))))
	}
	if charsetOK {
		(substrings) = (append(substrings, (fmt.Sprintf("charset=%s", c.Charset))))
	}
	(s) = (strings.Join(substrings, "; "))
	return s
}

// CookieHeader is a struct to prepare a Cookie HTTP header.
type CookieHeader struct {
	Cookies []*http.Cookie `json:"cookie"`
}

// String returns a string representation of a Cookie HTTP header value.
func (c CookieHeader) String() string {
	var cookie *http.Cookie
	var i int
	var substrings ([]string) = (make([]string, len(c.Cookies)))
	var s string
	for i, cookie = range c.Cookies {
		(substrings[i]) = (cookie.String())
	}
	(s) = (strings.Join(substrings, "; "))
	return s
}

// CrossOriginResourcePolicyHeader is a struct to prepare a Cross-Origin-Resource-Policy HTTP header.
type CrossOriginResourcePolicyHeader struct {
	CrossOrigin bool `json:"cross_origin"`
	SameOrigin  bool `json:"same_origin"`
	SameSite    bool `json:"same_site"`
}

// String returns a string representation of a Cross-Origin-Resource-Policy HTTP header.
func (c CrossOriginResourcePolicyHeader) String() string {
	if c.CrossOrigin {
		return "cross-origin"
	}
	if c.SameOrigin {
		return "same-origin"
	}
	if c.SameSite {
		return "same-site"
	}
	return "*"
}

// DNTHeader is a struct to prepare a DNT HTTP header.
type DNTHeader struct {
	DNT bool `json:"dnt"`
}

// String returns a string representation of a DNT HTTP header.
func (d DNTHeader) String() string {
	var s string = "0"
	if d.DNT {
		s = "1"
	}
	return s
}

// DPRHeader is a struct to prepare a DPR HTTP header.
type DPRHeader struct {
	DPR float32 `json:"dpr"`
}

// String returns a string representation of a DPR HTTP header.
func (d DPRHeader) String() string {
	return (fmt.Sprintf("%.1f", d.DPR))
}

// DateHeader is a struct to prepare a Date HTTP header.
type DateHeader struct {
	Time time.Time `json:"date"`
}

// String returns a string representation of a Date HTTP header.
func (d DateHeader) String() string {
	return (d.Time.Format(http.TimeFormat))
}

// DeviceMemoryHeader is a struct to prepare a Device-Memory HTTP header.
type DeviceMemoryHeader struct {
	Memory float32 `json:"memory"`
}

// String returns a string representation of a Device-Memory HTTP header.
func (d DeviceMemoryHeader) String() string {
	return (fmt.Sprintf("%.1f", d.Memory))
}

// DigestHeader is a struct to prepare a Digest HTTP header.
type DigestHeader struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

// String returns a string representation of a Digest HTTP header.
func (d DigestHeader) String() string {
	return (fmt.Sprintf("%s=%s", d.Algorithm, d.Value))
}

// ETagHeader is a struct to prepare a ETag HTTP header.
type ETagHeader struct {
	Value string `json:"value"`
	W     bool   `json:"w"`
}

// String returns a string representation of a ETag HTTP header.
func (e ETagHeader) String() string {
	var s string = "%s"
	if e.W {
		s = ("W/%s")
	}
	return (fmt.Sprintf(s, e.Value))
}

// EarlyDataHeader is a struct to prepare a Early-Header HTTP header.
type EarlyDataHeader struct {
	EarlyData bool `json:"early_data"`
}

// String returns a string value of a Early-Data HTTP header.
func (e EarlyDataHeader) String() string {
	var s string
	if e.EarlyData {
		s = "1"
	}
	return s
}
