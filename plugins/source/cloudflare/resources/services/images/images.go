package images

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_images",
		Resolver:  fetchImages,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&cloudflare.Image{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
