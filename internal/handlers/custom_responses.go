package handlers

// import (
// 	"encoding/json"

// 	"github.com/shock_wave/restapi/internal/web/tasks"
// )

// // TaskWrapper — структура-обёртка для tasks.Task
// type TaskWrapper struct {
// 	Task tasks.Task
// }

// // Реализация кастомного `MarshalJSON` для гарантированного порядка полей
// func (tw TaskWrapper) MarshalJSON() ([]byte, error) {
// 	var id uint
// 	var task string
// 	var isDone bool

// 	// Проверяем указатели, чтобы избежать nil-значений
// 	if tw.Task.Id != nil {
// 		id = *tw.Task.Id
// 	}
// 	if tw.Task.Task != nil {
// 		task = *tw.Task.Task
// 	}
// 	if tw.Task.IsDone != nil {
// 		isDone = *tw.Task.IsDone
// 	}

// 	// Возвращаем JSON с правильным порядком полей
// 	return json.Marshal(struct {
// 		Id     uint   `json:"id"`
// 		Task   string `json:"task"`
// 		IsDone bool   `json:"is_done"`
// 	}{
// 		Id:     id,
// 		Task:   task,
// 		IsDone: isDone,
// 	})
// }
