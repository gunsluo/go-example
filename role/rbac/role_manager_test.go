package rbac

import (
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestMemoryRoleManager(t *testing.T) {
	rm, err := NewRoleManager()
	if err != nil {
		t.Fatalf("Failed to new RoleManager: %v", err)
	}

	err = rm.Load()
	if err != nil {
		t.Fatalf("Failed to load: %v", err)
	}

	testcase(t, rm)
	testcase2(t, rm)
	testcase3(t, rm)
}

func TestDBRoleManager(t *testing.T) {
	db := createDB(t, func(mock sqlmock.Sqlmock) error {
		mock.ExpectQuery("^SELECT (.+) FROM role").
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}))
		mock.ExpectQuery("^SELECT (.+) FROM rule").
			WillReturnRows(sqlmock.NewRows([]string{"p_type", "v0", "v1", "v2", "v3", "v4", "v5"}))

		// test case1
		for i := 0; i < 4; i++ {
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO role").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			mock.ExpectCommit()
		}

		mock.ExpectBegin()
		mock.ExpectQuery("INSERT INTO role").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectQuery("INSERT INTO role").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		mock.ExpectQuery("^SELECT (.+) FROM role WHERE name=?").WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(0))
		//mock.ExpectQuery("^SELECT (.+) FROM role WHERE name=?").WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(1))

		mock.ExpectQuery("^SELECT (.+) FROM role").WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(1))
		mock.ExpectQuery("^SELECT (.+) FROM role (.+)").
			WithArgs(10, 0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
				AddRow(1, "KungFu-china:master", "test").
				AddRow(2, "KungFu-china:fan", "test").
				AddRow(3, "KungFu-japan:fst-class", "test").
				AddRow(4, "target-reach:admin", "test").
				AddRow(5, "sys:admin", "test").
				AddRow(6, "sys:anonymous", "test"))

		mock.ExpectQuery("^SELECT (.+) FROM role").WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(1))
		mock.ExpectQuery("^SELECT (.+) FROM role (.+)").
			WithArgs("%KungFu-china%", 10, 0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
				AddRow(1, "KungFu-china:master", "test").
				AddRow(2, "KungFu-china:fan", "test"))

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO rule").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO rule").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		for i := 0; i < 3; i++ {
			mock.ExpectExec("INSERT INTO rule").WillReturnResult(sqlmock.NewResult(1, 1))
		}

		for i := 0; i < 1; i++ {
			mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
		}

		for i := 0; i < 4; i++ {
			mock.ExpectBegin()
			mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
		}
		for i := 0; i < 2; i++ {
			mock.ExpectBegin()
			mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
		}

		//test case2
		for i := 0; i < 3; i++ {
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO role").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			mock.ExpectCommit()
		}

		for i := 0; i < 3; i++ {
			mock.ExpectExec("INSERT INTO rule").WillReturnResult(sqlmock.NewResult(1, 1))
		}
		mock.ExpectQuery("^SELECT (.+) FROM role").WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(0))

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		//test case3
		for i := 0; i < 6; i++ {
			mock.ExpectBegin()
			mock.ExpectQuery("INSERT INTO role").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			mock.ExpectCommit()
		}

		for i := 0; i < 9; i++ {
			mock.ExpectExec("INSERT INTO rule").WillReturnResult(sqlmock.NewResult(1, 1))
		}
		mock.ExpectQuery("^SELECT (.+) FROM role").WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(0))

		for i := 0; i < 6; i++ {
			mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
		}

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		return nil
	})

	rm, err := NewRoleManager(db)
	if err != nil {
		t.Fatalf("Failed to new RoleManager: %v", err)
	}

	err = rm.Load()
	if err != nil {
		t.Fatalf("Failed to load role: %v", err)
	}

	testcase(t, rm)
	testcase2(t, rm)
	testcase3(t, rm)

	rm.Close() // db.Close()
}

