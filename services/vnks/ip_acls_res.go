/*
 * vnks
 *
 * <br/>https://nks.apigw.ntruss.com/vnks/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnks

type IpAclsRes struct {
	DefaultAction *string `json:"defaultAction"`

	Entries []*IpAclsEntriesRes `json:"entries,omitempty"`
}
