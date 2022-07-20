/*
 * sourcebuild
 *
 * <br/>https://sourcebuild.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcebuild

type GetProjectDetailResponseEnv struct {
	Compute *GetProjectDetailResponseEnvCompute `json:"compute,omitempty"`

	Platform *GetProjectDetailResponseEnvPlatform `json:"platform,omitempty"`

	Docker *GetProjectDetailResponseEnvDocker `json:"docker,omitempty"`

	Timeout *int32 `json:"timeout,omitempty"`

	EnvVars []*ProjectEnvEnvVars `json:"envVars,omitempty"`
}
