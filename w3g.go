package w3g

import (
	"fmt"
	"net"
	"net/http"
	"net/mail"
	"net/url"
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

// CloudfrontForwardedProto HTTP request header is the AWS Cloudfront origin protocol that the request was forwarded from.
const CloudfrontForwardedProto string = "Cloudfront-Forwarded-Proto"

// CloudfrontIsDesktopViewer HTTP request header is the AWS Cloudfront header that identifies that the requesting client is a desktop device.
const CloudfrontIsDesktopViewer string = "Cloudfront-Is-Desktop-Viewer"

// CloudfrontIsMobileViewer HTTP request header is the AWS Cloudfront header that identifies that the requesting client is a mobile device.
const CloudfrontIsMobileViewer string = "Cloudfront-Is-Mobile-Viewer"

// CloudfrontIsSmartTvViewer HTTP request header is the AWS Cloudfront header that identifies that the requesting client is a smart tv.
const CloudfrontIsSmartTvViewer string = "Cloudfront-Is-Smarttv-Viewer"

// CloudfrontIsTabletViewer HTTP request header is the AWS Cloudfront header that identifies that the requesting client is a tablet device.
const CloudfrontIsTabletViewer string = "Cloudfront-Is-Tablet-Viewer"

// CloudfrontViewerCountry HTTP request headr is the AWS Cloudfront header that identifies the requesting clients country of origin.
const CloudfrontViewerCountry string = "Cloudfront-Viewer-Country"

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

// ContentMD5 HTTP header indicates the MD5 hash value of the HTTP request.
const ContentMD5 string = "Content-MD5"

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

// DeltaBase HTTP response header specifies the delta-encoding entity tag of the response.
const DeltaBase string = "Delta-Base"

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

// HTTP2Settings HTTP request header specifies that a request upgrades from HTTP/1.1 to HTTP/2.
const HTTP2Settings string = "HTTP2-Settings"

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

const IM string = "IM"

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

const MaxForwards string = "Max-Forwards"

// Origin request HTTP header indicates where a fetch originates from.
const Origin string = "Origin"

const P3P string = "P3P"

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

// Refresh response HTTP header controls when a new resource has been created and the time until redirection.
const Refresh string = "Refresh"

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

const Status string = "Status"

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

// XCorrelationID correlates HTTP requests between a client and server.
const XCorrelationID string = "X-Correlation-ID"

const XCSRFToken string = "X-CSRF-Token"

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

// XRealIP request HTTP header is the IP address of the origin client.
const XRealIP string = "X-Real-Ip"

// XRequestID correlates HTTP requests between a client and server.
const XRequestID string = "X-Request-ID"

// XRequestedWith  HTTP header is a request header used to indicate whether the HTTP request was an AJAX request.
const XRequestedWith string = "X-Requested-With"

const XUIDH string = "X-UIDH"

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

// ContentMD5Header is a struct to prepare a Content-MD5 HTTP header.
type ContentMD5Header struct {
	MD5 string `json:"md5"`
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

// String returns a string representation of a Early-Data HTTP header.
func (e EarlyDataHeader) String() string {
	var s string
	if e.EarlyData {
		s = "1"
	}
	return s
}

// ExpectHeader is a struct to prepare a Expect HTTP header.
type ExpectHeader struct{}

// String returns a string representation of a Expect HTTP header.
func (e ExpectHeader) String() string {
	return ("100-continue")
}

// ExpectCTHeader is a a struct to prepare a Expect-CT HTTP header.
type ExpectCTHeader struct {
	Enforce   bool   `json:"enforce"`
	MaxAge    int64  `json:"max_age"`
	ReportURI string `json:"report_uri"`
}

// String returns a string representation of a Expect-CT HTTP header.
func (e ExpectCTHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	if !reflect.ValueOf(e.MaxAge).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("%d", e.MaxAge)))
	}
	if e.Enforce {
		(substrings) = (append(substrings, "enforce"))
	}
	if !reflect.ValueOf(e.ReportURI).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("report-uri=\"%s\"", e.ReportURI)))
	}
	(s) = (strings.Join(substrings, ", "))
	return s
}

// ExpiresHeader is a struct to prepare a Expires HTTP header.
type ExpiresHeader struct {
	Expires time.Time
}

