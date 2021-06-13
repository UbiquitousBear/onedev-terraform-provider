package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"regexp"
	"strconv"
)

func validateName(v interface{}, _ string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs
}

func resourceOnedevProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		CustomizeDiff: customdiff.Sequence(
			customdiff.ComputedIf("slug", func(d *schema.ResourceDiff, meta interface{}) bool {
				return d.HasChange("name")
			}),
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateName,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"issuemanagementenabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default: false,
			},
			"forkedfromid": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func getOrDefault(d *schema.ResourceData, field string, defaultVal interface{}) interface{} {
	val, _ := d.GetOk(field)

	if val == nil {
		return defaultVal
	} else {
		return val
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	project := Project{
		ForkedFromId:           getOrDefault(d, "forkedfromid", 0).(int),
		Name:                   d.Get("name").(string),
		Description:            d.Get("description").(string),
		IssueManagementEnabled: getOrDefault(d, "issuemanagementenabled", false).(bool),
	}

	apiClient := m.(*Client)
	createResponse, err := apiClient.NewProject(project)

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(createResponse.Id))

	return nil
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*Client)

	itemId, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}


	item, err := apiClient.GetProject(itemId)
	if err != nil {
		return err
	}
	d.SetId(strconv.Itoa(item.Id))
	d.Set("name", item.Name)
	d.Set("description", item.Description)
	d.Set("forkedfromid", item.ForkedFromId)
	d.Set("issuemanagementenabled", item.IssueManagementEnabled)

	return nil
}

func resourceProjectUpdate (d *schema.ResourceData, m interface{}) error {
	itemId, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	apiClient := m.(*Client)
	project := Project{
		Id: 					itemId,
		ForkedFromId:           d.Get("forkedfromid").(int),
		Name:                   d.Get("name").(string),
		Description:            d.Get("description").(string),
		IssueManagementEnabled: d.Get("issuemanagementenabled").(bool),
	}

	updateResponse, err := apiClient.UpdateProject(project)
	if err != nil {
		return err
	}

	item, err := apiClient.GetProject(updateResponse.Id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(item.Id))
	d.Set("name", item.Name)
	d.Set("description", item.Description)
	d.Set("forkedfromid", item.ForkedFromId)
	d.Set("issuemanagementenabled", item.IssueManagementEnabled)

	return nil

}

func resourceProjectDelete (d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*Client)
	itemId, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	err = apiClient.DeleteProject(itemId)
	if err != nil {
		return err
	}

	return nil
}
