package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

var YearID uint8
var UserTitles []UserTitle
var UserTitleDomains []UserTitleDomain
var UserTitleCampuses []UserTitleCampus
var UserTitleSchools []UserTitleSchool
var UserTitleClasses []UserTitleClass
var UserTitleBranches []UserTitleBranch
var UserTitleLessons []UserTitleLesson
var ActiveUsers map[uint32]bool
var Users map[uint32]User
var Titles map[uint8]Title
var TitleTypes map[uint8]TitleType = map[uint8]TitleType{
	1: {1, "Admin"},
	2: {2, "Reseller"},
	3: {3, "Genel Müdürlük"},
	4: {4, "Kampüs"},
	5: {5, "İdari"},
	6: {6, "Akademik"},
	7: {7, "Veli"},
	8: {8, "Öğrenci"},
	9: {9, "Atanmamış"},
}

type UserTitle struct {
	UserID    uint32
	TitleID   uint8
	YearID    uint8
	IsDefault bool
}

type UserTitleDomain struct {
	UserID   uint32
	TitleID  uint8
	YearID   uint8
	DomainID uint8
}

type UserTitleCampus struct {
	UserID   uint32
	TitleID  uint8
	YearID   uint8
	CampusID uint32
}

type UserTitleSchool struct {
	UserID   uint32
	TitleID  uint8
	YearID   uint8
	SchoolID uint32
}

type UserTitleClass struct {
	UserID  uint32
	TitleID uint8
	YearID  uint8
	ClassID uint32
}

type UserTitleBranch struct {
	UserID   uint32
	TitleID  uint8
	YearID   uint8
	BranchID uint32
}

type UserTitleLesson struct {
	UserID   uint32
	TitleID  uint8
	YearID   uint8
	LessonID uint32
}

type Title struct {
	TitleID uint8
	TypeID  uint8
	Title   string
}

type TitleType struct {
	TitleTypeID uint8
	Title       string
}

type ActiveUser struct {
	UserID uint32
}

type User struct {
	UserID             uint32
	DefaultTitleID     uint8
	ComputeTitleID     uint8
	ComputeTitleTypeID uint8
	Titles             map[uint8]Title
	NetworkUserIDS     map[uint32]bool
	DomainIDs          map[uint8]bool
	CampusIDs          map[uint32]bool
	SchoolIDs          map[uint32]bool
	ClassIDs           map[uint32]bool
	BranchIDs          map[uint32]bool
	LessonIDs          map[uint32]bool
	DomainIDArray      []uint8
	CampusIDArray      []uint32
	SchoolIDArray      []uint32
	ClassIDArray       []uint32
	BranchIDArray      []uint32
	LessonIDArray      []uint32
}

func FillStore(db *sql.DB) {
	FetchAllTitle(db)
	FetchAllActiveUser(db)
	FetchAllUserTitle(db)
	FetchAllUserTitleDomain(db)
	FetchAllUserTitleCampus(db)
	FetchAllUserTitleSchool(db)
	FetchAllUserTitleClass(db)
	FetchAllUserTitleBranch(db)
	FetchAllUserTitleLesson(db)
}