// String returns a string representation of a Expires HTTP header.
func (e ExpiresHeader) String() string {
	return (e.Expires.Format(http.TimeFormat))
}

// FeaturePolicyHeader is a struct to prepare a Feature-Policy HTTP header.
type FeaturePolicyHeader struct {
	Accelerometer               string `json:"accelerometer"`
	AmbientLightSensor          string `json:"ambient_light_sensor"`
	Autoplay                    string `json:"autoplay"`
	Battery                     string `json:"battery"`
	Camera                      string `json:"camera"`
	DisplayCapture              string `json:"display_capture"`
	DocumentDomain              string `json:"document_domain"`
	EncryptedMedia              string `json:"encrypted_media"`
	ExecutionWhileNotRendered   string `json:"execution_while_not_rendered"`
	ExecutionWhileOutOfViewport string `json:"execution_while_out_of_viewport"`
	Fullscreen                  string `json:"fullscreen"`
	Geolocation                 string `json:"geolocation"`
	Gyroscope                   string `json:"gyroscope"`
	LayoutAnimations            string `json:"layout_animations"`
	LegacyImageFormats          string `json:"legacy_image_formats"`
	Magnetometer                string `json:"magnetometer"`
	Microphone                  string `json:"microphone"`
	Midi                        string `json:"midi"`
	NavigationOverride          string `json:"navigation_override"`
	OversizedImages             string `json:"oversized_images"`
	Payment                     string `json:"payment"`
	PictureInPicture            string `json:"picture_in_picture"`
	PublicKeyCredentials        string `json:"publickey_credentials"`
	SyncXHR                     string `json:"sync_xhr"`
	USB                         string `json:"usb"`
	VR                          string `json:"vr"`
	WakeLock                    string `json:"wake_lock"`
	XRSpatialTracking           string `json:"xr_spatial_tracking"`
}

// String returns a string representation of a Feature-Policy HTTP header.
func (f FeaturePolicyHeader) String() string {
	var regex *regexp.Regexp = regexp.MustCompile(`[A-Z][^A-Z]*`)
	var substrings ([]string) = (make([]string, 0))
	var s string
	var r reflect.Value = reflect.ValueOf(f)
	var v reflect.Type = (reflect.Indirect(r).Type())
	for i, n := 0, r.NumField(); i < n; i++ {
		var f reflect.Value = r.Field(i)
		if f.IsValid() && !f.IsZero() {
			var name string = strings.Join(regex.FindAllString(v.Field(i).Name, -1), "-")
			name = strings.ToLower(name)
			switch f.Kind() {
			case reflect.String:
				(substrings) = (append(substrings, fmt.Sprintf("%s '%s'", name, f.String())))
			}
		}
	}
	(s) = (strings.Join(substrings, "; "))
	return s
}

// ForwardedHeader is a struct to prepare a Forwarded HTTP header.
type ForwardedHeader struct {
	By         string `json:"by"`
	Host       string `json:"host"`
	Identifier net.IP `json:"identifier"`
	Proto      string `json:"proto"`
}

func (f ForwardedHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	if !reflect.ValueOf(f.By).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("by=%s", f.By)))
	}
	if f.Identifier != nil {
		if f.Identifier.To4() != nil {
			(substrings) = (append(substrings, fmt.Sprintf("for=%s", f.Identifier.String())))
		} else if f.Identifier.To16() != nil {
			(substrings) = (append(substrings, fmt.Sprintf("for=\"[%s]\"", f.Identifier.String())))
		}
	}
	if !reflect.ValueOf(f.Host).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("host=%s", f.Host)))
	}
	(s) = (strings.Join(substrings, "; "))
	return s
}

// FromHeader is a struct to prepare a From HTTP header.
type FromHeader struct {
	Email mail.Address `json:"email"`
}

// String returns a string representation of a From HTTP header.
func (f FromHeader) String() string {
	return (f.Email.String())
}

// HostHeader is a struct to prepare a Host HTTP header.
type HostHeader struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// String returns a string representation of a Host HTTP header.
func (h HostHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	if !reflect.ValueOf(h.Host).IsZero() {
		(substrings) = (append(substrings, h.Host))
	}
	if !reflect.ValueOf(h.Port).IsZero() {
		(substrings) = (append(substrings, h.Port))
	}
	(s) = (strings.Join(substrings, ":"))
	return s
}

