//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

const (
	//  No error.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NONE = 0
	//  An asynchronous IO operation is not yet complete. This usually does not
	//  indicate a fatal error. Typically this error will be generated as a
	//  notification to wait for some external notification that the IO operation
	//  finally completed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_IO_PENDING = -1
	//  A generic failure occurred.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_FAILED = -2
	//  An operation was aborted (due to user action).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ABORTED = -3
	//  An argument to the function is incorrect.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_ARGUMENT = -4
	//  The handle or file descriptor is invalid.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_HANDLE = -5
	//  The file or directory cannot be found.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_FILE_NOT_FOUND = -6
	//  An operation timed out.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_TIMED_OUT = -7
	//  The file is too large.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_FILE_TOO_BIG = -8
	//  An unexpected error. This may be caused by a programming mistake or an
	//  invalid assumption.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UNEXPECTED = -9
	//  Permission to access a resource, other than the network, was denied.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ACCESS_DENIED = -10
	//  The operation failed because of unimplemented functionality.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NOT_IMPLEMENTED = -11
	//  There were not enough resources to complete the operation.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INSUFFICIENT_RESOURCES = -12
	//  Memory allocation failed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_OUT_OF_MEMORY = -13
	//  The file upload failed because the file's modification time was different
	//  from the expectation.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UPLOAD_FILE_CHANGED = -14
	//  The socket is not connected.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SOCKET_NOT_CONNECTED = -15
	//  The file already exists.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_FILE_EXISTS = -16
	//  The path or file name is too long.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_FILE_PATH_TOO_LONG = -17
	//  Not enough room left on the disk.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_FILE_NO_SPACE = -18
	//  The file has a virus.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_FILE_VIRUS_INFECTED = -19
	//  The client chose to block the request.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOCKED_BY_CLIENT = -20
	//  The network changed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NETWORK_CHANGED = -21
	//  The request was blocked by the URL block list configured by the domain
	//  administrator.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOCKED_BY_ADMINISTRATOR = -22
	//  The socket is already connected.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SOCKET_IS_CONNECTED = -23
	//  The upload failed because the upload stream needed to be re-read, due to a
	//  retry or a redirect, but the upload stream doesn't support that operation.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UPLOAD_STREAM_REWIND_NOT_SUPPORTED = -25
	//  The request failed because the URLRequestContext is shutting down, or has
	//  been shut down.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CONTEXT_SHUT_DOWN = -26
	//  The request failed because the response was delivered along with requirements
	//  which are not met ('X-Frame-Options' and 'Content-Security-Policy' ancestor
	//  checks and 'Cross-Origin-Resource-Policy' for instance).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOCKED_BY_RESPONSE = -27
	//  The request was blocked by system policy disallowing some or all cleartext
	//  requests. Used for NetworkSecurityPolicy on Android.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CLEARTEXT_NOT_PERMITTED = -29
	//  The request was blocked by a Content Security Policy
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOCKED_BY_CSP = -30
	//  The request was blocked by CORB or ORB.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOCKED_BY_ORB = -32
	//  The request was blocked because it originated from a frame that has disabled
	//  network access.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NETWORK_ACCESS_REVOKED = -33
	//  The request was blocked by fingerprinting protections.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOCKED_BY_FINGERPRINTING_PROTECTION = -34
	//  The request was blocked by the Incognito Mode URL block list configured by
	//  the domain administrator.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOCKED_IN_INCOGNITO_BY_ADMINISTRATOR = -35
	//  The request was blocked because the local network permission is missing.
	//  Note that this is different from BLOCKED_BY_LOCAL_NETWORK_ACCESS_CHECKS
	//  which is specifically for a CORS error code.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_LOCAL_NETWORK_PERMISSION_MISSING = -36
	//  A connection was closed (corresponding to a TCP FIN).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CONNECTION_CLOSED = -100
	//  A connection was reset (corresponding to a TCP RST).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CONNECTION_RESET = -101
	//  A connection attempt was refused.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CONNECTION_REFUSED = -102
	//  A connection timed out as a result of not receiving an ACK for data sent.
	//  This can include a FIN packet that did not get ACK'd.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CONNECTION_ABORTED = -103
	//  A connection attempt failed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CONNECTION_FAILED = -104
	//  The host name could not be resolved.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NAME_NOT_RESOLVED = -105
	//  The Internet connection has been lost.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INTERNET_DISCONNECTED = -106
	//  An SSL protocol error occurred.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_PROTOCOL_ERROR = -107
	//  The IP address or port number is invalid (e.g., cannot connect to the IP
	//  address 0 or the port 0).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ADDRESS_INVALID = -108
	//  The IP address is unreachable. This usually means that there is no route to
	//  the specified host or network.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ADDRESS_UNREACHABLE = -109
	//  The server requested a client certificate for SSL client authentication.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_CLIENT_AUTH_CERT_NEEDED = -110
	//  A tunnel connection through the proxy could not be established. For more info
	//  see the comment on PROXY_UNABLE_TO_CONNECT_TO_DESTINATION.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_TUNNEL_CONNECTION_FAILED = -111
	//  The client and server don't support a common SSL protocol version or
	//  cipher suite.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_VERSION_OR_CIPHER_MISMATCH = -113
	//  The server requested a renegotiation (rehandshake).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_RENEGOTIATION_REQUESTED = -114
	//  The proxy requested authentication (for tunnel establishment) with an
	//  unsupported method.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PROXY_AUTH_UNSUPPORTED = -115
	//  The SSL handshake failed because of a bad or missing client certificate.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BAD_SSL_CLIENT_AUTH_CERT = -117
	//  A connection attempt timed out.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CONNECTION_TIMED_OUT = -118
	//  There are too many pending DNS resolves, so a request in the queue was
	//  aborted.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_HOST_RESOLVER_QUEUE_TOO_LARGE = -119
	//  Failed establishing a connection to the SOCKS proxy server for a target host.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SOCKS_CONNECTION_FAILED = -120
	//  The SOCKS proxy server failed establishing connection to the target host
	//  because that host is unreachable.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SOCKS_CONNECTION_HOST_UNREACHABLE = -121
	//  The request to negotiate an alternate protocol failed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ALPN_NEGOTIATION_FAILED = -122
	//  The peer sent an SSL no_renegotiation alert message.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_NO_RENEGOTIATION = -123
	//  Winsock sometimes reports more data written than passed. This is probably
	//  due to a broken LSP.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_WINSOCK_UNEXPECTED_WRITTEN_BYTES = -124
	//  An SSL peer sent us a fatal decompression_failure alert. This typically
	//  occurs when a peer selects DEFLATE compression in the mistaken belief that
	//  it supports it.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_DECOMPRESSION_FAILURE_ALERT = -125
	//  An SSL peer sent us a fatal bad_record_mac alert. This has been observed
	//  from servers with buggy DEFLATE support.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_BAD_RECORD_MAC_ALERT = -126
	//  The proxy requested authentication (for tunnel establishment).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PROXY_AUTH_REQUESTED = -127
	//  Could not create a connection to the proxy server. An error occurred
	//  either in resolving its name, or in connecting a socket to it.
	//  Note that this does NOT include failures during the actual "CONNECT" method
	//  of an HTTP proxy.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PROXY_CONNECTION_FAILED = -130
	//  A mandatory proxy configuration could not be used. Currently this means
	//  that a mandatory PAC script could not be fetched, parsed or executed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_MANDATORY_PROXY_CONFIGURATION_FAILED = -131
	//  We've hit the max socket limit for the socket pool while preconnecting. We
	//  don't bother trying to preconnect more sockets.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PRECONNECT_MAX_SOCKET_LIMIT = -133
	//  The permission to use the SSL client certificate's private key was denied.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_CLIENT_AUTH_PRIVATE_KEY_ACCESS_DENIED = -134
	//  The SSL client certificate has no private key.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_CLIENT_AUTH_CERT_NO_PRIVATE_KEY = -135
	//  The certificate presented by the HTTPS Proxy was invalid.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PROXY_CERTIFICATE_INVALID = -136
	//  An error occurred when trying to do a name resolution (DNS).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NAME_RESOLUTION_FAILED = -137
	//  Permission to access the network was denied. This is used to distinguish
	//  errors that were most likely caused by a firewall from other access denied
	//  errors. See also ERR_ACCESS_DENIED.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NETWORK_ACCESS_DENIED = -138
	//  The request throttler module cancelled this request to avoid DDOS.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_TEMPORARILY_THROTTLED = -139
	//  We were unable to sign the CertificateVerify data of an SSL client auth
	//  handshake with the client certificate's private key.
	//
	//  Possible causes for this include the user implicitly or explicitly
	//  denying access to the private key, the private key may not be valid for
	//  signing, the key may be relying on a cached handle which is no longer
	//  valid, or the CSP won't allow arbitrary data to be signed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_CLIENT_AUTH_SIGNATURE_FAILED = -141
	//  The message was too large for the transport. (for example a UDP message
	//  which exceeds size threshold).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_MSG_TOO_BIG = -142
	//  Websocket protocol error. Indicates that we are terminating the connection
	//  due to a malformed frame or other protocol violation.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_WS_PROTOCOL_ERROR = -145
	//  Returned when attempting to bind an address that is already in use.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ADDRESS_IN_USE = -147
	//  The certificate didn't match the built-in public key pins for the host name.
	//  The pins are set in net/http/transport_security_state.cc and require that
	//  one of a set of public keys exist on the path from the leaf to the root.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_PINNED_KEY_NOT_IN_CERT_CHAIN = -150
	//  Server request for client certificate did not contain any types we support.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CLIENT_AUTH_CERT_TYPE_UNSUPPORTED = -151
	//  An SSL peer sent us a fatal decrypt_error alert. This typically occurs when
	//  a peer could not correctly verify a signature (in CertificateVerify or
	//  ServerKeyExchange) or validate a Finished message.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_DECRYPT_ERROR_ALERT = -153
	//  There are too many pending WebSocketJob instances, so the new job was not
	//  pushed to the queue.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_WS_THROTTLE_QUEUE_TOO_LARGE = -154
	//  The SSL server certificate changed in a renegotiation.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_SERVER_CERT_CHANGED = -156
	//  The SSL server sent us a fatal unrecognized_name alert.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_UNRECOGNIZED_NAME_ALERT = -159
	//  Failed to set the socket's receive buffer size as requested.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SOCKET_SET_RECEIVE_BUFFER_SIZE_ERROR = -160
	//  Failed to set the socket's send buffer size as requested.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SOCKET_SET_SEND_BUFFER_SIZE_ERROR = -161
	//  Failed to set the socket's receive buffer size as requested, despite success
	//  return code from setsockopt.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SOCKET_RECEIVE_BUFFER_SIZE_UNCHANGEABLE = -162
	//  Failed to set the socket's send buffer size as requested, despite success
	//  return code from setsockopt.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SOCKET_SEND_BUFFER_SIZE_UNCHANGEABLE = -163
	//  Failed to import a client certificate from the platform store into the SSL
	//  library.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_CLIENT_AUTH_CERT_BAD_FORMAT = -164
	//  Resolving a hostname to an IP address list included the IPv4 address
	//  "127.0.53.53". This is a special IP address which ICANN has recommended to
	//  indicate there was a name collision, and alert admins to a potential
	//  problem.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ICANN_NAME_COLLISION = -166
	//  The SSL server presented a certificate which could not be decoded. This is
	//  not a certificate error code as no X509Certificate object is available. This
	//  error is fatal.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_SERVER_CERT_BAD_FORMAT = -167
	//  Certificate Transparency: Received a signed tree head that failed to parse.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CT_STH_PARSING_FAILED = -168
	//  Certificate Transparency: Received a signed tree head whose JSON parsing was
	//  OK but was missing some of the fields.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CT_STH_INCOMPLETE = -169
	//  The attempt to reuse a connection to send proxy auth credentials failed
	//  before the AuthController was used to generate credentials. The caller should
	//  reuse the controller with a new connection. This error is only used
	//  internally by the network stack.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UNABLE_TO_REUSE_CONNECTION_FOR_PROXY_AUTH = -170
	//  Certificate Transparency: Failed to parse the received consistency proof.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CT_CONSISTENCY_PROOF_PARSING_FAILED = -171
	//  The SSL server required an unsupported cipher suite that has since been
	//  removed. This error will temporarily be signaled on a fallback for one or two
	//  releases immediately following a cipher suite's removal, after which the
	//  fallback will be removed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_OBSOLETE_CIPHER = -172
	//  When a WebSocket handshake is done successfully and the connection has been
	//  upgraded, the URLRequest is cancelled with this error code.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_WS_UPGRADE = -173
	//  Socket ReadIfReady support is not implemented. This error should not be user
	//  visible, because the normal Read() method is used as a fallback.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_READ_IF_READY_NOT_IMPLEMENTED = -174
	//  No socket buffer space is available.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NO_BUFFER_SPACE = -176
	//  There were no common signature algorithms between our client certificate
	//  private key and the server's preferences.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_CLIENT_AUTH_NO_COMMON_ALGORITHMS = -177
	//  TLS 1.3 early data was rejected by the server. This will be received before
	//  any data is returned from the socket. The request should be retried with
	//  early data disabled.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_EARLY_DATA_REJECTED = -178
	//  TLS 1.3 early data was offered, but the server responded with TLS 1.2 or
	//  earlier. This is an internal error code to account for a
	//  backwards-compatibility issue with early data and TLS 1.2. It will be
	//  received before any data is returned from the socket. The request should be
	//  retried with early data disabled.
	//
	//  See https://tools.ietf.org/html/rfc8446#appendix-D.3 for details.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_WRONG_VERSION_ON_EARLY_DATA = -179
	//  TLS 1.3 was enabled, but a lower version was negotiated and the server
	//  returned a value indicating it supported TLS 1.3. This is part of a security
	//  check in TLS 1.3, but it may also indicate the user is behind a buggy
	//  TLS-terminating proxy which implemented TLS 1.2 incorrectly. (See
	//  https://crbug.com/boringssl/226.)
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_TLS13_DOWNGRADE_DETECTED = -180
	//  The server's certificate has a keyUsage extension incompatible with the
	//  negotiated TLS key exchange method.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SSL_KEY_USAGE_INCOMPATIBLE = -181
	//  The ECHConfigList fetched over DNS cannot be parsed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_ECH_CONFIG_LIST = -182
	//  ECH was enabled, but the server was unable to decrypt the encrypted
	//  ClientHello.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ECH_NOT_NEGOTIATED = -183
	//  ECH was enabled, the server was unable to decrypt the encrypted ClientHello,
	//  and additionally did not present a certificate valid for the public name.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ECH_FALLBACK_CERTIFICATE_INVALID = -184
	//  An attempt to proxy a request failed because the proxy wasn't able to
	//  successfully connect to the destination. This likely indicates an issue with
	//  the request itself (for instance, the hostname failed to resolve to an IP
	//  address or the destination server refused the connection). This error code
	//  is used to indicate that the error is outside the control of the proxy server
	//  and thus the proxy chain should not be marked as bad. This is in contrast to
	//  ERR_TUNNEL_CONNECTION_FAILED which is used for general purpose errors
	//  connecting to the proxy and by the proxy request response handling when a
	//  proxy delegate doesn't indicate via a different error code whether proxy
	//  fallback should occur. Note that for IP Protection proxies this error code
	//  causes the proxy to be marked as bad since the preference is to fail open for
	//  general purpose errors, but for other proxies this error does not cause the
	//  proxy to be marked as bad.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PROXY_UNABLE_TO_CONNECT_TO_DESTINATION = -186
	//  Some implementations of ProxyDelegate query a separate entity to know whether
	//  it should cancel tunnel prior to:
	//  - The HTTP CONNECT requests being sent out
	//  - The HTTP CONNECT response being parsed by //net
	//  An example is CronetProxyDelegate: Cronet allows developers to decide whether
	//  the tunnel being established should be canceled.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PROXY_DELEGATE_CANCELED_CONNECT_REQUEST = -187
	//  Some implementations of ProxyDelegate query a separate entity to know whether
	//  it should cancel tunnel prior to:
	//  - The HTTP CONNECT requests being sent out
	//  - The HTTP CONNECT response being parsed by //net
	//  An example is CronetProxyDelegate: Cronet allows developers to decide whether
	//  the tunnel being established should be canceled.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PROXY_DELEGATE_CANCELED_CONNECT_RESPONSE = -188
	//  The server responded with a certificate whose common name did not match
	//  the host name. This could mean:
	//
	//  1. An attacker has redirected our traffic to their server and is
	//  presenting a certificate for which they know the private key.
	//
	//  2. The server is misconfigured and responding with the wrong cert.
	//
	//  3. The user is on a wireless network and is being redirected to the
	//  network's login page.
	//
	//  4. The OS has used a DNS search suffix and the server doesn't have
	//  a certificate for the abbreviated name in the address bar.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_COMMON_NAME_INVALID = -200
	//  The server responded with a certificate that, by our clock, appears to
	//  either not yet be valid or to have expired. This could mean:
	//
	//  1. An attacker is presenting an old certificate for which they have
	//  managed to obtain the private key.
	//
	//  2. The server is misconfigured and is not presenting a valid cert.
	//
	//  3. Our clock is wrong.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_DATE_INVALID = -201
	//  The server responded with a certificate that is signed by an authority
	//  we don't trust. The could mean:
	//
	//  1. An attacker has substituted the real certificate for a cert that
	//  contains their public key and is signed by their cousin.
	//
	//  2. The server operator has a legitimate certificate from a CA we don't
	//  know about, but should trust.
	//
	//  3. The server is presenting a self-signed certificate, providing no
	//  defense against active attackers (but foiling passive attackers).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_AUTHORITY_INVALID = -202
	//  The server responded with a certificate that contains errors.
	//  This error is not recoverable.
	//
	//  MSDN describes this error as follows:
	//  "The SSL certificate contains errors."
	//  NOTE: It's unclear how this differs from ERR_CERT_INVALID. For consistency,
	//  use that code instead of this one from now on.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_CONTAINS_ERRORS = -203
	//  The certificate has no mechanism for determining if it is revoked. In
	//  effect, this certificate cannot be revoked.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_NO_REVOCATION_MECHANISM = -204
	//  Revocation information for the security certificate for this site is not
	//  available. This could mean:
	//
	//  1. An attacker has compromised the private key in the certificate and is
	//  blocking our attempt to find out that the cert was revoked.
	//
	//  2. The certificate is unrevoked, but the revocation server is busy or
	//  unavailable.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_UNABLE_TO_CHECK_REVOCATION = -205
	//  The server responded with a certificate has been revoked.
	//  We have the capability to ignore this error, but it is probably not the
	//  thing to do.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_REVOKED = -206
	//  The server responded with a certificate that is invalid.
	//  This error is not recoverable.
	//
	//  MSDN describes this error as follows:
	//  "The SSL certificate is invalid."
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_INVALID = -207
	//  The server responded with a certificate that is signed using a weak
	//  signature algorithm.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_WEAK_SIGNATURE_ALGORITHM = -208
	//  The host name specified in the certificate is not unique.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_NON_UNIQUE_NAME = -210
	//  The server responded with a certificate that contains a weak key (e.g.
	//  a too-small RSA key).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_WEAK_KEY = -211
	//  The certificate claimed DNS names that are in violation of name constraints.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_NAME_CONSTRAINT_VIOLATION = -212
	//  The certificate's validity period is too long.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_VALIDITY_TOO_LONG = -213
	//  Certificate Transparency was required for this connection, but the server
	//  did not provide CT information that complied with the policy.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERTIFICATE_TRANSPARENCY_REQUIRED = -214
	//  The certificate is known to be used for interception by an entity other
	//  the device owner.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_KNOWN_INTERCEPTION_BLOCKED = -217
	//  -218 was SSL_OBSOLETE_VERSION which is not longer used. TLS 1.0/1.1 instead
	//  cause SSL_VERSION_OR_CIPHER_MISMATCH now.
	//  The certificate is self signed and it's being used for either an RFC1918 IP
	//  literal URL, or a url ending in .local.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_SELF_SIGNED_LOCAL_NETWORK = -219
	//  The value immediately past the last certificate error code.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_END = -220
	//  The URL is invalid.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_URL = -300
	//  The scheme of the URL is disallowed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DISALLOWED_URL_SCHEME = -301
	//  The scheme of the URL is unknown.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UNKNOWN_URL_SCHEME = -302
	//  Attempting to load an URL resulted in a redirect to an invalid URL.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_REDIRECT = -303
	//  Attempting to load an URL resulted in too many redirects.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_TOO_MANY_REDIRECTS = -310
	//  Attempting to load an URL resulted in an unsafe redirect (e.g., a redirect
	//  to file:// is considered unsafe).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UNSAFE_REDIRECT = -311
	//  Attempting to load an URL with an unsafe port number. These are port
	//  numbers that correspond to services, which are not robust to spurious input
	//  that may be constructed as a result of an allowed web construct (e.g., HTTP
	//  looks a lot like SMTP, so form submission to port 25 is denied).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UNSAFE_PORT = -312
	//  The server's response was invalid.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_RESPONSE = -320
	//  Error in chunked transfer encoding.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_CHUNKED_ENCODING = -321
	//  The server did not support the request method.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_METHOD_NOT_SUPPORTED = -322
	//  The response was 407 (Proxy Authentication Required), yet we did not send
	//  the request to a proxy.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UNEXPECTED_PROXY_AUTH = -323
	//  The server closed the connection without sending any data.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_EMPTY_RESPONSE = -324
	//  The headers section of the response is too large.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_RESPONSE_HEADERS_TOO_BIG = -325
	//  The evaluation of the PAC script failed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PAC_SCRIPT_FAILED = -327
	//  The response was 416 (Requested range not satisfiable) and the server cannot
	//  satisfy the range requested.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_REQUEST_RANGE_NOT_SATISFIABLE = -328
	//  The identity used for authentication is invalid.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_MALFORMED_IDENTITY = -329
	//  Content decoding of the response body failed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CONTENT_DECODING_FAILED = -330
	//  An operation could not be completed because all network IO
	//  is suspended.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NETWORK_IO_SUSPENDED = -331
	//  There are no supported proxies in the provided list.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NO_SUPPORTED_PROXIES = -336
	//  There is an HTTP/2 protocol error.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SPDY_PROTOCOL_ERROR = -337
	//  Credentials could not be established during HTTP Authentication.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_AUTH_CREDENTIALS = -338
	//  An HTTP Authentication scheme was tried which is not supported on this
	//  machine.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UNSUPPORTED_AUTH_SCHEME = -339
	//  Detecting the encoding of the response failed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ENCODING_DETECTION_FAILED = -340
	//  (GSSAPI) No Kerberos credentials were available during HTTP Authentication.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_MISSING_AUTH_CREDENTIALS = -341
	//  An unexpected, but documented, SSPI or GSSAPI status code was returned.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UNEXPECTED_SECURITY_LIBRARY_STATUS = -342
	//  The environment was not set up correctly for authentication.
	//  For example, no KDC could be found or the principal is unknown.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_MISCONFIGURED_AUTH_ENVIRONMENT = -343
	//  An undocumented SSPI or GSSAPI status code was returned.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UNDOCUMENTED_SECURITY_LIBRARY_STATUS = -344
	//  The HTTP response was too big to drain.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_RESPONSE_BODY_TOO_BIG_TO_DRAIN = -345
	//  The HTTP response contained multiple distinct Content-Length headers.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_RESPONSE_HEADERS_MULTIPLE_CONTENT_LENGTH = -346
	//  HTTP/2 headers have been received, but not all of them - status or version
	//  headers are missing, so we're expecting additional frames to complete them.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INCOMPLETE_SPDY_HEADERS = -347
	//  No PAC URL configuration could be retrieved from DHCP. This can indicate
	//  either a failure to retrieve the DHCP configuration, or that there was no
	//  PAC URL configured in DHCP.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PAC_NOT_IN_DHCP = -348
	//  The HTTP response contained multiple Content-Disposition headers.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_RESPONSE_HEADERS_MULTIPLE_CONTENT_DISPOSITION = -349
	//  The HTTP response contained multiple Location headers.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_RESPONSE_HEADERS_MULTIPLE_LOCATION = -350
	//  HTTP/2 server refused the request without processing, and sent either a
	//  GOAWAY frame with error code NO_ERROR and Last-Stream-ID lower than the
	//  stream id corresponding to the request indicating that this request has not
	//  been processed yet, or a RST_STREAM frame with error code REFUSED_STREAM.
	//  Client MAY retry (on a different connection). See RFC7540 Section 8.1.4.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SPDY_SERVER_REFUSED_STREAM = -351
	//  HTTP/2 server didn't respond to the PING message.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SPDY_PING_FAILED = -352
	//  The HTTP response body transferred fewer bytes than were advertised by the
	//  Content-Length header when the connection is closed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CONTENT_LENGTH_MISMATCH = -354
	//  The HTTP response body is transferred with Chunked-Encoding, but the
	//  terminating zero-length chunk was never sent when the connection is closed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INCOMPLETE_CHUNKED_ENCODING = -355
	//  There is a QUIC protocol error.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_QUIC_PROTOCOL_ERROR = -356
	//  The HTTP headers were truncated by an EOF.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_RESPONSE_HEADERS_TRUNCATED = -357
	//  The QUIC crypto handshake failed. This means that the server was unable
	//  to read any requests sent, so they may be resent.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_QUIC_HANDSHAKE_FAILED = -358
	//  Transport security is inadequate for the HTTP/2 version.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SPDY_INADEQUATE_TRANSPORT_SECURITY = -360
	//  The peer violated HTTP/2 flow control.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SPDY_FLOW_CONTROL_ERROR = -361
	//  The peer sent an improperly sized HTTP/2 frame.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SPDY_FRAME_SIZE_ERROR = -362
	//  Decoding or encoding of compressed HTTP/2 headers failed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SPDY_COMPRESSION_ERROR = -363
	//  Proxy Auth Requested without a valid Client Socket Handle.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PROXY_AUTH_REQUESTED_WITH_NO_CONNECTION = -364
	//  HTTP_1_1_REQUIRED error code received on HTTP/2 session.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_HTTP_1_1_REQUIRED = -365
	//  HTTP_1_1_REQUIRED error code received on HTTP/2 session to proxy.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PROXY_HTTP_1_1_REQUIRED = -366
	//  The PAC script terminated fatally and must be reloaded.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PAC_SCRIPT_TERMINATED = -367
	//  The server was expected to return an HTTP/1.x response, but did not. Rather
	//  than treat it as HTTP/0.9, this error is returned.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_HTTP_RESPONSE = -370
	//  Initializing content decoding failed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CONTENT_DECODING_INIT_FAILED = -371
	//  Received HTTP/2 RST_STREAM frame with NO_ERROR error code. This error should
	//  be handled internally by HTTP/2 code, and should not make it above the
	//  SpdyStream layer.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SPDY_RST_STREAM_NO_ERROR_RECEIVED = -372
	//  An HTTP transaction was retried too many times due for authentication or
	//  invalid certificates. This may be due to a bug in the net stack that would
	//  otherwise infinite loop, or if the server or proxy continually requests fresh
	//  credentials or presents a fresh invalid certificate.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_TOO_MANY_RETRIES = -375
	//  Received an HTTP/2 frame on a closed stream.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SPDY_STREAM_CLOSED = -376
	//  The server returned a non-2xx HTTP response code.
	//
	//  Note that this error is only used by certain APIs that interpret the HTTP
	//  response itself. URLRequest for instance just passes most non-2xx
	//  response back as success.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_HTTP_RESPONSE_CODE_FAILURE = -379
	//  The certificate presented on a QUIC connection does not chain to a known root
	//  and the origin connected to is not on a list of domains where unknown roots
	//  are allowed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_QUIC_CERT_ROOT_NOT_KNOWN = -380
	//  A GOAWAY frame has been received indicating that the request has not been
	//  processed and is therefore safe to retry on a different connection.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_QUIC_GOAWAY_REQUEST_CAN_BE_RETRIED = -381
	//  The ACCEPT_CH restart has been triggered too many times
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_TOO_MANY_ACCEPT_CH_RESTARTS = -382
	//  The IP address space of the remote endpoint differed from the previous
	//  observed value during the same request. Any cache entry for the affected
	//  request should be invalidated.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INCONSISTENT_IP_ADDRESS_SPACE = -383
	//  The IP address space of the cached remote endpoint is blocked by private
	//  network access check.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHED_IP_ADDRESS_SPACE_BLOCKED_BY_LOCAL_NETWORK_ACCESS_POLICY = -384
	//  The connection is blocked by private network access checks.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOCKED_BY_LOCAL_NETWORK_ACCESS_CHECKS = -385
	//  Content decoding failed due to the zstd window size being too big (over 8MB).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ZSTD_WINDOW_SIZE_TOO_BIG = -386
	//  The compression dictionary cannot be loaded.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DICTIONARY_LOAD_FAILED = -387
	//  The header of dictionary compressed stream does not match the expected value.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_UNEXPECTED_CONTENT_DICTIONARY_HEADER = -388
	//  The cache does not have the requested entry.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_MISS = -400
	//  Unable to read from the disk cache.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_READ_FAILURE = -401
	//  Unable to write to the disk cache.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_WRITE_FAILURE = -402
	//  The operation is not supported for this entry.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_OPERATION_NOT_SUPPORTED = -403
	//  The disk cache is unable to open this entry.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_OPEN_FAILURE = -404
	//  The disk cache is unable to create this entry.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_CREATE_FAILURE = -405
	//  Multiple transactions are racing to create disk cache entries. This is an
	//  internal error returned from the HttpCache to the HttpCacheTransaction that
	//  tells the transaction to restart the entry-creation logic because the state
	//  of the cache has changed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_RACE = -406
	//  The cache was unable to read a checksum record on an entry. This can be
	//  returned from attempts to read from the cache. It is an internal error,
	//  returned by the SimpleCache backend, but not by any URLRequest methods
	//  or members.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_CHECKSUM_READ_FAILURE = -407
	//  The cache found an entry with an invalid checksum. This can be returned from
	//  attempts to read from the cache. It is an internal error, returned by the
	//  SimpleCache backend, but not by any URLRequest methods or members.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_CHECKSUM_MISMATCH = -408
	//  Internal error code for the HTTP cache. The cache lock timeout has fired.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_LOCK_TIMEOUT = -409
	//  Received a challenge after the transaction has read some data, and the
	//  credentials aren't available. There isn't a way to get them at that point.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_AUTH_FAILURE_AFTER_READ = -410
	//  Internal not-quite error code for the HTTP cache. In-memory hints suggest
	//  that the cache entry would not have been usable with the transaction's
	//  current configuration (e.g. load flags, mode, etc.)
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_ENTRY_NOT_SUITABLE = -411
	//  The disk cache is unable to doom this entry.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_DOOM_FAILURE = -412
	//  The disk cache is unable to open or create this entry.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CACHE_OPEN_OR_CREATE_FAILURE = -413
	//  The server's response was insecure (e.g. there was a cert error).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INSECURE_RESPONSE = -501
	//  An attempt to import a client certificate failed, as the user's key
	//  database lacked a corresponding private key.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NO_PRIVATE_KEY_FOR_CERT = -502
	//  An error adding a certificate to the OS certificate database.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_ADD_USER_CERT_FAILED = -503
	//  An error occurred while handling a signed exchange.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_SIGNED_EXCHANGE = -504
	//  An error occurred while handling a Web Bundle source.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_INVALID_WEB_BUNDLE = -505
	//  A Trust Tokens protocol operation-executing request failed for one of a
	//  number of reasons (precondition failure, internal error, bad response).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_TRUST_TOKEN_OPERATION_FAILED = -506
	//  When handling a Trust Tokens protocol operation-executing request, the system
	//  was able to execute the request's Trust Tokens operation without sending the
	//  request to its destination: for instance, the results could have been present
	//  in a local cache (for redemption) or the operation could have been diverted
	//  to a local provider (for "platform-provided" issuance).
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_TRUST_TOKEN_OPERATION_SUCCESS_WITHOUT_SENDING_REQUEST = -507
	//  This is a placeholder value that should never be used within net.
	//  When Cronet APIs are being backed by HttpEngine (i.e., HttpEngineProvider is
	//  being used), org.chromium.net.NetworkException#getCronetInternalErrorCode is
	//  not supported (android.net.http.NetworkException#getCronetInternalErrorCode
	//  does not exist). In this scenario, getCronetInternalErrorCode will always
	//  return this error. This is a first step towards the deprecation of
	//  getCronetInternalErrorCode.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_HTTPENGINE_PROVIDER_IN_USE = -508
	//  PKCS #12 import failed due to incorrect password.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PKCS12_IMPORT_BAD_PASSWORD = -701
	//  PKCS #12 import failed due to other error.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PKCS12_IMPORT_FAILED = -702
	//  CA import failed - not a CA cert.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_IMPORT_CA_CERT_NOT_CA = -703
	//  Import failed - certificate already exists in database.
	//  Note it's a little weird this is an error but reimporting a PKCS12 is ok
	//  (no-op). That's how Mozilla does it, though.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_IMPORT_CERT_ALREADY_EXISTS = -704
	//  CA import failed due to some other error.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_IMPORT_CA_CERT_FAILED = -705
	//  Server certificate import failed due to some internal error.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_IMPORT_SERVER_CERT_FAILED = -706
	//  PKCS #12 import failed due to invalid MAC.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PKCS12_IMPORT_INVALID_MAC = -707
	//  PKCS #12 import failed due to invalid/corrupt file.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PKCS12_IMPORT_INVALID_FILE = -708
	//  PKCS #12 import failed due to unsupported features.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PKCS12_IMPORT_UNSUPPORTED = -709
	//  Key generation failed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_KEY_GENERATION_FAILED = -710
	//  Failure to export private key.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_PRIVATE_KEY_EXPORT_FAILED = -712
	//  Self-signed certificate generation failed.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_SELF_SIGNED_CERT_GENERATION_FAILED = -713
	//  The certificate database changed in some way.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_DATABASE_CHANGED = -714
	//  The certificate verifier configuration changed in some way.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_CERT_VERIFIER_CHANGED = -716
	//  DNS resolver received a malformed response.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_MALFORMED_RESPONSE = -800
	//  DNS server requires TCP
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_SERVER_REQUIRES_TCP = -801
	//  DNS transaction timed out.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_TIMED_OUT = -803
	//  The entry was not found in cache or other local sources, for lookups where
	//  only local sources were queried.
	//  TODO(ericorth): Consider renaming to DNS_LOCAL_MISS or something like that as
	//  the cache is not necessarily queried either.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_NS_CACHE_MISS = -804
	//  Suffix search list rules prevent resolution of the given host name.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_SEARCH_EMPTY = -805
	//  Failed to sort addresses according to RFC3484.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_SORT_ERROR = -806
	//  Failed to resolve the hostname of a DNS-over-HTTPS server.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_SECURE_RESOLVER_HOSTNAME_RESOLUTION_FAILED = -808
	//  DNS identified the request as disallowed for insecure connection (http/ws).
	//  Error should be handled as if an HTTP redirect was received to redirect to
	//  https or wss.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_NAME_HTTPS_ONLY = -809
	//  All DNS requests associated with this job have been cancelled.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_REQUEST_CANCELLED = -810
	//  The hostname resolution of HTTPS record was expected to be resolved with
	//  alpn values of supported protocols, but did not.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_NO_MATCHING_SUPPORTED_ALPN = -811
	//  When checking whether secure DNS can be used, the response returned for the
	//  requested probe record either had no answer or was invalid.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_SECURE_PROBE_RECORD_INVALID = -814
	//  Returned when DNS cache invalidation is in progress. This is a
	//  transient error. Callers may want to retry later.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_CACHE_INVALIDATION_IN_PROGRESS = -815
	//  The DNS server responded with a format error response code.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_FORMAT_ERROR = -816
	//  The DNS server responded with a server failure response code.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_SERVER_FAILURE = -817
	//  The DNS server responded that the query type is not implemented.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_NOT_IMPLEMENTED = -818
	//  The DNS server responded that the request was refused.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_REFUSED = -819
	//  The DNS server responded with an rcode indicating that the request failed,
	//  but the rcode is not one that we have a specific error code for. In other
	//  words, the rcode was not one of the following:
	//  <code>
	//  - NOERR
	//  - FORMERR
	//  - SERVFAIL
	//  - NXDOMAIN
	//  - NOTIMP
	//  - REFUSED
	//  </code>
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_DNS_OTHER_FAILURE = -820
	//  The following errors are for mapped from a subset of invalid
	//  storage::BlobStatus.
	//  The construction arguments are invalid. This is considered a bad IPC.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOB_INVALID_CONSTRUCTION_ARGUMENTS = -900
	//  We don't have enough memory for the blob.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOB_OUT_OF_MEMORY = -901
	//  We couldn't create or write to a file. File system error, like a full disk.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOB_FILE_WRITE_FAILED = -902
	//  The renderer was destroyed while data was in transit.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOB_SOURCE_DIED_IN_TRANSIT = -903
	//  The renderer destructed the blob before it was done transferring, and there
	//  were no outstanding references (no one is waiting to read) to keep the
	//  blob alive.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOB_DEREFERENCED_WHILE_BUILDING = -904
	//  A blob that we referenced during construction is broken, or a browser-side
	//  builder tries to build a blob with a blob reference that isn't finished
	//  constructing.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOB_REFERENCED_BLOB_BROKEN = -905
	//  A file that we referenced during construction is not accessible to the
	//  renderer trying to create the blob.
	//  TCefErrorCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_errorcode_t)</see>
	ERR_BLOB_REFERENCED_FILE_UNAVAILABLE = -906
	//  command_id constants declared in cef_command_ids.h and used by some callbacks in ICefCommandHandler
	//  cef_command_ids.h is generated in /include/cef_command_ids.h
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/app/chrome_command_ids.h">The command_id values are also available in chrome/app/chrome_command_ids.h</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:ui/base/command_id_constants.h">COMMAND_ID_FIRST_UNBOUNDED is available in ui/base/command_id_constants.h</see>
	COMMAND_ID_FIRST_UNBOUNDED                                                      = 0xE000
	IDC_MinimumLabelValue                                                           = 4000
	IDC_BACK                                                                        = 33000
	IDC_FORWARD                                                                     = 33001
	IDC_RELOAD                                                                      = 33002
	IDC_HOME                                                                        = 33003
	IDC_OPEN_CURRENT_URL                                                            = 33004
	IDC_STOP                                                                        = 33006
	IDC_RELOAD_BYPASSING_CACHE                                                      = 33007
	IDC_RELOAD_CLEARING_CACHE                                                       = 33009
	IDC_NEW_WINDOW                                                                  = 34000
	IDC_NEW_INCOGNITO_WINDOW                                                        = 34001
	IDC_CLOSE_WINDOW                                                                = 34012
	IDC_ALWAYS_ON_TOP                                                               = 34013
	IDC_NEW_TAB                                                                     = 34014
	IDC_CLOSE_TAB                                                                   = 34015
	IDC_SELECT_NEXT_TAB                                                             = 34016
	IDC_SELECT_PREVIOUS_TAB                                                         = 34017
	IDC_SELECT_TAB_0                                                                = 34018
	IDC_SELECT_TAB_1                                                                = 34019
	IDC_SELECT_TAB_2                                                                = 34020
	IDC_SELECT_TAB_3                                                                = 34021
	IDC_SELECT_TAB_4                                                                = 34022
	IDC_SELECT_TAB_5                                                                = 34023
	IDC_SELECT_TAB_6                                                                = 34024
	IDC_SELECT_TAB_7                                                                = 34025
	IDC_SELECT_LAST_TAB                                                             = 34026
	IDC_DUPLICATE_TAB                                                               = 34027
	IDC_RESTORE_TAB                                                                 = 34028
	IDC_SHOW_AS_TAB                                                                 = 34029
	IDC_FULLSCREEN                                                                  = 34030
	IDC_EXIT                                                                        = 34031
	IDC_MOVE_TAB_NEXT                                                               = 34032
	IDC_MOVE_TAB_PREVIOUS                                                           = 34033
	IDC_SEARCH                                                                      = 34035
	IDC_WINDOW_MENU                                                                 = 34045
	IDC_MINIMIZE_WINDOW                                                             = 34046
	IDC_MAXIMIZE_WINDOW                                                             = 34047
	IDC_ALL_WINDOWS_FRONT                                                           = 34048
	IDC_NAME_WINDOW                                                                 = 34049
	IDC_TOGGLE_MULTITASK_MENU                                                       = 34050
	IDC_USE_SYSTEM_TITLE_BAR                                                        = 34051
	IDC_RESTORE_WINDOW                                                              = 34052
	IDC_OPEN_IN_PWA_WINDOW                                                          = 34053
	IDC_MOVE_TAB_TO_NEW_WINDOW                                                      = 34054
	IDC_NEW_SPLIT_TAB                                                               = 34055
	IDC_TOGGLE_VERTICAL_TABS                                                        = 34056
	IDC_VERTICAL_TABS_SEND_FEEDBACK                                                 = 34057
	IDC_COPY_URL                                                                    = 34060
	IDC_OPEN_IN_CHROME                                                              = 34061
	IDC_WEB_APP_SETTINGS                                                            = 34062
	IDC_WEB_APP_MENU_APP_INFO                                                       = 34063
	IDC_WEB_APP_UPGRADE_DIALOG                                                      = 34064
	IDC_VISIT_DESKTOP_OF_LRU_USER_2                                                 = 34080
	IDC_VISIT_DESKTOP_OF_LRU_USER_3                                                 = 34081
	IDC_VISIT_DESKTOP_OF_LRU_USER_4                                                 = 34082
	IDC_VISIT_DESKTOP_OF_LRU_USER_5                                                 = 34083
	IDC_ADD_NEW_TAB_TO_GROUP                                                        = 34100
	IDC_CREATE_NEW_TAB_GROUP                                                        = 34101
	IDC_FOCUS_NEXT_TAB_GROUP                                                        = 34102
	IDC_FOCUS_PREV_TAB_GROUP                                                        = 34103
	IDC_CLOSE_TAB_GROUP                                                             = 34104
	IDC_GROUP_UNGROUPED_TABS                                                        = 34105
	IDC_CREATE_NEW_TAB_GROUP_TOP_LEVEL                                              = 34106
	IDC_ADD_NEW_TAB_RECENT_GROUP                                                    = 34107
	IDC_UNFOCUS_TAB_GROUP                                                           = 34108
	IDC_BOOKMARK_THIS_TAB                                                           = 35000
	IDC_BOOKMARK_ALL_TABS                                                           = 35001
	IDC_VIEW_SOURCE                                                                 = 35002
	IDC_PRINT                                                                       = 35003
	IDC_SAVE_PAGE                                                                   = 35004
	IDC_EMAIL_PAGE_LOCATION                                                         = 35006
	IDC_BASIC_PRINT                                                                 = 35007
	IDC_SAVE_CREDIT_CARD_FOR_PAGE                                                   = 35008
	IDC_SHOW_TRANSLATE                                                              = 35009
	IDC_MANAGE_PASSWORDS_FOR_PAGE                                                   = 35010
	IDC_ROUTE_MEDIA                                                                 = 35011
	IDC_WINDOW_MUTE_SITE                                                            = 35012
	IDC_WINDOW_PIN_TAB                                                              = 35013
	IDC_WINDOW_GROUP_TAB                                                            = 35014
	IDC_MIGRATE_LOCAL_CREDIT_CARD_FOR_PAGE                                          = 35015
	IDC_SEND_TAB_TO_SELF                                                            = 35016
	IDC_FOCUS_THIS_TAB                                                              = 35017
	IDC_QRCODE_GENERATOR                                                            = 35021
	IDC_WINDOW_CLOSE_TABS_TO_RIGHT                                                  = 35022
	IDC_WINDOW_CLOSE_OTHER_TABS                                                     = 35023
	IDC_NEW_TAB_TO_RIGHT                                                            = 35024
	IDC_SAVE_AUTOFILL_ADDRESS                                                       = 35025
	IDC_OFFERS_AND_REWARDS_FOR_PAGE                                                 = 35026
	IDC_WEBAUTHN                                                                    = 35027
	IDC_SHARING_HUB                                                                 = 35028
	IDC_SHARING_HUB_MENU                                                            = 35029
	IDC_FILLED_CARD_INFORMATION                                                     = 35030
	IDC_SHARING_HUB_SCREENSHOT                                                      = 35031
	IDC_VIRTUAL_CARD_ENROLL                                                         = 35032
	IDC_SAVE_IBAN_FOR_PAGE                                                          = 35035
	IDC_AUTOFILL_MANDATORY_REAUTH                                                   = 35036
	IDC_PROFILE_MENU_IN_APP_MENU                                                    = 35039
	IDC_PASSWORDS_AND_AUTOFILL_MENU                                                 = 35040
	IDC_SHOW_PASSWORD_MANAGER                                                       = 35041
	IDC_SHOW_PAYMENT_METHODS                                                        = 35042
	IDC_SHOW_ADDRESSES                                                              = 35043
	IDC_ORGANIZE_TABS                                                               = 35044
	IDC_SEND_SHARED_TAB_GROUP_FEEDBACK                                              = 35046
	IDC_SHOW_IDENTITY_DOCS                                                          = 35047
	IDC_SHOW_TRAVEL                                                                 = 35048
	IDC_SHOW_CONTACT_INFO                                                           = 35049
	IDC_MUTE_TARGET_SITE                                                            = 35050
	IDC_PIN_TARGET_TAB                                                              = 35051
	IDC_GROUP_TARGET_TAB                                                            = 35052
	IDC_DUPLICATE_TARGET_TAB                                                        = 35053
	IDC_CUT                                                                         = 36000
	IDC_COPY                                                                        = 36001
	IDC_PASTE                                                                       = 36003
	IDC_EDIT_MENU                                                                   = 36004
	IDC_FIND                                                                        = 37000
	IDC_FIND_NEXT                                                                   = 37001
	IDC_FIND_PREVIOUS                                                               = 37002
	IDC_CLOSE_FIND_OR_STOP                                                          = 37003
	IDC_FIND_MENU                                                                   = 37100
	IDC_FIND_AND_EDIT_MENU                                                          = 37200
	IDC_SAVE_AND_SHARE_MENU                                                         = 37300
	IDC_CUSTOMIZE_CHROME                                                            = 37350
	IDC_CLOSE_PROFILE                                                               = 35351
	IDC_MANAGE_GOOGLE_ACCOUNT                                                       = 35352
	IDC_SHOW_SYNC_SETTINGS                                                          = 35353
	IDC_TURN_ON_SYNC                                                                = 35354
	IDC_SHOW_SIGNIN_WHEN_PAUSED                                                     = 35355
	IDC_OPEN_GUEST_PROFILE                                                          = 35356
	IDC_ADD_NEW_PROFILE                                                             = 35357
	IDC_MANAGE_CHROME_PROFILES                                                      = 35358
	IDC_SHOW_SIGNIN                                                                 = 35359
	IDC_SHOW_SYNC_PASSPHRASE_DIALOG                                                 = 35360
	IDC_ZOOM_MENU                                                                   = 38000
	IDC_ZOOM_PLUS                                                                   = 38001
	IDC_ZOOM_NORMAL                                                                 = 38002
	IDC_ZOOM_MINUS                                                                  = 38003
	IDC_FOCUS_TOOLBAR                                                               = 39000
	IDC_FOCUS_LOCATION                                                              = 39001
	IDC_FOCUS_SEARCH                                                                = 39002
	IDC_FOCUS_MENU_BAR                                                              = 39003
	IDC_FOCUS_NEXT_PANE                                                             = 39004
	IDC_FOCUS_PREVIOUS_PANE                                                         = 39005
	IDC_FOCUS_BOOKMARKS                                                             = 39006
	IDC_FOCUS_INACTIVE_POPUP_FOR_ACCESSIBILITY                                      = 39007
	IDC_FOCUS_WEB_CONTENTS_PANE                                                     = 39009
	IDC_OPEN_FILE                                                                   = 40000
	IDC_CREATE_SHORTCUT                                                             = 40002
	IDC_DEVELOPER_MENU                                                              = 40003
	IDC_DEV_TOOLS                                                                   = 40004
	IDC_DEV_TOOLS_CONSOLE                                                           = 40005
	IDC_TASK_MANAGER                                                                = 40006
	IDC_DEV_TOOLS_DEVICES                                                           = 40007
	IDC_FEEDBACK                                                                    = 40008
	IDC_SHOW_BOOKMARK_BAR                                                           = 40009
	IDC_SHOW_HISTORY                                                                = 40010
	IDC_SHOW_BOOKMARK_MANAGER                                                       = 40011
	IDC_SHOW_DOWNLOADS                                                              = 40012
	IDC_CLEAR_BROWSING_DATA                                                         = 40013
	IDC_IMPORT_SETTINGS                                                             = 40014
	IDC_OPTIONS                                                                     = 40015
	IDC_EDIT_SEARCH_ENGINES                                                         = 40016
	IDC_VIEW_PASSWORDS                                                              = 40017
	IDC_ABOUT                                                                       = 40018
	IDC_HELP_PAGE_VIA_KEYBOARD                                                      = 40019
	IDC_HELP_PAGE_VIA_MENU                                                          = 40020
	IDC_SHOW_APP_MENU                                                               = 40021
	IDC_MANAGE_EXTENSIONS                                                           = 40022
	IDC_DEV_TOOLS_INSPECT                                                           = 40023
	IDC_UPGRADE_DIALOG                                                              = 40024
	IDC_SHOW_HISTORY_CLUSTERS_SIDE_PANEL                                            = 40025
	IDC_PROFILING_ENABLED                                                           = 40028
	IDC_BOOKMARKS_MENU                                                              = 40029
	IDC_SAVED_TAB_GROUPS_MENU                                                       = 40030
	IDC_EXTENSION_ERRORS                                                            = 40031
	IDC_SHOW_SETTINGS_CHANGE_FIRST                                                  = 40033
	IDC_SHOW_SETTINGS_CHANGE_LAST                                                   = 40133
	IDC_SHOW_AVATAR_MENU                                                            = 40134
	IDC_EXTENSION_INSTALL_ERROR_FIRST                                               = 40135
	IDC_EXTENSION_INSTALL_ERROR_LAST                                                = 40235
	IDC_TOGGLE_REQUEST_TABLET_SITE                                                  = 40236
	IDC_DEV_TOOLS_TOGGLE                                                            = 40237
	IDC_RECENT_TABS_MENU                                                            = 40239
	IDC_RECENT_TABS_NO_DEVICE_TABS                                                  = 40240
	IDC_SHOW_SETTINGS_RESET_BUBBLE                                                  = 40241
	IDC_DISTILL_PAGE                                                                = 40243
	IDC_HELP_MENU                                                                   = 40244
	IDC_SHOW_SRT_BUBBLE                                                             = 40246
	IDC_ELEVATED_RECOVERY_DIALOG                                                    = 40247
	IDC_TAKE_SCREENSHOT                                                             = 40248
	IDC_MORE_TOOLS_MENU                                                             = 40249
	IDC_TOGGLE_FULLSCREEN_TOOLBAR                                                   = 40250
	IDC_CUSTOMIZE_TOUCH_BAR                                                         = 40251
	IDC_SHOW_BETA_FORUM                                                             = 40252
	IDC_TOGGLE_JAVASCRIPT_APPLE_EVENTS                                              = 40253
	IDC_INSTALL_PWA                                                                 = 40254
	IDC_SHOW_MANAGEMENT_PAGE                                                        = 40255
	IDC_PASTE_AND_GO                                                                = 40256
	IDC_SHOW_SAVE_LOCAL_CARD_SIGN_IN_PROMO_IF_APPLICABLE                            = 40257
	IDC_CLOSE_SIGN_IN_PROMO                                                         = 40258
	IDC_SHOW_FULL_URLS                                                              = 40259
	IDC_CARET_BROWSING_TOGGLE                                                       = 40260
	IDC_CHROME_TIPS                                                                 = 40263
	IDC_CHROME_WHATS_NEW                                                            = 40264
	IDC_PERFORMANCE                                                                 = 40266
	IDC_EXTENSIONS_SUBMENU                                                          = 40267
	IDC_EXTENSIONS_SUBMENU_MANAGE_EXTENSIONS                                        = 40268
	IDC_EXTENSIONS_SUBMENU_VISIT_CHROME_WEB_STORE                                   = 40269
	IDC_READING_LIST_MENU                                                           = 40270
	IDC_READING_LIST_MENU_ADD_TAB                                                   = 40271
	IDC_READING_LIST_MENU_SHOW_UI                                                   = 40272
	IDC_SHOW_READING_MODE_SIDE_PANEL                                                = 40273
	IDC_SHOW_BOOKMARK_SIDE_PANEL                                                    = 40274
	IDC_SHOW_CHROME_LABS                                                            = 40276
	IDC_RECENT_TABS_LOGIN_FOR_DEVICE_TABS                                           = 40277
	IDC_OPEN_RECENT_TAB                                                             = 40278
	IDC_OPEN_SAFETY_HUB                                                             = 40279
	IDC_SAFETY_HUB_SHOW_PASSWORD_CHECKUP                                            = 40280
	IDC_SAFETY_HUB_MANAGE_EXTENSIONS                                                = 40281
	IDC_SHOW_GOOGLE_LENS_SHORTCUT                                                   = 40282
	IDC_SHOW_CUSTOMIZE_CHROME_SIDE_PANEL                                            = 40283
	IDC_SHOW_CUSTOMIZE_CHROME_TOOLBAR                                               = 40284
	IDC_TASK_MANAGER_APP_MENU                                                       = 40285
	IDC_TASK_MANAGER_SHORTCUT                                                       = 40286
	IDC_TASK_MANAGER_CONTEXT_MENU                                                   = 40287
	IDC_TASK_MANAGER_MAIN_MENU                                                      = 40288
	IDC_SHOW_HISTORY_SIDE_PANEL                                                     = 40293
	IDC_OPEN_GLIC                                                                   = 40294
	IDC_FIND_EXTENSIONS                                                             = 40295
	IDC_SHOW_SEARCH_TOOLS                                                           = 40296
	IDC_SHOW_COMMENTS_SIDE_PANEL                                                    = 40297
	IDC_RECENT_TABS_SEE_DEVICE_TABS                                                 = 40298
	IDC_SHOW_AI_MODE_OMNIBOX_BUTTON                                                 = 40299
	IDC_CONTENT_CONTEXT_INSPECTELEMENT_WITH_GEMINI                                  = 40300
	IDC_CONTENT_CONTEXT_INSPECTELEMENT_WITH_DEVTOOLS                                = 40301
	IDC_REPORT_UNSAFE_SITE                                                          = 40302
	IDC_SPELLCHECK_SUGGESTION_0                                                     = 41000
	IDC_SPELLCHECK_SUGGESTION_1                                                     = 41001
	IDC_SPELLCHECK_SUGGESTION_2                                                     = 41002
	IDC_SPELLCHECK_SUGGESTION_3                                                     = 41003
	IDC_SPELLCHECK_SUGGESTION_4                                                     = 41004
	IDC_SPELLCHECK_MENU                                                             = 41005
	IDC_SPELLCHECK_LANGUAGES_FIRST                                                  = 41006
	IDC_SPELLCHECK_LANGUAGES_LAST                                                   = 41106
	IDC_CHECK_SPELLING_WHILE_TYPING                                                 = 41107
	IDC_SPELLPANEL_TOGGLE                                                           = 41109
	IDC_SPELLCHECK_ADD_TO_DICTIONARY                                                = 41110
	IDC_SPELLCHECK_MULTI_LINGUAL                                                    = 41111
	IDC_WRITING_DIRECTION_MENU                                                      = 41120
	IDC_WRITING_DIRECTION_DEFAULT                                                   = 41121
	IDC_WRITING_DIRECTION_LTR                                                       = 41122
	IDC_WRITING_DIRECTION_RTL                                                       = 41123
	IDC_TRANSLATE_ORIGINAL_LANGUAGE_BASE                                            = 42100
	IDC_TRANSLATE_TARGET_LANGUAGE_BASE                                              = 42400
	IDC_VIEW_MENU                                                                   = 44000
	IDC_FILE_MENU                                                                   = 44001
	IDC_CHROME_MENU                                                                 = 44002
	IDC_HIDE_APP                                                                    = 44003
	IDC_HISTORY_MENU                                                                = 46000
	IDC_TAB_MENU                                                                    = 46001
	IDC_PROFILE_MAIN_MENU                                                           = 46100
	IDC_INPUT_METHODS_MENU                                                          = 46300
	IDC_CONTENT_CONTEXT_CUSTOM_FIRST                                                = 47000
	IDC_CONTENT_CONTEXT_CUSTOM_LAST                                                 = 48000
	IDC_EXTENSIONS_CONTEXT_CUSTOM_FIRST                                             = 49000
	IDC_EXTENSIONS_CONTEXT_CUSTOM_LAST                                              = 50000
	IDC_CONTENT_CONTEXT_OPENLINKNEWTAB                                              = 50100
	IDC_CONTENT_CONTEXT_OPENLINKNEWWINDOW                                           = 50101
	IDC_CONTENT_CONTEXT_OPENLINKOFFTHERECORD                                        = 50102
	IDC_CONTENT_CONTEXT_SAVELINKAS                                                  = 50103
	IDC_CONTENT_CONTEXT_COPYLINKLOCATION                                            = 50104
	IDC_CONTENT_CONTEXT_COPYEMAILADDRESS                                            = 50105
	IDC_CONTENT_CONTEXT_OPENLINKWITH                                                = 50106
	IDC_CONTENT_CONTEXT_COPYLINKTEXT                                                = 50107
	IDC_CONTENT_CONTEXT_OPENLINKINPROFILE                                           = 50108
	IDC_CONTENT_CONTEXT_OPENLINKBOOKMARKAPP                                         = 50109
	IDC_CONTENT_CONTEXT_OPENLINKPREVIEW                                             = 50110
	IDC_CONTENT_CONTEXT_OPENLINKSPLITVIEW                                           = 50111
	IDC_CONTENT_CONTEXT_SAVEIMAGEAS                                                 = 50120
	IDC_CONTENT_CONTEXT_COPYIMAGELOCATION                                           = 50121
	IDC_CONTENT_CONTEXT_COPYIMAGE                                                   = 50122
	IDC_CONTENT_CONTEXT_OPENIMAGENEWTAB                                             = 50123
	IDC_CONTENT_CONTEXT_SEARCHWEBFORIMAGE                                           = 50124
	IDC_CONTENT_CONTEXT_OPEN_ORIGINAL_IMAGE_NEW_TAB                                 = 50125
	IDC_CONTENT_CONTEXT_LOAD_IMAGE                                                  = 50126
	IDC_CONTENT_CONTEXT_SEARCHLENSFORIMAGE                                          = 50127
	IDC_CONTENT_CONTEXT_GLICSHAREIMAGE                                              = 50128
	IDC_CONTENT_CONTEXT_SAVEVIDEOFRAMEAS                                            = 50130
	IDC_CONTENT_CONTEXT_SAVEAVAS                                                    = 50131
	IDC_CONTENT_CONTEXT_COPYAVLOCATION                                              = 50132
	IDC_CONTENT_CONTEXT_COPYVIDEOFRAME                                              = 50133
	IDC_CONTENT_CONTEXT_SEARCHLENSFORVIDEOFRAME                                     = 50134
	IDC_CONTENT_CONTEXT_SEARCHWEBFORVIDEOFRAME                                      = 50135
	IDC_CONTENT_CONTEXT_OPENAVNEWTAB                                                = 50136
	IDC_CONTENT_CONTEXT_PICTUREINPICTURE                                            = 50137
	IDC_CONTENT_CONTEXT_LOOP                                                        = 50140
	IDC_CONTENT_CONTEXT_CONTROLS                                                    = 50141
	IDC_CONTENT_CONTEXT_ROTATECW                                                    = 50142
	IDC_CONTENT_CONTEXT_ROTATECCW                                                   = 50143
	IDC_CONTENT_CONTEXT_COPY                                                        = 50150
	IDC_CONTENT_CONTEXT_CUT                                                         = 50151
	IDC_CONTENT_CONTEXT_PASTE                                                       = 50152
	IDC_CONTENT_CONTEXT_DELETE                                                      = 50153
	IDC_CONTENT_CONTEXT_UNDO                                                        = 50154
	IDC_CONTENT_CONTEXT_REDO                                                        = 50155
	IDC_CONTENT_CONTEXT_SELECTALL                                                   = 50156
	IDC_CONTENT_CONTEXT_PASTE_AND_MATCH_STYLE                                       = 50157
	IDC_CONTENT_CONTEXT_COPYLINKTOTEXT                                              = 50158
	IDC_CONTENT_CONTEXT_RESHARELINKTOTEXT                                           = 50159
	IDC_CONTENT_CONTEXT_REMOVELINKTOTEXT                                            = 50160
	IDC_CONTENT_CONTEXT_TRANSLATE                                                   = 50161
	IDC_CONTENT_CONTEXT_INSPECTELEMENT                                              = 50162
	IDC_CONTENT_CONTEXT_VIEWPAGEINFO                                                = 50163
	IDC_CONTENT_CONTEXT_LANGUAGE_SETTINGS                                           = 50164
	IDC_CONTENT_CONTEXT_LOOK_UP                                                     = 50165
	IDC_CONTENT_CONTEXT_NO_SPELLING_SUGGESTIONS                                     = 50166
	IDC_CONTENT_CONTEXT_SPELLING_SUGGESTION                                         = 50167
	IDC_CONTENT_CONTEXT_SPELLING_TOGGLE                                             = 50168
	IDC_CONTENT_CONTEXT_OPEN_IN_READING_MODE                                        = 50169
	IDC_CONTENT_CONTEXT_SAVEPLUGINAS                                                = 50170
	IDC_CONTENT_CONTEXT_INSPECTBACKGROUNDPAGE                                       = 50171
	IDC_CONTENT_CONTEXT_RELOAD_PACKAGED_APP                                         = 50172
	IDC_CONTENT_CONTEXT_RESTART_PACKAGED_APP                                        = 50173
	IDC_CONTENT_CONTEXT_LENS_REGION_SEARCH                                          = 50174
	IDC_CONTENT_CONTEXT_WEB_REGION_SEARCH                                           = 50175
	IDC_CONTENT_CONTEXT_GENERATEPASSWORD                                            = 50176
	IDC_CONTENT_CONTEXT_EXIT_FULLSCREEN                                             = 50177
	IDC_CONTENT_CONTEXT_SHOWALLSAVEDPASSWORDS                                       = 50178
	IDC_CONTENT_CONTEXT_PARTIAL_TRANSLATE                                           = 50179
	IDC_CONTENT_CONTEXT_RELOADFRAME                                                 = 50180
	IDC_CONTENT_CONTEXT_VIEWFRAMESOURCE                                             = 50181
	IDC_CONTENT_CONTEXT_VIEWFRAMEINFO                                               = 50182
	IDC_CONTENT_CONTEXT_ADD_A_NOTE                                                  = 50185
	IDC_CONTENT_CONTEXT_GOTOURL                                                     = 50190
	IDC_CONTENT_CONTEXT_SEARCHWEBFOR                                                = 50191
	IDC_CONTENT_CONTEXT_SEARCHWEBFORNEWTAB                                          = 50192
	IDC_CONTENT_CONTEXT_LENS_OVERLAY                                                = 50193
	IDC_CONTENT_CONTEXT_USE_PASSKEY_FROM_ANOTHER_DEVICE                             = 50194
	IDC_CONTENT_CONTEXT_OPEN_WITH1                                                  = 50200
	IDC_CONTENT_CONTEXT_OPEN_WITH2                                                  = 50201
	IDC_CONTENT_CONTEXT_OPEN_WITH3                                                  = 50202
	IDC_CONTENT_CONTEXT_OPEN_WITH4                                                  = 50203
	IDC_CONTENT_CONTEXT_OPEN_WITH5                                                  = 50204
	IDC_CONTENT_CONTEXT_OPEN_WITH6                                                  = 50205
	IDC_CONTENT_CONTEXT_OPEN_WITH7                                                  = 50206
	IDC_CONTENT_CONTEXT_OPEN_WITH8                                                  = 50207
	IDC_CONTENT_CONTEXT_OPEN_WITH9                                                  = 50208
	IDC_CONTENT_CONTEXT_OPEN_WITH10                                                 = 50209
	IDC_CONTENT_CONTEXT_OPEN_WITH11                                                 = 50210
	IDC_CONTENT_CONTEXT_OPEN_WITH12                                                 = 50211
	IDC_CONTENT_CONTEXT_OPEN_WITH13                                                 = 50212
	IDC_CONTENT_CONTEXT_OPEN_WITH14                                                 = 50213
	IDC_CONTENT_CONTEXT_EMOJI                                                       = 50220
	IDC_CONTEXT_COMPOSE                                                             = 50230
	IDC_CONTENT_CONTEXT_CLOSE_GLIC                                                  = 50231
	IDC_CONTENT_CONTEXT_RELOAD_GLIC                                                 = 50232
	IDC_CONTENT_CONTEXT_ARCHIVE_GLIC                                                = 50233
	IDC_BOOKMARK_BAR_OPEN_ALL                                                       = 51000
	IDC_BOOKMARK_BAR_OPEN_ALL_NEW_WINDOW                                            = 51001
	IDC_BOOKMARK_BAR_OPEN_ALL_INCOGNITO                                             = 51002
	IDC_BOOKMARK_BAR_OPEN_INCOGNITO                                                 = 51003
	IDC_BOOKMARK_BAR_OPEN_ALL_NEW_TAB_GROUP                                         = 51004
	IDC_BOOKMARK_BAR_RENAME_FOLDER                                                  = 51005
	IDC_BOOKMARK_BAR_EDIT                                                           = 51006
	IDC_BOOKMARK_BAR_REMOVE                                                         = 51007
	IDC_BOOKMARK_BAR_UNDO                                                           = 51008
	IDC_BOOKMARK_BAR_REDO                                                           = 51009
	IDC_BOOKMARK_BAR_ADD_NEW_BOOKMARK                                               = 51010
	IDC_BOOKMARK_BAR_NEW_FOLDER                                                     = 51011
	IDC_BOOKMARK_MANAGER                                                            = 51012
	IDC_BOOKMARK_BAR_ALWAYS_SHOW                                                    = 51013
	IDC_BOOKMARK_BAR_SHOW_APPS_SHORTCUT                                             = 51014
	IDC_BOOKMARK_BAR_SHOW_READING_LIST                                              = 51015
	IDC_BOOKMARK_BAR_SHOW_MANAGED_BOOKMARKS                                         = 51016
	IDC_BOOKMARK_BAR_TRACK_PRICE_FOR_SHOPPING_BOOKMARK                              = 51017
	IDC_BOOKMARK_BAR_UNTRACK_PRICE_FOR_SHOPPING_BOOKMARK                            = 51018
	IDC_BOOKMARK_BAR_ADD_TO_BOOKMARKS_BAR                                           = 51019
	IDC_BOOKMARK_BAR_REMOVE_FROM_BOOKMARKS_BAR                                      = 51020
	IDC_BOOKMARK_BAR_TOGGLE_SHOW_TAB_GROUPS                                         = 51021
	IDC_BOOKMARK_BAR_MOVE                                                           = 51022
	IDC_BOOKMARK_BAR_OPEN_SPLIT_VIEW                                                = 51023
	IDC_CONTENT_CONTEXT_SHARING_CLICK_TO_CALL_SINGLE_DEVICE                         = 51030
	IDC_CONTENT_CONTEXT_SHARING_CLICK_TO_CALL_MULTIPLE_DEVICES                      = 51031
	IDC_CONTENT_CONTEXT_SHARING_SHARED_CLIPBOARD_SINGLE_DEVICE                      = 51032
	IDC_CONTENT_CONTEXT_SHARING_SHARED_CLIPBOARD_MULTIPLE_DEVICES                   = 51033
	IDC_CONTENT_CONTEXT_GENERATE_QR_CODE                                            = 51034
	IDC_CONTENT_CONTEXT_SHARING_SUBMENU                                             = 51035
	IDC_CONTENT_PASTE_FROM_CLIPBOARD                                                = 51037
	IDC_STATUS_TRAY_KEEP_CHROME_RUNNING_IN_BACKGROUND                               = 51100
	IDC_STATUS_TRAY_KEEP_CHROME_RUNNING_IN_BACKGROUND_SETTING                       = 51101
	IDC_MEDIA_ROUTER_ABOUT                                                          = 51200
	IDC_MEDIA_ROUTER_HELP                                                           = 51201
	IDC_MEDIA_ROUTER_LEARN_MORE                                                     = 51202
	IDC_MEDIA_ROUTER_ALWAYS_SHOW_TOOLBAR_ACTION                                     = 51204
	IDC_MEDIA_ROUTER_SHOWN_BY_POLICY                                                = 51206
	IDC_MEDIA_ROUTER_SHOW_IN_TOOLBAR                                                = 51207
	IDC_MEDIA_ROUTER_TOGGLE_MEDIA_REMOTING                                          = 51208
	IDC_MEDIA_TOOLBAR_CONTEXT_REPORT_CAST_ISSUE                                     = 51209
	IDC_MEDIA_TOOLBAR_CONTEXT_SHOW_OTHER_SESSIONS                                   = 51210
	IDC_UPDATE_SIDE_PANEL_PIN_STATE                                                 = 51211
	IDC_MEDIA_STREAM_DEVICE_STATUS_TRAY                                             = 51300
	IDC_MEDIA_CONTEXT_MEDIA_STREAM_CAPTURE_LIST_FIRST                               = 51301
	IDC_MEDIA_CONTEXT_MEDIA_STREAM_CAPTURE_LIST_LAST                                = 51399
	IDC_MEDIA_STREAM_DEVICE_ALWAYS_ALLOW                                            = 51400
	IDC_CONTENT_CONTEXT_PROTOCOL_HANDLER_FIRST                                      = 52000
	IDC_CONTENT_CONTEXT_PROTOCOL_HANDLER_LAST                                       = 52199
	IDC_CONTENT_CONTEXT_PROTOCOL_HANDLER_SETTINGS                                   = 52200
	IDC_OPEN_LINK_IN_PROFILE_FIRST                                                  = 52300
	IDC_OPEN_LINK_IN_PROFILE_LAST                                                   = 52399
	IDC_CONTENT_CONTEXT_START_SMART_SELECTION_ACTION1                               = 52400
	IDC_CONTENT_CONTEXT_START_SMART_SELECTION_ACTION2                               = 52401
	IDC_CONTENT_CONTEXT_START_SMART_SELECTION_ACTION3                               = 52402
	IDC_CONTENT_CONTEXT_START_SMART_SELECTION_ACTION4                               = 52403
	IDC_CONTENT_CONTEXT_START_SMART_SELECTION_ACTION5                               = 52404
	IDC_CONTENT_CONTEXT_ACCESSIBILITY_LABELS_TOGGLE                                 = 52410
	IDC_CONTENT_CONTEXT_ACCESSIBILITY_LABELS                                        = 52411
	IDC_CONTENT_CONTEXT_ACCESSIBILITY_LABELS_TOGGLE_ONCE                            = 52412
	IDC_CONTENT_CONTEXT_QUICK_ANSWERS_INLINE_ANSWER                                 = 52413
	IDC_CONTENT_CONTEXT_QUICK_ANSWERS_INLINE_QUERY                                  = 52414
	IDC_TAB_SEARCH                                                                  = 52500
	IDC_TAB_SEARCH_CLOSE                                                            = 52501
	IDC_DEBUG_TOGGLE_TABLET_MODE                                                    = 52510
	IDC_DEBUG_PRINT_VIEW_TREE                                                       = 52511
	IDC_DEBUG_PRINT_VIEW_TREE_DETAILS                                               = 52512
	IDC_CONTENT_CONTEXT_AUTOFILL_FEEDBACK                                           = 52990
	IDC_CONTENT_CONTEXT_AUTOFILL_FALLBACK_PLUS_ADDRESS                              = 52994
	IDC_CONTENT_CONTEXT_AUTOFILL_FALLBACK_PASSWORDS_SELECT_PASSWORD                 = 52998
	IDC_CONTENT_CONTEXT_AUTOFILL_FALLBACK_PASSWORDS_IMPORT_PASSWORDS                = 52999
	IDC_CONTENT_CONTEXT_AUTOFILL_FALLBACK_PASSWORDS_SUGGEST_PASSWORD                = 53000
	IDC_CONTENT_CONTEXT_AUTOFILL_FALLBACK_PASSWORDS_USE_PASSKEY_FROM_ANOTHER_DEVICE = 53002
	IDC_CONTENT_CONTEXT_AUTOFILL_FALLBACK_AT_MEMORY                                 = 53003
	IDC_LIVE_CAPTION                                                                = 53251
	IDC_DEVICE_SYSTEM_TRAY_ICON_FIRST                                               = 53260
	IDC_DEVICE_SYSTEM_TRAY_ICON_LAST                                                = 53299
	IDC_SET_BROWSER_AS_DEFAULT                                                      = 53300
	IDC_GLIC_STATUS_ICON_MENU_SHOW                                                  = 53310
	IDC_GLIC_STATUS_ICON_MENU_CUSTOMIZE_KEYBOARD_SHORTCUT                           = 53311
	IDC_GLIC_STATUS_ICON_MENU_REMOVE_ICON                                           = 53312
	IDC_GLIC_STATUS_ICON_MENU_SETTINGS                                              = 53313
	IDC_GLIC_STATUS_ICON_MENU_EXIT                                                  = 53314
	IDC_GLIC_STATUS_ICON_MENU_CLOSE                                                 = 53315
	IDC_GLIC_STATUS_ICON_MENU_TOGGLE                                                = 53316
	IDC_GLIC_TOGGLE_PIN                                                             = 53320
	IDC_TAB_SEARCH_TOGGLE_PIN                                                       = 53321
	IDC_PROJECTS_PANEL_TOGGLE_PIN                                                   = 53322
	IDC_EVERYTHING_MENU_TOGGLE_PIN                                                  = 53323
	IDC_SHOW_CONTEXTUAL_TASKS_SIDE_PANEL                                            = 54000
	IDC_OMNIBOX_CONTEXT_ADD_IMAGE                                                   = 54010
	IDC_OMNIBOX_CONTEXT_ADD_FILE                                                    = 54011
	IDC_OMNIBOX_CONTEXT_CREATE_IMAGES                                               = 54012
	IDC_OMNIBOX_CONTEXT_DEEP_RESEARCH                                               = 54013
	//  Supported certificate status code values. See net\cert\cert_status_flags.h
	//  for more information. CERT_STATUS_NONE is new in CEF because we use an
	//  enum while cert_status_flags.h uses a typedef and static const variables.
	//  TCefCertStatus values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cert_status_t)</see>
	CERT_STATUS_NONE                       = 0
	CERT_STATUS_COMMON_NAME_INVALID        = 1 << 0
	CERT_STATUS_DATE_INVALID               = 1 << 1
	CERT_STATUS_AUTHORITY_INVALID          = 1 << 2
	CERT_STATUS_NO_REVOCATION_MECHANISM    = 1 << 4
	CERT_STATUS_UNABLE_TO_CHECK_REVOCATION = 1 << 5
	CERT_STATUS_REVOKED                    = 1 << 6
	CERT_STATUS_INVALID                    = 1 << 7
	CERT_STATUS_WEAK_SIGNATURE_ALGORITHM   = 1 << 8
	CERT_STATUS_NON_UNIQUE_NAME            = 1 << 10
	CERT_STATUS_WEAK_KEY                   = 1 << 11
	CERT_STATUS_PINNED_KEY_MISSING         = 1 << 13
	CERT_STATUS_NAME_CONSTRAINT_VIOLATION  = 1 << 14
	CERT_STATUS_VALIDITY_TOO_LONG          = 1 << 15
	CERT_STATUS_IS_EV                      = 1 << 16
	CERT_STATUS_REV_CHECKING_ENABLED       = 1 << 17
	CERT_STATUS_SHA1_SIGNATURE_PRESENT     = 1 << 19
	CERT_STATUS_CT_COMPLIANCE_FAILED       = 1 << 20
	CERT_STATUS_FIRST_ERROR                = CERT_STATUS_COMMON_NAME_INVALID
	CERT_STATUS_LAST_ERROR                 = CERT_STATUS_VALIDITY_TOO_LONG
	//  <summary>Writeable, Enumerable, Configurable</summary>
	//  TCefV8PropertyAttributes value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_v8_propertyattribute_t)</see>
	V8_PROPERTY_ATTRIBUTE_NONE = 0
	//  <summary>Not writeable</summary>
	//  TCefV8PropertyAttributes value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_v8_propertyattribute_t)</see>
	V8_PROPERTY_ATTRIBUTE_READONLY = 1 << 0
	//  <summary>Not enumerable</summary>
	//  TCefV8PropertyAttributes value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_v8_propertyattribute_t)</see>
	V8_PROPERTY_ATTRIBUTE_DONTENUM = 1 << 1
	//  <summary>Not configurable</summary>
	//  TCefV8PropertyAttributes value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_v8_propertyattribute_t)</see>
	V8_PROPERTY_ATTRIBUTE_DONTDELETE = 1 << 2
	//  Source is a link click or the JavaScript window.open function. This is
	//  also the default value for requests like sub-resource loads that are not
	//  navigations.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_LINK = 0
	//  Source is some other "explicit" navigation. This is the default value for
	//  navigations where the actual type is unknown. See also
	//  TT_DIRECT_LOAD_FLAG.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_EXPLICIT = 1
	//  User got to this page through a suggestion in the UI (for example, via the
	//  destinations page). Chrome style only.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_AUTO_BOOKMARK = 2
	//  Source is a subframe navigation. This is any content that is automatically
	//  loaded in a non-toplevel frame. For example, if a page consists of several
	//  frames containing ads, those ad URLs will have this transition type.
	//  The user may not even realize the content in these pages is a separate
	//  frame, so may not care about the URL.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_AUTO_SUBFRAME = 3
	//  Source is a subframe navigation explicitly requested by the user that will
	//  generate new navigation entries in the back/forward list. These are
	//  probably more important than frames that were automatically loaded in
	//  the background because the user probably cares about the fact that this
	//  link was loaded.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_MANUAL_SUBFRAME = 4
	//  User got to this page by typing in the URL bar and selecting an entry
	//  that did not look like a URL. For example, a match might have the URL
	//  of a Google search result page, but appear like "Search Google for ...".
	//  These are not quite the same as EXPLICIT navigations because the user
	//  didn't type or see the destination URL. Chrome style only.
	//  See also TT_KEYWORD.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_GENERATED = 5
	//  This is a toplevel navigation. This is any content that is automatically
	//  loaded in a toplevel frame. For example, opening a tab to show the ASH
	//  screen saver, opening the devtools window, opening the NTP after the safe
	//  browsing warning, opening web-based dialog boxes are examples of
	//  AUTO_TOPLEVEL navigations. Chrome style only.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_AUTO_TOPLEVEL = 6
	//  Source is a form submission by the user. NOTE: In some situations
	//  submitting a form does not result in this transition type. This can happen
	//  if the form uses a script to submit the contents.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_FORM_SUBMIT = 7
	//  Source is a "reload" of the page via the Reload function or by re-visiting
	//  the same URL. NOTE: This is distinct from the concept of whether a
	//  particular load uses "reload semantics" (i.e. bypasses cached data).
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_RELOAD = 8
	//  The url was generated from a replaceable keyword other than the default
	//  search provider. If the user types a keyword (which also applies to
	//  tab-to-search) in the omnibox this qualifier is applied to the transition
	//  type of the generated url. TemplateURLModel then may generate an
	//  additional visit with a transition type of TT_KEYWORD_GENERATED against
	//  the url 'http://' + keyword. For example, if you do a tab-to-search
	//  against wikipedia the generated url has a transition qualifer of
	//  TT_KEYWORD, and TemplateURLModel generates a visit for 'wikipedia.org'
	//  with a transition type of TT_KEYWORD_GENERATED. Chrome style only.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_KEYWORD = 9
	//  Corresponds to a visit generated for a keyword. See description of
	//  TT_KEYWORD for more details. Chrome style only.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_KEYWORD_GENERATED = 10
	//  Number of TCefTransitionType values.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_NUM_VALUES = TT_KEYWORD_GENERATED + 1
	//  General mask defining the bits used for the source values.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_SOURCE_MASK = 0x000000FF
	//  Attempted to visit a URL but was blocked.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_BLOCKED_FLAG = 0x00800000
	//  Used the Forward or Back function to navigate among browsing history.
	//  Will be ORed to the transition type for the original load.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_FORWARD_BACK_FLAG = 0x01000000
	//  Loaded a URL directly via CreateBrowser, LoadURL or LoadRequest.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_DIRECT_LOAD_FLAG = 0x02000000
	//  User is navigating to the home page. Chrome style only.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_HOME_PAGE_FLAG = 0x04000000
	//  The transition originated from an external application; the exact
	//  definition of this is embedder dependent. Chrome style only.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_FROM_API_FLAG = 0x08000000
	//  The beginning of a navigation chain.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_CHAIN_START_FLAG = 0x10000000
	//  The last transition in a redirect chain.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_CHAIN_END_FLAG = 0x20000000
	//  Redirects caused by JavaScript or a meta refresh tag on the page.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_CLIENT_REDIRECT_FLAG = 0x40000000
	//  Redirects sent from the server by HTTP headers.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_SERVER_REDIRECT_FLAG = 0x80000000
	//  Used to test whether a transition involves a redirect.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_IS_REDIRECT_MASK = 0xC0000000
	//  General mask defining the bits used for the qualifiers.
	//  TCefTransitionType value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_transition_type_t)</see>
	TT_QUALIFIER_MASK = 0xFFFFFF00
	//  Default behavior.
	//  TCefUrlRequestFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see>
	UR_FLAG_NONE = 0
	//  If set the cache will be skipped when handling the request. Setting this
	//  value is equivalent to specifying the "Cache-Control: no-cache" request
	//  header. Setting this value in combination with UR_FLAG_ONLY_FROM_CACHE
	//  will cause the request to fail.
	//  TCefUrlRequestFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see>
	UR_FLAG_SKIP_CACHE = 1 << 0
	//  If set the request will fail if it cannot be served from the cache (or
	//  some equivalent local store). Setting this value is equivalent to
	//  specifying the "Cache-Control: only-if-cached" request header. Setting
	//  this value in combination with UR_FLAG_SKIP_CACHE or UR_FLAG_DISABLE_CACHE
	//  will cause the request to fail.
	//  TCefUrlRequestFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see>
	UR_FLAG_ONLY_FROM_CACHE = 1 << 1
	//  If set the cache will not be used at all. Setting this value is equivalent
	//  to specifying the "Cache-Control: no-store" request header. Setting this
	//  value in combination with UR_FLAG_ONLY_FROM_CACHE will cause the request
	//  to fail.
	//  TCefUrlRequestFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see>
	UR_FLAG_DISABLE_CACHE = 1 << 2
	//  If set user name, password, and cookies may be sent with the request, and
	//  cookies may be saved from the response.
	//  TCefUrlRequestFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see>
	UR_FLAG_ALLOW_STORED_CREDENTIALS = 1 << 3
	//  If set upload progress events will be generated when a request has a body.
	//  TCefUrlRequestFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see>
	UR_FLAG_REPORT_UPLOAD_PROGRESS = 1 << 4
	//  If set the ICefURLRequestClient.OnDownloadData method will not be called.
	//  TCefUrlRequestFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see>
	UR_FLAG_NO_DOWNLOAD_DATA = 1 << 5
	//  If set 5XX redirect errors will be propagated to the observer instead of
	//  automatically re-tried. This currently only applies for requests
	//  originated in the browser process.
	//  TCefUrlRequestFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see>
	UR_FLAG_NO_RETRY_ON_5XX = 1 << 6
	//  If set 3XX responses will cause the fetch to halt immediately rather than
	//  continue through the redirect.
	//  TCefUrlRequestFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlrequest_flags_t)</see>
	UR_FLAG_STOP_ON_REDIRECT = 1 << 7
	//  No options.
	//  TCefSchemeOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scheme_options_t)</see>
	CEF_SCHEME_OPTION_NONE = 0
	//  If CEF_SCHEME_OPTION_STANDARD is set the scheme will be treated as a
	//  standard scheme. Standard schemes are subject to URL canonicalization and
	//  parsing rules as defined in the Common Internet Scheme Syntax RFC 1738
	//  Section 3.1 available at http://www.ietf.org/rfc/rfc1738.txt
	//
	//  In particular, the syntax for standard scheme URLs must be of the form:
	//  <pre>
	//  [scheme]://[username]:[password]@[host]:[port]/[url-path]
	//  </pre> Standard scheme URLs must have a host component that is a fully
	//  qualified domain name as defined in Section 3.5 of RFC 1034 [13] and
	//  Section 2.1 of RFC 1123. These URLs will be canonicalized to
	//  "scheme://host/path" in the simplest case and
	//  "scheme://username:password@host:port/path" in the most explicit case. For
	//  example, "scheme:host/path" and "scheme:///host/path" will both be
	//  canonicalized to "scheme://host/path". The origin of a standard scheme URL
	//  is the combination of scheme, host and port (i.e., "scheme://host:port" in
	//  the most explicit case).
	//
	//  For non-standard scheme URLs only the "scheme:" component is parsed and
	//  canonicalized. The remainder of the URL will be passed to the handler as-
	//  is. For example, "scheme:///some%20text" will remain the same.
	//  Non-standard scheme URLs cannot be used as a target for form submission.
	//  TCefSchemeOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scheme_options_t)</see>
	CEF_SCHEME_OPTION_STANDARD = 1 << 0
	//  If CEF_SCHEME_OPTION_LOCAL is set the scheme will be treated with the same
	//  security rules as those applied to "file" URLs. Normal pages cannot link
	//  to or access local URLs. Also, by default, local URLs can only perform
	//  XMLHttpRequest calls to the same URL (origin + path) that originated the
	//  request. To allow XMLHttpRequest calls from a local URL to other URLs with
	//  the same origin set the CefSettings.file_access_from_file_urls_allowed
	//  value to true (1). To allow XMLHttpRequest calls from a local URL to all
	//  origins set the CefSettings.universal_access_from_file_urls_allowed value
	//  to true (1).
	//  TCefSchemeOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scheme_options_t)</see>
	CEF_SCHEME_OPTION_LOCAL = 1 << 1
	//  If CEF_SCHEME_OPTION_DISPLAY_ISOLATED is set the scheme can only be
	//  displayed from other content hosted with the same scheme. For example,
	//  pages in other origins cannot create iframes or hyperlinks to URLs with
	//  the scheme. For schemes that must be accessible from other schemes don't
	//  set this, set CEF_SCHEME_OPTION_CORS_ENABLED, and use CORS
	//  "Access-Control-Allow-Origin" headers to further restrict access.
	//  TCefSchemeOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scheme_options_t)</see>
	CEF_SCHEME_OPTION_DISPLAY_ISOLATED = 1 << 2
	//  If CEF_SCHEME_OPTION_SECURE is set the scheme will be treated with the
	//  same security rules as those applied to "https" URLs. For example, loading
	//  this scheme from other secure schemes will not trigger mixed content
	//  warnings.
	//  TCefSchemeOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scheme_options_t)</see>
	CEF_SCHEME_OPTION_SECURE = 1 << 3
	//  If CEF_SCHEME_OPTION_CORS_ENABLED is set the scheme can be sent CORS
	//  requests. This value should be set in most cases where
	//  CEF_SCHEME_OPTION_STANDARD is set.
	//  TCefSchemeOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scheme_options_t)</see>
	CEF_SCHEME_OPTION_CORS_ENABLED = 1 << 4
	//  If CEF_SCHEME_OPTION_CSP_BYPASSING is set the scheme can bypass Content-
	//  Security-Policy (CSP) checks. This value should not be set in most cases
	//  where CEF_SCHEME_OPTION_STANDARD is set.
	//  TCefSchemeOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scheme_options_t)</see>
	CEF_SCHEME_OPTION_CSP_BYPASSING = 1 << 5
	//  If CEF_SCHEME_OPTION_FETCH_ENABLED is set the scheme can perform Fetch API
	//  requests.
	//  TCefSchemeOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_scheme_options_t)</see>
	CEF_SCHEME_OPTION_FETCH_ENABLED = 1 << 6
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_UNKNOWN = 0
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_UI = 1 << 0
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_MOUSE = 1 << 1
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_MUTATION = 1 << 2
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_KEYBOARD = 1 << 3
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_TEXT = 1 << 4
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_COMPOSITION = 1 << 5
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_DRAG = 1 << 6
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_CLIPBOARD = 1 << 7
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_MESSAGE = 1 << 8
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_WHEEL = 1 << 9
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_BEFORE_TEXT_INSERTED = 1 << 10
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_OVERFLOW = 1 << 11
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_PAGE_TRANSITION = 1 << 12
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_POPSTATE = 1 << 13
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_PROGRESS = 1 << 14
	//  DOM event category flag.
	//  TCefDomEventCategory value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_dom_event_category_t)</see>
	DOM_EVENT_CATEGORY_XMLHTTPREQUEST_PROGRESS = 1 << 15
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_NONE = 0
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_CAPS_LOCK_ON = 1 << 0
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_SHIFT_DOWN = 1 << 1
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_CONTROL_DOWN = 1 << 2
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_ALT_DOWN = 1 << 3
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_LEFT_MOUSE_BUTTON = 1 << 4
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_MIDDLE_MOUSE_BUTTON = 1 << 5
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_RIGHT_MOUSE_BUTTON = 1 << 6
	//  Supported event bit flag. Mac OS-X command key.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_COMMAND_DOWN = 1 << 7
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_NUM_LOCK_ON = 1 << 8
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_IS_KEY_PAD = 1 << 9
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_IS_LEFT = 1 << 10
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_IS_RIGHT = 1 << 11
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_ALTGR_DOWN = 1 << 12
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_IS_REPEAT = 1 << 13
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_PRECISION_SCROLLING_DELTA = 1 << 14
	//  Supported event bit flag.
	//  TCefEventFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_event_flags_t)</see>
	EVENTFLAG_SCROLL_BY_PAGE = 1 << 15
	//  "Verb" of a drag-and-drop operation as negotiated between the source and destination.
	//  TCefDragOperation value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_drag_operations_mask_t)</see>
	DRAG_OPERATION_NONE = 0
	//  "Verb" of a drag-and-drop operation as negotiated between the source and destination.
	//  TCefDragOperation value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_drag_operations_mask_t)</see>
	DRAG_OPERATION_COPY = 1 << 0
	//  "Verb" of a drag-and-drop operation as negotiated between the source and destination.
	//  TCefDragOperation value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_drag_operations_mask_t)</see>
	DRAG_OPERATION_LINK = 1 << 1
	//  "Verb" of a drag-and-drop operation as negotiated between the source and destination.
	//  TCefDragOperation value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_drag_operations_mask_t)</see>
	DRAG_OPERATION_GENERIC = 1 << 2
	//  "Verb" of a drag-and-drop operation as negotiated between the source and destination.
	//  TCefDragOperation value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_drag_operations_mask_t)</see>
	DRAG_OPERATION_PRIVATE = 1 << 3
	//  "Verb" of a drag-and-drop operation as negotiated between the source and destination.
	//  TCefDragOperation value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_drag_operations_mask_t)</see>
	DRAG_OPERATION_MOVE = 1 << 4
	//  "Verb" of a drag-and-drop operation as negotiated between the source and destination.
	//  TCefDragOperation value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_drag_operations_mask_t)</see>
	DRAG_OPERATION_DELETE = 1 << 5
	//  "Verb" of a drag-and-drop operation as negotiated between the source and destination.
	//  TCefDragOperation value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_drag_operations_mask_t)</see>
	DRAG_OPERATION_EVERY = 0xFFFFFFFF
	//  Requires that the file exists before allowing the user to pick it.
	//  TCefFileDialogMode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_file_dialog_mode_t)</see>
	FILE_DIALOG_OPEN = 0x00000000
	//  Like Open, but allows picking multiple files to open.
	//  TCefFileDialogMode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_file_dialog_mode_t)</see>
	FILE_DIALOG_OPEN_MULTIPLE = 0x00000001
	//  Like Open, but selects a folder to open.
	//  TCefFileDialogMode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_file_dialog_mode_t)</see>
	FILE_DIALOG_OPEN_FOLDER = 0x00000002
	//  Allows picking a nonexistent file, and prompts to overwrite if the file
	//  already exists.
	//  TCefFileDialogMode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_file_dialog_mode_t)</see>
	FILE_DIALOG_SAVE = 0x00000003
	//  Number of TCefFileDialogMode values
	//  TCefFileDialogMode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_file_dialog_mode_t)</see>
	FILE_DIALOG_NUM_VALUES = FILE_DIALOG_SAVE + 1
	//  Don't unescape anything at all.
	//  TCefUriUnescapeRule values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see>
	UU_NONE = 0
	//  Don't unescape anything special, but all normal unescaping will happen.
	//  This is a placeholder and can't be combined with other flags (since it's
	//  just the absence of them). All other unescape rules imply "normal" in
	//  addition to their special meaning. Things like escaped letters, digits,
	//  and most symbols will get unescaped with this mode.
	//  TCefUriUnescapeRule values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see>
	UU_NORMAL = 1 << 0
	//  Convert %20 to spaces. In some places where we're showing URLs, we may
	//  want this. In places where the URL may be copied and pasted out, then
	//  you wouldn't want this since it might not be interpreted in one piece
	//  by other applications.
	//  TCefUriUnescapeRule values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see>
	UU_SPACES = 1 << 1
	//  Unescapes '/' and '\\'. If these characters were unescaped, the resulting
	//  URL won't be the same as the source one. Moreover, they are dangerous to
	//  unescape in strings that will be used as file paths or names. This value
	//  should only be used when slashes don't have special meaning, like data
	//  URLs.
	//  TCefUriUnescapeRule values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see>
	UU_PATH_SEPARATORS = 1 << 2
	//  Unescapes various characters that will change the meaning of URLs,
	//  including '%', '+', '&', '#'. Does not unescape path separators.
	//  If these characters were unescaped, the resulting URL won't be the same
	//  as the source one. This flag is used when generating final output like
	//  filenames for URLs where we won't be interpreting as a URL and want to do
	//  as much unescaping as possible.
	//  TCefUriUnescapeRule values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see>
	UU_URL_SPECIAL_CHARS_EXCEPT_PATH_SEPARATORS = 1 << 3
	//  URL queries use "+" for space. This flag controls that replacement.
	//  TCefUriUnescapeRule values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_uri_unescape_rule_t)</see>
	UU_REPLACE_PLUS_WITH_SPACE = 1 << 4
	//  <summary>Navigation.</summary>
	//  TCefMenuId value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_id_t)</see>
	MENU_ID_BACK           = 100
	MENU_ID_FORWARD        = 101
	MENU_ID_RELOAD         = 102
	MENU_ID_RELOAD_NOCACHE = 103
	MENU_ID_STOPLOAD       = 104
	//  <summary>Editing.</summary>
	//  TCefMenuId value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_id_t)</see>
	MENU_ID_UNDO              = 110
	MENU_ID_REDO              = 111
	MENU_ID_CUT               = 112
	MENU_ID_COPY              = 113
	MENU_ID_PASTE             = 114
	MENU_ID_PASTE_MATCH_STYLE = 115
	MENU_ID_DELETE            = 116
	MENU_ID_SELECT_ALL        = 117
	//  <summary>Miscellaneous.</summary>
	//  TCefMenuId value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_id_t)</see>
	MENU_ID_FIND        = 130
	MENU_ID_PRINT       = 131
	MENU_ID_VIEW_SOURCE = 132
	//  <summary>Spell checking word correction suggestions.</summary>
	//  TCefMenuId value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_id_t)</see>
	MENU_ID_SPELLCHECK_SUGGESTION_0    = 200
	MENU_ID_SPELLCHECK_SUGGESTION_1    = 201
	MENU_ID_SPELLCHECK_SUGGESTION_2    = 202
	MENU_ID_SPELLCHECK_SUGGESTION_3    = 203
	MENU_ID_SPELLCHECK_SUGGESTION_4    = 204
	MENU_ID_SPELLCHECK_SUGGESTION_LAST = 204
	MENU_ID_NO_SPELLING_SUGGESTIONS    = 205
	MENU_ID_ADD_TO_DICTIONARY          = 206
	//  <summary>Custom menu items originating from the renderer process.</summary>
	//  TCefMenuId value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_id_t)</see>
	MENU_ID_CUSTOM_FIRST = 220
	MENU_ID_CUSTOM_LAST  = 250
	//  All user-defined menu IDs should come between MENU_ID_USER_FIRST and
	//  MENU_ID_USER_LAST to avoid overlapping the Chromium and CEF ID ranges
	//  defined in the tools/gritsettings/resource_ids file.
	//  TCefMenuId value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_menu_id_t)</see>
	MENU_ID_USER_FIRST = 26500
	MENU_ID_USER_LAST  = 28500
	//  No node is selected.
	//  TCefContextMenuTypeFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_type_flags_t)</see>
	CM_TYPEFLAG_NONE = 0
	//  The top page is selected.
	//  TCefContextMenuTypeFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_type_flags_t)</see>
	CM_TYPEFLAG_PAGE = 1 << 0
	//  A subframe page is selected.
	//  TCefContextMenuTypeFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_type_flags_t)</see>
	CM_TYPEFLAG_FRAME = 1 << 1
	//  A link is selected.
	//  TCefContextMenuTypeFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_type_flags_t)</see>
	CM_TYPEFLAG_LINK = 1 << 2
	//  A media node is selected.
	//  TCefContextMenuTypeFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_type_flags_t)</see>
	CM_TYPEFLAG_MEDIA = 1 << 3
	//  There is a textual or mixed selection that is selected.
	//  TCefContextMenuTypeFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_type_flags_t)</see>
	CM_TYPEFLAG_SELECTION = 1 << 4
	//  An editable element is selected.
	//  TCefContextMenuTypeFlags value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_type_flags_t)</see>
	CM_TYPEFLAG_EDITABLE = 1 << 5
	//  Supported context menu media state bit flags. These constants match their
	//  equivalents in Chromium's ContextMenuData::MediaFlags and should not be
	//  renumbered.
	//  <TCefContextMenuMediaStateFlags values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_media_state_flags_t)</see>
	CM_MEDIAFLAG_NONE                   = 0
	CM_MEDIAFLAG_IN_ERROR               = 1 << 0
	CM_MEDIAFLAG_PAUSED                 = 1 << 1
	CM_MEDIAFLAG_MUTED                  = 1 << 2
	CM_MEDIAFLAG_LOOP                   = 1 << 3
	CM_MEDIAFLAG_CAN_SAVE               = 1 << 4
	CM_MEDIAFLAG_HAS_AUDIO              = 1 << 5
	CM_MEDIAFLAG_CAN_TOGGLE_CONTROLS    = 1 << 6
	CM_MEDIAFLAG_CONTROLS               = 1 << 7
	CM_MEDIAFLAG_CAN_PRINT              = 1 << 8
	CM_MEDIAFLAG_CAN_ROTATE             = 1 << 9
	CM_MEDIAFLAG_CAN_PICTURE_IN_PICTURE = 1 << 10
	CM_MEDIAFLAG_PICTURE_IN_PICTURE     = 1 << 11
	CM_MEDIAFLAG_CAN_LOOP               = 1 << 12
	//  Supported context menu edit state bit flags. These constants match their
	//  equivalents in Chromium's ContextMenuDataEditFlags and should not be
	//  renumbered.
	//  TCefContextMenuEditStateFlags values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_context_menu_edit_state_flags_t)</see>
	CM_EDITFLAG_NONE            = 0
	CM_EDITFLAG_CAN_UNDO        = 1 << 0
	CM_EDITFLAG_CAN_REDO        = 1 << 1
	CM_EDITFLAG_CAN_CUT         = 1 << 2
	CM_EDITFLAG_CAN_COPY        = 1 << 3
	CM_EDITFLAG_CAN_PASTE       = 1 << 4
	CM_EDITFLAG_CAN_DELETE      = 1 << 5
	CM_EDITFLAG_CAN_SELECT_ALL  = 1 << 6
	CM_EDITFLAG_CAN_TRANSLATE   = 1 << 7
	CM_EDITFLAG_CAN_EDIT_RICHLY = 1 << 8
	//  Supported SSL version values.
	//  TCefSSLVersion values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_ssl_version_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:net/ssl/ssl_connection_status_flags.h">See net/ssl/ssl_connection_status_flags.h for more information.</see>
	SSL_CONNECTION_VERSION_UNKNOWN    = 0
	SSL_CONNECTION_VERSION_SSL2       = 1
	SSL_CONNECTION_VERSION_SSL3       = 2
	SSL_CONNECTION_VERSION_TLS1       = 3
	SSL_CONNECTION_VERSION_TLS1_1     = 4
	SSL_CONNECTION_VERSION_TLS1_2     = 5
	SSL_CONNECTION_VERSION_TLS1_3     = 6
	SSL_CONNECTION_VERSION_QUIC       = 7
	SSL_CONNECTION_VERSION_NUM_VALUES = SSL_CONNECTION_VERSION_QUIC + 1
	//  Supported SSL content status flags. See content/public/common/ssl_status.h
	//  for more information.
	//  TCefSSLContentStatus values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_ssl_content_status_t)</see>
	SSL_CONTENT_NORMAL_CONTENT             = 0
	SSL_CONTENT_DISPLAYED_INSECURE_CONTENT = 1 << 0
	SSL_CONTENT_RAN_INSECURE_CONTENT       = 1 << 1
	//  Default behavior.
	//  TCefJsonWriterOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_writer_options_t)</see>
	JSON_WRITER_DEFAULT = 0
	//  This option instructs the writer that if a Binary value is encountered,
	//  the value (and key if within a dictionary) will be omitted from the
	//  output, and success will be returned. Otherwise, if a binary value is
	//  encountered, failure will be returned.
	//  TCefJsonWriterOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_writer_options_t)</see>
	JSON_WRITER_OMIT_BINARY_VALUES = 1 << 0
	//  This option instructs the writer to write doubles that have no fractional
	//  part as a normal integer (i.e., without using exponential notation
	//  or appending a '.0') as long as the value is within the range of a
	//  64-bit int.
	//  TCefJsonWriterOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_writer_options_t)</see>
	JSON_WRITER_OMIT_DOUBLE_TYPE_PRESERVATION = 1 << 1
	//  Return a slightly nicer formatted json string (pads with whitespace to
	//  help with readability).
	//  TCefJsonWriterOptions value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_json_writer_options_t)</see>
	JSON_WRITER_PRETTY_PRINT = 1 << 2
	//  Default logging (currently INFO logging).
	//  TCefLogSeverity value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_severity_t)</see>
	LOGSEVERITY_DEFAULT = 0
	//  Verbose logging.
	//  TCefLogSeverity value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_severity_t)</see>
	LOGSEVERITY_VERBOSE = 1
	//  DEBUG logging.
	//  TCefLogSeverity value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_severity_t)</see>
	LOGSEVERITY_DEBUG = LOGSEVERITY_VERBOSE
	//  INFO logging.
	//  TCefLogSeverity value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_severity_t)</see>
	LOGSEVERITY_INFO = 2
	//  WARNING logging.
	//  TCefLogSeverity value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_severity_t)</see>
	LOGSEVERITY_WARNING = 3
	//  ERROR logging.
	//  TCefLogSeverity value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_severity_t)</see>
	LOGSEVERITY_ERROR = 4
	//  FATAL logging.
	//  TCefLogSeverity value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_severity_t)</see>
	LOGSEVERITY_FATAL = 5
	//  Disable logging to file for all messages, and to stderr for messages with
	//  severity less than FATAL.
	//  TCefLogSeverity value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_severity_t)</see>
	LOGSEVERITY_DISABLE = 99
	//  Print job duplex mode values.
	//  TCefDuplexMode values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_duplex_mode_t)</see>
	DUPLEX_MODE_UNKNOWN    = -1
	DUPLEX_MODE_SIMPLEX    = 0
	DUPLEX_MODE_LONG_EDGE  = 1
	DUPLEX_MODE_SHORT_EDGE = 2
	DUPLEX_MODE_NUM_VALUES = DUPLEX_MODE_SHORT_EDGE + 1
	//  Result codes for ICefMediaRouter.CreateRoute. Should be kept in sync with
	//  Chromium's media_router::mojom::RouteRequestResultCode type.
	//  TCefMediaRouterCreateResult values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_route_create_result_t)</see>
	CEF_MRCR_UNKNOWN_ERROR                      = 0
	CEF_MRCR_OK                                 = 1
	CEF_MRCR_TIMED_OUT                          = 2
	CEF_MRCR_ROUTE_NOT_FOUND                    = 3
	CEF_MRCR_SINK_NOT_FOUND                     = 4
	CEF_MRCR_INVALID_ORIGIN                     = 5
	CEF_MRCR_OFF_THE_RECORD_MISMATCH_DEPRECATED = 6
	CEF_MRCR_NO_SUPPORTED_PROVIDER              = 7
	CEF_MRCR_CANCELLED                          = 8
	CEF_MRCR_ROUTE_ALREADY_EXISTS               = 9
	CEF_MRCR_DESKTOP_PICKER_FAILED              = 10
	CEF_MRCR_ROUTE_ALREADY_TERMINATED           = 11
	CEF_MRCR_REDUNDANT_REQUEST                  = 12
	CEF_MRCR_USER_NOT_ALLOWED                   = 13
	CEF_MRCR_NOTIFICATION_DISABLED              = 14
	CEF_MRCR_NUM_VALUES                         = CEF_MRCR_NOTIFICATION_DISABLED + 1
	//  Connection state for a MediaRoute object. Should be kept in sync with
	//  Chromium's blink::mojom::PresentationConnectionState type.
	//  TCefMediaRouteConnectionState values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_route_connection_state_t)</see>
	CEF_MRCS_UNKNOWN    = -1
	CEF_MRCS_CONNECTING = 0
	CEF_MRCS_CONNECTED  = 1
	CEF_MRCS_CLOSED     = 2
	CEF_MRCS_TERMINATED = 3
	CEF_MRCS_NUM_VALUES = CEF_MRCS_TERMINATED + 1
	//  Cookie priority values.
	//  TCefCookiePriority values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cookie_priority_t)</see>
	CEF_COOKIE_PRIORITY_LOW    = -1
	CEF_COOKIE_PRIORITY_MEDIUM = 0
	CEF_COOKIE_PRIORITY_HIGH   = 1
	//  Represents commands available to TextField. Should be kept in sync with
	//  Chromium's views::TextField::MenuCommands type.
	//  TCefTextFieldCommands values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_text_field_commands_t)</see>
	CEF_TFC_UNKNOWN     = 0
	CEF_TFC_CUT         = 1
	CEF_TFC_COPY        = 2
	CEF_TFC_PASTE       = 3
	CEF_TFC_SELECT_ALL  = 4
	CEF_TFC_SELECT_WORD = 5
	CEF_TFC_UNDO        = 6
	CEF_TFC_DELETE      = 7
	CEF_TFC_NUM_VALUES  = 8
	//  Chrome toolbar types.
	//  TCefChromeToolbarType values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_chrome_toolbar_type_t)</see>
	CEF_CTT_UNKNOWN    = 0
	CEF_CTT_NONE       = 1
	CEF_CTT_NORMAL     = 2
	CEF_CTT_LOCATION   = 3
	CEF_CTT_NUM_VALUES = 4
	//  Docking modes supported by ICefWindow.AddOverlay.
	//  TCefDockingMode values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_docking_mode_t)</see>
	CEF_DOCKING_MODE_TOP_LEFT     = 0
	CEF_DOCKING_MODE_TOP_RIGHT    = 1
	CEF_DOCKING_MODE_BOTTOM_LEFT  = 2
	CEF_DOCKING_MODE_BOTTOM_RIGHT = 3
	CEF_DOCKING_MODE_CUSTOM       = 4
	CEF_DOCKING_MODE_NUM_VALUES   = CEF_DOCKING_MODE_CUSTOM + 1
	//  Show the window as normal.
	//  TCefShowState value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_show_state_t)</see>
	CEF_SHOW_STATE_NORMAL = 0
	//  Show the window as minimized.
	//  TCefShowState value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_show_state_t)</see>
	CEF_SHOW_STATE_MINIMIZED = 1
	//  Show the window as maximized.
	//  TCefShowState value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_show_state_t)</see>
	CEF_SHOW_STATE_MAXIMIZED = 2
	//  Show the window as fullscreen.
	//  TCefShowState value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_show_state_t)</see>
	CEF_SHOW_STATE_FULLSCREEN = 3
	//  Show the window as hidden (no dock thumbnail).
	//  Only supported on MacOS..
	//  TCefShowState value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_show_state_t)</see>
	CEF_SHOW_STATE_HIDDEN = 4
	//  Number of TCefShowState values
	//  TCefShowState value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_show_state_t)</see>
	CEF_SHOW_STATE_NUM_VALUES = CEF_SHOW_STATE_HIDDEN + 1
	//  Supported quick menu state bit flags.
	//  TCefQuickMenuEditStateFlags values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_quick_menu_edit_state_flags_t)</see>
	QM_EDITFLAG_NONE         = 0
	QM_EDITFLAG_CAN_ELLIPSIS = 1 << 0
	QM_EDITFLAG_CAN_CUT      = 1 << 1
	QM_EDITFLAG_CAN_COPY     = 1 << 2
	QM_EDITFLAG_CAN_PASTE    = 1 << 3
	//  Values indicating what state of the touch handle is set.
	//  TCefTouchHandleStateFlags values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_touch_handle_state_flags_t)</see>
	CEF_THS_FLAG_NONE        = 0
	CEF_THS_FLAG_ENABLED     = 1 << 0
	CEF_THS_FLAG_ORIENTATION = 1 << 1
	CEF_THS_FLAG_ORIGIN      = 1 << 2
	CEF_THS_FLAG_ALPHA       = 1 << 3
	//  No permission.
	//  TCefMediaAccessPermissionTypes value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t)</see>
	CEF_MEDIA_PERMISSION_NONE = 0
	//  Device audio capture permission.
	//  TCefMediaAccessPermissionTypes value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t)</see>
	CEF_MEDIA_PERMISSION_DEVICE_AUDIO_CAPTURE = 1 << 0
	//  Device video capture permission.
	//  TCefMediaAccessPermissionTypes value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t)</see>
	CEF_MEDIA_PERMISSION_DEVICE_VIDEO_CAPTURE = 1 << 1
	//  Desktop audio capture permission.
	//  TCefMediaAccessPermissionTypes value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t)</see>
	CEF_MEDIA_PERMISSION_DESKTOP_AUDIO_CAPTURE = 1 << 2
	//  Desktop video capture permission.
	//  TCefMediaAccessPermissionTypes value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_access_permission_types_t)</see>
	CEF_MEDIA_PERMISSION_DESKTOP_VIDEO_CAPTURE = 1 << 3
	//  Permission types used with OnShowPermissionPrompt. Some types are
	//  platform-specific or only supported with Chrome style. Should be kept
	//  in sync with Chromium's permissions::RequestType type.
	//  TCefPermissionRequestTypes values.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_permission_request_types_t)</see>
	CEF_PERMISSION_TYPE_NONE                       = 0
	CEF_PERMISSION_TYPE_AR_SESSION                 = 1 << 0
	CEF_PERMISSION_TYPE_CAMERA_PAN_TILT_ZOOM       = 1 << 1
	CEF_PERMISSION_TYPE_CAMERA_STREAM              = 1 << 2
	CEF_PERMISSION_TYPE_CAPTURED_SURFACE_CONTROL   = 1 << 3
	CEF_PERMISSION_TYPE_CLIPBOARD                  = 1 << 4
	CEF_PERMISSION_TYPE_TOP_LEVEL_STORAGE_ACCESS   = 1 << 5
	CEF_PERMISSION_TYPE_DISK_QUOTA                 = 1 << 6
	CEF_PERMISSION_TYPE_LOCAL_FONTS                = 1 << 7
	CEF_PERMISSION_TYPE_GEOLOCATION                = 1 << 8
	CEF_PERMISSION_TYPE_HAND_TRACKING              = 1 << 9
	CEF_PERMISSION_TYPE_IDENTITY_PROVIDER          = 1 << 10
	CEF_PERMISSION_TYPE_IDLE_DETECTION             = 1 << 11
	CEF_PERMISSION_TYPE_MIC_STREAM                 = 1 << 12
	CEF_PERMISSION_TYPE_MIDI_SYSEX                 = 1 << 13
	CEF_PERMISSION_TYPE_MULTIPLE_DOWNLOADS         = 1 << 14
	CEF_PERMISSION_TYPE_NOTIFICATIONS              = 1 << 15
	CEF_PERMISSION_TYPE_KEYBOARD_LOCK              = 1 << 16
	CEF_PERMISSION_TYPE_POINTER_LOCK               = 1 << 17
	CEF_PERMISSION_TYPE_PROTECTED_MEDIA_IDENTIFIER = 1 << 18
	CEF_PERMISSION_TYPE_REGISTER_PROTOCOL_HANDLER  = 1 << 19
	CEF_PERMISSION_TYPE_STORAGE_ACCESS             = 1 << 20
	CEF_PERMISSION_TYPE_VR_SESSION                 = 1 << 21
	CEF_PERMISSION_TYPE_WEB_APP_INSTALLATION       = 1 << 22
	CEF_PERMISSION_TYPE_WINDOW_MANAGEMENT          = 1 << 23
	CEF_PERMISSION_TYPE_FILE_SYSTEM_ACCESS         = 1 << 24
	CEF_PERMISSION_TYPE_LOCAL_NETWORK_ACCESS       = 1 << 25
	CEF_PERMISSION_TYPE_LOCAL_NETWORK              = 1 << 26
	CEF_PERMISSION_TYPE_LOOPBACK_NETWORK           = 1 << 27
	CEF_PERMISSION_TYPE_SENSORS                    = 1 << 28
	//  Platform API hash.
	//  ucef_api_hash parameter.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/cef_api_hash.h">CEF source file: /include/cef_api_hash.h</see>
	CEF_API_HASH_PLATFORM = 0
	//  Universal API hash. (deprecated, same as CEF_API_HASH_PLATFORM)
	//  ucef_api_hash parameter.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/cef_api_hash.h">CEF source file: /include/cef_api_hash.h</see>
	CEF_API_HASH_UNIVERSAL = 1
	//  Commit hash.
	//  ucef_api_hash parameter.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/cef_api_hash.h">CEF source file: /include/cef_api_hash.h</see>
	CEF_COMMIT_HASH = 2
	//  Sandbox hash.
	//  ucef_api_hash parameter.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/cef_api_hash.h">CEF source file: /include/cef_api_hash.h</see>
	CEF_SANDBOX_COMPAT_HASH = 3
	//  No interrupt reason.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_NONE = 0
	//  Generic file operation failure.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_FAILED = 1
	//  The file cannot be accessed due to security restrictions.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_ACCESS_DENIED = 2
	//  There is not enough room on the drive.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_NO_SPACE = 3
	//  The directory or file name is too long.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_NAME_TOO_LONG = 5
	//  The file is too large for the file system to handle.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_LARGE = 6
	//  The file contains a virus.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_VIRUS_INFECTED = 7
	//  The file was in use. Too many files are opened at once. We have run out of
	//  memory.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_TRANSIENT_ERROR = 10
	//  The file was blocked due to local policy.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_BLOCKED = 11
	//  An attempt to check the safety of the download failed due to unexpected
	//  reasons. See http://crbug.com/153212.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_SECURITY_CHECK_FAILED = 12
	//  An attempt was made to seek past the end of a file in opening
	//  a file (as part of resuming a previously interrupted download).
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT = 13
	//  The partial file didn't match the expected hash.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH = 14
	//  The source and the target of the download were the same.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_FILE_SAME_AS_SOURCE = 15
	//  Generic network failure.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_NETWORK_FAILED = 20
	//  The network operation timed out.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_NETWORK_TIMEOUT = 21
	//  The network connection has been lost.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_NETWORK_DISCONNECTED = 22
	//  The server has gone down.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_NETWORK_SERVER_DOWN = 23
	//  The network request was invalid. This may be due to the original URL or a
	//  redirected URL:
	//  - Having an unsupported scheme.
	//  - Being an invalid URL.
	//  - Being disallowed by policy.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_NETWORK_INVALID_REQUEST = 24
	//  The server indicates that the operation has failed (generic).
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_FAILED = 30
	//  The server does not support range requests.
	//  Internal use only: must restart from the beginning.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE = 31
	//  The server does not have the requested data.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_BAD_CONTENT = 33
	//  Server didn't authorize access to resource.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_UNAUTHORIZED = 34
	//  Server certificate problem.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_CERT_PROBLEM = 35
	//  Server access forbidden.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_FORBIDDEN = 36
	//  Unexpected server response. This might indicate that the responding server
	//  may not be the intended server.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_UNREACHABLE = 37
	//  The server sent fewer bytes than the content-length header. It may
	//  indicate that the connection was closed prematurely, or the Content-Length
	//  header was invalid. The download is only interrupted if strong validators
	//  are present. Otherwise, it is treated as finished.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_CONTENT_LENGTH_MISMATCH = 38
	//  An unexpected cross-origin redirect happened.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_SERVER_CROSS_ORIGIN_REDIRECT = 39
	//  The user canceled the download.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_USER_CANCELED = 40
	//  The user shut down the browser.
	//  Internal use only: resume pending downloads if possible.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_USER_SHUTDOWN = 41
	//  The browser crashed.
	//  Internal use only: resume pending downloads if possible.
	//  TCefDownloadInterruptReason value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_download_interrupt_reason_t)</see>
	CEF_DOWNLOAD_INTERRUPT_REASON_CRASH = 50
	//  Prepend the default list of items.
	//  TCefLogItems value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_items_t)</see>
	LOG_ITEMS_DEFAULT = 0
	//  Prepend no items.
	//  TCefLogItems value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_items_t)</see>
	LOG_ITEMS_NONE = 1 << 0
	//  Prepend the process ID.
	//  TCefLogItems value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_items_t)</see>
	LOG_ITEMS_FLAG_PROCESS_ID = 1 << 1
	//  Prepend the thread ID.
	//  TCefLogItems value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_items_t)</see>
	LOG_ITEMS_FLAG_THREAD_ID = 1 << 2
	//  Prepend the timestamp.
	//  TCefLogItems value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_items_t)</see>
	LOG_ITEMS_FLAG_TIME_STAMP = 1 << 3
	//  Prepend the tickcount.
	//  TCefLogItems value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_log_items_t)</see>
	LOG_ITEMS_FLAG_TICK_COUNT = 1 << 4
	//  Normal exit code.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/result_codes.h">See Chromium's content::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NORMAL_EXIT = 0
	//  Process was killed by user or system.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/result_codes.h">See Chromium's content::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_KILLED = 1
	//  Process hung.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/result_codes.h">See Chromium's content::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_HUNG = 2
	//  A bad message caused the process termination.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/result_codes.h">See Chromium's content::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_KILLED_BAD_MESSAGE = 3
	//  The GPU process exited because initialization failed.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:content/public/common/result_codes.h">See Chromium's content::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_GPU_DEAD_ON_ARRIVAL = 4
	//  First Chrome result code.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_CHROME_FIRST = 5
	//  An invalid command line url was given.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_INVALID_CMDLINE_URL = CEF_RESULT_CODE_CHROME_FIRST
	//  The process is of an unknown type.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_BAD_PROCESS_TYPE = 6
	//  A critical chrome file is missing.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_MISSING_DATA = 7
	//  Failed to make Chrome default browser (not used?).
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SHELL_INTEGRATION_FAILED = 8
	//  Machine level install exists.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_MACHINE_LEVEL_INSTALL_EXISTS = 9
	//  Uninstall detected another chrome instance.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_UNINSTALL_CHROME_ALIVE = 10
	//  The user changed their mind.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_UNINSTALL_USER_CANCEL = 11
	//  Delete profile as well during uninstall.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_UNINSTALL_DELETE_PROFILE = 12
	//  Command line parameter is not supported.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_UNSUPPORTED_PARAM = 13
	//  Browser import hung and was killed.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_IMPORTER_HUNG = 14
	//  Trying to restart the browser we crashed.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_RESPAWN_FAILED = 15
	//  Generic code used to communicate some
	//  simple outcome back to the process that launched us. This is used for
	//  experiments and the actual meaning depends on the experiment.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NORMAL_EXIT_EXP1 = 16
	//  Generic code used to communicate some
	//  simple outcome back to the process that launched us. This is used for
	//  experiments and the actual meaning depends on the experiment.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NORMAL_EXIT_EXP2 = 17
	//  Generic code used to communicate some
	//  simple outcome back to the process that launched us. This is used for
	//  experiments and the actual meaning depends on the experiment.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NORMAL_EXIT_EXP3 = 18
	//  Generic code used to communicate some
	//  simple outcome back to the process that launched us. This is used for
	//  experiments and the actual meaning depends on the experiment.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NORMAL_EXIT_EXP4 = 19
	//  For experiments this return code means that the user canceled causes the
	//  did_run "dr" signal to be reset soi this chrome run does not count as
	//  active chrome usage.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NORMAL_EXIT_CANCEL = 20
	//  The profile was in use on another host.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_PROFILE_IN_USE = 21
	//  Failed to pack an extension via the cmd line.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_PACK_EXTENSION_ERROR = 22
	//  Failed to silently uninstall an extension.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_UNINSTALL_EXTENSION_ERROR = 23
	//  The browser process exited early by passing the command line to another
	//  running browser.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NORMAL_EXIT_PROCESS_NOTIFIED = 24
	//  A dummy value we should not use. See crbug.com/152285.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NOTUSED_1 = 25
	//  Failed to install an item from the webstore when the
	//  kInstallEphemeralAppFromWebstore command line flag was present.
	//  As this flag is no longer supported, this return code should never be
	//  returned.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_INSTALL_FROM_WEBSTORE_ERROR_2 = 26
	//  A dummy value we should not use. See crbug.com/152285.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NOTUSED_2 = 27
	//  Returned when the user has not yet accepted the EULA.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_EULA_REFUSED = 28
	//  Failed to migrate user data directory for side-by-side package support
	//  (Linux-only).
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SXS_MIGRATION_FAILED_NOT_USED = 29
	//  The action is not allowed by a policy.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_ACTION_DISALLOWED_BY_POLICY = 30
	//  A browser process was sandboxed. This should never happen.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_INVALID_SANDBOX_STATE = 31
	//  Cloud policy enrollment is failed or given up by user.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_CLOUD_POLICY_ENROLLMENT_FAILED = 32
	//  Chrome was downgraded since the last launch. Perform downgrade processing
	//  and relaunch.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_DOWNGRADE_AND_RELAUNCH = 33
	//  The GPU process was terminated due to context lost.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_GPU_EXIT_ON_CONTEXT_LOST = 34
	//  Chrome detected that there was a new version waiting to launch and renamed
	//  the files and launched the new version. This result code is never returned
	//  from the main process, but is instead used as a signal for early
	//  termination of browser. See `IsNormalResultCode` below.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NORMAL_EXIT_UPGRADE_RELAUNCHED = 35
	//  An early startup command was executed and the browser must exit.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NORMAL_EXIT_PACK_EXTENSION_SUCCESS = 36
	//  The browser process exited because system resource are exhausted. The
	//  system state can't be recovered and will be unstable.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SYSTEM_RESOURCE_EXHAUSTED = 37
	//  The browser process exited because it was re-launched without elevation.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NORMAL_EXIT_AUTO_DE_ELEVATED = 38
	//  Upon encountering a commit failure in a process, PartitionAlloc terminated
	//  another process deemed less important.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_TERMINATED_BY_OTHER_PROCESS_ON_COMMIT_FAILURE = 39
	//  The isolated browser process launched but it was not possible to wait on
	//  the exit of the process, so the browser must exit.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_INVALID_ISOLATED_BROWSER_PROCESS = 40
	//  Last Chrome result code.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:chrome/common/chrome_result_codes.h">See chrome::ResultCode type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_CHROME_LAST = 41
	//  First Sandbox result code.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_FIRST = 7006
	//  Windows sandbox could not set the integrity level.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_INTEGRITY = CEF_RESULT_CODE_SANDBOX_FATAL_FIRST
	//  Windows sandbox could not lower the token.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_DROPTOKEN = 7007
	//  Windows sandbox failed to flush registry handles.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_FLUSHANDLES = 7008
	//  Windows sandbox failed to forbid HCKU caching.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_CACHEDISABLE = 7009
	//  Windows sandbox failed to close pending handles.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_CLOSEHANDLES = 7010
	//  Windows sandbox could not set the mitigation policy.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_MITIGATION = 7011
	//  Windows sandbox exceeded the job memory limit.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_MEMORY_EXCEEDED = 7012
	//  Windows sandbox failed to warmup.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_WARMUP = 7013
	//  Windows sandbox broker terminated in shutdown.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_BROKER_SHUTDOWN_HUNG = 7014
	//  Last Sandbox result code.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_SANDBOX_FATAL_LAST = 7015
	//  Last TCefResultCode value.
	//  TCefResultCode value.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_resultcode_t)</see>
	//  <see href="https://source.chromium.org/chromium/chromium/src/+/main:sandbox/win/src/sandbox_types.h">See sandbox::TerminationCodes type.</see>
	//  Use CefResultCodeToString in uCEFMiscFunctions to convert this value into a human readable message
	CEF_RESULT_CODE_NUM_VALUES = CEF_RESULT_CODE_SANDBOX_FATAL_LAST + 1
	//  *
	//  ******************************************************
	//  ****************** OTHER CONSTANTS *******************
	//  ******************************************************
	//  *
	ABOUTBLANK_URI      = "about:blank"
	DEVTOOLS_WINDOWNAME = "DevTools"
	//  Direct proxy type: Never use a proxy.
	CEF_PROXYTYPE_DIRECT = 0
	//  Auto_detect proxy type: Auto detect proxy settings.
	CEF_PROXYTYPE_AUTODETECT = 1
	//  System proxy type: Use system proxy settings.
	CEF_PROXYTYPE_SYSTEM = 2
	//  Fixed_servers proxy type: Use fixed proxy servers.
	CEF_PROXYTYPE_FIXED_SERVERS = 3
	//  Pac_script proxy type: Use a .pac proxy script.
	CEF_PROXYTYPE_PAC_SCRIPT = 4
	//  Used in the severity parameter in the 'cef_log' function, also known as 'CefLog' in CEF4Delphi.
	//  The log severities are used to index into the array of names, see log_severity_names.
	//  /base/allocator/partition_allocator/partition_alloc_base/logging.h
	//  /base/logging.cc
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_logging.h">CEF source file: /include/base/cef_logging.h (LogSeverity)</see>
	CEF_LOG_SEVERITY_VERBOSE = -1
	CEF_LOG_SEVERITY_INFO    = 0
	CEF_LOG_SEVERITY_WARNING = 1
	CEF_LOG_SEVERITY_ERROR   = 2
	//  <summary>This severity log level causes a crash.</summary>
	CEF_LOG_SEVERITY_FATAL                      = 3
	CEF_MAX_CONNECTIONS_PER_PROXY_DEFAULT_VALUE = 32
	CEF_MAX_CONNECTIONS_PER_PROXY_MIN_VALUE     = 7
	CEF_MAX_CONNECTIONS_PER_PROXY_MAX_VALUE     = 99
	CEF_COOKIE_PREF_DEFAULT                     = 0
	CEF_COOKIE_PREF_ALLOW                       = 1
	CEF_COOKIE_PREF_BLOCK                       = 2
	CEF_DEFAULT_ENABLEFOCUSDELAY                = 500
	//  YouTube restrict mode.
	//  <see href="https://chromium.googlesource.com/chromium/src/+/refs/tags/77.0.3865.90/chrome/common/net/safe_search_util.h">Chromium source file: /chrome/common/net/safe_search_util.h (YouTubeRestrictMode)</see>
	//  <see href="https://www.chromium.org/administrators/policy-list-3#ForceYouTubeRestrict">Chromium policy list: https://www.chromium.org/administrators/policy-list-3#ForceYouTubeRestrict</see>
	YOUTUBE_RESTRICT_OFF      = 0
	YOUTUBE_RESTRICT_MODERATE = 1
	YOUTUBE_RESTRICT_STRICT   = 2
	ZOOM_STEP_25              = 0
	ZOOM_STEP_33              = 1
	ZOOM_STEP_50              = 2
	ZOOM_STEP_67              = 3
	ZOOM_STEP_75              = 4
	ZOOM_STEP_90              = 5
	ZOOM_STEP_100             = 6
	ZOOM_STEP_110             = 7
	ZOOM_STEP_125             = 8
	ZOOM_STEP_150             = 9
	ZOOM_STEP_175             = 10
	ZOOM_STEP_200             = 11
	ZOOM_STEP_250             = 12
	ZOOM_STEP_300             = 13
	ZOOM_STEP_400             = 14
	ZOOM_STEP_500             = 15
	ZOOM_STEP_UNK             = 16
	ZOOM_STEP_MIN             = ZOOM_STEP_25
	ZOOM_STEP_MAX             = ZOOM_STEP_500
	ZOOM_STEP_DEF             = ZOOM_STEP_100
	ZOOM_PCT_DELTA            = 5
	CEF_PREFERENCES_SAVED     = WM_APP + 0xA00
	CEF_DOONCLOSE             = WM_APP + 0xA01
	CEF_STARTDRAGGING         = WM_APP + 0xA02
	CEF_AFTERCREATED          = WM_APP + 0xA03
	CEF_PENDINGRESIZE         = WM_APP + 0xA04
	CEF_PUMPHAVEWORK          = WM_APP + 0xA05
	CEF_DESTROY               = WM_APP + 0xA06
	CEF_DOONBEFORECLOSE       = WM_APP + 0xA07
	CEF_PENDINGINVALIDATE     = WM_APP + 0xA08
	CEF_IMERANGECHANGED       = WM_APP + 0xA09
	CEF_SENTINEL_START        = WM_APP + 0xA0A
	CEF_SENTINEL_DOCLOSE      = WM_APP + 0xA0B
	CEF_BEFORECLOSE           = WM_APP + 0xA0C
	CEF_INVALIDATE            = WM_APP + 0xA0D
	CEF_FOCUSENABLED          = WM_APP + 0xA0E
	//  Default values for the Windowsless framerate setting in TChromiumOptions
	//  The values are frames per second.
	CEF_OSR_FRAMERATE_DEFAULT                 = 30
	CEF_OSR_SHARED_TEXTURES_FRAMERATE_DEFAULT = 60
	CEF_TIMER_MINIMUM                         = 0x0000000A
	CEF_TIMER_MAXIMUM                         = 0x7FFFFFFF
	CEF_TIMER_MAXDELAY                        = 1000 / CEF_OSR_FRAMERATE_DEFAULT
	CEF_TIMER_DEPLETEWORK_CYCLES              = 10
	CEF_TIMER_DEPLETEWORK_DELAY               = 50
	CRLF                                      = "\r" + "\n"
	//  These contants are declared in the "Windows" unit but
	//  some old Delphi versions don't have them.
	//  We have to add "CEF_" to be compatible with C++ Builder.
	CEF_IMAGE_FILE_MACHINE_I386  = 0x014C
	CEF_IMAGE_FILE_MACHINE_AMD64 = 0x8664
	//  Modifier values used in the Input.dispatchTouchEvent and Input.dispatchMouseEvent DevTools methods.
	//  Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	//  https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchTouchEvent
	CEF_MOUSETOUCH_EVENT_MODIFIERS_NONE    = 0
	CEF_MOUSETOUCH_EVENT_MODIFIERS_ALT     = 1 << 0
	CEF_MOUSETOUCH_EVENT_MODIFIERS_CTRL    = 1 << 1
	CEF_MOUSETOUCH_EVENT_MODIFIERS_METACMD = 1 << 2
	CEF_MOUSETOUCH_EVENT_MODIFIERS_SHIFT   = 1 << 3
	//  Modifier values used in the Input.dispatchMouseEvent DevTools method.
	//  A number indicating which buttons are pressed on the mouse when a mouse event is triggered. Left=1, Right=2, Middle=4, Back=8, Forward=16, None=0.
	//  https://chromedevtools.github.io/devtools-protocol/tot/Input/#method-dispatchMouseEvent
	CEF_PRESSED_MOUSE_BUTTONS_NONE    = 0
	CEF_PRESSED_MOUSE_BUTTONS_LEFT    = 1 << 0
	CEF_PRESSED_MOUSE_BUTTONS_RIGHT   = 1 << 1
	CEF_PRESSED_MOUSE_BUTTONS_MIDDLE  = 1 << 2
	CEF_PRESSED_MOUSE_BUTTONS_BACK    = 1 << 3
	CEF_PRESSED_MOUSE_BUTTONS_FORWARD = 1 << 4
	//  This constant is defined by Chromium in chrome/app/main_dll_loader_win.cc
	//  It's used with SetProcessShutdownParameters to set a shutdown priority for the
	//  subprocesses. $280 is the default value for applications.
	CHROMIUM_NONBROWSERSHUTDOWNPRIORITY = 0x280
	//  Maximum number of accelerated paint planes used in TCefAcceleratedPaintInfo.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_linux.h">CEF source file: /include/internal/cef_types_linux.h (kAcceleratedPaintMaxPlanes)</see>
	CEF_KACCELERATEDPAINTMAXPLANES = 4
	//  This value may be used with the mseconds_between_dumps parameter in GlobalCEFApp.DumpWithoutCrashing.
	//  <see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/base/cef_dump_without_crashing.h">CEF source file: /include/base/cef_dump_without_crashing.h (kOneDayInMilliseconds)</see>
	CEF_ONEDAYINMILLISECONDS = 86400000
)
