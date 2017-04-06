type ReverseProxy struct {
        // Director must be a function which modifies
        // the request into a new request to be sent
        // using Transport. Its response is then copied
        // back to the original client unmodified.
        // Director must not access the provided Request
        // after returning.
        Director func(*http.Request)

        // The transport used to perform proxy requests.
        // If nil, http.DefaultTransport is used.
        Transport http.RoundTripper

        // FlushInterval specifies the flush interval
        // to flush to the client while copying the
        // response body.
        // If zero, no periodic flushing is done.
        FlushInterval time.Duration

        // ErrorLog specifies an optional logger for errors
        // that occur when attempting to proxy the request.
        // If nil, logging goes to os.Stderr via the log package's
        // standard logger.
        ErrorLog *log.Logger

        // BufferPool optionally specifies a buffer pool to
        // get byte slices for use by io.CopyBuffer when
        // copying HTTP response bodies.
        BufferPool BufferPool

        // ModifyResponse is an optional function that
        // modifies the Response from the backend.
        // If it returns an error, the proxy returns a StatusBadGateway error.
        ModifyResponse func(*http.Response) error
}