func FillUsers() {
	Users = make(map[uint32]User)
	for _, userTitle := range UserTitles {
		if !ActiveUsers[userTitle.UserID] {
			continue
		}

		if _, ok := Users[userTitle.UserID]; ok {
			user := Users[userTitle.UserID]

			if userTitle.IsDefault {
				user.DefaultTitleID = userTitle.TitleID
			}

			if TitleTypes[Titles[userTitle.TitleID].TypeID].TitleTypeID < user.ComputeTitleTypeID {
				user.ComputeTitleID = userTitle.TitleID
				user.ComputeTitleTypeID = TitleTypes[Titles[userTitle.TitleID].TypeID].TitleTypeID
			}

			user.Titles[userTitle.TitleID] = Titles[userTitle.TitleID]

			Users[userTitle.UserID] = user
		} else {
			Users[userTitle.UserID] = User{
				UserID:             userTitle.UserID,
				DefaultTitleID:     userTitle.TitleID,
				ComputeTitleID:     userTitle.TitleID,
				ComputeTitleTypeID: TitleTypes[Titles[userTitle.TitleID].TypeID].TitleTypeID,
				Titles:             map[uint8]Title{userTitle.TitleID: Titles[userTitle.TitleID]},
				DomainIDs:          make(map[uint8]bool),
				CampusIDs:          make(map[uint32]bool),
				SchoolIDs:          make(map[uint32]bool),
				ClassIDs:           make(map[uint32]bool),
				BranchIDs:          make(map[uint32]bool),
				LessonIDs:          make(map[uint32]bool),
				DomainIDArray:      make([]uint8, 0),
				CampusIDArray:      make([]uint32, 0),
				SchoolIDArray:      make([]uint32, 0),
				ClassIDArray:       make([]uint32, 0),
				BranchIDArray:      make([]uint32, 0),
				LessonIDArray:      make([]uint32, 0),
				NetworkUserIDS:     make(map[uint32]bool),
			}
		}
	}

	for _, row := range UserTitleDomains {
		if !ActiveUsers[row.UserID] || Users[row.UserID].DomainIDs == nil {
			continue
		}

		user := Users[row.UserID]
		if !user.DomainIDs[row.DomainID] {
			user.DomainIDs[row.DomainID] = true
			user.DomainIDArray = append(user.DomainIDArray, row.DomainID)
			Users[row.UserID] = user
		}

	}

	for _, row := range UserTitleCampuses {
		if !ActiveUsers[row.UserID] || Users[row.UserID].CampusIDs == nil {
			continue
		}

		user := Users[row.UserID]
		if !user.CampusIDs[row.CampusID] {

			user.CampusIDs[row.CampusID] = true
			user.CampusIDArray = append(user.CampusIDArray, row.CampusID)
			Users[row.UserID] = user
		}
	}

	for _, row := range UserTitleSchools {
		if !ActiveUsers[row.UserID] || Users[row.UserID].SchoolIDs == nil {
			continue
		}

		user := Users[row.UserID]
		if !user.SchoolIDs[row.SchoolID] {
			user.SchoolIDs[row.SchoolID] = true
			user.SchoolIDArray = append(user.SchoolIDArray, row.SchoolID)
			Users[row.UserID] = user
		}
	}

	for _, row := range UserTitleClasses {
		if !ActiveUsers[row.UserID] || Users[row.UserID].ClassIDs == nil {
			continue
		}

		user := Users[row.UserID]
		if !user.ClassIDs[row.ClassID] {
			user.ClassIDs[row.ClassID] = true
			user.ClassIDArray = append(user.ClassIDArray, row.ClassID)
			Users[row.UserID] = user
		}
	}

	for _, row := range UserTitleBranches {
		if !ActiveUsers[row.UserID] || Users[row.UserID].BranchIDs == nil {
			continue
		}

		user := Users[row.UserID]
		if !user.BranchIDs[row.BranchID] {
			user.BranchIDs[row.BranchID] = true
			user.BranchIDArray = append(user.BranchIDArray, row.BranchID)
			Users[row.UserID] = user
		}
	}

	for _, row := range UserTitleLessons {
		if !ActiveUsers[row.UserID] || Users[row.UserID].LessonIDs == nil {
			continue
		}

		user := Users[row.UserID]
		if !user.LessonIDs[row.LessonID] {
			user.LessonIDs[row.LessonID] = true
			user.LessonIDArray = append(user.LessonIDArray, row.LessonID)
			Users[row.UserID] = user
		}
	}
}

