package query

const (
	StoreJobType   = `INSERT INTO job_types (job_type_name) VALUES ($1) RETURNING id, job_type_name`
	GetAllJobTypes = `SELECT * FROM job_types`
)
