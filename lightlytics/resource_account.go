package lightlytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Account struct {
	_id string
	display_name string
	account_type string
	cloud_account_id string
	cloud_regions []string
	stack_region string
	template_url string
}

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAccountCreate,
		ReadContext:   resourceAccountRead,
		UpdateContext: resourceAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Schema: map[string]*schema.Schema{
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_regions": &schema.Schema{
				Type:     schema.TypeList,
				Elem: 	  &schema.Schema{
				    			Type: schema.TypeString,
						},
				Required: true,
			},
			"stack_region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"template_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lightlytics_collection_token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_auth_token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	account_type := d.Get("account_type").(string)
	display_name := d.Get("display_name").(string)
	cloud_account_id := d.Get("cloud_account_id").(string)
	cloud_regions := d.Get("cloud_regions").([]interface{})
	stack_region := d.Get("stack_region").(string)

	query := `
		mutation CreateAccount($account_type: CloudProvider!, $cloud_account_id: String!, $display_name: String, $cloud_regions: [String], $stack_region: String) {
			createAccount(account: {
				account_type: $account_type,
				cloud_account_id: $cloud_account_id,
				display_name: $display_name,
				cloud_regions: $cloud_regions,
				stack_region: $stack_region
			  })
			{
				_id
				template_url
				external_id
				lightlytics_collection_token
				account_auth_token
			}
	}`

	variables := map[string]interface{}{
        "account_type": account_type,
    	"cloud_account_id": cloud_account_id,
    	"display_name": display_name,
    	"cloud_regions": cloud_regions,
    	"stack_region": stack_region}

	data, err := c.doRequest(query, variables)

	if err != nil {
		return diag.FromErr(err)
	}

    account := data["createAccount"].(map[string]interface{})

	id := account["_id"].(string)

    d.SetId(id)
    d.Set("template_url", account["template_url"])
    d.Set("external_id", account["external_id"])
    d.Set("lightlytics_collection_token", account["lightlytics_collection_token"])
    d.Set("account_auth_token", account["account_auth_token"])

	return diags
}

func resourceAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	accountID := d.Id()

	query := `
		query {
			accounts {
				_id
				account_type
				cloud_account_id
				display_name
				cloud_regions
				stack_region
				template_url
				external_id
				lightlytics_collection_token
				account_auth_token
			}
		}`

	data, err := c.doRequest(query, nil)

	if err != nil {
		return diag.FromErr(err)
	}

	accounts := data["accounts"].([]interface{})

	for _, acc := range accounts {

		account := acc.(map[string]interface{})

		if account["_id"] == accountID {
			d.Set("display_name", account["display_name"])
			d.Set("account_type", account["account_type"])
			d.Set("cloud_account_id", account["cloud_account_id"])
			d.Set("cloud_regions", account["cloud_regions"])
			d.Set("stack_region", account["stack_region"])
			d.Set("template_url", account["template_url"])
			d.Set("external_id", account["external_id"])
			d.Set("lightlytics_collection_token", account["lightlytics_collection_token"])
			d.Set("account_auth_token", account["account_auth_token"])
		}
	}

	return diags
}

func resourceAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	accountID := d.Id()

	if d.HasChange("cloud_regions") || d.HasChange("display_name") {
		query := `
			mutation UpdateAccount($id: ID!, $account: AccountUpdateInput) {
				updateAccount(id: $id, account: $account) {
					_id
				}
			}`

		display_name := d.Get("display_name").(string)
		cloud_regions := d.Get("cloud_regions").([]interface{})

		variables := map[string]interface{}{
			"id": accountID,
			"account": map[string]interface{}{
				"cloud_regions": cloud_regions,
				"display_name": display_name }}

		_, err := c.doRequest(query, variables)

		if err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	accountID := d.Id()

	query := `
		mutation DeleteAccount($id: ID!) {
			deleteAccount(id: $id)
		}`

	variables := map[string]interface{}{
        "id": accountID}

	_, err := c.doRequest(query, variables)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

