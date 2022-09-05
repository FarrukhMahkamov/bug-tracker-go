package query

const (
	StoreJobType   = `INSERT INTO job_types (job_type_name) VALUES ($1) RETURNING id, job_type_name`
	GetAllJobTypes = `SELECT id, job_type_name FROM job_types`
	ShowJobType    = `SELECT id, job_type_name FROM job_types WHERE id=$1`
	DeleteJobType  = `DELETE FROM job_types WHERE id=$1`
	UpdatedJobType = `UPDATE job_types SET job_type_name=$1 WHERE id=$2`
)
