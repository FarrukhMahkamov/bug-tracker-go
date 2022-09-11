package dto

type Bug struct {
	Id             int64  `json:"id" db:"id"`
	BugTitle       string `json:"bug_title" db:"bug_title"`
	BugDescription string `json:"bug_description" db:"bug_description"`
	IsCompleted    bool   `json:"is_comleted" db:"is_completed"`
	CategoryId     int    `json:"category_id" db:"category_id"`
	StatusId       int    `json:"status_id" db:"status_id"`
}

type AllBugs struct {
	Id             int64  `json:"id" db:"id"`
	BugTitle       string `json:"bug_title" db:"bug_title"`
	BugDescription string `json:"bug_description" db:"bug_description"`
	IsCompleted    bool   `json:"is_comleted" db:"is_completed"`
	StatusId       int    `json:"status_id" db:"status_id"`
	StatusName     string `json:"status_name" db:"status_name"`
}

type BugUpdate struct {
	Id             int64  `json:"id" db:"id"`
	BugTitle       string `json:"bug_title" db:"bug_title"`
	BugDescription string `json:"bug_description" db:"bug_description"`
	IsCompleted    bool   `json:"is_comleted" db:"is_completed"`
	CategoryId     int    `json:"category_id" db:"category_id"`
	StatusId       int    `json:"status_id" db:"status_id"`
}

type BugTag struct {
	TagId []int `json:"tag_id"`
}
