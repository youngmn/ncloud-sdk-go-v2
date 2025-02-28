/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type AccessControlGroupRule struct {

	// ACG번호
AccessControlGroupNo *string `json:"accessControlGroupNo,omitempty"`

	// 프로토콜유형
ProtocolType *CommonCode `json:"protocolType,omitempty"`

	// IP블록
IpBlock *string `json:"ipBlock,omitempty"`

	// 포트범위
PortRange *string `json:"portRange,omitempty"`

	// 접근소스ACG
AccessControlGroupSequence *string `json:"accessControlGroupSequence,omitempty"`

	// ACGRule유형
AccessControlGroupRuleType *CommonCode `json:"accessControlGroupRuleType,omitempty"`

	// ACGRule설명
AccessControlGroupRuleDescription *string `json:"accessControlGroupRuleDescription,omitempty"`
}
