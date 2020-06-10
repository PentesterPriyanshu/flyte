// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/workflow.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Failure Handling Strategy
type WorkflowMetadata_OnFailurePolicy int32

const (
	// FAIL_IMMEDIATELY instructs the system to fail as soon as a node fails in the workflow. It'll automatically
	// abort all currently running nodes and clean up resources before finally marking the workflow executions as
	// failed.
	WorkflowMetadata_FAIL_IMMEDIATELY WorkflowMetadata_OnFailurePolicy = 0
	// FAIL_AFTER_EXECUTABLE_NODES_COMPLETE instructs the system to make as much progress as it can. The system will
	// not alter the dependencies of the execution graph so any node that depend on the failed node will not be run.
	// Other nodes that will be executed to completion before cleaning up resources and marking the workflow
	// execution as failed.
	WorkflowMetadata_FAIL_AFTER_EXECUTABLE_NODES_COMPLETE WorkflowMetadata_OnFailurePolicy = 1
)

var WorkflowMetadata_OnFailurePolicy_name = map[int32]string{
	0: "FAIL_IMMEDIATELY",
	1: "FAIL_AFTER_EXECUTABLE_NODES_COMPLETE",
}

var WorkflowMetadata_OnFailurePolicy_value = map[string]int32{
	"FAIL_IMMEDIATELY":                     0,
	"FAIL_AFTER_EXECUTABLE_NODES_COMPLETE": 1,
}

func (x WorkflowMetadata_OnFailurePolicy) String() string {
	return proto.EnumName(WorkflowMetadata_OnFailurePolicy_name, int32(x))
}

func (WorkflowMetadata_OnFailurePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{8, 0}
}

// Defines a condition and the execution unit that should be executed if the condition is satisfied.
type IfBlock struct {
	Condition            *BooleanExpression `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	ThenNode             *Node              `protobuf:"bytes,2,opt,name=then_node,json=thenNode,proto3" json:"then_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IfBlock) Reset()         { *m = IfBlock{} }
func (m *IfBlock) String() string { return proto.CompactTextString(m) }
func (*IfBlock) ProtoMessage()    {}
func (*IfBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{0}
}

func (m *IfBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfBlock.Unmarshal(m, b)
}
func (m *IfBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfBlock.Marshal(b, m, deterministic)
}
func (m *IfBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfBlock.Merge(m, src)
}
func (m *IfBlock) XXX_Size() int {
	return xxx_messageInfo_IfBlock.Size(m)
}
func (m *IfBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_IfBlock.DiscardUnknown(m)
}

var xxx_messageInfo_IfBlock proto.InternalMessageInfo

func (m *IfBlock) GetCondition() *BooleanExpression {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *IfBlock) GetThenNode() *Node {
	if m != nil {
		return m.ThenNode
	}
	return nil
}

// Defines a series of if/else blocks. The first branch whose condition evaluates to true is the one to execute.
// If no conditions were satisfied, the else_node or the error will execute.
type IfElseBlock struct {
	//+required. First condition to evaluate.
	Case *IfBlock `protobuf:"bytes,1,opt,name=case,proto3" json:"case,omitempty"`
	//+optional. Additional branches to evaluate.
	Other []*IfBlock `protobuf:"bytes,2,rep,name=other,proto3" json:"other,omitempty"`
	//+required.
	//
	// Types that are valid to be assigned to Default:
	//	*IfElseBlock_ElseNode
	//	*IfElseBlock_Error
	Default              isIfElseBlock_Default `protobuf_oneof:"default"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IfElseBlock) Reset()         { *m = IfElseBlock{} }
func (m *IfElseBlock) String() string { return proto.CompactTextString(m) }
func (*IfElseBlock) ProtoMessage()    {}
func (*IfElseBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{1}
}

func (m *IfElseBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfElseBlock.Unmarshal(m, b)
}
func (m *IfElseBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfElseBlock.Marshal(b, m, deterministic)
}
func (m *IfElseBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfElseBlock.Merge(m, src)
}
func (m *IfElseBlock) XXX_Size() int {
	return xxx_messageInfo_IfElseBlock.Size(m)
}
func (m *IfElseBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_IfElseBlock.DiscardUnknown(m)
}

var xxx_messageInfo_IfElseBlock proto.InternalMessageInfo

func (m *IfElseBlock) GetCase() *IfBlock {
	if m != nil {
		return m.Case
	}
	return nil
}

func (m *IfElseBlock) GetOther() []*IfBlock {
	if m != nil {
		return m.Other
	}
	return nil
}

type isIfElseBlock_Default interface {
	isIfElseBlock_Default()
}

type IfElseBlock_ElseNode struct {
	ElseNode *Node `protobuf:"bytes,3,opt,name=else_node,json=elseNode,proto3,oneof"`
}

type IfElseBlock_Error struct {
	Error *Error `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*IfElseBlock_ElseNode) isIfElseBlock_Default() {}

func (*IfElseBlock_Error) isIfElseBlock_Default() {}

func (m *IfElseBlock) GetDefault() isIfElseBlock_Default {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *IfElseBlock) GetElseNode() *Node {
	if x, ok := m.GetDefault().(*IfElseBlock_ElseNode); ok {
		return x.ElseNode
	}
	return nil
}

func (m *IfElseBlock) GetError() *Error {
	if x, ok := m.GetDefault().(*IfElseBlock_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*IfElseBlock) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*IfElseBlock_ElseNode)(nil),
		(*IfElseBlock_Error)(nil),
	}
}

