/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-08-07T06:55:28Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type Raid struct {

	// RAID구분이름
RaidTypeName *string `json:"raidTypeName,omitempty"`

	// RAID이름
RaidName *string `json:"raidName,omitempty"`
}