// IfMatchHeader is a struct to prepare a If-Match HTTP header.
type IfMatchHeader struct {
	Value string `json:"value"`
}

// String returns a string representation of a If-Match HTTP header.
func (i IfMatchHeader) String() string {
	if reflect.ValueOf(i.Value).IsZero() {
		return "*"
	}
	return i.Value
}

// IfModifiedSinceHeader is a struct to prepare a If-Modified-Since HTTP header.
type IfModifiedSinceHeader struct {
	Time time.Time `json:"time"`
}

// String returns a string representation of a If-Modified-Since HTTP header.
func (i IfModifiedSinceHeader) String() string {
	return (i.Time.Format(http.TimeFormat))
}

// IfNoneMatchHeader is a struct to prepare a If-None-Match HTTP header.
type IfNoneMatchHeader struct {
	Value string `json:"value"`
}

// String returns a string representation of a If-None-Match HTTP header.
func (i IfNoneMatchHeader) String() string {
	if reflect.ValueOf(i.Value).IsZero() {
		return "*"
	}
	return i.Value
}

// IfRangeHeader is a struct to prepare a If-Range HTTP header.
type IfRangeHeader struct {
	ETag string    `json:"etag"`
	Time time.Time `json:"time"`
}

// String returns a string representation of a If-Range HTTP header.
func (i IfRangeHeader) String() string {
	var s string
	if !reflect.ValueOf(i.ETag).IsZero() {
		return i.ETag
	}
	if !reflect.ValueOf(i.Time).IsZero() {
		return i.Time.Format(http.TimeFormat)
	}
	return s
}

// IfUnmodifiedSinceHeader is a struct to prepare a If-Unmodified-Since HTTP header.
type IfUnmodifiedSinceHeader struct {
	Time time.Time `json:"time"`
}

// String return a string representation of a If-Unmodified-Since HTTP header.
func (i IfUnmodifiedSinceHeader) String() string {
	return (i.Time.Format(http.TimeFormat))
}

// KeepAliveHeader is a struct to prepare a Keep-Alive HTTP header.
type KeepAliveHeader struct {
	Max     int `json:"max"`
	Timeout int `json:"timeout"`
}

// String returns a string representation of a Keep-Alive HTTP header.
func (k KeepAliveHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	if !reflect.ValueOf(k.Max).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("max=%d", k.Max)))
	}
	if !reflect.ValueOf(k.Timeout).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("timeout=%d", k.Timeout)))
	}
	(s) = (strings.Join(substrings, ", "))
	return s
}

// LargeAllocationHeader is a struct to prepare a Large-Allocation HTTP header.
type LargeAllocationHeader struct {
	Megabytes int64 `json:"metabytes"`
}

// String returns a string representation of a Large-Allocation HTTP header.
func (l LargeAllocationHeader) String() string {
	return (fmt.Sprintf("%d", l.Megabytes))
}

// LastModifiedHeader is a struct to prepare a Last-Modified HTTP header.
type LastModifiedHeader struct {
	Time time.Time `json:"time"`
}

// String returns a string representation of a Last-Modified HTTP header.
func (l LastModifiedHeader) String() string {
	return (l.Time.Format(http.TimeFormat))
}

// LinkHeader is a struct to prepare a Link HTTP header.
type LinkHeader struct {
	URL url.URL `json:"url"`
}

// String returns a string representation of a Link HTTP header.
func (l LinkHeader) String() string {
	return (l.URL.String())
}

// LocationHeader is a struct to prepare a Location HTTP header.
type LocationHeader struct {
	URL url.URL `json:"url"`
}

// String returns a string representation of a Location HTTP header.
func (l LocationHeader) String() string {
	return (l.URL.Path)
}

// OriginHeader is a struct to prepare a Origin HTTP header.
type OriginHeader struct {
	URL url.URL `json:"url"`
}

// String returns a string representation of a Origin HTTP header.
func (o OriginHeader) String() string {
	return (o.URL.String())
}

// PragmaHeader is a struct to prepare a Pragma HTTP header.
type PragmaHeader struct{}

// String returns a string representation of a Pragma HTTP header.
func (p PragmaHeader) String() string {
	return "no-cache"
}

