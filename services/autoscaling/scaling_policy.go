/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-08-07T06:47:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type ScalingPolicy struct {

PolicyName *string `json:"policyName,omitempty"`

AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

AdjustmentType *CommonCode `json:"adjustmentType,omitempty"`

ScalingAdjustment *int32 `json:"scalingAdjustment,omitempty"`

Cooldown *int32 `json:"cooldown,omitempty"`

MinAdjustmentStep *int32 `json:"minAdjustmentStep,omitempty"`
}
