/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Holds metadata around how a task was executed. TODO(katrogan): Extend to include freeform fields (https://github.com/flyteorg/flyte/issues/325).
type EventTaskExecutionMetadata struct {
	InstanceClass *TaskExecutionMetadataInstanceClass `json:"instance_class,omitempty"`
}
