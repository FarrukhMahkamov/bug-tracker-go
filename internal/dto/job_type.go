package dto

type JobType struct {
	Id          int16  `json:"id" db:"id"`
	JobTypeName string `json:"job_type_name" db:"job_type_name"`
}

type JobTypeUpdate struct {
	JobTypeName string `json:"job_type_name" db:"job_type_name"`
}
