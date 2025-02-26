package projects

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/xanzy/go-gitlab"
)

func ProjectBranches() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_project_branches",
		Resolver:  fetchProjectBranches,
		Transform: client.TransformWithStruct(&gitlab.Branch{}),
		Columns: []schema.Column{
			{
				Name:     "base_url",
				Type:     schema.TypeString,
				Resolver: client.ResolveURL,
			},
			{
				Name:     "project_id",
				Type:     schema.TypeInt,
				Resolver: resolveProjectID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