// ProxyAuthenticateHeader is a struct to prepare a Proxy-Authenticate HTTP header.
type ProxyAuthenticateHeader struct {
	Realm string `json:"realm"`
	Type  string `json:"string"`
}

// String returns a string representation of a Proxy-Authenticate HTTP header.
func (p ProxyAuthenticateHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	if !reflect.ValueOf(p.Type).IsZero() {
		(substrings) = (append(substrings, p.Type))
	}
	if !reflect.ValueOf(p.Realm).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("realm=%s", p.Realm)))
	}
	(s) = (strings.Join(substrings, " "))
	return s
}

// ProxyAuthorizationHeader is a struct to prepare a Proxy-Authorization HTTP header.
type ProxyAuthorizationHeader struct {
	Credentials string `json:"credentials"`
	Type        string `json:"type"`
}

// String returns a string representation of a Proxy-Authorization HTTP header.
func (p ProxyAuthorizationHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	if !reflect.ValueOf(p.Type).IsZero() {
		(substrings) = (append(substrings, p.Type))
	}
	if !reflect.ValueOf(p.Credentials).IsZero() {
		(substrings) = (append(substrings, p.Credentials))
	}
	(s) = (strings.Join(substrings, " "))
	return s
}

// PublicKeyPinsHeader is a struct to prepare a Public-Key-Pins HTTP header.
type PublicKeyPinsHeader struct {
	IncludeSubDomains bool    `json:"include_subdomains"`
	MaxAge            int64   `json:"max_age"`
	PinSHA256         string  `json:"pin_sha256"`
	ReportURI         url.URL `json:"report_uri"`
}

func (p PublicKeyPinsHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	if p.IncludeSubDomains {
		(substrings) = (append(substrings, "includeSubDomains"))
	}
	if !reflect.ValueOf(p.MaxAge).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("max-age=%d", p.MaxAge)))
	}
	if !reflect.ValueOf(p.PinSHA256).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("pin-sha256=\"%s\"", p.PinSHA256)))
	}
	if !reflect.ValueOf(p.ReportURI).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("report-uri=\"%s\"", p.ReportURI.String())))
	}
	(s) = (strings.Join(substrings, "; "))
	return s
}

// PublicKeyPinsReporyOnlyHeader is a struct to prepare a Public-Key-Pins-Report-Only HTTP header.
type PublicKeyPinsReporyOnlyHeader struct {
	IncludeSubDomains bool    `json:"include_subdomains"`
	MaxAge            int64   `json:"max_age"`
	PinSHA256         string  `json:"pin_sha256"`
	ReportURI         url.URL `json:"report_uri"`
}

func (p PublicKeyPinsReporyOnlyHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	if p.IncludeSubDomains {
		(substrings) = (append(substrings, "includeSubDomains"))
	}
	if !reflect.ValueOf(p.MaxAge).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("max-age=%d", p.MaxAge)))
	}
	if !reflect.ValueOf(p.PinSHA256).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("pin-sha256=\"%s\"", p.PinSHA256)))
	}
	if !reflect.ValueOf(p.ReportURI).IsZero() {
		(substrings) = (append(substrings, fmt.Sprintf("report-uri=\"%s\"", p.ReportURI.String())))
	}
	(s) = (strings.Join(substrings, "; "))
	return s
}

// RangeHeader is a struct to prepare a Range HTTP header.
type RangeHeader struct {
	RangeEnd     int64  `json:"range_end"`
	RangeStart   int64  `json:"range_start"`
	SuffixLength int64  `json:"suffix_length"`
	Unit         string `json:"unit"`
}

func (r RangeHeader) String() string {
	var rangeStartOK, rangeEndOK, suffixLengthOK, unitOK = !reflect.ValueOf(r.RangeEnd).IsZero(), !reflect.ValueOf(r.RangeStart).IsZero(), !reflect.ValueOf(r.SuffixLength).IsZero(), !reflect.ValueOf(r.Unit).IsZero()
	var substrings ([]string) = (make([]string, 0))
	var s string
	if unitOK {
		(substrings) = (append(substrings, r.Unit))
	}
	if rangeEndOK || rangeStartOK {
		var ss ([]string) = (make([]string, 0))
		if rangeStartOK {
			substrings = append(ss, fmt.Sprintf("%d", r.RangeStart))
		}
		if rangeEndOK {
			substrings = append(ss, fmt.Sprintf("%d", r.RangeEnd))
		}
		(substrings) = (append(substrings, strings.Join(ss, "-")))
	} else if suffixLengthOK {
		(substrings) = (append(substrings, fmt.Sprintf("-%d", r.SuffixLength)))
	}
	(s) = (strings.Join(substrings, ""))
	return s
}

