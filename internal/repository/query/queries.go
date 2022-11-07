package query

const (
	RegisterUser = `INSERT INTO users (name, username, email, password, image) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, username, email, image`
	GetUser      = `SELECT id, username, email, image FROM users WHERE email=$1 AND password=$2`

	StoreJobType   = `INSERT INTO job_types (job_type_name) VALUES ($1) RETURNING id, job_type_name`
	GetAllJobTypes = `SELECT id, job_type_name FROM job_types`
	ShowJobType    = `SELECT id, job_type_name FROM job_types WHERE id=$1`
	DeleteJobType  = `DELETE FROM job_types WHERE id=$1`
	UpdatedJobType = `UPDATE job_types SET job_type_name=$1 WHERE id=$2`

	StoreTeam           = `INSERT INTO teams (team_name) VALUES ($1) RETURNING id, team_name`
	GetAllTeams         = `SELECT id, team_name FROM teams`
	ShowTeam            = `SELECT id, team_name FROM teams WHERE id=$1`
	DeleteTeam          = `DELETE FROM teams WHERE id=$1`
	UpdatedTeam         = `UPDATE teams SET team_name=$1 WHERE id=$2`
	AddUserToTeam       = `INSERT INTO teams_users(team_id, user_id) VALUES ($1, $2)`
	RemoveUsersFromTeam = `DELETE FROM teams_users WHERE team_id = $1 AND user_id = $2`

	StoreTag            = `INSERT INTO tags (tag_name) VALUES ($1) RETURNING id, tag_name`
	GetAllTags          = `SELECT id, tag_name FROM tags`
	ShowTag             = `SELECT id, tag_name FROM tags WHERE id=$1`
	DeleteTag           = `DELETE FROM tags WHERE id=$1`
	UpdatedTag          = `UPDATE tags SET tag_name=$1 WHERE id=$2`
	AttachTeamToBug     = `INSERT INTO bugs_teams(bug_id, team_id) VALUEST ($1, $2)`
	DetachTeamFromBug   = `DELETE FROM bugs_teams WHERE bug_id = $1 AND team_id = $2`
	AttachUserToBug     = `INSERT INTO bugs_users(bug_id, user_id) VALUES ($1, $2)`
	DeattachUserFromBug = `DELETE FROM bugs_users WHERE bug_id = $1 AND user_id = $2`

	StoreCategory    = `INSERT INTO categories (category_name) VALUES ($1) RETURNING id, category_name`
	GetAllCategories = `SELECT id, category_name FROM categories`
	ShowCategory     = `SELECT id, category_name FROM categories WHERE id=$1`
	DeleteCategory   = `DELETE FROM categories WHERE id=$1`
	UpdatedCategory  = `UPDATE categories SET category_name=$1 WHERE id=$2`

	StoreStatus    = `INSERT INTO statuses (status_name) VALUES ($1) RETURNING id, status_name`
	GetAllStatuses = `SELECT id, status_name FROM statuses`
	ShowStatus     = `SELECT id, status_name FROM statuses WHERE id=$1`
	DeleteStatus   = `DELETE FROM statuses WHERE id=$1`
	UpdatedStatus  = `UPDATE statuses SET status_name=$1 WHERE id=$2`

	StoreProject                = `INSERT INTO projects (project_name) VALUES ($1) RETURNING id, project_name`
	GetAllProjects              = `SELECT id, project_name FROM projects`
	ShowProject                 = `SELECT id, project_name FROM projects WHERE id=$1`
	DeleteProject               = `DELETE FROM projects WHERE id=$1`
	UpdatedProject              = `UPDATE projects SET project_name=$1 WHERE id=$2`
	AddUserToProject            = `INSERT INTO projects_users(poroject_id, user_id) VALUES ($1, $2)`
	AddTeamToProject            = `INSERT INTO projects_teams(project_id, team_id) VALUES ($1, $2)`
	SelectUsersIdFromTeamsUsers = `SELECT user_id FROM teams_users WHERE team_id=$1`
	GetBugsByProjectId          = `SELECT b.id, b.bug_title, b.bug_description, b.is_completed, b.status_id FROM bugs b INNER JOIN bugs_users u ON b.id = u.bug_id WHERE user_id = $1 AND b.project_id = $2`
	GetProjectThatAttacedUser   = `SELECT p.id, p.project_name FROM projects p INNER JOIN projects_users pu ON p.id = pu.poroject_id WHERE user_id=$1`

	StoreBug = `INSERT INTO bugs (bug_title, bug_description, status_id, category_id, project_id) VALUES ($1, $2, $3, $4, $5)
				RETURNING id, bug_title, bug_description, is_completed, status_id, category_id, project_id`
	GetAllBugs     = `SELECT s.status_name, b.id, b.bug_title, b.bug_description, b.is_completed, b.status_id FROM statuses s INNER JOIN bugs b ON s.id = b.id`
	CloseIssue     = `UPDATE bugs SET is_completed=true WHERE id=$1`
	AddTag         = `INSERT INTO bugs_tags(bug_id, tag_id) VALUES ($1, $2)`
	GetTagsByBugId = `SELECT t.id, t.tag_name FROM tags t  INNER JOIN bugs_tags on t.id = bugs_tags.tag_id WHERE bug_id = $1`
	GetTeamUsers   = `SELECT u.id, u.name, u.username, u.email, u.image, u.role FROM users u INNER JOIN teams_users t ON u.id = t.user_id WHERE team_id = $1`
)
