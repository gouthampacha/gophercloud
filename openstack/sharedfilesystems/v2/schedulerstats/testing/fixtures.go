package testing

import (
	"fmt"
	"math"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/sharedfilesystems/v2/schedulerstats"
	"github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

const StoragePoolsListBody = `
{
    "pools": [
        {
            "host": "ostkstoragecontroller01",
            "name": "ostkstoragecontroller01@vienna#lvm-single-pool",
            "pool": "lvm-single-pool",
            "backend": "vienna"
        },
        {
            "host": "ostkstoragecontroller01",
            "name": "ostkstoragecontroller01@gamma#gamma_data_pool",
            "pool": "gamma_data_pool",
            "backend": "gamma"
        },
        {
            "host": "ostkstoragecontroller01",
            "name": "ostkstoragecontroller01@theta#theta_data_pool",
            "pool": "theta_data_pool",
            "backend": "theta"
        },
        {
            "host": "ostkstoragecontroller01",
            "name": "ostkstoragecontroller01@prague#lvm-single-pool",
            "pool": "lvm-single-pool",
            "backend": "prague"
        },
        {
            "host": "ostkstoragecontroller01",
            "name": "ostkstoragecontroller01@delta#delta_data_pool",
            "pool": "delta_data_pool",
            "backend": "delta"
        }
    ]
}
`

const StoragePoolsListBodyDetail = `
{
    "pools": [
        {
            "host": "ostkstoragecontroller01",
            "capabilities": {
                "pool_name": "lvm-single-pool",
                "vendor_name": "Open Source",
                "compression": false,
                "driver_version": "1.0",
                "snapshot_support": true,
                "replication_type": null,
                "ipv6_support": false,
                "share_backend_name": "Vienna",
                "driver_handles_share_servers": false,
                "reserved_percentage": 0,
                "total_capacity_gb": 100.0,
                "revert_to_snapshot_support": true,
                "ipv4_support": true,
                "sg_consistent_snapshot_support": null,
                "timestamp": "2019-03-28T23:53:48.343538",
                "create_share_from_snapshot_support": true,
                "dedupe": false,
                "mount_snapshot_support": true,
                "free_capacity_gb": 100.0,
                "storage_protocol": "NFS_CIFS",
                "replication_domain": null
            },
            "name": "ostkstoragecontroller01@vienna#lvm-single-pool",
            "pool": "lvm-single-pool",
            "backend": "vienna"
        },
        {
            "host": "ostkstoragecontroller01",
            "capabilities": {
                "pool_name": "gamma_data_pool",
                "vendor_name": "Open Source",
                "compression": false,
                "timestamp": "2019-03-28T23:53:31.179279",
                "driver_version": "1.0",
                "replication_type": "readable",
                "mount_snapshot_support": true,
                "share_backend_name": "GAMMA",
                "driver_handles_share_servers": true,
                "reserved_percentage": 0,
                "total_capacity_gb": 1230.0,
                "revert_to_snapshot_support": true,
                "ipv4_support": true,
                "sg_consistent_snapshot_support": "pool",
                "ipv6_support": false,
                "create_share_from_snapshot_support": true,
                "dedupe": false,
                "snapshot_support": true,
                "free_capacity_gb": 1210.0,
                "storage_protocol": "NFS_CIFS",
                "replication_domain": "mirror_domain_01"
            },
            "name": "ostkstoragecontroller01@gamma#gamma_data_pool",
            "pool": "gamma_data_pool",
            "backend": "gamma"
        },
        {
            "host": "ostkstoragecontroller01",
            "capabilities": {
                "pool_name": "theta_data_pool",
                "vendor_name": "Open Source",
                "compression": false,
                "timestamp": "2019-03-28T23:53:42.411802",
                "driver_version": "1.0",
                "replication_type": "readable",
                "mount_snapshot_support": true,
                "share_backend_name": "THETA",
                "driver_handles_share_servers": false,
                "reserved_percentage": 0,
                "total_capacity_gb": 1230.0,
                "revert_to_snapshot_support": true,
                "ipv4_support": true,
                "sg_consistent_snapshot_support": "pool",
                "ipv6_support": false,
                "create_share_from_snapshot_support": true,
                "dedupe": false,
                "snapshot_support": true,
                "free_capacity_gb": 1210.0,
                "storage_protocol": "NFS_CIFS",
                "replication_domain": "mirror_domain_01"
            },
            "name": "ostkstoragecontroller01@theta#theta_data_pool",
            "pool": "theta_data_pool",
            "backend": "theta"
        },
        {
            "host": "ostkstoragecontroller01",
            "capabilities": {
                "pool_name": "lvm-single-pool",
                "vendor_name": "Open Source",
                "compression": false,
                "driver_version": "1.0",
                "snapshot_support": true,
                "replication_type": null,
                "ipv6_support": false,
                "share_backend_name": "Prague",
                "driver_handles_share_servers": false,
                "reserved_percentage": 0,
                "total_capacity_gb": 100.0,
                "revert_to_snapshot_support": true,
                "ipv4_support": true,
                "sg_consistent_snapshot_support": null,
                "timestamp": "2019-03-28T23:53:06.242294",
                "create_share_from_snapshot_support": true,
                "dedupe": false,
                "mount_snapshot_support": true,
                "free_capacity_gb": 100.0,
                "storage_protocol": "NFS_CIFS",
                "replication_domain": null
            },
            "name": "ostkstoragecontroller01@prague#lvm-single-pool",
            "pool": "lvm-single-pool",
            "backend": "prague"
        },
        {
            "host": "ostkstoragecontroller01",
            "capabilities": {
                "pool_name": "delta_data_pool",
                "vendor_name": "Open Source",
                "compression": false,
                "timestamp": "2019-03-28T23:53:48.406598",
                "driver_version": "1.0",
                "replication_type": "readable",
                "mount_snapshot_support": true,
                "share_backend_name": "DELTA",
                "driver_handles_share_servers": false,
                "reserved_percentage": 0,
                "total_capacity_gb": 1230.0,
                "revert_to_snapshot_support": true,
                "ipv4_support": true,
                "sg_consistent_snapshot_support": "pool",
                "ipv6_support": false,
                "create_share_from_snapshot_support": true,
                "dedupe": false,
                "snapshot_support": true,
                "free_capacity_gb": 1210.0,
                "storage_protocol": "NFS_CIFS",
                "replication_domain": "mirror_domain_01"
            },
            "name": "ostkstoragecontroller01@delta#delta_data_pool",
            "pool": "delta_data_pool",
            "backend": "delta"
        }
    ]
}
`

var (
	StoragePoolFake1 = schedulerstats.StoragePool{
		Name: "rbd:cinder.volumes.ssd@cinder.volumes.ssd#cinder.volumes.ssd",
		Capabilities: schedulerstats.Capabilities{
			DriverVersion:     "1.2.0",
			FreeCapacityGB:    64765,
			StorageProtocol:   "ceph",
			TotalCapacityGB:   787947.93,
			VendorName:        "Open Source",
			VolumeBackendName: "cinder.volumes.ssd",
		},
	}

	StoragePoolFake2 = schedulerstats.StoragePool{
		Name: "rbd:cinder.volumes.hdd@cinder.volumes.hdd#cinder.volumes.hdd",
		Capabilities: schedulerstats.Capabilities{
			DriverVersion:     "1.2.0",
			FreeCapacityGB:    0.0,
			StorageProtocol:   "ceph",
			TotalCapacityGB:   math.Inf(1),
			VendorName:        "Open Source",
			VolumeBackendName: "cinder.volumes.hdd",
		},
	}
)

func HandleStoragePoolsListSuccessfully(t *testing.T) {
	testhelper.Mux.HandleFunc("/scheduler-stats/pools", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "GET")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")

		r.ParseForm()
		if r.FormValue("detail") == "true" {
			fmt.Fprintf(w, StoragePoolsListBodyDetail)
		} else {
			fmt.Fprintf(w, StoragePoolsListBody)
		}
	})
}
