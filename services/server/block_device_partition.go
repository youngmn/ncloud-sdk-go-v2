/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type BlockDevicePartition struct {

	// 마운트포인트
MountPoint *string `json:"mountPoint,omitempty"`

	// 파티션사이즈
PartitionSize *string `json:"partitionSize,omitempty"`
}
