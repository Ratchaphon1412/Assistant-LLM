package entities

type Chat struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	WorkflowID uint   `json:"workflow_id" gorm:""`
	Prompt     string `json:"prompt" gorm:""`
}
