/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-30T02:46:34Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type RecreateServerInstanceRequest struct {

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// 서버인스턴스이름
ServerInstanceName *string `json:"serverInstanceName,omitempty"`

	// 서버이미지상품코드
ServerImageProductCode *string `json:"serverImageProductCode,omitempty"`

	// 사용자데이터
UserData *string `json:"userData,omitempty"`
}
