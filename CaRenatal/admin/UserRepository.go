package user

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/amalica/CarRental/entities"
)

// Repository is a struct that define the Repository type.
type Repository struct {
	connection *sql.DB
}

// MatchTag is a struct that define a type of a match tag returned from database.
type MatchTag struct {
	UID         string
	Category    string
	Subcategory string
	WorkType    int
}

// IRepository is an interface that specifies database operations on User type.
type IRepository interface {
	AddUser(*entities.User) error
	SearchUser(string) *entities.User
	UpdateUser(*entities.User) error
	RegisterTPUsers(string, string, string, string) error
	UpdateTPUsers(string, string) error
	SearchTPUser(string) (*entities.User, string)
	CountMember(string) int
	RemoveFileDB(string, string) string
	RemoveFile(string, string) error
	AddMatchTag(uid string, category string, subcategory string, worktype int) error
	RemoveMatchTag(uid string, category string, subcategory string, worktype int) error
	GetUserMatchTags(uid string) []*MatchTag
	SearchProjectWMatchTag(matchTag *MatchTag) []*entities.Project
	SearchMember(tableName string, columnValue string) bool
	GetOwner(string) string
}

// NewRepository is a function that return new IRepository type.
func NewRepository(conn *sql.DB) IRepository {
	return &Repository{connection: conn}
}

