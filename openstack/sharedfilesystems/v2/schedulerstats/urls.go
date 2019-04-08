package schedulerstats

import "github.com/gophercloud/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "pools")
}

func listDetailURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "pools", "detail")
}
