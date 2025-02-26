package elastic

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Monitors() *schema.Table {
	return &schema.Table{
		Name:        "azure_elastic_monitors",
		Resolver:    fetchMonitors,
		Description: "https://learn.microsoft.com/en-us/rest/api/elastic/monitors/list?tabs=HTTP#elasticmonitorresource",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_elastic_monitors", client.Namespacemicrosoft_elastic),
		Transform:   transformers.TransformWithStruct(&armelastic.MonitorResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchMonitors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armelastic.NewMonitorsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