// RefererHeader is a struct to prepare a Referer HTTP header.
type RefererHeader struct {
	URL url.URL
}

// String returns a string representation of a Referer HTTP header.
func (r RefererHeader) String() string {
	return (r.URL.String())
}

// ReferrerPolicyHeader is a struct to prepare a Referrer Policy HTTP header.
type ReferrerPolicyHeader struct {
	NoReferrer                  bool `json:"no_referrer"`
	NoReferrerWhenDowngrade     bool `json:"no_referrer_when_downgrade"`
	Origin                      bool `json:"origin"`
	OriginWhenCrossOrigin       bool `json:"origin_when_cross_origin"`
	SameOrigin                  bool `json:"same_origin"`
	StrictOrigin                bool `json:"strict_origin"`
	StrictOriginWhenCrossOrigin bool `json:"strict_origin_when_cross_origin"`
}

// String returns a string representation of a Referrer-Policy HTTP header.
func (r ReferrerPolicyHeader) String() string {
	var regex *regexp.Regexp = regexp.MustCompile(`[A-Z][^A-Z]*`)
	var substrings ([]string) = (make([]string, 0))
	var s string
	var x reflect.Value = reflect.ValueOf(r)
	var v reflect.Type = (reflect.Indirect(x).Type())
	for i, n := 0, x.NumField(); i < n; i++ {
		var f reflect.Value = x.Field(i)
		if f.IsValid() && !f.IsZero() {
			var name string = strings.Join(regex.FindAllString(v.Field(i).Name, -1), "-")
			name = strings.ToLower(name)
			switch f.Kind() {
			case reflect.Bool:
				(substrings) = (append(substrings, fmt.Sprintf("%s", name)))
			}
		}
	}
	(s) = (strings.Join(substrings, ", "))
	return s
}

// RetryAfterHeader is a struct to prepare a Retry-After HTTP header.
type RetryAfterHeader struct {
	Seconds int64     `json:"seconds"`
	Time    time.Time `json:"time"`
}

// String returns a string representation of a Retry-After HTTP header.
func (r RetryAfterHeader) String() string {
	if !reflect.ValueOf(r.Time).IsZero() {
		return (r.Time.Format(http.TimeFormat))
	}
	return (fmt.Sprintf("%d", r.Seconds))
}

// SaveDataHeader is a struct to prepare a Save-Data HTTP header.
type SaveDataHeader struct {
	On bool `json:"on"`
}

// String is a struct to prepare a Save-Data HTTP header.
func (s SaveDataHeader) String() string {
	if s.On {
		return "on"
	}
	return "off"
}

// SecFetchDestHeader is a struct to prepare a Sec-Fetch-Dest HTTP header.
type SecFetchDestHeader struct {
	Audio          bool `json:"audio"`
	Audioworklet   bool `json:"audioworklet"`
	Document       bool `json:"document"`
	Embed          bool `json:"embed"`
	Empty          bool `json:"empty"`
	Font           bool `json:"font"`
	Image          bool `json:"image"`
	Manifest       bool `json:"manifest"`
	NestedDocument bool `json:"nested_document"`
	Object         bool `json:"object"`
	Paintworklet   bool `json:"paintworklet"`
	Report         bool `json:"report"`
	Script         bool `json:"script"`
	Serviceworker  bool `json:"serviceworker"`
	Sharedworker   bool `json:"sharedworker"`
	Style          bool `json:"style"`
	Track          bool `json:"track"`
	Video          bool `json:"video"`
	Worker         bool `json:"worker"`
	Xslt           bool `json:"xslt"`
}