func ComputeUserNetworksIDs() {
	for _, user := range Users {
		if user.ComputeTitleTypeID == 0 {
			fmt.Println(user)
			panic("ComputeTitleTypeID is 0 for userID: " + strconv.Itoa(int(user.UserID)))
		}

		// Super Admin
		if Titles[user.ComputeTitleID].TypeID == 1 {
			continue
		}

		user.NetworkUserIDS[user.UserID] = true

		// Domain Level
		if user.ComputeTitleTypeID == 2 || user.ComputeTitleTypeID == 3 {
			for _, searchUser := range Users {
				if searchUser.ComputeTitleTypeID == 1 {
					user.NetworkUserIDS[searchUser.UserID] = true
				} else {
					for _, searchUserDomainID := range searchUser.DomainIDArray {
						if user.DomainIDs[searchUserDomainID] {
							user.NetworkUserIDS[searchUser.UserID] = true
						}
					}
				}
			}
		}

		// Campus Level
		if user.ComputeTitleTypeID == 4 {
			for _, searchUser := range Users {
				if searchUser.ComputeTitleTypeID == 1 {
					user.NetworkUserIDS[searchUser.UserID] = true
				} else if searchUser.ComputeTitleTypeID == 2 || searchUser.ComputeTitleTypeID == 3 {
					for _, searchUserDomainID := range searchUser.DomainIDArray {
						if user.DomainIDs[searchUserDomainID] {
							user.NetworkUserIDS[searchUser.UserID] = true
						}
					}
				} else {
					for _, searchUserCampusID := range searchUser.CampusIDArray {
						if user.CampusIDs[searchUserCampusID] {
							user.NetworkUserIDS[searchUser.UserID] = true
						}
					}
				}
			}
		}

		// School Level
		if user.ComputeTitleTypeID == 5 {
			for _, searchUser := range Users {
				if searchUser.ComputeTitleTypeID == 1 {
					user.NetworkUserIDS[searchUser.UserID] = true
				} else if searchUser.ComputeTitleTypeID == 2 || searchUser.ComputeTitleTypeID == 3 {
					for _, searchUserDomainID := range searchUser.DomainIDArray {
						if user.DomainIDs[searchUserDomainID] {
							user.NetworkUserIDS[searchUser.UserID] = true
						}
					}
				} else if searchUser.ComputeTitleTypeID == 4 {
					for _, searchUserCampusID := range searchUser.CampusIDArray {
						if user.CampusIDs[searchUserCampusID] {
							user.NetworkUserIDS[searchUser.UserID] = true
						}
					}
				} else {
					for _, searchUserSchoolID := range searchUser.SchoolIDArray {
						if user.SchoolIDs[searchUserSchoolID] {
							user.NetworkUserIDS[searchUser.UserID] = true
						}
					}
				}
			}
		}

		// Branch Level
		if user.ComputeTitleTypeID == 6 {
			for _, searchUser := range Users {
				if searchUser.ComputeTitleTypeID == 1 {
					user.NetworkUserIDS[searchUser.UserID] = true
				} else if searchUser.ComputeTitleTypeID == 2 || searchUser.ComputeTitleTypeID == 3 {
					for _, searchUserDomainID := range searchUser.DomainIDArray {
						if user.DomainIDs[searchUserDomainID] {
							user.NetworkUserIDS[searchUser.UserID] = true
						}
					}
				} else if searchUser.ComputeTitleTypeID == 4 {
					for _, searchUserCampusID := range searchUser.CampusIDArray {
						if user.CampusIDs[searchUserCampusID] {
							user.NetworkUserIDS[searchUser.UserID] = true
						}
					}
				} else {
					for _, searchUserBranchID := range searchUser.BranchIDArray {
						if user.BranchIDs[searchUserBranchID] {
							user.NetworkUserIDS[searchUser.UserID] = true
						}
					}
				}
			}
		}

		Users[user.UserID] = user
	}
}

func FetchAllUserTitle(db *sql.DB) (*[]UserTitle, error) {
	rowCount, rowCountErr := FetchAllUserTitleCount(db)
	if rowCountErr != nil {
		return nil, rowCountErr
	}

	// convert rowCount to string

	UserTitles = make([]UserTitle, rowCount)

	results, err := db.Query("SELECT user_id, title_id, year_id, is_default FROM user_title WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return nil, err
	}

	var i = 0
	for results.Next() {
		var userTitle UserTitle
		err = results.Scan(&userTitle.UserID, &userTitle.TitleID, &userTitle.YearID, &userTitle.IsDefault)
		if err != nil {
			return nil, err
		}

		UserTitles[i] = userTitle
		i++
	}

	return &UserTitles, err
}

func FetchAllUserTitleCount(db *sql.DB) (int, error) {
	rowCountResult, err := db.Query("SELECT COUNT(*) as count FROM user_title WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return 0, err
	}

	rowCountResult.Next()
	var rowCount int
	err = rowCountResult.Scan(&rowCount)
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}

