/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-30T02:46:34Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type CreatePublicIpInstanceRequest struct {

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// 공인IP설명
PublicIpDescription *string `json:"publicIpDescription,omitempty"`

	// 인터넷라인구분코드
InternetLineTypeCode *string `json:"internetLineTypeCode,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`

	// ZONE번호
ZoneNo *string `json:"zoneNo,omitempty"`
}
