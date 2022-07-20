/*
 * sourcebuild
 *
 * <br/>https://sourcebuild.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcebuild

type EnvPlatformConfigRequest struct {
	Os *EnvPlatformConfigRequestOs `json:"os,omitempty"`

	Runtime *EnvPlatformConfigRequestRuntime `json:"runtime,omitempty"`

	Registry *string `json:"registry,omitempty"`

	Image *string `json:"image,omitempty"`

	Tag *string `json:"tag,omitempty"`
}
