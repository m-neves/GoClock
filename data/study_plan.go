package data

type StudyPlan struct {
	Id            int    `json:"id"`
	Name          string `json:"study_plan_name"`
	CurrentCourse uint   `json:"-"`
	UserId        uint   `json:"user_id,omitempty"`
}

type StudyPlanCourse struct {
	StudyPlanId uint
	CourseId    uint
	Order       uint
}

type StudyPlanCourses struct {
	StudyPlan StudyPlan
	Courses   []*Course
}
