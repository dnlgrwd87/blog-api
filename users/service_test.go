package users_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dnlgrwd87/blog-api/models"
	"github.com/dnlgrwd87/blog-api/users"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	Context("#GetAllUsers", func() {
		It("will return a list of users", func() {
			expectedUser := models.User{ID: 1, FirstName: "Daniel", LastName: "Garwood", Email: "dnlgrwd@gmail.com"}
			db, mock, _ := sqlmock.New()
			mockDb, _ := gorm.Open("postgres", db)
			defer mockDb.Close()

			mock.ExpectQuery(`SELECT (.+) FROM "users"`).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"}).
						AddRow(1, "Daniel", "Garwood", "dnlgrwd@gmail.com"))

			service := users.UserService{DB: mockDb}
			users := service.GetAllUsers()

			if err := mock.ExpectationsWereMet(); err != nil {
				Fail("Expectations were not met")
			}

			Expect(users[0].ID).To(Equal(expectedUser.ID))
			Expect(users[0].FirstName).To(Equal(expectedUser.FirstName))
			Expect(users[0].LastName).To(Equal(expectedUser.LastName))
			Expect(users[0].Email).To(Equal(expectedUser.Email))
		})
	})

	Context("#GetUserById", func() {
		It("will return a single user with posts", func() {
			expectedUser := models.User{ID: 1, FirstName: "Daniel", LastName: "Garwood", Email: "dnlgrwd@gmail.com"}
			expectedPost := models.Post{ID: 1, Title: "My first post"}
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

			service := users.UserService{DB: mockDb}
			user, _ := service.GetUserById("1")

			if err := mock.ExpectationsWereMet(); err != nil {
				Fail("Expectations were not met")
			}

			Expect(user.ID).To(Equal(expectedUser.ID))
			Expect(user.FirstName).To(Equal(expectedUser.FirstName))
			Expect(user.LastName).To(Equal(expectedUser.LastName))
			Expect(user.Email).To(Equal(expectedUser.Email))
			Expect(user.Posts[0].ID).To(Equal(expectedPost.ID))
			Expect(user.Posts[0].Title).To(Equal(expectedPost.Title))
		})
	})
})