// BranchNode is a special node that alter the flow of the workflow graph. It allows the control flow to branch at
// runtime based on a series of conditions that get evaluated on various parameters (e.g. inputs, primtives).
type BranchNode struct {
	//+required
	IfElse               *IfElseBlock `protobuf:"bytes,1,opt,name=if_else,json=ifElse,proto3" json:"if_else,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BranchNode) Reset()         { *m = BranchNode{} }
func (m *BranchNode) String() string { return proto.CompactTextString(m) }
func (*BranchNode) ProtoMessage()    {}
func (*BranchNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{2}
}

func (m *BranchNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchNode.Unmarshal(m, b)
}
func (m *BranchNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchNode.Marshal(b, m, deterministic)
}
func (m *BranchNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchNode.Merge(m, src)
}
func (m *BranchNode) XXX_Size() int {
	return xxx_messageInfo_BranchNode.Size(m)
}
func (m *BranchNode) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchNode.DiscardUnknown(m)
}

var xxx_messageInfo_BranchNode proto.InternalMessageInfo

func (m *BranchNode) GetIfElse() *IfElseBlock {
	if m != nil {
		return m.IfElse
	}
	return nil
}

// Refers to the task that the Node is to execute.
type TaskNode struct {
	// Types that are valid to be assigned to Reference:
	//	*TaskNode_ReferenceId
	Reference            isTaskNode_Reference `protobuf_oneof:"reference"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskNode) Reset()         { *m = TaskNode{} }
func (m *TaskNode) String() string { return proto.CompactTextString(m) }
func (*TaskNode) ProtoMessage()    {}
func (*TaskNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{3}
}

func (m *TaskNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskNode.Unmarshal(m, b)
}
func (m *TaskNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskNode.Marshal(b, m, deterministic)
}
func (m *TaskNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskNode.Merge(m, src)
}
func (m *TaskNode) XXX_Size() int {
	return xxx_messageInfo_TaskNode.Size(m)
}
func (m *TaskNode) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskNode.DiscardUnknown(m)
}

var xxx_messageInfo_TaskNode proto.InternalMessageInfo

type isTaskNode_Reference interface {
	isTaskNode_Reference()
}

type TaskNode_ReferenceId struct {
	ReferenceId *Identifier `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3,oneof"`
}

func (*TaskNode_ReferenceId) isTaskNode_Reference() {}

func (m *TaskNode) GetReference() isTaskNode_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *TaskNode) GetReferenceId() *Identifier {
	if x, ok := m.GetReference().(*TaskNode_ReferenceId); ok {
		return x.ReferenceId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TaskNode) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TaskNode_ReferenceId)(nil),
	}
}

// Refers to a the workflow the node is to execute.
type WorkflowNode struct {
	// Types that are valid to be assigned to Reference:
	//	*WorkflowNode_LaunchplanRef
	//	*WorkflowNode_SubWorkflowRef
	Reference            isWorkflowNode_Reference `protobuf_oneof:"reference"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WorkflowNode) Reset()         { *m = WorkflowNode{} }
func (m *WorkflowNode) String() string { return proto.CompactTextString(m) }
func (*WorkflowNode) ProtoMessage()    {}
func (*WorkflowNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{4}
}

func (m *WorkflowNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowNode.Unmarshal(m, b)
}
func (m *WorkflowNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowNode.Marshal(b, m, deterministic)
}
func (m *WorkflowNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowNode.Merge(m, src)
}
func (m *WorkflowNode) XXX_Size() int {
	return xxx_messageInfo_WorkflowNode.Size(m)
}
func (m *WorkflowNode) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowNode.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowNode proto.InternalMessageInfo

type isWorkflowNode_Reference interface {
	isWorkflowNode_Reference()
}

type WorkflowNode_LaunchplanRef struct {
	LaunchplanRef *Identifier `protobuf:"bytes,1,opt,name=launchplan_ref,json=launchplanRef,proto3,oneof"`
}

type WorkflowNode_SubWorkflowRef struct {
	SubWorkflowRef *Identifier `protobuf:"bytes,2,opt,name=sub_workflow_ref,json=subWorkflowRef,proto3,oneof"`
}

func (*WorkflowNode_LaunchplanRef) isWorkflowNode_Reference() {}

func (*WorkflowNode_SubWorkflowRef) isWorkflowNode_Reference() {}

func (m *WorkflowNode) GetReference() isWorkflowNode_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *WorkflowNode) GetLaunchplanRef() *Identifier {
	if x, ok := m.GetReference().(*WorkflowNode_LaunchplanRef); ok {
		return x.LaunchplanRef
	}
	return nil
}

func (m *WorkflowNode) GetSubWorkflowRef() *Identifier {
	if x, ok := m.GetReference().(*WorkflowNode_SubWorkflowRef); ok {
		return x.SubWorkflowRef
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WorkflowNode) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WorkflowNode_LaunchplanRef)(nil),
		(*WorkflowNode_SubWorkflowRef)(nil),
	}
}

// Defines extra information about the Node.
type NodeMetadata struct {
	// A friendly name for the Node
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The overall timeout of a task.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Number of retries per task.
	Retries *RetryStrategy `protobuf:"bytes,5,opt,name=retries,proto3" json:"retries,omitempty"`
	// Identify whether node is interruptible
	//
	// Types that are valid to be assigned to InterruptibleValue:
	//	*NodeMetadata_Interruptible
	InterruptibleValue   isNodeMetadata_InterruptibleValue `protobuf_oneof:"interruptible_value"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *NodeMetadata) Reset()         { *m = NodeMetadata{} }
