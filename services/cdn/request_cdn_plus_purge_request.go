/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-08-07T06:43:44Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type RequestCdnPlusPurgeRequest struct {

	// CDN인스턴스번호
CdnInstanceNo *string `json:"cdnInstanceNo"`

	// 도메인ID리스트
DomainIdList []*string `json:"domainIdList,omitempty"`

	// 전체퍼지여부
IsWholePurge *bool `json:"isWholePurge"`

	// 전체도메인여부
IsWholeDomain *bool `json:"isWholeDomain"`

	// 대상파일리스트
TargetFileList []*string `json:"targetFileList,omitempty"`

	// 대상디렉토리명
TargetDirectoryName *string `json:"targetDirectoryName,omitempty"`
}
