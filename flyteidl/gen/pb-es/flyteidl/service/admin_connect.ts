// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts"
// @generated from file flyteidl/service/admin.proto (package flyteidl.service, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Task, TaskCreateRequest, TaskCreateResponse, TaskList } from "../admin/task_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { NamedEntity, NamedEntityGetRequest, NamedEntityIdentifierList, NamedEntityIdentifierListRequest, NamedEntityList, NamedEntityListRequest, NamedEntityUpdateRequest, NamedEntityUpdateResponse, ObjectGetRequest, ResourceListRequest } from "../admin/common_pb.js";
import { Workflow, WorkflowCreateRequest, WorkflowCreateResponse, WorkflowList } from "../admin/workflow_pb.js";
import { ActiveLaunchPlanListRequest, ActiveLaunchPlanRequest, LaunchPlan, LaunchPlanCreateRequest, LaunchPlanCreateResponse, LaunchPlanList, LaunchPlanUpdateRequest, LaunchPlanUpdateResponse } from "../admin/launch_plan_pb.js";
import { Execution, ExecutionCountRequest, ExecutionCountResponse, ExecutionCreateRequest, ExecutionCreateResponse, ExecutionList, ExecutionRecoverRequest, ExecutionRelaunchRequest, ExecutionTerminateRequest, ExecutionTerminateResponse, ExecutionUpdateRequest, ExecutionUpdateResponse, WorkflowExecutionGetDataRequest, WorkflowExecutionGetDataResponse, WorkflowExecutionGetMetricsRequest, WorkflowExecutionGetMetricsResponse, WorkflowExecutionGetRequest } from "../admin/execution_pb.js";
import { DynamicNodeWorkflowResponse, GetDynamicNodeWorkflowRequest, NodeExecution, NodeExecutionForTaskListRequest, NodeExecutionGetDataRequest, NodeExecutionGetDataResponse, NodeExecutionGetRequest, NodeExecutionList, NodeExecutionListRequest } from "../admin/node_execution_pb.js";
import { Project, ProjectListRequest, ProjectRegisterRequest, ProjectRegisterResponse, Projects, ProjectUpdateResponse } from "../admin/project_pb.js";
import { NodeExecutionEventRequest, NodeExecutionEventResponse, TaskExecutionEventRequest, TaskExecutionEventResponse, WorkflowExecutionEventRequest, WorkflowExecutionEventResponse } from "../admin/event_pb.js";
import { TaskExecution, TaskExecutionGetDataRequest, TaskExecutionGetDataResponse, TaskExecutionGetRequest, TaskExecutionList, TaskExecutionListRequest } from "../admin/task_execution_pb.js";
import { ProjectDomainAttributesDeleteRequest, ProjectDomainAttributesDeleteResponse, ProjectDomainAttributesGetRequest, ProjectDomainAttributesGetResponse, ProjectDomainAttributesUpdateRequest, ProjectDomainAttributesUpdateResponse } from "../admin/project_domain_attributes_pb.js";
import { ProjectAttributesDeleteRequest, ProjectAttributesDeleteResponse, ProjectAttributesGetRequest, ProjectAttributesGetResponse, ProjectAttributesUpdateRequest, ProjectAttributesUpdateResponse } from "../admin/project_attributes_pb.js";
import { WorkflowAttributesDeleteRequest, WorkflowAttributesDeleteResponse, WorkflowAttributesGetRequest, WorkflowAttributesGetResponse, WorkflowAttributesUpdateRequest, WorkflowAttributesUpdateResponse } from "../admin/workflow_attributes_pb.js";
import { ListMatchableAttributesRequest, ListMatchableAttributesResponse } from "../admin/matchable_resource_pb.js";
import { GetVersionRequest, GetVersionResponse } from "../admin/version_pb.js";
import { DescriptionEntity, DescriptionEntityList, DescriptionEntityListRequest } from "../admin/description_entity_pb.js";

/**
 * The following defines an RPC service that is also served over HTTP via grpc-gateway.
 * Standard response codes for both are defined here: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
 *
 * @generated from service flyteidl.service.AdminService
 */
