package build

import (
	"sync"
)

type Info struct{}

var mu = &sync.Mutex{}

func (i *Info) Version() string {
	mu.Lock()
	defer mu.Unlock()
	return Version
}

func (i *Info) SHA() string {
	mu.Lock()
	defer mu.Unlock()
	return SHA
}
func (i *Info) BuildDate() string {
	mu.Lock()
	defer mu.Unlock()
	return BuildDate
}
func (i *Info) License() string {
	mu.Lock()
	defer mu.Unlock()
	return License
}
func (i *Info) HasTLS() bool {
	mu.Lock()
	defer mu.Unlock()
	return HasTLS()
}
func (i *Info) MaxBrokerClients() int {
	mu.Lock()
	defer mu.Unlock()
	return MaxBrokerClients()
}
func (i *Info) ProvisionSecurity() bool {
	mu.Lock()
	defer mu.Unlock()
	return ProvisionSecurity()
}
func (i *Info) ProvisionDefault() bool {
	mu.Lock()
	defer mu.Unlock()
	return ProvisionDefault()
}
func (i *Info) ProvisionBrokerURLs() string {
	mu.Lock()
	defer mu.Unlock()
	return ProvisionBrokerURLs
}
func (i *Info) ProvisionBrokerSRVDomain() string {
	mu.Lock()
	defer mu.Unlock()
	return ProvisionBrokerSRVDomain
}
func (i *Info) ProvisionAgent() bool {
	mu.Lock()
	defer mu.Unlock()
	return ProvisionAgent == "true"
}
func (i *Info) ProvisionRegistrationData() string {
	mu.Lock()
	defer mu.Unlock()
	return ProvisionRegistrationData
}
func (i *Info) ProvisionFacts() string {
	mu.Lock()
	defer mu.Unlock()
	return ProvisionFacts
}
func (i *Info) ProvisionToken() string {
	mu.Lock()
	defer mu.Unlock()
	return ProvisionToken
}
func (i *Info) ProvisionJWTFile() string {
	mu.Lock()
	defer mu.Unlock()
	return ProvisionJWTFile
}
func (i *Info) ProvisionStatusFile() string {
	mu.Lock()
	defer mu.Unlock()
	return ProvisionStatusFile
}
func (i *Info) AgentProviders() []string {
	mu.Lock()
	defer mu.Unlock()
	return AgentProviders
}

func (i *Info) RegisterAgentProvider(p string) {
	mu.Lock()
	defer mu.Unlock()

	AgentProviders = append(AgentProviders, p)
}

func (i *Info) SetProvisionBrokerURLs(u string) {
	mu.Lock()
	defer mu.Unlock()

	ProvisionBrokerURLs = u
}

func (i *Info) SetProvisionToken(t string) {
	mu.Lock()
	defer mu.Unlock()

	ProvisionToken = t
}

func (i *Info) SetProvisionBrokerSRVDomain(d string) {
	mu.Lock()
	defer mu.Unlock()

	ProvisionBrokerSRVDomain = d
}

func (i *Info) EnableProvisionModeAsDefault() {
	mu.Lock()
	defer mu.Unlock()

	ProvisionModeDefault = "true"
}

func (i *Info) DisableProvisionModeAsDefault() {
	mu.Lock()
	defer mu.Unlock()

	ProvisionModeDefault = "false"
}

func (i *Info) EnableProvisionModeSecurity() {
	mu.Lock()
	defer mu.Unlock()

	ProvisionSecure = "true"
}

func (i *Info) DisableProvisionModeSecurity() {
	mu.Lock()
	defer mu.Unlock()

	ProvisionSecure = "false"
}

func (i *Info) SetProvisionFacts(f string) {
	mu.Lock()
	defer mu.Unlock()

	ProvisionFacts = f
}

func (i *Info) SetProvisionRegistrationData(f string) {
	mu.Lock()
	defer mu.Unlock()

	ProvisionRegistrationData = f
}
