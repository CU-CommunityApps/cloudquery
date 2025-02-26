package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func pipelineJobs() *schema.Table {
	return &schema.Table{
		Name:        "aws_elastictranscoder_pipeline_jobs",
		Description: `https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-jobs-by-pipeline.html`,
		Resolver:    fetchElastictranscoderPipelineJobs,
		Transform:   transformers.TransformWithStruct(&types.Job{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchElastictranscoderPipelineJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Elastictranscoder

	p := elastictranscoder.NewListJobsByPipelinePaginator(
		svc,
		&elastictranscoder.ListJobsByPipelineInput{
			PipelineId: aws.String(parent.Get("id").String()),
		},
	)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- response.Jobs
	}

	return nil
}
