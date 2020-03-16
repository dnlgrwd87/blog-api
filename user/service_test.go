package user_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dnlgrwd87/blog-api/models"
	"github.com/dnlgrwd87/blog-api/user"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var expectedUser = models.User{
	ID:        1,
	FirstName: "Daniel",
	LastName:  "Garwood",
	Email:     "dnlgrwd@gmail.com",
	Posts: []models.Post{
		{ID: 1, Title: "My first post"},
	},
}

var expectedUserList = []models.UserListDTO{
	{
		ID:        1,
		FirstName: "Daniel",
		LastName:  "Garwood",
		Email:     "dnlgrwd@gmail.com",
	},
}

var _ = Describe("Service", func() {
	Context("#GetAllUsers", func() {
		It("will return a list of users", func() {
			db, mock, _ := sqlmock.New()
			mockDb, _ := gorm.Open("postgres", db)
			defer mockDb.Close()

			mock.ExpectQuery(`SELECT (.+) FROM "users"`).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"}).
						AddRow(1, "Daniel", "Garwood", "dnlgrwd@gmail.com"))

			service := user.UserService{DB: mockDb}
			users, err := service.GetAllUsers()

			if err != nil {
				Fail("Error should not have occurred")
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				Fail("Expectations were not met")
			}

			Expect(users[0]).To(Equal(expectedUserList[0]))
		})
	})

	Context("#GetUserById", func() {
		It("will return a single user with post", func() {
			db, mock, _ := sqlmock.New()
			mockDb, _ := gorm.Open("postgres", db)
			defer mockDb.Close()

			mock.ExpectQuery(`SELECT (.+) FROM "users"`).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"}).
						AddRow(1, "Daniel", "Garwood", "dnlgrwd@gmail.com"))

			mock.ExpectQuery(`SELECT (.+) FROM "posts"`).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "title"}).
						AddRow(1, "My first post"))

			service := user.UserService{DB: mockDb}
			user, _ := service.GetUserById(expectedUser.ID)

			if err := mock.ExpectationsWereMet(); err != nil {
				Fail("Expectations were not met")
			}

			Expect(user).To(Equal(expectedUser))
		})
		It("will return an error when no user is found", func() {
			db, mock, _ := sqlmock.New()
			mockDb, _ := gorm.Open("postgres", db)
			defer mockDb.Close()

			mock.ExpectQuery(`SELECT (.+) FROM "users"`).
				WillReturnError(gorm.ErrRecordNotFound)

			service := user.UserService{DB: mockDb}
			user, err := service.GetUserById(1)

			if err := mock.ExpectationsWereMet(); err != nil {
				Fail("Expectations were not met")
			}

			Expect(err).To(HaveOccurred())
			Expect(user.ID).To(Equal(0))
		})
	})
})
