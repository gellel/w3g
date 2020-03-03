package w3g

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

// Link entity-header field provides a means for serialising one or more links in HTTP headers.
const Link string = "Link"

// SetCookie HTTP response header is used to send cookies from the server to the user agent, so the user agent can send them back to the server later.
const SetCookie string = "Set-Cookie"
