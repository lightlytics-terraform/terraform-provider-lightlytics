package lightlytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type KubernetesAccount struct {
	_id string
	display_name string
	status string
	collection_token string
	creation_date string
}

func kubernetesResourceAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: kubernetesResourceAccountCreate,
		ReadContext:   kubernetesResourceAccountRead,
		UpdateContext: kubernetesResourceAccountUpdate,
		DeleteContext: kubernetesResourceAccountDelete,
		Schema: map[string]*schema.Schema{
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"eks_arn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"collection_token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_date": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func kubernetesResourceAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	display_name := d.Get("display_name").(string)
	eks_arn := d.Get("eks_arn").(string)

	query := `
		mutation CreateKubernetes($display_name: String, $eks_arn: String) {
			createKubernetes(kubernetes: {
				display_name: $display_name,
				eks_arn: $eks_arn,
			  })
			{
				_id
				status
				collection_token
				creation_date
			}
	}`

	variables := map[string]interface{}{
    	"display_name": display_name,
    	"eks_arn": eks_arn}

	data, err := c.doRequest(query, variables)

	if err != nil {
		return diag.FromErr(err)
	}

	kubernetesaccount := data["createKubernetes"].(map[string]interface{})

	id := kubernetesaccount["_id"].(string)

    d.SetId(id)
    d.Set("status", kubernetesaccount["status"])
    d.Set("collection_token", kubernetesaccount["collection_token"])
    d.Set("creation_date", kubernetesaccount["creation_date"])

	return diags
}

func kubernetesResourceAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	kubernetesAccountID := d.Id()

	query := `
		query {
			kubernetes {
				_id
				display_name
				status
				collection_token
				creation_date
			}
		}`

	data, err := c.doRequest(query, nil)

	if err != nil {
		return diag.FromErr(err)
	}

	kubernetesAccounts := data["kubernetes"].([]interface{})

	for _, acc := range kubernetesAccounts {

		kubernetesAccount := acc.(map[string]interface{})

		if kubernetesAccount["_id"] == kubernetesAccountID {
			d.Set("display_name", kubernetesAccount["display_name"])
			d.Set("status", kubernetesAccount["status"])
			d.Set("collection_token", kubernetesAccount["collection_token"])
			d.Set("creation_date", kubernetesAccount["creation_date"])
		}
	}

	return diags
}

func kubernetesResourceAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	kubernetesAccountID := d.Id()

	if d.HasChange("display_name") {
		query := `
			mutation UpdateKubernetes($id: ID!, $kubernetes: EditKubernetesInput) {
				updateKubernetes(id: $id, kubernetes: $kubernetes) {
					_id
				}
			}`

		display_name := d.Get("display_name").(string)

		variables := map[string]interface{}{
			"id": kubernetesAccountID,
			"kubernetes": map[string]interface{}{
			"display_name": display_name }}

		_, err := c.doRequest(query, variables)

		if err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func kubernetesResourceAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	kubernetesAccountID := d.Id()

	query := `
		mutation DeleteKubernetes($id: ID!) {
			deleteKubernetes(id: $id)
		}`

	variables := map[string]interface{}{
        "id": kubernetesAccountID}

	_, err := c.doRequest(query, variables)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

