package schedulerstats

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToStoragePoolsListQuery() (string, error)
}

// ListOpts are query parameters that help filter the data
type ListOpts struct {
	// Host name to filter matching pools by
	Host string `q:"host"`

    // Backend name to filter matching pools by
	Backend string `q:"backend"`

	// Pool name to filter matching pools by
	Pool string `q:"pool"`

	// Share type name or ID to filter matching pools by
	ShareType string `q:"share_type"`
}

// ToStoragePoolsListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToStoragePoolsListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List makes a request against the API to list storage pool information.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToStoragePoolsListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return StoragePoolPage{pagination.SinglePageBase(r)}
	})
}

// List Details makes a request against the API to retrieve detailed storage
// pool information.
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listDetailURL(client)
	if opts != nil {
		query, err := opts.ToStoragePoolsListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return StoragePoolPage{pagination.SinglePageBase(r)}
	})
}