// String returns a string representation of a Sec-Fetch-Dest HTTP header.
func (s SecFetchDestHeader) String() string {
	var regex *regexp.Regexp = regexp.MustCompile(`[A-Z][^A-Z]*`)
	var substrings ([]string) = (make([]string, 0))
	var ss string
	var x reflect.Value = reflect.ValueOf(s)
	var v reflect.Type = (reflect.Indirect(x).Type())
	for i, n := 0, x.NumField(); i < n; i++ {
		var f reflect.Value = x.Field(i)
		if f.IsValid() && !f.IsZero() {
			var name string = strings.Join(regex.FindAllString(v.Field(i).Name, -1), "-")
			name = strings.ToLower(name)
			switch f.Kind() {
			case reflect.Bool:
				(substrings) = (append(substrings, fmt.Sprintf("%s", name)))
			}
		}
	}
	(ss) = (strings.Join(substrings, ", "))
	return ss
}

// SecFetchModeHeader is a struct to preparea Sec-Fetch-Mode HTTP header.
type SecFetchModeHeader struct {
	Cors           bool `json:"cors"`
	Navigate       bool `json:"navigate"`
	NestedNavigate bool `json:"nested_navigate"`
	NoCors         bool `json:"no_cors"`
	SameOrigin     bool `json:"same_origin"`
	WebSocket      bool `json:"websocket"`
}

// String returns a string representation of a Sec-Fetch-Mode HTTP header.
func (s SecFetchModeHeader) String() string {
	var regex *regexp.Regexp = regexp.MustCompile(`[A-Z][^A-Z]*`)
	var substrings ([]string) = (make([]string, 0))
	var ss string
	var x reflect.Value = reflect.ValueOf(s)
	var v reflect.Type = (reflect.Indirect(x).Type())
	for i, n := 0, x.NumField(); i < n; i++ {
		var f reflect.Value = x.Field(i)
		if f.IsValid() && !f.IsZero() {
			var name string = strings.Join(regex.FindAllString(v.Field(i).Name, -1), "-")
			name = strings.ToLower(name)
			switch f.Kind() {
			case reflect.Bool:
				(substrings) = (append(substrings, fmt.Sprintf("%s", name)))
			}
		}
	}
	(ss) = (strings.Join(substrings, ", "))
	return ss
}

// SecFetchSiteHeader is a struct to prepare a Sec-Fetch-Site HTTP header.
type SecFetchSiteHeader struct {
	CrossSite  bool `json:"cross_site"`
	None       bool `json:"none"`
	SameOrigin bool `json:"same_origin"`
	SameSite   bool `json:"same_site"`
}

// String returns a string representation of a Sec-Fetch-Site HTTP header.
func (s SecFetchSiteHeader) String() string {
	var regex *regexp.Regexp = regexp.MustCompile(`[A-Z][^A-Z]*`)
	var substrings ([]string) = (make([]string, 0))
	var ss string
	var x reflect.Value = reflect.ValueOf(s)
	var v reflect.Type = (reflect.Indirect(x).Type())
	for i, n := 0, x.NumField(); i < n; i++ {
		var f reflect.Value = x.Field(i)
		if f.IsValid() && !f.IsZero() {
			var name string = strings.Join(regex.FindAllString(v.Field(i).Name, -1), "-")
			name = strings.ToLower(name)
			switch f.Kind() {
			case reflect.Bool:
				(substrings) = (append(substrings, fmt.Sprintf("%s", name)))
			}
		}
	}
	(ss) = (strings.Join(substrings, ", "))
	return ss
}

// SecFetchUserHeader is a struct to prepare a Sec-Fetch-User HTTP header.
type SecFetchUserHeader struct {
	Activated bool `json:"activated"`
}

// String returns a string representation of a Sec-Fetch-User HTTP header.
func (s SecFetchUserHeader) String() string {
	if s.Activated {
		return "?1"
	}
	return "?0"
}

// SecWebSocketAcceptHeader is a struct to prepare Sec-Web-Socket-Accept HTTP header.
type SecWebSocketAcceptHeader struct {
	HashedKey string `json:"hashed_key"`
}

// String returns a string representation of a Sec-Web-Socket-Accept HTTP header.
func (s SecWebSocketAcceptHeader) String() string {
	return (s.HashedKey)
}

// ServerHeader is a struct to prepare a Server HTTP header.
type ServerHeader struct {
	Server string `json:"server"`
}