func (m *NodeMetadata) String() string { return proto.CompactTextString(m) }
func (*NodeMetadata) ProtoMessage()    {}
func (*NodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{5}
}

func (m *NodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetadata.Unmarshal(m, b)
}
func (m *NodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetadata.Marshal(b, m, deterministic)
}
func (m *NodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetadata.Merge(m, src)
}
func (m *NodeMetadata) XXX_Size() int {
	return xxx_messageInfo_NodeMetadata.Size(m)
}
func (m *NodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetadata proto.InternalMessageInfo

func (m *NodeMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeMetadata) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *NodeMetadata) GetRetries() *RetryStrategy {
	if m != nil {
		return m.Retries
	}
	return nil
}

type isNodeMetadata_InterruptibleValue interface {
	isNodeMetadata_InterruptibleValue()
}

type NodeMetadata_Interruptible struct {
	Interruptible bool `protobuf:"varint,6,opt,name=interruptible,proto3,oneof"`
}

func (*NodeMetadata_Interruptible) isNodeMetadata_InterruptibleValue() {}

func (m *NodeMetadata) GetInterruptibleValue() isNodeMetadata_InterruptibleValue {
	if m != nil {
		return m.InterruptibleValue
	}
	return nil
}

func (m *NodeMetadata) GetInterruptible() bool {
	if x, ok := m.GetInterruptibleValue().(*NodeMetadata_Interruptible); ok {
		return x.Interruptible
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NodeMetadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NodeMetadata_Interruptible)(nil),
	}
}

// Links a variable to an alias.
type Alias struct {
	// Must match one of the output variable names on a node.
	Var string `protobuf:"bytes,1,opt,name=var,proto3" json:"var,omitempty"`
	// A workflow-level unique alias that downstream nodes can refer to in their input.
	Alias                string   `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}
func (*Alias) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{6}
}

func (m *Alias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alias.Unmarshal(m, b)
}
func (m *Alias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alias.Marshal(b, m, deterministic)
}
func (m *Alias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alias.Merge(m, src)
}
func (m *Alias) XXX_Size() int {
	return xxx_messageInfo_Alias.Size(m)
}
func (m *Alias) XXX_DiscardUnknown() {
	xxx_messageInfo_Alias.DiscardUnknown(m)
}

var xxx_messageInfo_Alias proto.InternalMessageInfo

func (m *Alias) GetVar() string {
	if m != nil {
		return m.Var
	}
	return ""
}

func (m *Alias) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

// A Workflow graph Node. One unit of execution in the graph. Each node can be linked to a Task, a Workflow or a branch
// node.
type Node struct {
	// A workflow-level unique identifier that identifies this node in the workflow. "inputs" and "outputs" are reserved
	// node ids that cannot be used by other nodes.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Extra metadata about the node.
	Metadata *NodeMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specifies how to bind the underlying interface's inputs. All required inputs specified in the underlying interface
	// must be fullfilled.
	Inputs []*Binding `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty"`
	//+optional Specifies execution depdendency for this node ensuring it will only get scheduled to run after all its
	// upstream nodes have completed. This node will have an implicit depdendency on any node that appears in inputs
	// field.
	UpstreamNodeIds []string `protobuf:"bytes,4,rep,name=upstream_node_ids,json=upstreamNodeIds,proto3" json:"upstream_node_ids,omitempty"`
	//+optional. A node can define aliases for a subset of its outputs. This is particularly useful if different nodes
	// need to conform to the same interface (e.g. all branches in a branch node). Downstream nodes must refer to this
	// nodes outputs using the alias if one's specified.
	OutputAliases []*Alias `protobuf:"bytes,5,rep,name=output_aliases,json=outputAliases,proto3" json:"output_aliases,omitempty"`
	// Information about the target to execute in this node.
	//
	// Types that are valid to be assigned to Target:
	//	*Node_TaskNode
	//	*Node_WorkflowNode
	//	*Node_BranchNode
	Target               isNode_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{7}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetMetadata() *NodeMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetInputs() []*Binding {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Node) GetUpstreamNodeIds() []string {
	if m != nil {
		return m.UpstreamNodeIds
	}
	return nil
}

func (m *Node) GetOutputAliases() []*Alias {
	if m != nil {
		return m.OutputAliases
	}
	return nil
}

type isNode_Target interface {
	isNode_Target()
}

type Node_TaskNode struct {
	TaskNode *TaskNode `protobuf:"bytes,6,opt,name=task_node,json=taskNode,proto3,oneof"`
}

type Node_WorkflowNode struct {
	WorkflowNode *WorkflowNode `protobuf:"bytes,7,opt,name=workflow_node,json=workflowNode,proto3,oneof"`
}

type Node_BranchNode struct {
	BranchNode *BranchNode `protobuf:"bytes,8,opt,name=branch_node,json=branchNode,proto3,oneof"`
}

func (*Node_TaskNode) isNode_Target() {}

func (*Node_WorkflowNode) isNode_Target() {}

func (*Node_BranchNode) isNode_Target() {}

func (m *Node) GetTarget() isNode_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Node) GetTaskNode() *TaskNode {
	if x, ok := m.GetTarget().(*Node_TaskNode); ok {
		return x.TaskNode
	}
	return nil
}