func FetchAllUserTitleDomain(db *sql.DB) (*[]UserTitleDomain, error) {
	rowCount, rowCountErr := FetchAllUserTitleDomainCount(db)
	if rowCountErr != nil {
		return nil, rowCountErr
	}

	UserTitleDomains = make([]UserTitleDomain, rowCount)

	results, err := db.Query("SELECT user_id, title_id, year_id, domain_id FROM user_title_domain WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return nil, err
	}

	var i = 0
	for results.Next() {
		var row UserTitleDomain
		err = results.Scan(&row.UserID, &row.TitleID, &row.YearID, &row.DomainID)
		if err != nil {
			return nil, err
		}

		UserTitleDomains[i] = row
		i++
	}

	return &UserTitleDomains, err
}

func FetchAllUserTitleDomainCount(db *sql.DB) (int, error) {
	rowCountResult, err := db.Query("SELECT COUNT(*) as count FROM user_title_domain WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return 0, err
	}

	rowCountResult.Next()
	var rowCount int
	err = rowCountResult.Scan(&rowCount)
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}

func FetchAllUserTitleCampus(db *sql.DB) (*[]UserTitleCampus, error) {
	rowCount, rowCountErr := FetchAllUserTitleCampusCount(db)
	if rowCountErr != nil {
		return nil, rowCountErr
	}

	UserTitleCampuses = make([]UserTitleCampus, rowCount)

	results, err := db.Query("SELECT user_id, title_id, year_id, campus_id FROM user_title_campus WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return nil, err
	}

	var i = 0
	for results.Next() {
		var row UserTitleCampus
		err = results.Scan(&row.UserID, &row.TitleID, &row.YearID, &row.CampusID)
		if err != nil {
			return nil, err
		}

		UserTitleCampuses[i] = row
		i++
	}

	return &UserTitleCampuses, err
}

func FetchAllUserTitleCampusCount(db *sql.DB) (int, error) {
	rowCountResult, err := db.Query("SELECT COUNT(*) as count FROM user_title_campus WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return 0, err
	}

	rowCountResult.Next()
	var rowCount int
	err = rowCountResult.Scan(&rowCount)
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}

func FetchAllUserTitleSchool(db *sql.DB) (*[]UserTitleSchool, error) {
	rowCount, rowCountErr := FetchAllUserTitleSchoolCount(db)
	if rowCountErr != nil {
		return nil, rowCountErr
	}

	UserTitleSchools = make([]UserTitleSchool, rowCount)

	results, err := db.Query("SELECT user_id, title_id, year_id, school_id FROM user_title_school WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return nil, err
	}

	var i = 0
	for results.Next() {
		var row UserTitleSchool
		err = results.Scan(&row.UserID, &row.TitleID, &row.YearID, &row.SchoolID)
		if err != nil {
			return nil, err
		}

		UserTitleSchools[i] = row
		i++
	}

	return &UserTitleSchools, err
}

func FetchAllUserTitleSchoolCount(db *sql.DB) (int, error) {
	rowCountResult, err := db.Query("SELECT COUNT(*) as count FROM user_title_school WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return 0, err
	}

	rowCountResult.Next()
	var rowCount int
	err = rowCountResult.Scan(&rowCount)
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}

func FetchAllUserTitleClass(db *sql.DB) (*[]UserTitleClass, error) {
	rowCount, rowCountErr := FetchAllUserTitleClassCount(db)
	if rowCountErr != nil {
		return nil, rowCountErr
	}

	UserTitleClasses = make([]UserTitleClass, rowCount)

	results, err := db.Query("SELECT user_id, title_id, year_id, class_id FROM user_title_class WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return nil, err
	}

	var i = 0
	for results.Next() {
		var row UserTitleClass
		err = results.Scan(&row.UserID, &row.TitleID, &row.YearID, &row.ClassID)
		if err != nil {
			return nil, err
		}

		UserTitleClasses[i] = row
		i++
	}

	return &UserTitleClasses, err
}

func FetchAllUserTitleClassCount(db *sql.DB) (int, error) {
	rowCountResult, err := db.Query("SELECT COUNT(*) as count FROM user_title_class WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return 0, err
	}

	rowCountResult.Next()
	var rowCount int
	err = rowCountResult.Scan(&rowCount)
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}

