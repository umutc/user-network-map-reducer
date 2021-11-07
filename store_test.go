package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	mysqlDB, err := sql.Open("mysql", "root:1002@tcp(localhost:3306)/app")
	if err != nil {
		defer mysqlDB.Close()
	}

	db = mysqlDB

	if err != nil {
		panic(err.Error())
	}
}

func BenchmarkStoreTitle(b *testing.B) {
	b.ReportAllocs()
	rows, err := FetchAllUserTitle(db)
	if err != nil {
		panic(err.Error())
	}

	recordLength := len(*rows)
	b.ReportMetric(float64(recordLength), "item")
	if recordLength < 1 {
		b.Errorf("Expected at least 1 user, got %d", recordLength)
	}

}

func BenchmarkStoreTitleDomain(b *testing.B) {
	b.ReportAllocs()
	rows, err := FetchAllUserTitleDomain(db)
	if err != nil {
		panic(err.Error())
	}
	recordLength := len(*rows)
	b.ReportMetric(float64(recordLength), "item")

	if recordLength < 1 {
		b.Errorf("Expected at least 1 user, got %d", recordLength)
	}

}

func BenchmarkStoreTitleCampus(b *testing.B) {
	b.ReportAllocs()
	rows, err := FetchAllUserTitleCampus(db)
	if err != nil {
		panic(err.Error())
	}
	recordLength := len(*rows)
	b.ReportMetric(float64(recordLength), "item")

	if recordLength < 1 {
		b.Errorf("Expected at least 1 user, got %d", recordLength)
	}

}

func BenchmarkStoreTitleSchool(b *testing.B) {
	b.ReportAllocs()
	rows, err := FetchAllUserTitleSchool(db)
	if err != nil {
		panic(err.Error())
	}
	recordLength := len(*rows)
	b.ReportMetric(float64(recordLength), "item")

	if recordLength < 1 {
		b.Errorf("Expected at least 1 user, got %d", recordLength)
	}

}

func BenchmarkStoreTitleClass(b *testing.B) {
	b.ReportAllocs()
	rows, err := FetchAllUserTitleClass(db)
	if err != nil {
		panic(err.Error())
	}
	recordLength := len(*rows)
	b.ReportMetric(float64(recordLength), "item")

	if recordLength < 1 {
		b.Errorf("Expected at least 1 user, got %d", recordLength)
	}

}

func BenchmarkStoreTitleBranch(b *testing.B) {
	b.ReportAllocs()
	rows, err := FetchAllUserTitleBranch(db)
	if err != nil {
		panic(err.Error())
	}
	recordLength := len(*rows)
	b.ReportMetric(float64(recordLength), "item")

	if recordLength < 1 {
		b.Errorf("Expected at least 1 user, got %d", recordLength)
	}

}

func BenchmarkStoreTitleLesson(b *testing.B) {
	b.ReportAllocs()
	rows, err := FetchAllUserTitleLesson(db)
	if err != nil {
		panic(err.Error())
	}
	recordLength := len(*rows)
	b.ReportMetric(float64(recordLength), "item")

	if recordLength < 1 {
		b.Errorf("Expected at least 1 user, got %d", recordLength)
	}

}

func BenchmarkFillStore(b *testing.B) {
	b.ReportAllocs()
	FillStore(db)

	TitlesLen := len(Titles)
	UserTitlesLen := len(UserTitles)
	UserTitleDomainsLen := len(UserTitleDomains)
	UserTitleCampusesLen := len(UserTitleCampuses)
	UserTitleSchoolsLen := len(UserTitleSchools)
	UserTitleClassesLen := len(UserTitleClasses)
	UserTitleBranchesLen := len(UserTitleBranches)
	UserTitleLessonsLen := len(UserTitleLessons)

	b.ReportMetric(float64(TitlesLen), "TitlesLen")
	b.ReportMetric(float64(UserTitlesLen), "UTitlesLen")
	b.ReportMetric(float64(UserTitleDomainsLen), "DomainsLen")
	b.ReportMetric(float64(UserTitleCampusesLen), "CampusesLen")
	b.ReportMetric(float64(UserTitleSchoolsLen), "SchoolsLen")
	b.ReportMetric(float64(UserTitleClassesLen), "ClassesLen")
	b.ReportMetric(float64(UserTitleBranchesLen), "BranchesLen")
	b.ReportMetric(float64(UserTitleLessonsLen), "LessonsLen")
}

func BenchmarkFillUsers(b *testing.B) {
	b.ReportAllocs()
	FillStore(db)
	FillUsers()
}

func BenchmarkComputeUserNetwork(b *testing.B) {
	FillStore(db)
	FillUsers()
	ComputeUserNetworksIDs()
}