// AddUser is a method that adds a user to the provided database.
func (psql *Repository) AddUser(user *entities.User) error {

	totalNumOfUsers := psql.CountMember("users")

	user.UID = fmt.Sprintf("UID%d", totalNumOfUsers+1)

	stmt, _ := psql.connection.Prepare(`INSERT INTO Users (uid,first_name, last_name, password, phonenumber, email, 
		job_title, country, city, gender,cv, profile_pic, bio, rating) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	_, err := stmt.Exec(
		user.UID,
		user.Firstname,
		user.Lastname,
		user.Password,
		user.Phonenumber,
		user.Email,
		user.JobTitle,
		user.Country,
		user.City,
		user.Gender,
		user.CV,
		user.ProfilePic,
		user.Bio,
		user.Rating)

	if err != nil {
		return err
	}
	return nil
}

// SearchUser is a method that search for a User type from the 'users' table in the database using the provided Email address.
func (psql *Repository) SearchUser(identifier string) *entities.User {
	stmt, err := psql.connection.Prepare("SELECT * FROM Users WHERE email=? || uid=?")
	row := stmt.QueryRow(identifier, identifier)
	if err != nil {
		panic(err)
	}
	var user entities.User
	row.Scan(
		&user.UID,
		&user.Firstname,
		&user.Lastname,
		&user.Password,
		&user.Phonenumber,
		&user.Email,
		&user.JobTitle,
		&user.Country,
		&user.City,
		&user.Gender,
		&user.CV,
		&user.ProfilePic,
		&user.Bio,
		&user.Rating)

	return &user
}

// UpdateUser is a method that updates the user data in the database using the provided parameters.
func (psql *Repository) UpdateUser(user *entities.User) error {
	stmt, _ := psql.connection.Prepare(`UPDATE Users SET first_name = ?, last_name = ?, phonenumber = ?, email = ?, job_title = ?,
		country = ?, city = ?, gender = ?, cv = ?, profile_pic = ?, bio = ? WHERE uid = ?`)
	_, err := stmt.Exec(
		user.Firstname,
		user.Lastname,
		user.Phonenumber,
		user.Email,
		user.JobTitle,
		user.Country,
		user.City,
		user.Gender,
		user.CV,
		user.ProfilePic,
		user.Bio,
		user.UID)

	if err != nil {
		return err
	}
	return nil

}

// RegisterTPUsers is a method that adds a third party user to the provided database.
func (psql *Repository) RegisterTPUsers(uid, email, password, from string) error {
	stmt, _ := psql.connection.Prepare(`INSERT INTO tp_users (uid, email, password, origin) VALUES (?,?,?,?)`)
	_, err := stmt.Exec(uid, email, password, from)

	if err != nil {
		return err
	}
	return nil
}

// UpdateTPUsers is a method that updates the third party user data in the database using the provided parameters.
func (psql *Repository) UpdateTPUsers(uid, email string) error {
	stmt, _ := psql.connection.Prepare(`UPDATE tp_users SET email=? WHERE uid=?`)
	_, err := stmt.Exec(email, uid)

	if err != nil {
		return err
	}
	return nil
}

// SearchTPUser is a method that search for a User type from the 'tp_users' table in database using the provided Email address.
func (psql *Repository) SearchTPUser(identifier string) (*entities.User, string) {

	stmt, err := psql.connection.Prepare("SELECT * FROM tp_users WHERE email=? || uid=?")
	row := stmt.QueryRow(identifier, identifier)
	if err != nil {
		panic(err)
	}
	var from string
	var user entities.User
	row.Scan(
		&user.UID,
		&user.Email,
		&user.Password,
		&from)
	return &user, from
}

// CountMember is a method that is used for counting the member of a table where our table name is provided as an argument.
func (psql *Repository) CountMember(tableName string) (totalNumOfMembers int) {

	stmt, err := psql.connection.Prepare("SELECT COUNT(*) FROM " + tableName)
	if err != nil {
		return
	}
	row := stmt.QueryRow()
	row.Scan(&totalNumOfMembers)
	return

}

// AddMatchTag is a method that adds a matchTag to the match tag table.
func (psql *Repository) AddMatchTag(uid string, category string, subcategory string, worktype int) error {

	stmt, _ := psql.connection.Prepare(`INSERT INTO match_tag_table (uid,category,subcategory,worktype
		) VALUES (?,?,?,?)`)
	_, err := stmt.Exec(
		uid,
		category,
		subcategory,
		worktype,
	)

	if err != nil {
		panic(err)
	}
	return nil
}

// RemoveMatchTag is a method that removes a matchTag from the match tag table.
func (psql *Repository) RemoveMatchTag(uid string, category string, subcategory string, worktype int) error {

	stmt, _ := psql.connection.Prepare(`DELETE FROM match_tag_table WHERE uid=? && category=? && 
	subcategory=? && worktype=?`)
	_, err := stmt.Exec(
		uid,
		category,
		subcategory,
		worktype,
	)

	if err != nil {
		panic(err)
	}
	return nil
}

// GetUserMatchTags is a method that returns all the match tag the user defined.
func (psql *Repository) GetUserMatchTags(uid string) []*MatchTag {

	stmt, err := psql.connection.Prepare("SELECT * FROM match_tag_table WHERE uid=?")
	rows, err := stmt.Query(uid)
	if err != nil {
		panic(err)
	}

	var matchTagStrore []*MatchTag
	for rows.Next() {

		matchTag := new(MatchTag)

		rows.Scan(
			&matchTag.UID,
			&matchTag.Category,
			&matchTag.Subcategory,
			&matchTag.WorkType)
		matchTagStrore = append(matchTagStrore, matchTag)
	}

	return matchTagStrore
}

// SearchProjectWMatchTag is a method that returns all the projects that match the provided tag.
func (psql *Repository) SearchProjectWMatchTag(matchTag *MatchTag) []*entities.Project {

	query := `SELECT * FROM projects WHERE`
	if matchTag.Category != "" {
		query += fmt.Sprintf(` category='%s'`, matchTag.Category)
	}
	if matchTag.Subcategory != "" {
		query += fmt.Sprintf(` && subcategory='%s'`, matchTag.Subcategory)
	}
	if matchTag.WorkType != 4 {
		query += fmt.Sprintf(` && worktype=%d`, matchTag.WorkType)
	}

	stmt, err := psql.connection.Prepare(query)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	var projects []*entities.Project

	for rows.Next() {
		project := new(entities.Project)
		err = rows.Scan(&project.ID, &project.Title,
			&project.Description, &project.Details,
			&project.Category, &project.Subcategory,
			&project.Budget, &project.WorkType,
			&project.Closed, &project.CreatedAt)

		if err != nil {
			panic(err)
		}

		projects = append(projects, project)
	}

	return projects

}

// RemoveFileDB is a method that is used for clearing a certain column.
func (psql *Repository) RemoveFileDB(uid string, columnName string) string {

	ctx := context.Background()
	var value string

	stmt := `SELECT ` + columnName + ` FROM users WHERE uid = ?`
	row := psql.connection.QueryRowContext(ctx, stmt, uid)

	row.Scan(&value)

	stmt = `UPDATE users SET ` + columnName + `="" WHERE uid = ?`
	_, err := psql.connection.QueryContext(ctx, stmt, uid)

	if err != nil {
		return ""
	}

	return value

}

// RemoveFile is a method that removes a given file path from the assets folder.
func (psql *Repository) RemoveFile(filename string, folder string) error {

	if err := os.Remove("./assets/" + folder + "/" + filename); err != nil {
		return err
	}

	return nil

}

// SearchMember is a method that is used for searching the member of a table where our table name is provided as an argument.
func (psql *Repository) SearchMember(tableName string, columnValue string) bool {

	stmt, _ := psql.connection.Prepare("")
	if tableName == "attached_files" || tableName == "categories" || tableName == "subcategories" {
		stmt, _ = psql.connection.Prepare("SELECT COUNT(*) FROM " + tableName + " WHERE name=?")
	}

	if tableName == "projects" {
		stmt, _ = psql.connection.Prepare("SELECT COUNT(*) FROM " + tableName + " WHERE id=?")
	}

	var totalNumOfMembers int
	row := stmt.QueryRow(columnValue)
	row.Scan(&totalNumOfMembers)

	if totalNumOfMembers > 0 {
		return true
	}
	return false
}

// GetOwner is a method that returns the owner of a project by searching through the project user table.
func (psql *Repository) GetOwner(pid string) string {
	stmt, _ := psql.connection.Prepare("SELECT uid FROM user_project_table WHERE pid=?")

	var uid string
	row := stmt.QueryRow(pid)
	row.Scan(&uid)

	return uid
}
