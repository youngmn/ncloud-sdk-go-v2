/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type NasVolumeInstance struct {

	// NAS볼륨인스턴스번호
NasVolumeInstanceNo *string `json:"nasVolumeInstanceNo,omitempty"`

	// NAS볼륨인스턴스상태
NasVolumeInstanceStatus *CommonCode `json:"nasVolumeInstanceStatus,omitempty"`

	// NAS볼륨인스턴스OP
NasVolumeInstanceOperation *CommonCode `json:"nasVolumeInstanceOperation,omitempty"`

	// 볼륨인스턴스상태명
NasVolumeInstanceStatusName *string `json:"nasVolumeInstanceStatusName,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// NAS볼륨인스턴스설명
NasVolumeInstanceDescription *string `json:"nasVolumeInstanceDescription,omitempty"`

	// 마운트정보
MountInformation *string `json:"mountInformation,omitempty"`

	// 볼륨할당프로토콜구분
VolumeAllotmentProtocolType *CommonCode `json:"volumeAllotmentProtocolType,omitempty"`

	// 볼륨명
VolumeName *string `json:"volumeName,omitempty"`

	// 볼륨총사이즈
VolumeTotalSize *int64 `json:"volumeTotalSize,omitempty"`

	// 볼륨사이즈
VolumeSize *int64 `json:"volumeSize,omitempty"`

	// 볼륨사용사이즈
VolumeUseSize *int64 `json:"volumeUseSize,omitempty"`

	// 볼륨사용비율
VolumeUseRatio *float32 `json:"volumeUseRatio,omitempty"`

	// 스냅샷볼륨설정비율
SnapshotVolumeConfigurationRatio *float32 `json:"snapshotVolumeConfigurationRatio,omitempty"`

	// 스냅샷볼륨설정기간구분
SnapshotVolumeConfigPeriodType *CommonCode `json:"snapshotVolumeConfigPeriodType,omitempty"`

	// 스냅샷볼륨설정시간
SnapshotVolumeConfigTime *int32 `json:"snapshotVolumeConfigTime,omitempty"`

	// 스냅샷사이즈
SnapshotVolumeSize *int64 `json:"snapshotVolumeSize,omitempty"`

	// 스냅사용사이즈
SnapshotVolumeUseSize *int64 `json:"snapshotVolumeUseSize,omitempty"`

	// 스냅샷사용비율
SnapshotVolumeUseRatio *float32 `json:"snapshotVolumeUseRatio,omitempty"`

	// 스냅샷설정여부
IsSnapshotConfiguration *bool `json:"isSnapshotConfiguration,omitempty"`

	// 이벤트설정여부
IsEventConfiguration *bool `json:"isEventConfiguration,omitempty"`

	// 리전
Region *Region `json:"region,omitempty"`

	// ZONE
Zone *Zone `json:"zone,omitempty"`

	// 반납보호여부
IsReturnProtection *bool `json:"isReturnProtection,omitempty"`

	// NAS볼륨커스텀IP리스트
NasVolumeInstanceCustomIpList []*NasVolumeInstanceCustomIp `json:"nasVolumeInstanceCustomIpList,omitempty"`

	// NAS볼륨서버인스턴스리스트
NasVolumeServerInstanceList []*ServerInstance `json:"nasVolumeServerInstanceList,omitempty"`
}
