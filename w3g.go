package w3g

import (
	"fmt"
	"reflect"
	"strings"
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

// AcceptHTTPHeader is a struct to prepare a Accept HTTP request header.
type AcceptHTTPHeader struct {
	MIMESubType string  `json:"mime_subtype"`
	MIMEType    string  `json:"mime_type"`
	Q           float32 `json:"q"`
}

// Value returns a string representation of a Accept HTTP request header value.
func (a AcceptHTTPHeader) Value() string {
	var mimeSubTypeLengthOK, mimeTypeLengthOK, qOK bool = (len(a.MIMEType) != 0), (len(a.MIMESubType) != 0), (a.Q != 0)
	var mimeSubType, mimeType string = "*", "*"
	var substrings []string = (make([]string, 2))
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

func newHTTPHeaderValue(h ...interface{ Value() string }) string {
	var header interface{ Value() string }
	var i int
	var substrings []string = (make([]string, len(h)))
	var s string
	for i, header = range h {
		substrings[i] = (header.Value())
	}
	s = strings.Join(substrings, ",")
	return s
}

// NewAcceptHeader creates a new Accept HTTP request header as a formatted HTTP header value string.
func NewAcceptHeader(h ...AcceptHTTPHeader) string {
	var i int
	var s = make([]interface{ Value() string }, len(h))
	var v interface{ Value() string }
	for i, v = range h {
		(s[i]) = v
	}
	return (newHTTPHeaderValue(s...))
}

// AcceptCHHTTPHeader is a struct to prepare a Accept-CH HTTP response header.
type AcceptCHHTTPHeader struct {
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

// Value returns a string representation of a Accept-CH HTTP request header value.
func (a AcceptCHHTTPHeader) Value() string {
	var substrings []string = (make([]string, 0))
	var substring string
	var s string
	var r reflect.Value = reflect.ValueOf(a)
	for _, substring = range []string{AcceptCH, AcceptCHLifetime, ContentDPR, DPR, EarlyData, SaveData, ViewportWidth, Width} {
		var f reflect.Value = (r.FieldByName(substring))
		if f.IsValid() && f.Bool() {
			substrings = (append(substrings, substring))
		}
	}
	s = (strings.Join(substrings, ","))
	return s
}

// NewAcceptCHHeader creates a new Accept-CH HTTP response header as a formatted HTTP header value string.
func NewAcceptCHHeader(h ...AcceptCHHTTPHeader) string {
	var i int
	var s = make([]interface{ Value() string }, len(h))
	var v interface{ Value() string }
	for i, v = range h {
		(s[i]) = v
	}
	return (newHTTPHeaderValue(s...))
}
