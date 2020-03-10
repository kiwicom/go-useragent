package useragent

// UserAgent contains all available fields in UserAgent according to RFC-22 rules.
type UserAgent struct {
	// Name is the service name
	Name string
	// Version is the version of the remote service, can be tag or commit hash
	Version string
	// Environment decribes where the service is run (e.g. sandbox, production)
	Environment string
	// SystemInfo describes extra metadata (e.g. node-fetch 1.2)
	SystemInfo string
}
