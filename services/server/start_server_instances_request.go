/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-30T02:46:34Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type StartServerInstancesRequest struct {

	// 서버인스턴스번호리스트
ServerInstanceNoList []*string `json:"serverInstanceNoList"`
}
