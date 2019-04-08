/*
Package schedulerstats returns the storage service catalog information from the
shared file systems service. All configured shared file systems backends expose
one or more storage pools to the scheduler service. The scheduler service
compiles this into a storage service catalog, which it exposes via an API that
this module retrieves information from. The storage service catalog is useful
to the scheduler to decide where to place resources such as shares, snapshot
clones, replicas, share groups as they are provisioned. Each pool in the
storage service catalog has "capabilities", some of which are common across all
storage backends, and some of which are unique to some storage backends.


Example Usage:

	listOpts := schedulerstats.ListOpts{
		Detail: true,
	}

	allPages, err := schedulerstats.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allStats, err := schedulerstats.ExtractStoragePools(allPages)
	if err != nil {
		panic(err)
	}

	for _, stat := range allStats {
		fmt.Printf("%+v\n", stat)
	}
*/
package schedulerstats
