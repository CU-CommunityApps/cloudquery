package cloudiot

import (
	pb "cloud.google.com/go/iot/apiv1/iotpb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Devices() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudiot_devices",
		Description: `https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices#Device`,
		Resolver:    fetchDevices,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudiot.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Device{}, transformers.WithPrimaryKeys("Name")),
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
			DeviceConfigs(), DeviceStates(),
		},
	}
}
