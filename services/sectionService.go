package services

import (
	"onez19/config"
	"onez19/models"
)

func GetAllSectionsByWorkspaceID(workspaceID string) ([]models.Section, error) {
	var sections []models.Section

	// ดึงข้อมูล sections จากฐานข้อมูลตาม workspace_id
	rows, err := config.DB.Query("SELECT id, workspace_id, name FROM section WHERE workspace_id = ?", workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var section models.Section
		if err := rows.Scan(&section.ID, &section.WorkspaceID, &section.Name); err != nil {
			return nil, err
		}
		sections = append(sections, section)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sections, nil
}

func CreateSection(section models.Section) error {
	_, err := config.DB.Exec("INSERT INTO section (workspace_id, name) VALUES (?, ?)", section.WorkspaceID, section.Name)
	return err
}

func EditSectionName(sectionID int, newName string) error {
	_, err := config.DB.Exec("UPDATE section SET name = ? WHERE id = ?", newName, sectionID)
	return err
}
