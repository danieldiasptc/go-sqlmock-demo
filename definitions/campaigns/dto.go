package campaigns

type ListFilter struct {
	Query  string `form:"query"`
	Status string `form:"status" binding:"omitempty,oneof=enabled disabled complete"`
	Type   string `form:"type" binding:"omitempty,oneof=one_time recurring"`
}