func FetchAllUserTitleBranch(db *sql.DB) (*[]UserTitleBranch, error) {
	rowCount, rowCountErr := FetchAllUserTitleBranchCount(db)
	if rowCountErr != nil {
		return nil, rowCountErr
	}

	UserTitleBranches = make([]UserTitleBranch, rowCount)

	results, err := db.Query("SELECT user_id, title_id, year_id, branch_id FROM user_title_branch WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return nil, err
	}

	var i = 0
	for results.Next() {
		var row UserTitleBranch
		err = results.Scan(&row.UserID, &row.TitleID, &row.YearID, &row.BranchID)
		if err != nil {
			return nil, err
		}

		UserTitleBranches[i] = row
		i++
	}

	return &UserTitleBranches, err
}

func FetchAllUserTitleBranchCount(db *sql.DB) (int, error) {
	rowCountResult, err := db.Query("SELECT COUNT(*) as count FROM user_title_branch WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return 0, err
	}

	rowCountResult.Next()
	var rowCount int
	err = rowCountResult.Scan(&rowCount)
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}

func FetchAllUserTitleLesson(db *sql.DB) (*[]UserTitleLesson, error) {
	rowCount, rowCountErr := FetchAllUserTitleLessonCount(db)
	if rowCountErr != nil {
		return nil, rowCountErr
	}

	UserTitleLessons = make([]UserTitleLesson, rowCount)

	results, err := db.Query("SELECT user_id, title_id, year_id, lesson_id FROM user_title_lesson WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return nil, err
	}

	var i = 0
	for results.Next() {
		var row UserTitleLesson
		err = results.Scan(&row.UserID, &row.TitleID, &row.YearID, &row.LessonID)
		if err != nil {
			return nil, err
		}

		UserTitleLessons[i] = row
		i++
	}

	return &UserTitleLessons, err
}

func FetchAllUserTitleLessonCount(db *sql.DB) (int, error) {
	rowCountResult, err := db.Query("SELECT COUNT(*) as count FROM user_title_lesson WHERE year_id = " + strconv.Itoa(int(YearID)))
	if err != nil {
		return 0, err
	}

	rowCountResult.Next()
	var rowCount int
	err = rowCountResult.Scan(&rowCount)
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}

func FetchAllTitle(db *sql.DB) (*map[uint8]Title, error) {

	rowCount, rowCountErr := FetchAllTitleCount(db)
	if rowCountErr != nil {
		return nil, rowCountErr
	}

	Titles = make(map[uint8]Title, uint8(rowCount))
	results, err := db.Query("SELECT id, type_id, title FROM title WHERE status_id = 1")

	if err != nil {
		return nil, err
	}

	for results.Next() {
		var row Title
		err = results.Scan(&row.TitleID, &row.TypeID, &row.Title)

		if err != nil {
			return nil, err
		}

		Titles[row.TitleID] = row
	}

	return &Titles, err
}

func FetchAllTitleCount(db *sql.DB) (int, error) {
	rowCountResult, err := db.Query("SELECT COUNT(*) as count FROM title WHERE status_id = 1")
	if err != nil {
		return 0, err
	}

	rowCountResult.Next()
	var rowCount int
	err = rowCountResult.Scan(&rowCount)
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}

func FetchAllActiveUser(db *sql.DB) (*map[uint32]bool, error) {

	rowCount, rowCountErr := FetchAllActiveUserCount(db)
	if rowCountErr != nil {
		return nil, rowCountErr
	}

	ActiveUsers = make(map[uint32]bool, uint32(rowCount))
	results, err := db.Query("SELECT id FROM user WHERE status_id = 1")

	if err != nil {
		return nil, err
	}

	for results.Next() {
		var row ActiveUser
		err = results.Scan(&row.UserID)

		if err != nil {
			return nil, err
		}

		ActiveUsers[uint32(row.UserID)] = true
	}

	return &ActiveUsers, err
}

func FetchAllActiveUserCount(db *sql.DB) (int, error) {
	rowCountResult, err := db.Query("SELECT COUNT(*) as count FROM user WHERE status_id = 1")
	if err != nil {
		return 0, err
	}

	rowCountResult.Next()
	var rowCount int
	err = rowCountResult.Scan(&rowCount)
	if err != nil {
		return 0, err
	}

	return rowCount, nil
}
