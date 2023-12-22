package gormimpl

import (
	"context"
	"errors"
	"github.com/flyteorg/flyte/flyteadmin/pkg/common"

	"gorm.io/gorm"

	flyteAdminDbErrors "github.com/flyteorg/flyte/flyteadmin/pkg/repositories/errors"
	"github.com/flyteorg/flyte/flyteadmin/pkg/repositories/interfaces"
	"github.com/flyteorg/flyte/flyteadmin/pkg/repositories/models"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyte/flytestdlib/promutils"
)

// Implementation of TaskExecutionInterface.
type TaskExecutionRepo struct {
	db               *gorm.DB
	errorTransformer flyteAdminDbErrors.ErrorTransformer
	metrics          gormMetrics
}

func (r *TaskExecutionRepo) Create(ctx context.Context, id *core.TaskExecutionIdentifier, input models.TaskExecution) error {
	timer := r.metrics.CreateDuration.Start()
	tx := r.db.WithContext(ctx).Omit("id").Create(&input)
	timer.Stop()
	if tx.Error != nil {
		return r.errorTransformer.ToFlyteAdminError(tx.Error)
	}
	return nil
}

func (r *TaskExecutionRepo) Get(ctx context.Context, id *core.TaskExecutionIdentifier) (models.TaskExecution, error) {
	var taskExecution models.TaskExecution
	timer := r.metrics.GetDuration.Start()
	tx := r.db.WithContext(ctx).Where(&models.TaskExecution{
		TaskExecutionKey: models.TaskExecutionKey{
			TaskKey: models.TaskKey{
				Project: id.GetTaskId().Project,
				Domain:  id.GetTaskId().Domain,
				Name:    id.GetTaskId().Name,
				Version: id.GetTaskId().Version,
			},
			NodeExecutionKey: models.NodeExecutionKey{
				NodeID: id.GetNodeExecutionId().NodeId,
				ExecutionKey: models.ExecutionKey{
					Project: id.GetNodeExecutionId().GetExecutionId().Project,
					Domain:  id.GetNodeExecutionId().GetExecutionId().Domain,
					Name:    id.GetNodeExecutionId().GetExecutionId().Name,
				},
			},
			RetryAttempt: &id.RetryAttempt,
		},
	}).Preload("ChildNodeExecution").Take(&taskExecution)
	timer.Stop()

	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return models.TaskExecution{},
			flyteAdminDbErrors.GetMissingEntityError("task execution", id)
	} else if tx.Error != nil {
		return models.TaskExecution{}, r.errorTransformer.ToFlyteAdminError(tx.Error)
	}
	return taskExecution, nil
}

func (r *TaskExecutionRepo) Update(ctx context.Context, id *core.TaskExecutionIdentifier, execution models.TaskExecution) error {
	timer := r.metrics.UpdateDuration.Start()
	tx := r.db.WithContext(ctx).WithContext(ctx).Save(&execution) // TODO @hmaersaw - need to add WithContext to all db calls to link otel spans
	timer.Stop()

	if err := tx.Error; err != nil {
		return r.errorTransformer.ToFlyteAdminError(err)
	}
	return nil
}

func (r *TaskExecutionRepo) List(ctx context.Context, input interfaces.ListResourceInput) (interfaces.TaskExecutionCollectionOutput, error) {
	if err := ValidateListInput(input); err != nil {
		return interfaces.TaskExecutionCollectionOutput{}, err
	}

	var taskExecutions []models.TaskExecution
	tx := r.db.WithContext(ctx).Limit(input.Limit).Offset(input.Offset).Preload("ChildNodeExecution")

	// And add three join conditions (joining multiple tables is fine even we only filter on a subset of table attributes).
	// We are joining on task -> taskExec -> NodeExec -> Exec.
	// NOTE: the order in which the joins are called below are important because postgres will only know about certain
	// tables as they are joined. So we should do it in the order specified above.
	tx = tx.Joins(leftJoinTaskToTaskExec)
	tx = tx.Joins(innerJoinNodeExecToTaskExec)
	tx = tx.Joins(innerJoinExecToNodeExec)

	// Apply filters
	tx, err := applyScopedFilters(tx, common.TaskExecution, input.IdentifierScope, input.InlineFilters, input.MapFilters)
	if err != nil {
		return interfaces.TaskExecutionCollectionOutput{}, err
	}

	// Apply sort ordering.
	if input.SortParameter != nil {
		tx = tx.Order(input.SortParameter.GetGormOrderExpr())
	}

	timer := r.metrics.ListDuration.Start()
	tx = tx.Find(&taskExecutions)
	timer.Stop()
	if tx.Error != nil {
		return interfaces.TaskExecutionCollectionOutput{}, r.errorTransformer.ToFlyteAdminError(tx.Error)
	}

	return interfaces.TaskExecutionCollectionOutput{
		TaskExecutions: taskExecutions,
	}, nil
}

func (r *TaskExecutionRepo) Count(ctx context.Context, input interfaces.CountResourceInput) (int64, error) {
	var err error
	tx := r.db.WithContext(ctx).Model(&models.TaskExecution{})

	// Add three join conditions (joining multiple tables is fine even we only filter on a subset of table attributes).
	// We are joining on task -> taskExec -> NodeExec -> Exec.
	// NOTE: the order in which the joins are called below are important because postgres will only know about certain
	// tables as they are joined. So we should do it in the order specified above.
	tx = tx.Joins(leftJoinTaskToTaskExec)
	tx = tx.Joins(innerJoinNodeExecToTaskExec)
	tx = tx.Joins(innerJoinExecToNodeExec)

	// Apply filters
	tx, err = applyScopedFilters(tx, common.TaskExecution, input.IdentifierScope, input.InlineFilters, input.MapFilters)
	if err != nil {
		return 0, err
	}

	// Run the query
	timer := r.metrics.CountDuration.Start()
	var count int64
	tx = tx.Count(&count)
	timer.Stop()
	if tx.Error != nil {
		return 0, r.errorTransformer.ToFlyteAdminError(tx.Error)
	}
	return count, nil
}

// Returns an instance of TaskExecutionRepoInterface
func NewTaskExecutionRepo(
	db *gorm.DB, errorTransformer flyteAdminDbErrors.ErrorTransformer, scope promutils.Scope) interfaces.TaskExecutionRepoInterface {
	metrics := newMetrics(scope)
	return &TaskExecutionRepo{
		db:               db,
		errorTransformer: errorTransformer,
		metrics:          metrics,
	}
}