// String returns a string representation of a Server HTTP header.
func (s ServerHeader) String() string {
	return (s.Server)
}

// ServerTimingHeader is a struct to prepare a Server-Timing HTTP header.
type ServerTimingHeader struct {
	Cache       bool `json:"cache"`
	CPU         bool `json:"cpu"`
	MissedCache bool `json:"missed_cache"`
}

// SetCookieHeader is a struct to prepare a Set-Cookie HTTP header.
type SetCookieHeader struct {
	Cookie http.Cookie `json:"cookie"`
}

// String returns a strings representation of a Set-Cookie HTTP header.
func (s SetCookieHeader) String() string {
	return s.Cookie.String()
}

// SetCookie2Header is a struct to prepare a Set-Cookie-2 HTTP header.
type SetCookie2Header struct {
	Cookie http.Cookie `json:"cookie"`
}

// String reutrns a string representation of a Set-Cookie-2 HTTP header.
func (s SetCookie2Header) String() string {
	return s.Cookie.String()
}

// SourceMapHeader is a struct to prepare a Source-Map HTTP header.
type SourceMapHeader struct {
	URL url.URL `json:"url"`
}

// String returns a string representation of a Source-Map HTTP header.
func (s SourceMapHeader) String() string {
	return (s.URL.String())
}

// StrictTransportSecurityHeader is a struct to prepare a Strict-Transport-Security HTTP header.
type StrictTransportSecurityHeader struct {
	IncludeSubDomains bool  `json:"include_subdomains"`
	MaxAge            int64 `json:"max_age"`
	Preload           bool  `json:"preload"`
}

// String returns a string representation of a Strict-Transport-Security HTTP header.
func (s StrictTransportSecurityHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	(substrings) = (append(substrings, fmt.Sprintf("%d", s.MaxAge)))
	if !reflect.ValueOf(s.IncludeSubDomains).IsZero() {
		(substrings) = (append(substrings, "includeSubDomains"))
	}
	if !reflect.ValueOf(s.Preload).IsZero() {
		(substrings) = (append(substrings, "preload"))
	}
	return strings.Join(substrings, "; ")
}

// TEHeader is a struct to prepare a TE HTTP header.
type TEHeader struct {
	Compress bool    `json:"compress"`
	Deflate  bool    `json:"deflate"`
	GZip     bool    `json:"gzip"`
	Q        float32 `json:"q"`
	Trailers bool    `json:"trailers"`
}

// String returns a string representation of TE HTTP header.
func (t TEHeader) String() string {
	var qOK bool = !reflect.ValueOf(t.Q).IsZero()
	var s string
	if t.Compress {
		s = "compress"
	} else if t.Deflate {
		s = "deflate"
	} else if t.GZip {
		s = "gzip"
	} else if t.Trailers {
		s = "trailers"
	}
	if qOK {
		s = fmt.Sprintf("%s;%.2f", s, t.Q)
	}
	return s
}

// TimingAllowOriginHeader is a struct to prepare a Timing-Allow-Origin HTTP header.
type TimingAllowOriginHeader struct {
	Origins []url.URL `json:"origins"`
}

// String returns a string representation of a Timing-Allow-Origin HTTP header.
func (t TimingAllowOriginHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	var s string
	for _, u := range t.Origins {
		substrings = append(substrings, u.String())
	}
	s = strings.Join(substrings, ", ")
	return s
}

// TkHeader is a struct to prepare a Tk HTTP header.
type TkHeader struct {
	DisregardingDoNotTrack bool `json:"disregarding_do_not_track"`
	Dynamic                bool `json:"dynamic"`
	Gateway                bool `json:"gateway"`
	NotTracking            bool `json:"not_tracking"`
	PotentialConsent       bool `json:"potential_consent"`
	Tracking               bool `json:"tracking"`
	TrackingWithConsent    bool `json:"tracking_with_consent"`
	UnderConstruction      bool `json:"under_construction"`
	Updated                bool `json:"updated"`
}

// String returns a string representation of a Tk HTTP header.
func (t TkHeader) String() string {
	if t.DisregardingDoNotTrack {
		return "D"
	} else if t.Dynamic {
		return "?"
	} else if t.Gateway {
		return "G"
	} else if t.NotTracking {
		return "N"
	} else if t.PotentialConsent {
		return "P"
	} else if t.Tracking {
		return "T"
	} else if t.TrackingWithConsent {
		return "C"
	} else if t.Updated {
		return "U"
	}
	return "!"
}

