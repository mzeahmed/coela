package project

// WebServer identifies the HTTP server used by a project.
type WebServer string

const (
	WebServerNginx  WebServer = "nginx"
	WebServerApache WebServer = "apache"
)

// String returns the human-readable label for w.
func (w WebServer) String() string {
	switch w {
	case WebServerNginx:
		return "Nginx"
	case WebServerApache:
		return "Apache"
	default:
		return string(w)
	}
}