func (m *Node) GetWorkflowNode() *WorkflowNode {
	if x, ok := m.GetTarget().(*Node_WorkflowNode); ok {
		return x.WorkflowNode
	}
	return nil
}

func (m *Node) GetBranchNode() *BranchNode {
	if x, ok := m.GetTarget().(*Node_BranchNode); ok {
		return x.BranchNode
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Node) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Node_TaskNode)(nil),
		(*Node_WorkflowNode)(nil),
		(*Node_BranchNode)(nil),
	}
}

// Metadata for the entire workflow. Defines execution behavior that does not change the final outputs of the workflow.
type WorkflowMetadata struct {
	// Total wait time a workflow can be delayed by queueing.
	QueuingBudget *duration.Duration `protobuf:"bytes,1,opt,name=queuing_budget,json=queuingBudget,proto3" json:"queuing_budget,omitempty"`
	// Defines how the system should behave when a failure is detected in the workflow execution.
	OnFailure            WorkflowMetadata_OnFailurePolicy `protobuf:"varint,2,opt,name=on_failure,json=onFailure,proto3,enum=flyteidl.core.WorkflowMetadata_OnFailurePolicy" json:"on_failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WorkflowMetadata) Reset()         { *m = WorkflowMetadata{} }
func (m *WorkflowMetadata) String() string { return proto.CompactTextString(m) }
func (*WorkflowMetadata) ProtoMessage()    {}
func (*WorkflowMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{8}
}

func (m *WorkflowMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowMetadata.Unmarshal(m, b)
}
func (m *WorkflowMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowMetadata.Marshal(b, m, deterministic)
}
func (m *WorkflowMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowMetadata.Merge(m, src)
}
func (m *WorkflowMetadata) XXX_Size() int {
	return xxx_messageInfo_WorkflowMetadata.Size(m)
}
func (m *WorkflowMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowMetadata proto.InternalMessageInfo

func (m *WorkflowMetadata) GetQueuingBudget() *duration.Duration {
	if m != nil {
		return m.QueuingBudget
	}
	return nil
}

func (m *WorkflowMetadata) GetOnFailure() WorkflowMetadata_OnFailurePolicy {
	if m != nil {
		return m.OnFailure
	}
	return WorkflowMetadata_FAIL_IMMEDIATELY
}

// Default Workflow Metadata for the entire workflow.
type WorkflowMetadataDefaults struct {
	// Identify whether workflow is interruptible.
	// The value set at the workflow level will be the defualt value used for nodes
	// unless explicitly set at the node level.
	Interruptible        bool     `protobuf:"varint,1,opt,name=interruptible,proto3" json:"interruptible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowMetadataDefaults) Reset()         { *m = WorkflowMetadataDefaults{} }
func (m *WorkflowMetadataDefaults) String() string { return proto.CompactTextString(m) }
func (*WorkflowMetadataDefaults) ProtoMessage()    {}
func (*WorkflowMetadataDefaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{9}
}

func (m *WorkflowMetadataDefaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowMetadataDefaults.Unmarshal(m, b)
}
func (m *WorkflowMetadataDefaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowMetadataDefaults.Marshal(b, m, deterministic)
}
func (m *WorkflowMetadataDefaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowMetadataDefaults.Merge(m, src)
}
func (m *WorkflowMetadataDefaults) XXX_Size() int {
	return xxx_messageInfo_WorkflowMetadataDefaults.Size(m)
}
func (m *WorkflowMetadataDefaults) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowMetadataDefaults.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowMetadataDefaults proto.InternalMessageInfo

func (m *WorkflowMetadataDefaults) GetInterruptible() bool {
	if m != nil {
		return m.Interruptible
	}
	return false
}

// Flyte Workflow Structure that encapsulates task, branch and subworkflow nodes to form a statically analyzable,
// directed acyclic graph.
type WorkflowTemplate struct {
	// A globally unique identifier for the workflow.
	Id *Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Extra metadata about the workflow.
	Metadata *WorkflowMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Defines a strongly typed interface for the Workflow. This can include some optional parameters.
	Interface *TypedInterface `protobuf:"bytes,3,opt,name=interface,proto3" json:"interface,omitempty"`
	// A list of nodes. In addition, "globals" is a special reserved node id that can be used to consume workflow inputs.
	Nodes []*Node `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// A list of output bindings that specify how to construct workflow outputs. Bindings can pull node outputs or
	// specify literals. All workflow outputs specified in the interface field must be bound in order for the workflow
	// to be validated. A workflow has an implicit dependency on all of its nodes to execute successfully in order to
	// bind final outputs.
	// Most of these outputs will be Binding's with a BindingData of type OutputReference.  That is, your workflow can
	// just have an output of some constant (`Output(5)`), but usually, the workflow will be pulling
	// outputs from the output of a task.
	Outputs []*Binding `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	//+optional A catch-all node. This node is executed whenever the execution engine determines the workflow has failed.
	// The interface of this node must match the Workflow interface with an additional input named "error" of type
	// pb.lyft.flyte.core.Error.
	FailureNode *Node `protobuf:"bytes,6,opt,name=failure_node,json=failureNode,proto3" json:"failure_node,omitempty"`
	// workflow defaults
	MetadataDefaults     *WorkflowMetadataDefaults `protobuf:"bytes,7,opt,name=metadata_defaults,json=metadataDefaults,proto3" json:"metadata_defaults,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *WorkflowTemplate) Reset()         { *m = WorkflowTemplate{} }
func (m *WorkflowTemplate) String() string { return proto.CompactTextString(m) }
func (*WorkflowTemplate) ProtoMessage()    {}
func (*WorkflowTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_fccede37486c456e, []int{10}
}

func (m *WorkflowTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTemplate.Unmarshal(m, b)
}
func (m *WorkflowTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTemplate.Marshal(b, m, deterministic)
}
func (m *WorkflowTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTemplate.Merge(m, src)
}
func (m *WorkflowTemplate) XXX_Size() int {
	return xxx_messageInfo_WorkflowTemplate.Size(m)
}
func (m *WorkflowTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTemplate proto.InternalMessageInfo

func (m *WorkflowTemplate) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *WorkflowTemplate) GetMetadata() *WorkflowMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *WorkflowTemplate) GetInterface() *TypedInterface {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *WorkflowTemplate) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *WorkflowTemplate) GetOutputs() []*Binding {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *WorkflowTemplate) GetFailureNode() *Node {
	if m != nil {
		return m.FailureNode
	}
	return nil
}

func (m *WorkflowTemplate) GetMetadataDefaults() *WorkflowMetadataDefaults {
	if m != nil {
		return m.MetadataDefaults
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.core.WorkflowMetadata_OnFailurePolicy", WorkflowMetadata_OnFailurePolicy_name, WorkflowMetadata_OnFailurePolicy_value)
	proto.RegisterType((*IfBlock)(nil), "flyteidl.core.IfBlock")
	proto.RegisterType((*IfElseBlock)(nil), "flyteidl.core.IfElseBlock")
	proto.RegisterType((*BranchNode)(nil), "flyteidl.core.BranchNode")
	proto.RegisterType((*TaskNode)(nil), "flyteidl.core.TaskNode")
	proto.RegisterType((*WorkflowNode)(nil), "flyteidl.core.WorkflowNode")
	proto.RegisterType((*NodeMetadata)(nil), "flyteidl.core.NodeMetadata")
	proto.RegisterType((*Alias)(nil), "flyteidl.core.Alias")
	proto.RegisterType((*Node)(nil), "flyteidl.core.Node")
	proto.RegisterType((*WorkflowMetadata)(nil), "flyteidl.core.WorkflowMetadata")
	proto.RegisterType((*WorkflowMetadataDefaults)(nil), "flyteidl.core.WorkflowMetadataDefaults")
	proto.RegisterType((*WorkflowTemplate)(nil), "flyteidl.core.WorkflowTemplate")
}

func init() { proto.RegisterFile("flyteidl/core/workflow.proto", fileDescriptor_fccede37486c456e) }

var fileDescriptor_fccede37486c456e = []byte{
	// 1036 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x5f, 0x6f, 0xdb, 0xb6,
	0x17, 0xb5, 0xe3, 0xc4, 0xb6, 0xae, 0x63, 0xd7, 0x65, 0xf3, 0xfb, 0xcd, 0xc9, 0xda, 0x2e, 0x30,
	0x8a, 0x2d, 0x2d, 0x36, 0x3b, 0x48, 0x81, 0xec, 0x21, 0x43, 0x51, 0xab, 0x51, 0x10, 0x03, 0xf9,
	0xd3, 0xb1, 0x1e, 0xba, 0xed, 0x45, 0xa0, 0xad, 0x2b, 0x87, 0x88, 0x2c, 0x6a, 0x14, 0xd5, 0xcc,
	0xd8, 0xb7, 0xd9, 0xc3, 0x9e, 0xf6, 0x59, 0xf6, 0xb4, 0xef, 0xb2, 0xd7, 0x41, 0xd4, 0x1f, 0xc7,
	0xaa, 0xd3, 0xec, 0x8d, 0xe4, 0x3d, 0xf7, 0xf2, 0xf2, 0xf0, 0x1c, 0x51, 0xf0, 0xd8, 0xf5, 0xe6,
	0x0a, 0xb9, 0xe3, 0xf5, 0x27, 0x42, 0x62, 0xff, 0x46, 0xc8, 0x6b, 0xd7, 0x13, 0x37, 0xbd, 0x40,
	0x0a, 0x25, 0x48, 0x33, 0x8b, 0xf6, 0xe2, 0xe8, 0xce, 0x93, 0x65, 0xf0, 0x44, 0xf8, 0x0e, 0x57,
	0x5c, 0xf8, 0x09, 0x7a, 0xe7, 0xe9, 0x72, 0x98, 0x3b, 0xe8, 0x2b, 0xee, 0x72, 0x94, 0x69, 0xbc,
	0x90, 0xce, 0x7d, 0x85, 0xd2, 0x65, 0x13, 0x4c, 0xc3, 0x85, 0x56, 0x3c, 0xae, 0x50, 0x32, 0x2f,
	0x4c, 0xa3, 0xdb, 0xcb, 0x51, 0x35, 0x0f, 0x30, 0x0b, 0x3d, 0x9d, 0x0a, 0x31, 0xf5, 0xb0, 0xaf,
	0x67, 0xe3, 0xc8, 0xed, 0x3b, 0x91, 0x64, 0x8b, 0xbe, 0xba, 0xbf, 0x41, 0x6d, 0xe8, 0x9a, 0x9e,
	0x98, 0x5c, 0x93, 0x57, 0x60, 0xe4, 0x5d, 0x77, 0xca, 0xbb, 0xe5, 0xbd, 0xc6, 0xc1, 0x6e, 0x6f,
	0xe9, 0x90, 0x3d, 0x53, 0x08, 0x0f, 0x99, 0x6f, 0xfd, 0x1a, 0x48, 0x0c, 0x43, 0x2e, 0x7c, 0xba,
	0x48, 0x21, 0xfb, 0x60, 0xa8, 0x2b, 0xf4, 0x6d, 0x5f, 0x38, 0xd8, 0x59, 0xd3, 0xf9, 0x8f, 0x0a,
	0xf9, 0x17, 0xc2, 0x41, 0x5a, 0x8f, 0x51, 0xf1, 0xa8, 0xfb, 0x77, 0x19, 0x1a, 0x43, 0xd7, 0xf2,
	0x42, 0x4c, 0x3a, 0x78, 0x01, 0xeb, 0x13, 0x16, 0x62, 0xba, 0xf9, 0xff, 0x0b, 0xc9, 0x69, 0x9f,
	0x54, 0x63, 0xc8, 0xd7, 0xb0, 0x21, 0xd4, 0x15, 0xca, 0xce, 0xda, 0x6e, 0xe5, 0x13, 0xe0, 0x04,
	0x44, 0x0e, 0xc0, 0x40, 0x2f, 0xc4, 0xa4, 0xb7, 0xca, 0x9d, 0xbd, 0x9d, 0x96, 0x68, 0x3d, 0xc6,
	0xc5, 0xe3, 0x78, 0x07, 0x94, 0x52, 0xc8, 0xce, 0xba, 0xc6, 0x6f, 0x15, 0xf0, 0x56, 0x1c, 0x3b,
	0x2d, 0xd1, 0x04, 0x64, 0x1a, 0x50, 0x73, 0xd0, 0x65, 0x91, 0xa7, 0xba, 0x03, 0x00, 0x53, 0x32,
	0x7f, 0x72, 0xa5, 0xcb, 0xbc, 0x84, 0x1a, 0x77, 0xed, 0xb8, 0x6a, 0x7a, 0xae, 0x9d, 0x8f, 0x5a,
	0xcd, 0x19, 0xa0, 0x55, 0xae, 0x27, 0xdd, 0xf7, 0x50, 0x1f, 0xb1, 0xf0, 0x5a, 0x17, 0x78, 0x05,
	0x9b, 0x12, 0x5d, 0x94, 0xe8, 0x4f, 0xd0, 0xe6, 0x4e, 0x5a, 0x65, 0xbb, 0x58, 0x25, 0x57, 0xd4,
	0x69, 0x89, 0x36, 0xf2, 0x84, 0xa1, 0x63, 0x36, 0xc0, 0xc8, 0xa7, 0xdd, 0x3f, 0xca, 0xb0, 0xf9,
	0x3e, 0x15, 0xb2, 0xae, 0x6e, 0x42, 0xcb, 0x63, 0x91, 0x3f, 0xb9, 0x0a, 0x3c, 0xe6, 0xdb, 0x12,
	0xdd, 0xff, 0x52, 0xbf, 0xb9, 0x48, 0xa1, 0xe8, 0x12, 0x0b, 0xda, 0x61, 0x34, 0xb6, 0x33, 0x83,
	0xe8, 0x2a, 0x6b, 0xf7, 0x57, 0x69, 0x85, 0xd1, 0x38, 0xeb, 0x85, 0xa2, 0xbb, 0xdc, 0xe8, 0x5f,
	0x65, 0xd8, 0x8c, 0x1b, 0x3c, 0x47, 0xc5, 0x1c, 0xa6, 0x18, 0x21, 0xb0, 0xee, 0xb3, 0x59, 0x42,
	0xa2, 0x41, 0xf5, 0x38, 0xe6, 0x56, 0xf1, 0x19, 0x8a, 0x48, 0xa5, 0x97, 0xb4, 0xdd, 0x4b, 0xf4,
	0xde, 0xcb, 0xf4, 0xde, 0x3b, 0x4e, 0xf5, 0x4e, 0x33, 0x24, 0x39, 0x84, 0x9a, 0x44, 0x25, 0x39,
	0x86, 0x9d, 0x0d, 0x9d, 0xf4, 0xb8, 0xd0, 0x24, 0x45, 0x25, 0xe7, 0xef, 0x94, 0x64, 0x0a, 0xa7,
	0x73, 0x9a, 0x81, 0xc9, 0x97, 0xd0, 0xd4, 0xb6, 0x94, 0x51, 0xa0, 0xf8, 0xd8, 0xc3, 0x4e, 0x75,
	0xb7, 0xbc, 0x57, 0x8f, 0xd9, 0x58, 0x5a, 0x36, 0xff, 0x07, 0x8f, 0x96, 0x16, 0xec, 0x0f, 0xcc,
	0x8b, 0xb0, 0xdb, 0x87, 0x8d, 0x81, 0xc7, 0x59, 0x48, 0xda, 0x50, 0xf9, 0xc0, 0x64, 0x7a, 0x8e,
	0x78, 0x48, 0xb6, 0x60, 0x83, 0xc5, 0x21, 0x4d, 0x9a, 0x41, 0x93, 0x49, 0xf7, 0xcf, 0x0a, 0xac,
	0xeb, 0x2b, 0x6a, 0xc1, 0x5a, 0x7a, 0xed, 0x06, 0x5d, 0xe3, 0x0e, 0xf9, 0x16, 0xea, 0xb3, 0x94,
	0x95, 0x94, 0xe6, 0xcf, 0x57, 0x68, 0x39, 0x23, 0x8e, 0xe6, 0x60, 0xd2, 0x83, 0x2a, 0xf7, 0x83,
	0x48, 0x85, 0x9d, 0xca, 0x4a, 0xd3, 0x98, 0xdc, 0x77, 0xb8, 0x3f, 0xa5, 0x29, 0x8a, 0xbc, 0x80,
	0x87, 0x51, 0x10, 0x2a, 0x89, 0x6c, 0xa6, 0x9d, 0x63, 0x73, 0x27, 0xec, 0xac, 0xef, 0x56, 0xf6,
	0x0c, 0xfa, 0x20, 0x0b, 0xc4, 0x5b, 0x0d, 0x9d, 0x90, 0x1c, 0x41, 0x4b, 0x44, 0x2a, 0x88, 0x94,
	0xad, 0xbb, 0xd7, 0xe4, 0x56, 0x56, 0xd8, 0x46, 0x73, 0x40, 0x9b, 0x09, 0x76, 0x90, 0x40, 0xc9,
	0x21, 0x18, 0x8a, 0x85, 0xd7, 0x89, 0x3d, 0xab, 0xfa, 0x48, 0x9f, 0x15, 0xf2, 0x32, 0x3b, 0xc4,
	0x16, 0x55, 0x99, 0x35, 0x4c, 0x68, 0xe6, 0xa2, 0xd3, 0xb9, 0xb5, 0x95, 0x74, 0xdc, 0x16, 0xfc,
	0x69, 0x89, 0x6e, 0xde, 0xdc, 0x36, 0xc0, 0x77, 0xd0, 0x18, 0x6b, 0xb7, 0x26, 0x15, 0xea, 0x2b,
	0x75, 0xbb, 0xf0, 0xf3, 0x69, 0x89, 0xc2, 0x38, 0x9f, 0x99, 0x75, 0xa8, 0x2a, 0x26, 0xa7, 0xa8,
	0xba, 0xff, 0x94, 0xa1, 0x9d, 0x6d, 0x94, 0x8b, 0xf6, 0x35, 0xb4, 0x7e, 0x89, 0x30, 0xe2, 0xfe,
	0xd4, 0x1e, 0x47, 0xce, 0x14, 0x55, 0xee, 0xae, 0x3b, 0x75, 0xda, 0x4c, 0x13, 0x4c, 0x8d, 0x27,
	0x17, 0x00, 0xc2, 0xb7, 0x5d, 0xc6, 0xbd, 0x48, 0x26, 0x9f, 0xd5, 0xd6, 0x41, 0xff, 0x8e, 0xf3,
	0x65, 0xdb, 0xf6, 0x2e, 0xfd, 0x93, 0x24, 0xe1, 0xad, 0xf0, 0xf8, 0x64, 0x4e, 0x0d, 0x91, 0x2d,
	0x74, 0xbf, 0x87, 0x07, 0x85, 0x28, 0xd9, 0x82, 0xf6, 0xc9, 0x60, 0x78, 0x66, 0x0f, 0xcf, 0xcf,
	0xad, 0xe3, 0xe1, 0x60, 0x64, 0x9d, 0xfd, 0xd4, 0x2e, 0x91, 0x3d, 0x78, 0xa6, 0x57, 0x07, 0x27,
	0x23, 0x8b, 0xda, 0xd6, 0x8f, 0xd6, 0x9b, 0x1f, 0x46, 0x03, 0xf3, 0xcc, 0xb2, 0x2f, 0x2e, 0x8f,
	0xad, 0x77, 0xf6, 0x9b, 0xcb, 0xf3, 0xb7, 0x67, 0xd6, 0xc8, 0x6a, 0x97, 0xbb, 0xaf, 0xa1, 0x53,
	0xec, 0xe0, 0x38, 0xf9, 0x14, 0x86, 0xe4, 0x59, 0xd1, 0x34, 0xf1, 0xf9, 0xeb, 0x05, 0xcb, 0x74,
	0x7f, 0xaf, 0x2c, 0xb8, 0x1b, 0xe1, 0x2c, 0xf0, 0x98, 0x42, 0xf2, 0x3c, 0x97, 0xfd, 0xa7, 0xbe,
	0x23, 0xda, 0x11, 0x47, 0x1f, 0x39, 0xe2, 0x8b, 0x7b, 0x28, 0xba, 0xe5, 0x8a, 0x23, 0x30, 0xf2,
	0xe7, 0x36, 0x7d, 0x1b, 0x9e, 0x14, 0xc5, 0x37, 0x0f, 0xd0, 0x19, 0x66, 0x20, 0xba, 0xc0, 0x93,
	0xe7, 0xb0, 0x11, 0xcb, 0x26, 0xb1, 0xc5, 0x1d, 0x0f, 0x5e, 0x82, 0x20, 0xfb, 0x50, 0x4b, 0x54,
	0x9f, 0x59, 0xe3, 0x2e, 0xfb, 0x65, 0x30, 0x72, 0x08, 0x9b, 0xe9, 0xc5, 0xdf, 0x76, 0xc6, 0xca,
	0x3d, 0x1a, 0x29, 0x50, 0x4b, 0x7a, 0x04, 0x0f, 0xb3, 0xd3, 0xd9, 0xe9, 0xa3, 0x14, 0xa6, 0xd6,
	0xf8, 0xea, 0x1e, 0x5e, 0xb2, 0x8b, 0xa3, 0xed, 0x59, 0x61, 0xc5, 0x3c, 0xf8, 0x79, 0x7f, 0xca,
	0xd5, 0x55, 0x34, 0xee, 0x4d, 0xc4, 0xac, 0xef, 0xcd, 0x5d, 0xd5, 0xcf, 0xff, 0x3b, 0xa6, 0xe8,
	0xf7, 0x83, 0xf1, 0x37, 0x53, 0xd1, 0x5f, 0xfa, 0x15, 0x19, 0x57, 0xb5, 0xbe, 0x5f, 0xfe, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0xab, 0x53, 0xa5, 0x4b, 0x09, 0x00, 0x00,
}
