package query

const (
	StoreJobType   = `INSERT INTO job_types (job_type_name) VALUES ($1) RETURNING id, job_type_name`
	GetAllJobTypes = `SELECT id, job_type_name FROM job_types`
	ShowJobType    = `SELECT id, job_type_name FROM job_types WHERE id=$1`
	DeleteJobType  = `DELETE FROM job_types WHERE id=$1`
	UpdatedJobType = `UPDATE job_types SET job_type_name=$1 WHERE id=$2`

	StoreTeam   = `INSERT INTO teams (team_name) VALUES ($1) RETURNING id, team_name`
	GetAllTeams = `SELECT id, team_name FROM teams`
	ShowTeam    = `SELECT id, team_name FROM teams WHERE id=$1`
	DeleteTeam  = `DELETE FROM teams WHERE id=$1`
	UpdatedTeam = `UPDATE teams SET team_name=$1 WHERE id=$2`

	StoreTag   = `INSERT INTO tags (tag_name) VALUES ($1) RETURNING id, tag_name`
	GetAllTags = `SELECT id, tag_name FROM tags`
	ShowTag    = `SELECT id, tag_name FROM tags WHERE id=$1`
	DeleteTag  = `DELETE FROM tags WHERE id=$1`
	UpdatedTag = `UPDATE tags SET tag_name=$1 WHERE id=$2`

	StoreCategory    = `INSERT INTO categories (category_name) VALUES ($1) RETURNING id, category_name`
	GetAllCategories = `SELECT id, category_name FROM categories`
	ShowCategory     = `SELECT id, category_name FROM categories WHERE id=$1`
	DeleteCategory   = `DELETE FROM categories WHERE id=$1`
	UpdatedCategory  = `UPDATE categories SET category_name=$1 WHERE id=$2`

	StoreStatus    = `INSERT INTO statuses (category_name) VALUES ($1) RETURNING id, status_name`
	GetAllStatuses = `SELECT id, status_name FROM statuses`
	ShowStatus     = `SELECT id, status_name FROM statuses WHERE id=$1`
	DeleteStatus   = `DELETE FROM statuses WHERE id=$1`
	UpdatedStatus  = `UPDATE statuses SET status_name=$1 WHERE id=$2`
)
