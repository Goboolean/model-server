package task_test

import (
	"os"
	"testing"

	"github.com/Goboolean/model-server/internal/domain/entity"
	task "github.com/Goboolean/model-server/internal/domain/service/taskqueue"
	"github.com/stretchr/testify/assert"
)



var instance *task.Queue

func TestMain(t *testing.M) {

	instance = task.New()

	code := t.Run()
	os.Exit(code)
}


/*

*/


func TestTaskQueue(t *testing.T) {


	t.Run("Enqueue", func(t *testing.T) {
		var err error

		len := instance.Length()

		err = instance.Enqueue(&entity.Task{})
		assert.NoError(t, err)
		assert.Equal(t, len+1, instance.Length())
	})

	t.Run("Dequeue", func(t *testing.T) {
		var err error
		
		len := instance.Length()

		_, err = instance.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, len-1, instance.Length())
	})

	t.Run("Dequeue_Empty", func(t *testing.T) {
		// Precondition: Queue is empty
		var err error

		instance.Dequeue()

		_, err = instance.Dequeue()
		assert.Error(t, err, task.ErrQueueEmpty)
	})

	t.Run("Remove", func(t *testing.T) {
		var err error
		var _task *entity.Task

		// Precondition: Queue have many elements

		instance.Enqueue(entity.NewTask(13))
		instance.Enqueue(entity.NewTask(22))

		_task, err = instance.Remove(13)
		assert.NoError(t, err)
		assert.Equal(t, int64(13), _task.ID)

		_task, err = instance.Remove(21)
		assert.Error(t, err, task.ErrTaskNotFound)
	})

	t.Run("Exists", func(t *testing.T) {

		instance.Enqueue(entity.NewTask(31))
		assert.True(t, instance.Exists(31))
	})

	t.Run("IsEmpty", func(t *testing.T) {

		for instance.Length() > 0 {
			instance.Dequeue()
		}

		assert.True(t, instance.IsEmpty())
	})
}