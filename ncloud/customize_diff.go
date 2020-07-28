package ncloud

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ncloudVpcCommonCustomizeDiff(diff *schema.ResourceDiff, v interface{}) error {
	if diff.HasChange("name") {
		old, new := diff.GetChange("name")
		if len(old.(string)) > 0 {
			return fmt.Errorf("Change 'name' is not support, Please set `name` as a old value = [%s -> %s]", new, old)
		}
	}

	if diff.HasChange("description") {
		old, new := diff.GetChange("description")
		if len(old.(string)) > 0 {
			return fmt.Errorf("Change 'description' is not support, Please set `description` as a old value = [%s -> %s]", new, old)
		}
	}

	return nil
}
