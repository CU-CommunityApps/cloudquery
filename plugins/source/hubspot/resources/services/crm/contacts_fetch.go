package crm

import (
	"context"

	"github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/contacts"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func fetchContacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := contacts.NewConfiguration()

	hubspotClient := contacts.NewAPIClient(config)
	cqClient := meta.(*client.Client)

	var after string
	for {
		if err := cqClient.RateLimiter.Wait(ctx); err != nil {
			return nil
		}

		req := hubspotClient.BasicApi.GetPage(hubspot.WithAuthorizer(ctx, cqClient.Authorizer)).Limit(client.DefaultPageSize)

		if len(after) > 0 {
			req = req.After(after)
		}
		out, _, err := req.Execute()
		if err != nil {
			return err
		}

		res <- out.Results

		if !out.HasPaging() {
			break
		}
		paging := out.GetPaging()
		if !paging.HasNext() {
			break
		}
		next := paging.GetNext()
		after = next.After
		if after == "" {
			break
		}
	}

	return nil
}
