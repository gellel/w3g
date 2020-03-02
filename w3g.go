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

// AccessControlAllowOrigin response header indicates whether the response can be shared with requesting code from the given origin.
const AccessControlAllowOrigin string = "Access-Control-Allow-Origin"

// AccessControlExposeHeaders response header indicates which headers can be exposed as part of the response by listing their names.
const AccessControlExposeHeaders string = "Access-Control-Expose-Headers"
