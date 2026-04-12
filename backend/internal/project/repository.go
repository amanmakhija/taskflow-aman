package project

import (
	"context"
	"taskflow/internal/db"
)

func CreateProject(p *Project) error {
	query := `
	INSERT INTO projects (id, name, description, owner_id)
	VALUES ($1, $2, $3, $4)
	`

	_, err := db.Pool.Exec(context.Background(),
		query,
		p.ID,
		p.Name,
		p.Description,
		p.OwnerID,
	)

	return err
}

func GetProjectsByUser(userID string) ([]Project, error) {
	query := `
	SELECT id, name, description, owner_id, created_at
	FROM projects
	WHERE owner_id = $1
	`

	rows, err := db.Pool.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var p Project
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.OwnerID, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func GetProjectByID(id string) (*Project, error) {
	query := `
	SELECT id, name, description, owner_id, created_at
	FROM projects WHERE id = $1
	`

	row := db.Pool.QueryRow(context.Background(), query, id)

	var p Project
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.OwnerID, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func UpdateProject(p *Project) error {
	query := `
	UPDATE projects
	SET name = $1, description = $2
	WHERE id = $3
	`

	_, err := db.Pool.Exec(context.Background(),
		query,
		p.Name,
		p.Description,
		p.ID,
	)

	return err
}

func DeleteProject(id string) error {
	query := `DELETE FROM projects WHERE id = $1`

	_, err := db.Pool.Exec(context.Background(), query, id)
	return err
}

func GetProjectWithTasks(projectID string) (*Project, []map[string]interface{}, error) {
	project, err := GetProjectByID(projectID)
	if err != nil {
		return nil, nil, err
	}

	query := `
	SELECT id, title, status, priority, assignee_id, due_date
	FROM tasks WHERE project_id = $1
	`

	rows, err := db.Pool.Query(context.Background(), query, projectID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var tasks []map[string]interface{}

	for rows.Next() {
		var (
			id, title, status, priority string
			assigneeID                  *string
			dueDate                     *string
		)

		err := rows.Scan(&id, &title, &status, &priority, &assigneeID, &dueDate)
		if err != nil {
			return nil, nil, err
		}

		tasks = append(tasks, map[string]interface{}{
			"id":          id,
			"title":       title,
			"status":      status,
			"priority":    priority,
			"assignee_id": assigneeID,
			"due_date":    dueDate,
		})
	}

	return project, tasks, nil
}