func testcase(t *testing.T, rm *RoleManager) {
	if _, err := rm.AddRole(&Role{Name: "sys:admin", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRole(&Role{Name: "sys:anonymous", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRole(&Role{Name: "KungFu-china:master", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRole(&Role{Name: "KungFu-china:fan", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRoles([]*Role{
		&Role{Name: "KungFu-japan:fst-class", Description: "test"},
		&Role{Name: "target-reach:admin", Description: "test"},
	}); err != nil {
		t.Fatalf("Failed to AddRoles %v", err)
	}

	if exist, err := rm.HasRole("KungFu-china:master"); err != nil || !exist {
		t.Fatalf("Failed to HasRole %v", err)
	}

	if exist, err := rm.HasRole("KungFu-china:abc"); err != nil || exist {
		t.Fatalf("Failed to HasRole %v", err)
	}

	roles, total, err := rm.GetRoles(10, 0)
	if err != nil {
		t.Fatalf("Failed to GetRoles %v", err)
	}
	if len(roles) == 0 || total == 0 {
		t.Fatalf("GetRoles failed")
	}
	if !containRoles(roles, "KungFu-china:master", "KungFu-china:fan",
		"KungFu-japan:fst-class", "target-reach:admin", "sys:admin", "sys:anonymous") {
		t.Fatalf("GetRoles failed")
	}

	roles, total, err = rm.GetRoles(10, 0, "KungFu-china")
	if err != nil {
		t.Fatalf("Failed to GetRoles %v", err)
	}
	if len(roles) == 0 || total == 0 {
		t.Fatalf("GetRoles failed")
	}
	if !containRoles(roles, "KungFu-china:master", "KungFu-china:fan") {
		t.Fatalf("GetRoles failed")
	}

	if err := rm.AddRoleForUsers("KungFu-china:master", []string{"Bruce Lee", "YIP Man"}); err != nil {
		t.Fatalf("Failed to AddRoleForUser %v", err)
	}
	if err := rm.AddRoleForUser("KungFu-china:fan", "Peter"); err != nil {
		t.Fatalf("Failed to AddRoleForUser %v", err)
	}
	if err := rm.AddRoleForUser("KungFu-japan:fst-class", "Tom"); err != nil {
		t.Fatalf("Failed to AddRoleForUser %v", err)
	}
	if err := rm.AddRoleForUser("target-reach:admin", "Zac"); err != nil {
		t.Fatalf("Failed to AddRoleForUser %v", err)
	}
	if err := rm.AddRoleForUser("target-reach:abc", "Zac"); err == nil {
		t.Fatalf("Failed to AddRoleForUser %v", err)
	}

	roles, err = rm.GetRolesForUser("Bruce Lee")
	if err != nil {
		t.Fatalf("Failed to GetRolesForUser %v", err)
	}
	if len(roles) != 1 || roles[0].Name != "KungFu-china:master" {
		t.Fatal("Failed to GetRolesForUser")
	}

	roles, err = rm.GetRolesForUser("Bruce Lee", "KungFu-china:master")
	if err != nil {
		t.Fatalf("Failed to GetRolesForUser %v", err)
	}
	if len(roles) != 1 || roles[0].Name != "KungFu-china:master" {
		t.Fatal("Failed to GetRolesForUser")
	}

	users, err := rm.GetUsersForRole("KungFu-china:master")
	if err != nil {
		t.Fatalf("Failed to GetUsersForRole %v", err)
	}
	if len(users) != 2 {
		t.Fatalf("GetUsersForRole failed")
	}
	if !isAllInArray(users, "Bruce Lee", "YIP Man") {
		t.Fatalf("GetUsersForRole failed")
	}

	has, err := rm.HasRoleForUser("KungFu-china:master", "Bruce Lee")
	if err != nil {
		t.Fatalf("Failed to HasRoleForUser %v", err)
	}
	if !has {
		t.Fatalf("HasRoleForUser failed")
	}

	has, err = rm.HasRoleForUser("KungFu-china:fan", "Bruce Lee")
	if err != nil {
		t.Fatalf("Failed to HasRoleForUser %v", err)
	}
	if has {
		t.Fatalf("HasRoleForUser failed")
	}

	if err := rm.RemoveRoleForUser("KungFu-china:master", "Bruce Lee"); err != nil {
		t.Fatalf("Failed to RemoveRoleForUser %v", err)
	}
	if err := rm.RemoveRoleForUser("KungFu-china:master", "Bruce Lee"); err != nil {
		t.Fatalf("Failed to RemoveRoleForUser %v", err)
	}

	if err := rm.RemoveRole("KungFu-china:master"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("KungFu-china:fan"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("KungFu-japan:fst-class"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("target-reach:admin"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("sys:admin"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}

	if err := rm.RemoveRole("sys:anonymous"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
}

func testcase2(t *testing.T, rm *RoleManager) {
	// Current role relation tree:
	// admin has the permission to member1 and member
	//          admin
	//       	/   \
	//      member1 member2

	// Current role inheritance tree:
	//      member1 member2
	//       	\   /
	//       	admin
	if _, err := rm.AddRole(&Role{Name: "demo:admin", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRole(&Role{Name: "demo:member1", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRole(&Role{Name: "demo:member2", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}

	if err := rm.AssignRole("demo:member1", "demo:admin"); err != nil {
		t.Fatalf("Failed to AssignRole %v", err)
	}
	if err := rm.AssignRole("demo:member2", "demo:admin"); err != nil {
		t.Fatalf("Failed to AssignRole %v", err)
	}

	// Current subject relation tree:
	//      member1 member2
	//       	\   /
	//       	admin
	//       	/
	//        tom
	if err := rm.AddRoleForUser("demo:admin", "tom"); err != nil {
		t.Fatalf("Failed to AddRoleForUser %v", err)
	}

	users, err := rm.GetUsersForRole("demo:admin")
	if err != nil {
		t.Fatalf("Failed to GetUsersForRole %v", err)
	}
	if len(users) != 1 {
		t.Fatalf("GetUsersForRole failed")
	}
	if !isAllInArray(users, "tom") {
		t.Fatalf("GetUsersForRole failed")
	}

	us, err := rm.GetUsersForRoleHierarchy("demo:admin")
	if err != nil {
		t.Fatalf("Failed to GetUsersForRoleHierarchy %v", err)
	}
	if len(us) != 3 {
		t.Fatalf("GetUsersForRoleHierarchy failed")
	}
	for _, u := range us {
		var matched bool
		if u.Name == "tom" && u.Type == SubjectUserType {
			matched = true
		}
		if u.Name == "demo:member1" && u.Type == RoleUserType {
			matched = true
		}
		if u.Name == "demo:member2" && u.Type == RoleUserType {
			matched = true
		}
		if !matched {
			t.Fatalf("GetUsersForRoleHierarchy failed")
		}
	}

	roles, err := rm.GetRolesForUser("tom")
	if err != nil {
		t.Fatalf("Failed to GetRolesForUser %v", err)
	}
	if len(roles) == 0 {
		t.Fatalf("GetRolesForUser failed")
	}
	if !containRoles(roles, "demo:admin") {
		t.Fatalf("GetRolesForUser failed")
	}

	roles, err = rm.GetRolesForUser("demo:member1")
	if err != nil {
		t.Fatalf("Failed to GetRolesForUser %v", err)
	}
	if len(roles) != 0 {
		t.Fatalf("GetRolesForUser failed")
	}

	roles, err = rm.GetRolesForUser("demo:admin")
	if err != nil {
		t.Fatalf("Failed to GetRolesForUser %v", err)
	}
	if len(roles) != 2 {
		t.Fatalf("GetRolesForUser failed")
	}
	if !containRoles(roles, "demo:member1", "demo:member2") {
		t.Fatalf("GetRolesForUser failed")
	}

	roles, err = rm.GetRolesHierarchyForUser("tom")
	if err != nil {
		t.Fatalf("Failed to GetRolesHierarchyForUser %v", err)
	}
	if len(roles) == 0 {
		t.Fatalf("GetRolesHierarchyForUser failed")
	}
	if !containRoles(roles, "demo:member1", "demo:member2", "demo:admin") {
		t.Fatalf("GetRolesHierarchyForUser failed")
	}

	roles, err = rm.GetRolesHierarchyForUser("demo:member1")
	if err != nil {
		t.Fatalf("Failed to GetRolesHierarchyForUser %v", err)
	}
	if len(roles) != 0 {
		t.Fatalf("GetRolesHierarchyForUser failed")
	}

	roles, err = rm.GetRolesHierarchyForUser("demo:admin")
	if err != nil {
		t.Fatalf("Failed to GetRolesHierarchyForUser %v", err)
	}
	if len(roles) != 2 {
		t.Fatalf("GetRolesHierarchyForUser failed")
	}
	if !containRoles(roles, "demo:member1", "demo:member2") {
		t.Fatalf("GetRolesHierarchyForUser failed")
	}

	if err := rm.RemoveRole("demo:admin"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("demo:member1"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("demo:member2"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
}

func testcase3(t *testing.T, rm *RoleManager) {
	// Current role relation tree:
	// The upper layer has the permission to lower it.
	// chief has the permission to fm and hrm
	//              chief
	//           /    \
	//     	   fm    hrm
	//           \   /
	//      	director
	//       	 /
	//       leader
	//        /
	//     member

	// Current role inheritance tree:
	//     member
	//         \
	//       leader
	//           \
	//      	director
	//           /    \
	//     	   fm    hrm
	//           \   /
	//           chief

	if _, err := rm.AddRole(&Role{Name: "demo:chief", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRole(&Role{Name: "demo:finical-manager", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRole(&Role{Name: "demo:hr-manager", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRole(&Role{Name: "demo:director", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRole(&Role{Name: "demo:leader", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}
	if _, err := rm.AddRole(&Role{Name: "demo:member", Description: "test"}); err != nil {
		t.Fatalf("Failed to AddRole %v", err)
	}

	if err := rm.AssignRole("demo:finical-manager", "demo:chief"); err != nil {
		t.Fatalf("Failed to AssignRole %v", err)
	}
	if err := rm.AssignRole("demo:hr-manager", "demo:chief"); err != nil {
		t.Fatalf("Failed to AssignRole %v", err)
	}
	if err := rm.AssignRole("demo:director", "demo:hr-manager"); err != nil {
		t.Fatalf("Failed to AssignRole %v", err)
	}
	if err := rm.AssignRole("demo:director", "demo:finical-manager"); err != nil {
		t.Fatalf("Failed to AssignRole %v", err)
	}
	if err := rm.AssignRole("demo:leader", "demo:director"); err != nil {
		t.Fatalf("Failed to AssignRole %v", err)
	}
	if err := rm.AssignRole("demo:member", "demo:leader"); err != nil {
		t.Fatalf("Failed to AssignRole %v", err)
	}

	// Current subject relation tree:
	//       member
	//          \
	//        leader
	//         /   \
	// jerry(sub) director
	//           /    \   \
	//     	   fm    hrm peter(sub)
	//           \   /
	//           chief
	//           /
	//		   tom(sub)
	if err := rm.AddRoleForUser("demo:chief", "tom"); err != nil {
		t.Fatalf("Failed to AddRoleForUser %v", err)
	}
	if err := rm.AddRoleForUser("demo:leader", "jerry"); err != nil {
		t.Fatalf("Failed to AddRoleForUser %v", err)
	}
	if err := rm.AddRoleForUser("demo:director", "peter"); err != nil {
		t.Fatalf("Failed to AddRoleForUser %v", err)
	}

	roles, err := rm.GetRolesForUser("jerry")
	if err != nil {
		t.Fatalf("Failed to GetRolesForUser %v", err)
	}
	if len(roles) == 0 {
		t.Fatalf("GetRolesForUser failed")
	}
	if !containRoles(roles, "demo:leader") {
		t.Fatalf("GetRolesForUser failed")
	}

	us, err := rm.GetUsersForRoleHierarchy("jerry")
	if err != nil {
		t.Fatalf("Failed to GetUsersForRoleHierarchy %v", err)
	}
	if len(us) == 0 {
		t.Fatalf("GetUsersForRoleHierarchy failed")
	}
	for _, u := range us {
		var matched bool
		if u.Name == "demo:leader" && u.Type == RoleUserType {
			matched = true
		}
		if !matched {
			t.Fatalf("GetUsersForRoleHierarchy failed")
		}
	}

	roles, err = rm.GetRolesForUser("tom")
	if err != nil {
		t.Fatalf("Failed to GetRolesForUser %v", err)
	}
	if len(roles) == 0 {
		t.Fatalf("GetRolesForUser failed")
	}
	if !containRoles(roles, "demo:chief") {
		t.Fatalf("GetRolesForUser failed")
	}

	roles, err = rm.GetRolesHierarchyForUser("tom")
	if err != nil {
		t.Fatalf("Failed to GetRolesForUser %v", err)
	}
	if len(roles) == 0 {
		t.Fatalf("GetRolesHierarchyForUsersForUser failed")
	}
	if !containRoles(roles, "demo:leader", "demo:member",
		"demo:director", "demo:finical-manager",
		"demo:hr-manager", "demo:chief") {
		t.Fatalf("GetRolesHierarchyForUser failed")
	}

	us, err = rm.GetUsersForRoleHierarchy("tom")
	if err != nil {
		t.Fatalf("Failed to GetUsersForRoleHierarchy %v", err)
	}
	if len(us) == 0 {
		t.Fatalf("GetUsersForRoleHierarchy failed")
	}
	for _, u := range us {
		var matched bool
		if u.Name == "demo:chief" && u.Type == RoleUserType {
			matched = true
		}
		if !matched {
			t.Fatalf("GetUsersForRoleHierarchy failed")
		}
	}

	roles, err = rm.GetRolesForUser("peter")
	if err != nil {
		t.Fatalf("Failed to GetRolesForUser %v", err)
	}
	if len(roles) == 0 {
		t.Fatalf("GetRolesForUser failed")
	}
	if !containRoles(roles, "demo:director") {
		t.Fatalf("GetRolesForUser failed")
	}

	roles, err = rm.GetRolesHierarchyForUser("peter")
	if err != nil {
		t.Fatalf("Failed to GetRolesHierarchyForUser %v", err)
	}
	if len(roles) == 0 {
		t.Fatalf("GetRolesHierarchyForUser failed")
	}
	if !containRoles(roles, "demo:leader", "demo:member",
		"demo:director") {
		t.Fatalf("GetRolesHierarchyForUser failed")
	}

	us, err = rm.GetUsersForRoleHierarchy("peter")
	if err != nil {
		t.Fatalf("Failed to GetUsersForRoleHierarchy %v", err)
	}
	if len(us) == 0 {
		t.Fatalf("GetUsersForRoleHierarchy failed")
	}
	for _, u := range us {
		var matched bool
		if u.Name == "demo:director" && u.Type == RoleUserType {
			matched = true
		}
		if !matched {
			t.Fatalf("GetUsersForRoleHierarchy failed")
		}
	}

	us, err = rm.GetUsersForRoleHierarchy("demo:hr-manager")
	if err != nil {
		t.Fatalf("Failed to GetUsersForRoleHierarchy %v", err)
	}
	if len(us) == 0 {
		t.Fatalf("GetUsersForRoleHierarchy failed")
	}
	for _, u := range us {
		var matched bool
		if u.Name == "demo:director" && u.Type == RoleUserType {
			matched = true
		}
		if !matched {
			t.Fatalf("GetUsersForRoleHierarchy failed")
		}
	}

	us, err = rm.GetUsersForRoleHierarchy("demo:chief")
	if err != nil {
		t.Fatalf("Failed to GetUsersForRoleHierarchy %v", err)
	}
	if len(us) != 3 {
		t.Fatalf("GetUsersForRoleHierarchy failed")
	}
	for _, u := range us {
		var matched bool
		if u.Name == "tom" && u.Type == SubjectUserType {
			matched = true
		}
		if u.Name == "demo:hr-manager" && u.Type == RoleUserType {
			matched = true
		}
		if u.Name == "demo:finical-manager" && u.Type == RoleUserType {
			matched = true
		}
		if !matched {
			t.Fatalf("GetUsersForRoleHierarchy failed")
		}
	}

	if err := rm.CancelAssignRole("demo:hr-manager", "demo:chief"); err != nil {
		t.Fatalf("Failed to CancelAssignRole %v", err)
	}
	if err := rm.CancelAssignRole("demo:finical-manager", "demo:chief"); err != nil {
		t.Fatalf("Failed to CancelAssignRole %v", err)
	}
	if err := rm.CancelAssignRole("demo:director", "demo:hr-manager"); err != nil {
		t.Fatalf("Failed to CancelAssignRole %v", err)
	}
	if err := rm.CancelAssignRole("demo:director", "demo:finical-manager"); err != nil {
		t.Fatalf("Failed to CancelAssignRole %v", err)
	}
	if err := rm.CancelAssignRole("demo:leader", "demo:director"); err != nil {
		t.Fatalf("Failed to CancelAssignRole %v", err)
	}
	if err := rm.CancelAssignRole("demo:member", "demo:leader"); err != nil {
		t.Fatalf("Failed to CancelAssignRole %v", err)
	}

	if err := rm.RemoveRole("demo:chief"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("demo:finical-manager"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("demo:hr-manager"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("demo:director"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("demo:leader"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
	if err := rm.RemoveRole("demo:member"); err != nil {
		t.Fatalf("Failed to RemoveRole %v", err)
	}
}
