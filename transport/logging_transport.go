package transport

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

var (
	DefaultRequestFormat  = "REQUEST:\n--------\n%s\n--------\n"
	DefaultResponseFormat = "RESPONSE:\n========\n%s\n========\n"
	DefaultLoggingOutput  = os.Stderr
)

func RequestFormat(fmt string) func(*LoggingTransport) {
	return func(t *LoggingTransport) { t.requestFormat = fmt }
}

func ResponseFormat(fmt string) func(*LoggingTransport) {
	return func(t *LoggingTransport) { t.responseFormat = fmt }
}

func LoggingOutput(w io.Writer) func(*LoggingTransport) {
	return func(t *LoggingTransport) { t.output = w }
}

func NewLoggingTransport(inner http.RoundTripper, opts ...func(*LoggingTransport)) *LoggingTransport {
	xport := &LoggingTransport{
		inner:          inner,
		output:         DefaultLoggingOutput,
		requestFormat:  DefaultRequestFormat,
		responseFormat: DefaultResponseFormat,
	}

	for _, opt := range opts {
		opt(xport)
	}

	return xport
}

type LoggingTransport struct {
	inner http.RoundTripper

	output         io.Writer
	requestFormat  string
	responseFormat string
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.output == nil {
		return t.inner.RoundTrip(req)
	}

	var (
		err  error
		data []byte
		resp *http.Response
	)

	if data, err = httputil.DumpRequestOut(req, true); err != nil {
		return nil, fmt.Errorf("logging request: %w", err)
	}

	fmt.Fprintf(t.output, t.requestFormat, string(data))

	if resp, err = http.DefaultTransport.RoundTrip(req); err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	if data, err = httputil.DumpResponse(resp, true); err != nil {
		return nil, fmt.Errorf("logging response: %w", err)
	}

	fmt.Fprintf(t.output, t.responseFormat, string(data))

	return resp, nil
}