export const AdminService = {
  typeName: "flyteidl.service.AdminService",
  methods: {
    /**
     * Create and upload a :ref:`ref_flyteidl.admin.Task` definition
     *
     * @generated from rpc flyteidl.service.AdminService.CreateTask
     */
    createTask: {
      name: "CreateTask",
      I: TaskCreateRequest,
      O: TaskCreateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a :ref:`ref_flyteidl.admin.Task` definition.
     *
     * @generated from rpc flyteidl.service.AdminService.GetTask
     */
    getTask: {
      name: "GetTask",
      I: ObjectGetRequest,
      O: Task,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a list of :ref:`ref_flyteidl.admin.NamedEntityIdentifier` of task objects.
     *
     * @generated from rpc flyteidl.service.AdminService.ListTaskIds
     */
    listTaskIds: {
      name: "ListTaskIds",
      I: NamedEntityIdentifierListRequest,
      O: NamedEntityIdentifierList,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a list of :ref:`ref_flyteidl.admin.Task` definitions.
     *
     * @generated from rpc flyteidl.service.AdminService.ListTasks
     */
    listTasks: {
      name: "ListTasks",
      I: ResourceListRequest,
      O: TaskList,
      kind: MethodKind.Unary,
    },
    /**
     * Create and upload a :ref:`ref_flyteidl.admin.Workflow` definition
     *
     * @generated from rpc flyteidl.service.AdminService.CreateWorkflow
     */
    createWorkflow: {
      name: "CreateWorkflow",
      I: WorkflowCreateRequest,
      O: WorkflowCreateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a :ref:`ref_flyteidl.admin.Workflow` definition.
     *
     * @generated from rpc flyteidl.service.AdminService.GetWorkflow
     */
    getWorkflow: {
      name: "GetWorkflow",
      I: ObjectGetRequest,
      O: Workflow,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a list of :ref:`ref_flyteidl.admin.NamedEntityIdentifier` of workflow objects.
     *
     * @generated from rpc flyteidl.service.AdminService.ListWorkflowIds
     */
    listWorkflowIds: {
      name: "ListWorkflowIds",
      I: NamedEntityIdentifierListRequest,
      O: NamedEntityIdentifierList,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a list of :ref:`ref_flyteidl.admin.Workflow` definitions.
     *
     * @generated from rpc flyteidl.service.AdminService.ListWorkflows
     */
    listWorkflows: {
      name: "ListWorkflows",
      I: ResourceListRequest,
      O: WorkflowList,
      kind: MethodKind.Unary,
    },
    /**
     * Create and upload a :ref:`ref_flyteidl.admin.LaunchPlan` definition
     *
     * @generated from rpc flyteidl.service.AdminService.CreateLaunchPlan
     */
    createLaunchPlan: {
      name: "CreateLaunchPlan",
      I: LaunchPlanCreateRequest,
      O: LaunchPlanCreateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a :ref:`ref_flyteidl.admin.LaunchPlan` definition.
     *
     * @generated from rpc flyteidl.service.AdminService.GetLaunchPlan
     */
    getLaunchPlan: {
      name: "GetLaunchPlan",
      I: ObjectGetRequest,
      O: LaunchPlan,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch the active version of a :ref:`ref_flyteidl.admin.LaunchPlan`.
     *
     * @generated from rpc flyteidl.service.AdminService.GetActiveLaunchPlan
     */
    getActiveLaunchPlan: {
      name: "GetActiveLaunchPlan",
      I: ActiveLaunchPlanRequest,
      O: LaunchPlan,
      kind: MethodKind.Unary,
    },
    /**
     * List active versions of :ref:`ref_flyteidl.admin.LaunchPlan`.
     *
     * @generated from rpc flyteidl.service.AdminService.ListActiveLaunchPlans
     */
    listActiveLaunchPlans: {
      name: "ListActiveLaunchPlans",
      I: ActiveLaunchPlanListRequest,
      O: LaunchPlanList,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a list of :ref:`ref_flyteidl.admin.NamedEntityIdentifier` of launch plan objects.
     *
     * @generated from rpc flyteidl.service.AdminService.ListLaunchPlanIds
     */
    listLaunchPlanIds: {
      name: "ListLaunchPlanIds",
      I: NamedEntityIdentifierListRequest,
      O: NamedEntityIdentifierList,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a list of :ref:`ref_flyteidl.admin.LaunchPlan` definitions.
     *
     * @generated from rpc flyteidl.service.AdminService.ListLaunchPlans
     */
    listLaunchPlans: {
      name: "ListLaunchPlans",
      I: ResourceListRequest,
      O: LaunchPlanList,
      kind: MethodKind.Unary,
    },
    /**
     * Updates the status of a registered :ref:`ref_flyteidl.admin.LaunchPlan`.
     *
     * @generated from rpc flyteidl.service.AdminService.UpdateLaunchPlan
     */
    updateLaunchPlan: {
      name: "UpdateLaunchPlan",
      I: LaunchPlanUpdateRequest,
      O: LaunchPlanUpdateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Triggers the creation of a :ref:`ref_flyteidl.admin.Execution`
     *
     * @generated from rpc flyteidl.service.AdminService.CreateExecution
     */
    createExecution: {
      name: "CreateExecution",
      I: ExecutionCreateRequest,
      O: ExecutionCreateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Triggers the creation of an identical :ref:`ref_flyteidl.admin.Execution`
     *
     * @generated from rpc flyteidl.service.AdminService.RelaunchExecution
     */
    relaunchExecution: {
      name: "RelaunchExecution",
      I: ExecutionRelaunchRequest,
      O: ExecutionCreateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Recreates a previously-run workflow execution that will only start executing from the last known failure point.
     * In Recover mode, users cannot change any input parameters or update the version of the execution.
     * This is extremely useful to recover from system errors and byzantine faults like - Loss of K8s cluster, bugs in platform or instability, machine failures,
     * downstream system failures (downstream services), or simply to recover executions that failed because of retry exhaustion and should complete if tried again.
     * See :ref:`ref_flyteidl.admin.ExecutionRecoverRequest` for more details.
     *
     * @generated from rpc flyteidl.service.AdminService.RecoverExecution
     */
    recoverExecution: {
      name: "RecoverExecution",
      I: ExecutionRecoverRequest,
      O: ExecutionCreateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches a :ref:`ref_flyteidl.admin.Execution`.
     *
     * @generated from rpc flyteidl.service.AdminService.GetExecution
     */
    getExecution: {
      name: "GetExecution",
      I: WorkflowExecutionGetRequest,
      O: Execution,
      kind: MethodKind.Unary,
    },
    /**
     * Update execution belonging to project domain :ref:`ref_flyteidl.admin.Execution`.
     *
     * @generated from rpc flyteidl.service.AdminService.UpdateExecution
     */
    updateExecution: {
      name: "UpdateExecution",
      I: ExecutionUpdateRequest,
      O: ExecutionUpdateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches input and output data for a :ref:`ref_flyteidl.admin.Execution`.
     *
     * @generated from rpc flyteidl.service.AdminService.GetExecutionData
     */
    getExecutionData: {
      name: "GetExecutionData",
      I: WorkflowExecutionGetDataRequest,
      O: WorkflowExecutionGetDataResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a list of :ref:`ref_flyteidl.admin.Execution`.
     *
     * @generated from rpc flyteidl.service.AdminService.ListExecutions
     */
    listExecutions: {
      name: "ListExecutions",
      I: ResourceListRequest,
      O: ExecutionList,
      kind: MethodKind.Unary,
    },
    /**
     * Terminates an in-progress :ref:`ref_flyteidl.admin.Execution`.
     *
     * @generated from rpc flyteidl.service.AdminService.TerminateExecution
     */
    terminateExecution: {
      name: "TerminateExecution",
      I: ExecutionTerminateRequest,
      O: ExecutionTerminateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches a :ref:`ref_flyteidl.admin.NodeExecution`.
     *
     * @generated from rpc flyteidl.service.AdminService.GetNodeExecution
     */
    getNodeExecution: {
      name: "GetNodeExecution",
      I: NodeExecutionGetRequest,
      O: NodeExecution,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches a :ref:`ref_flyteidl.admin.DynamicNodeWorkflowResponse`.
     *
     * @generated from rpc flyteidl.service.AdminService.GetDynamicNodeWorkflow
     */
    getDynamicNodeWorkflow: {
      name: "GetDynamicNodeWorkflow",
      I: GetDynamicNodeWorkflowRequest,
      O: DynamicNodeWorkflowResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a list of :ref:`ref_flyteidl.admin.NodeExecution`.
     *
     * @generated from rpc flyteidl.service.AdminService.ListNodeExecutions
     */
    listNodeExecutions: {
      name: "ListNodeExecutions",
      I: NodeExecutionListRequest,
      O: NodeExecutionList,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a list of :ref:`ref_flyteidl.admin.NodeExecution` launched by the reference :ref:`ref_flyteidl.admin.TaskExecution`.
     *
     * @generated from rpc flyteidl.service.AdminService.ListNodeExecutionsForTask
     */
    listNodeExecutionsForTask: {
      name: "ListNodeExecutionsForTask",
      I: NodeExecutionForTaskListRequest,
      O: NodeExecutionList,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches input and output data for a :ref:`ref_flyteidl.admin.NodeExecution`.
     *
     * @generated from rpc flyteidl.service.AdminService.GetNodeExecutionData
     */
    getNodeExecutionData: {
      name: "GetNodeExecutionData",
      I: NodeExecutionGetDataRequest,
      O: NodeExecutionGetDataResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Registers a :ref:`ref_flyteidl.admin.Project` with the Flyte deployment.
     *
     * @generated from rpc flyteidl.service.AdminService.RegisterProject
     */
    registerProject: {
      name: "RegisterProject",
      I: ProjectRegisterRequest,
      O: ProjectRegisterResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Updates an existing :ref:`ref_flyteidl.admin.Project`
     * flyteidl.admin.Project should be passed but the domains property should be empty;
     * it will be ignored in the handler as domains cannot be updated via this API.
     *
     * @generated from rpc flyteidl.service.AdminService.UpdateProject
     */
    updateProject: {
      name: "UpdateProject",
      I: Project,
      O: ProjectUpdateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches a list of :ref:`ref_flyteidl.admin.Project`
     *
     * @generated from rpc flyteidl.service.AdminService.ListProjects
     */
    listProjects: {
      name: "ListProjects",
      I: ProjectListRequest,
      O: Projects,
      kind: MethodKind.Unary,
    },
    /**
     * Indicates a :ref:`ref_flyteidl.event.WorkflowExecutionEvent` has occurred.
     *
     * @generated from rpc flyteidl.service.AdminService.CreateWorkflowEvent
     */
    createWorkflowEvent: {
      name: "CreateWorkflowEvent",
      I: WorkflowExecutionEventRequest,
      O: WorkflowExecutionEventResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Indicates a :ref:`ref_flyteidl.event.NodeExecutionEvent` has occurred.
     *
     * @generated from rpc flyteidl.service.AdminService.CreateNodeEvent
     */
    createNodeEvent: {
      name: "CreateNodeEvent",
      I: NodeExecutionEventRequest,
      O: NodeExecutionEventResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Indicates a :ref:`ref_flyteidl.event.TaskExecutionEvent` has occurred.
     *
     * @generated from rpc flyteidl.service.AdminService.CreateTaskEvent
     */
    createTaskEvent: {
      name: "CreateTaskEvent",
      I: TaskExecutionEventRequest,
      O: TaskExecutionEventResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches a :ref:`ref_flyteidl.admin.TaskExecution`.
     *
     * @generated from rpc flyteidl.service.AdminService.GetTaskExecution
     */
    getTaskExecution: {
      name: "GetTaskExecution",
      I: TaskExecutionGetRequest,
      O: TaskExecution,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches a list of :ref:`ref_flyteidl.admin.TaskExecution`.
     *
     * @generated from rpc flyteidl.service.AdminService.ListTaskExecutions
     */
    listTaskExecutions: {
      name: "ListTaskExecutions",
      I: TaskExecutionListRequest,
      O: TaskExecutionList,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches input and output data for a :ref:`ref_flyteidl.admin.TaskExecution`.
     *
     * @generated from rpc flyteidl.service.AdminService.GetTaskExecutionData
     */
    getTaskExecutionData: {
      name: "GetTaskExecutionData",
      I: TaskExecutionGetDataRequest,
      O: TaskExecutionGetDataResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Creates or updates custom :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration` for a project and domain.
     *
     * @generated from rpc flyteidl.service.AdminService.UpdateProjectDomainAttributes
     */
    updateProjectDomainAttributes: {
      name: "UpdateProjectDomainAttributes",
      I: ProjectDomainAttributesUpdateRequest,
      O: ProjectDomainAttributesUpdateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches custom :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration` for a project and domain.
     *
     * @generated from rpc flyteidl.service.AdminService.GetProjectDomainAttributes
     */
    getProjectDomainAttributes: {
      name: "GetProjectDomainAttributes",
      I: ProjectDomainAttributesGetRequest,
      O: ProjectDomainAttributesGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Deletes custom :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration` for a project and domain.
     *
     * @generated from rpc flyteidl.service.AdminService.DeleteProjectDomainAttributes
     */
    deleteProjectDomainAttributes: {
      name: "DeleteProjectDomainAttributes",
      I: ProjectDomainAttributesDeleteRequest,
      O: ProjectDomainAttributesDeleteResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Creates or updates custom :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration` at the project level
     *
     * @generated from rpc flyteidl.service.AdminService.UpdateProjectAttributes
     */
    updateProjectAttributes: {
      name: "UpdateProjectAttributes",
      I: ProjectAttributesUpdateRequest,
      O: ProjectAttributesUpdateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches custom :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration` for a project and domain.
     *
     * @generated from rpc flyteidl.service.AdminService.GetProjectAttributes
     */
    getProjectAttributes: {
      name: "GetProjectAttributes",
      I: ProjectAttributesGetRequest,
      O: ProjectAttributesGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Deletes custom :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration` for a project and domain.
     *
     * @generated from rpc flyteidl.service.AdminService.DeleteProjectAttributes
     */
    deleteProjectAttributes: {
      name: "DeleteProjectAttributes",
      I: ProjectAttributesDeleteRequest,
      O: ProjectAttributesDeleteResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Creates or updates custom :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration` for a project, domain and workflow.
     *
     * @generated from rpc flyteidl.service.AdminService.UpdateWorkflowAttributes
     */
    updateWorkflowAttributes: {
      name: "UpdateWorkflowAttributes",
      I: WorkflowAttributesUpdateRequest,
      O: WorkflowAttributesUpdateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches custom :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration` for a project, domain and workflow.
     *
     * @generated from rpc flyteidl.service.AdminService.GetWorkflowAttributes
     */
    getWorkflowAttributes: {
      name: "GetWorkflowAttributes",
      I: WorkflowAttributesGetRequest,
      O: WorkflowAttributesGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Deletes custom :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration` for a project, domain and workflow.
     *
     * @generated from rpc flyteidl.service.AdminService.DeleteWorkflowAttributes
     */
    deleteWorkflowAttributes: {
      name: "DeleteWorkflowAttributes",
      I: WorkflowAttributesDeleteRequest,
      O: WorkflowAttributesDeleteResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Lists custom :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration` for a specific resource type.
     *
     * @generated from rpc flyteidl.service.AdminService.ListMatchableAttributes
     */
    listMatchableAttributes: {
      name: "ListMatchableAttributes",
      I: ListMatchableAttributesRequest,
      O: ListMatchableAttributesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Returns a list of :ref:`ref_flyteidl.admin.NamedEntity` objects.
     *
     * @generated from rpc flyteidl.service.AdminService.ListNamedEntities
     */
    listNamedEntities: {
      name: "ListNamedEntities",
      I: NamedEntityListRequest,
      O: NamedEntityList,
      kind: MethodKind.Unary,
    },
    /**
     * Returns a :ref:`ref_flyteidl.admin.NamedEntity` object.
     *
     * @generated from rpc flyteidl.service.AdminService.GetNamedEntity
     */
    getNamedEntity: {
      name: "GetNamedEntity",
      I: NamedEntityGetRequest,
      O: NamedEntity,
      kind: MethodKind.Unary,
    },
    /**
     * Updates a :ref:`ref_flyteidl.admin.NamedEntity` object.
     *
     * @generated from rpc flyteidl.service.AdminService.UpdateNamedEntity
     */
    updateNamedEntity: {
      name: "UpdateNamedEntity",
      I: NamedEntityUpdateRequest,
      O: NamedEntityUpdateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc flyteidl.service.AdminService.GetVersion
     */
    getVersion: {
      name: "GetVersion",
      I: GetVersionRequest,
      O: GetVersionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a :ref:`ref_flyteidl.admin.DescriptionEntity` object.
     *
     * @generated from rpc flyteidl.service.AdminService.GetDescriptionEntity
     */
    getDescriptionEntity: {
      name: "GetDescriptionEntity",
      I: ObjectGetRequest,
      O: DescriptionEntity,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch a list of :ref:`ref_flyteidl.admin.DescriptionEntity` definitions.
     *
     * @generated from rpc flyteidl.service.AdminService.ListDescriptionEntities
     */
    listDescriptionEntities: {
      name: "ListDescriptionEntities",
      I: DescriptionEntityListRequest,
      O: DescriptionEntityList,
      kind: MethodKind.Unary,
    },
    /**
     * Fetches runtime metrics for a :ref:`ref_flyteidl.admin.Execution`.
     *
     * @generated from rpc flyteidl.service.AdminService.GetExecutionMetrics
     */
    getExecutionMetrics: {
      name: "GetExecutionMetrics",
      I: WorkflowExecutionGetMetricsRequest,
      O: WorkflowExecutionGetMetricsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Fetch the count of :ref:`ref_flyteidl.admin.Execution`.
     *
     * @generated from rpc flyteidl.service.AdminService.GetExecutionCount
     */
    getExecutionCount: {
      name: "GetExecutionCount",
      I: ExecutionCountRequest,
      O: ExecutionCountResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

