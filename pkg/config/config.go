package config

type DaemonConfig struct {
	Providers Providers `yaml:"providers"`
	Version   string    `yaml:"version"`
}

type Providers struct {
	Aws        []Aws        `yaml:"aws"`
	Vsphere    []Vsphere    `yaml:"vsphere"`
	Virtualbox []Virtualbox `yaml:"virtualbox"`
	Qemu       []Qemu       `yaml:"qemu"`
	Openstack  []Openstack  `yaml:"openstack"`
}

type Aws struct {
	Name   string `yaml:"name"`
	Region string `yaml:"region"`
	Zone   string `yaml:"zone"`
}

type Vsphere struct {
	Name            string `yaml:"name"`
	VsphereUser     string `yaml:"vsphere_user"`
	VspherePassword string `yaml:"vsphere_password"`
	VsphereURL      string `yaml:"vsphere_url"`
	Datastore       string `yaml:"datastore"`
	Datacenter      string `yaml:"datacenter"`
	NetworkLabel    string `yaml:"network"`
}

type Virtualbox struct {
	Name                  string                `yaml:"name"`
	AdapterName           string                `yaml:"adapter_name"`
	VirtualboxAdapterType VirtualboxAdapterType `yaml:"adapter_type"`
}

type Qemu struct {
	Name         string `yaml:"name"`
	NoGraphic    bool   `yaml:"no_graphic"`
	DebuggerPort int    `yaml:"debugger_port"`
}

type Openstack struct {
	Name          string `yaml:"name"`
	OSUser        string `yaml:"os_username"`
	OSPassword    string `yaml:"os_password"`
	OSAuthUrl     string `yaml:"os_auth_url"`
	OSTenantId    string `yaml:"os_tenant_id"`
	OSTenantName  string `yaml:"os_tenant_name"`
	OSProjectName string `yaml:"os_project_name"`
	OSRegion      string `yaml:"os_region"`
}

type VirtualboxAdapterType string

const (
	BridgedAdapter  = VirtualboxAdapterType("bridged")
	HostOnlyAdapter = VirtualboxAdapterType("host_only")
)

type ClientConfig struct {
	Host string `yaml:"host"`
}
