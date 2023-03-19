package activities

type ActivityQueryDto struct {
	Skip uint `form:"skip,default=0" json:"skip"`
	Take uint `form:"take,default=10" json:"take"`
}
