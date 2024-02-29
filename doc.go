/*
Package stacktrace provides functions for wrapping an error to include line
number and/or error code information.

A stacktrace produced by this package looks like this:

Failed to register for villain discovery
--- at github.com/houmanka/shield/agent/discovery.go:265 (ShieldAgent.reallyRegister) ---
--- at github.com/houmanka/shield/connector/impl.go:89 (Connector.Register) ---
Caused by: Failed to load S.H.I.E.L.D.config from /opt/shield/conf/shield.yaml
--- at github.com/houmanka/shield/connector/config.go:44 (withShieldConfig) ---
Caused by: There isn't enough time (4 picoseconds required)
--- at github.com/houmanka/shield/axiom/pseudo/resource.go:46 (PseudoResource.Adjust) ---
--- at github.com/houmanka/shield/axiom/pseudo/growth.go:110 (reciprocatingPseudo.growDown) ---
--- at github.com/houmanka/shield/axiom/pseudo/growth.go:121 (reciprocatingPseudo.verify) ---
Caused by: Inverse tachyon pulse failed
--- at github.com/houmanka/shield/metaphysic/tachyon.go:72 (TryPulse) ---

Note that stack traces are not designed to be user-visible.They can be valuable
in a log file of a server application, but nobody wants to see one of them in
CLI output or a web interface or a return value from library code.
*/
package stacktrace
