package elasticsearch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildElasticSearchVersions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticsearchserviceClient(ctrl)

	var versions []string
	if err := faker.FakeObject(&versions); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListElasticsearchVersions(gomock.Any(), gomock.Any()).Return(
		&elasticsearchservice.ListElasticsearchVersionsOutput{
			ElasticsearchVersions: versions,
		},
		nil,
	)

	var instanceTypes []types.ESPartitionInstanceType
	if err := faker.FakeObject(&versions); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListElasticsearchInstanceTypes(gomock.Any(), gomock.Any()).Return(
		&elasticsearchservice.ListElasticsearchInstanceTypesOutput{
			ElasticsearchInstanceTypes: instanceTypes,
		},
		nil,
	)

	return client.Services{Elasticsearchservice: m}
}

func TestElasticSearchVersions(t *testing.T) {
	client.AwsMockTestHelper(t, Versions(), buildElasticSearchVersions, client.TestOptions{})
}
