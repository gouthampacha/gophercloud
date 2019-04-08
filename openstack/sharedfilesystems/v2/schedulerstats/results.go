package schedulerstats

import (
	"encoding/json"
	"math"
    "github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Capabilities represents the characteristics of an individual StoragePool.
type Capabilities struct {
	// Common capabilities that should be present in all storage pools
	DriverHandlesShareServers string     `json:"driver_handles_share_servers"`
	DriverVersion             string     `json:"driver_version"`
	StorageProtocol           string     `json:"storage_protocol"`
	FreeCapacityGB            float64    `json:"-"`
	TotalCapacityGB           float64    `json:"-"`
	MaxOverSubscriptionRatio  float64    `json:"max_over_subscription_ratio"`
	ReservedPercentage        float64    `json:"reserved_percentage"`
	VendorName                string     `json:"vendor_name"`
	ShareBackendName          string     `json:"share_backend_name"`
	UpdatedAt                 time.Time  `json:"-"`

	// Common capabilities present since Liberty release
	SnapshotSupport           bool       `json:"snapshot_support"`

	// Common capabilities present since Mitaka release
	QoS                       bool       `json:"qos"`
	// Following values can be null where capability is not supported
	ReplicationType           string     `json:"replication_type"`
	ReplicationDomain         string     `json:"replication_domain"`
	// Following options need to be un-marshalled for special behavior
	// Dedupe                    []bool     `json:"-"`
	// Compression               []bool     `json:"-"`
	// ThinProvisioning          []bool     `json:"-"`

	// Common capabilities present since Ocata release
	CreateShareFromSnapshotSupport bool  `json:"create_share_from_snapshot_support"`
    RevertToSnapshotSupport        bool  `json:"revert_to_snapshot_support"`
    MountSnapshotSupport           bool  `json:"mount_snapshot_support"`

	// Common capabilities present since Pike release
	ShareGroupConsistentSnapshotSupport string  `json:"sg_consistent_snapshot_support"`
	IPv4Support                         bool    `json:"ipv4_support"`
	IPv6Support                         bool    `json:"ipv6_support"`


	// Optional capabilities presented
    // Are there any??
}

// StoragePool represents an individual StoragePool retrieved from the
// schedulerstats API.
type StoragePool struct {
	Name         string       `json:"name"`
	Host         string       `json:"host"`
	Backend      string       `json:"backend"`
	Pool         string       `json:"pool"`
	Capabilities Capabilities `json:"capabilities"`
}

func (r *Capabilities) UnmarshalJSON(b []byte) error {
	type tmp Capabilities
    type timestamp = gophercloud.JSONRFC3339MilliNoZ


	var s struct {
		tmp
		FreeCapacityGB       interface{}    `json:"free_capacity_gb"`
		TotalCapacityGB      interface{}    `json:"total_capacity_gb"`
		UpdatedAt            timestamp      `json:"updated"`

		// Dedupe            interface{}    `json:"dedupe"`
		// Compression       interface{}    `json:"compression"`
		// ThinProvisioning  interface{}    `json:"thin_provisioning"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Capabilities(s.tmp)

	// Generic function to parse a capacity value which may be a numeric
	// value, "unknown", or "infinite"
	parseCapacity := func(capacity interface{}) float64 {
		if capacity != nil {
			switch capacity.(type) {
			case float64:
				return capacity.(float64)
			case string:
				if capacity.(string) == "infinite" {
					return math.Inf(1)
				}
			}
		}
		return 0.0
	}

	r.FreeCapacityGB = parseCapacity(s.FreeCapacityGB)
	r.TotalCapacityGB = parseCapacity(s.TotalCapacityGB)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	return nil
}

// StoragePoolPage is a single page of all List results.
type StoragePoolPage struct {
	pagination.SinglePageBase
}

// IsEmpty satisfies the IsEmpty method of the Page interface. It returns true
// if a List contains no results.
func (page StoragePoolPage) IsEmpty() (bool, error) {
	va, err := ExtractStoragePools(page)
	return len(va) == 0, err
}

// ExtractStoragePools takes a List result and extracts the collection of
// StoragePools returned by the API.
func ExtractStoragePools(p pagination.Page) ([]StoragePool, error) {
	var s struct {
		StoragePools []StoragePool `json:"pools"`
	}
	err := (p.(StoragePoolPage)).ExtractInto(&s)
	return s.StoragePools, err
}
