package in

import "github.com/Goboolean/model-server/internal/domain/entity"


type TaskHandler interface {
	RegisterTask(*entity.Task) error
	CancelTask(int64) error
}
