package computeoptimizer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func EcsServiceRecommendations() *schema.Table {
	tableName := "aws_computeoptimizer_ecs_service_recommendations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_ECSServiceRecommendation.html`,
		Resolver:    fetchEcsServiceRecommendations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "compute-optimizer"),
		Transform:   transformers.TransformWithStruct(&types.ECSServiceRecommendation{}, transformers.WithPrimaryKeys("ServiceArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
		},
	}
}

func fetchEcsServiceRecommendations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	s := c.Services()
	svc := s.Computeoptimizer

	input := computeoptimizer.GetECSServiceRecommendationsInput{
		MaxResults: aws.Int32(1000),
	}
	// No paginator available
	for {
		response, err := svc.GetECSServiceRecommendations(ctx, &input)
		if err != nil {
			return err
		}

		if len(response.Errors) > 0 {
			c.Logger().Error().Str("table", "aws_computeoptimizer_ecs_service_recommendations").Msgf("Errors in response: %v", response.Errors)
		}

		if response.EcsServiceRecommendations != nil {
			res <- response.EcsServiceRecommendations
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}

		input.NextToken = response.NextToken
	}

	return nil
}