// TrailerHeader is a struct to prepare a Trailer HTTP header.
type TrailerHeader struct {
	Headers []string `json:"headers"`
}

// String returns a string representation of a Trailer HTTP header.
func (t TrailerHeader) String() string {
	return strings.Join(t.Headers, ",")
}

// TransferHeader is a struct to prepare a Transfer HTTP header.
type TransferHeader struct{}

// TransferEncodingHeader is a struct to prepare a Transfer-Encoding HTTP header.
type TransferEncodingHeader struct {
	Chunked  bool `json:"chunked"`
	Compress bool `json:"compress"`
	Deflate  bool `json:"delate"`
	GZip     bool `json:"gzip"`
	Identity bool `json:"identity"`
}

// String returns a string representation of a Transfer-Encoding HTTP header.
func (t TransferEncodingHeader) String() string {
	var substrings ([]string) = (make([]string, 0))
	if t.Chunked {
		substrings = append(substrings, "chunked")
	}
	if t.Compress {
		substrings = append(substrings, "compress")
	}
	if t.Deflate {
		substrings = append(substrings, "deflate")
	}
	if t.GZip {
		substrings = append(substrings, "gzip")
	}
	if t.Identity {
		substrings = append(substrings, "identity")
	}
	return strings.Join(substrings, ", ")
}

// UpgradeInsecureRequestsHeader is a struct to prepare a Upgrade-Insecure-Request HTTP header.
type UpgradeInsecureRequestsHeader struct {
	Upgrade bool `json:"upgrade"`
}

// String returns a string representation of a Upgrade-Insecure-Request HTTP header.
func (u UpgradeInsecureRequestsHeader) String() string {
	if u.Upgrade {
		return "1"
	}
	return "0"
}

// UserAgentHeader is a struct to prepare a User-Agent HTTP header.
type UserAgentHeader struct {
	UserAgent string `json:"user_agent"`
}

// String returns a string representation of a User-Agent HTTP header.
func (u UserAgentHeader) String() string {
	return u.UserAgent
}

// VaryHeader is a struct to prepare a Vary HTTP header.
type VaryHeader struct {
	Headers []string `json:"headers"`
}

// String returns a string representation of a Vary HTTP header.
func (v VaryHeader) String() string {
	return strings.Join(v.Headers, ", ")
}

// ViaHeader is a struct to prepare a Via HTTP header.
type ViaHeader struct {
	ProtocolName    string `json:"protocol_name"`
	ProtocolVersion string `json:"protocol_version"`
	Host            string `json:"host"`
	Port            string `json:"port"`
	Pseudonym       string `json:"pseudonym"`
}

// String returns a string representation of a Via HTTP header.
func (v ViaHeader) String() string {
	var substrings = (make([]string, 0))
	if !reflect.ValueOf(v.ProtocolName).IsZero() {
		substrings = append(substrings, v.ProtocolName)
	}
	if !reflect.ValueOf(v.ProtocolVersion).IsZero() {
		substrings = append(substrings, v.ProtocolVersion)
	}
	if !reflect.ValueOf(v.Host).IsZero() || !reflect.ValueOf(v.Port).IsZero() {
		var s = (make([]string, 0))
		if !reflect.ValueOf(v.Host).IsZero() {
			s = append(s, v.Host)
		}
		if !reflect.ValueOf(v.Port).IsZero() {
			s = append(s, v.Port)
		}
		substrings = append(substrings, strings.Join(s, ":"))
	}
	if !reflect.ValueOf(v.Pseudonym).IsZero() {
		substrings = append(substrings, v.Pseudonym)
	}
	return strings.Join(substrings, " ")
}

type WWWAuthenticateHeader struct{
	Charset string
	Realm string 
}

// XRealIPHeader is a struct to prepare a X-Real-Ip HTTP header.
type XRealIPHeader struct {
	IP net.IP `json:"ip"`
}

// String returns a string representation of a X-Real-Ip HTTP header.
func (x XRealIPHeader) String() string {
	return x.IP.String()
}
