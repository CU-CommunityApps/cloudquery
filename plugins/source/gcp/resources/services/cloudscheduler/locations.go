package cloudscheduler

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/cloudscheduler/v1"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudscheduler_locations",
		Description: `https://cloud.google.com/scheduler/docs/reference/rest/v1/projects.locations#Location`,
		Resolver:    fetchLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudscheduler.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Location{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			Jobs(),
		},
	}
}
