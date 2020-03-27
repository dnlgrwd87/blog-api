package user_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dnlgrwd87/blog-api/models"
	"github.com/dnlgrwd87/blog-api/user"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	var expectedUser = models.User{
		ID:        1,
		FirstName: "Daniel",
		LastName:  "Garwood",
		Email:     "dnlgrwd@gmail.com",
	}

	var expectedUserWithPosts = models.UserPostsDTO{
		ID:        1,
		FirstName: "Daniel",
		LastName:  "Garwood",
		Email:     "dnlgrwd@gmail.com",
		Posts: []models.Post{
			{ID: 1, Title: "My first post"},
		},
	}

	Context("#GetAllUsers", func() {
		It("should return a list of users", func() {
			db, mock, _ := sqlmock.New()
			mockDb, _ := gorm.Open("mysql", db)
			defer mockDb.Close()

			mock.ExpectQuery(`SELECT (.+) FROM .users`).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"}).
						AddRow(1, "Daniel", "Garwood", "dnlgrwd@gmail.com"))

			service := user.UserService{DB: mockDb}
			users, err := service.GetAllUsers()

			if err := mock.ExpectationsWereMet(); err != nil {
				Fail("Expectations were not met")
			}

			Expect(err).ToNot(HaveOccurred())
			Expect(users[0]).To(Equal(expectedUser))
		})
	})

	Context("#GetUserById", func() {
		It("should return a single user with post", func() {
			db, mock, _ := sqlmock.New()
			mockDb, _ := gorm.Open("mysql", db)
			defer mockDb.Close()

			mock.ExpectQuery(`SELECT (.+) FROM .users`).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"}).
						AddRow(1, "Daniel", "Garwood", "dnlgrwd@gmail.com"))

			mock.ExpectQuery(`SELECT (.+) FROM .posts`).
				WillReturnRows(
					sqlmock.NewRows([]string{"id", "title"}).
						AddRow(1, "My first post"))

			service := user.UserService{DB: mockDb}
			user, err := service.GetUserById(expectedUserWithPosts.ID)

			if err := mock.ExpectationsWereMet(); err != nil {
				Fail("Expectations were not met")
			}

			Expect(err).ToNot(HaveOccurred())
			Expect(user).To(Equal(expectedUserWithPosts))
		})
		It("should return an error when no user is found", func() {
			db, mock, _ := sqlmock.New()
			mockDb, _ := gorm.Open("mysql", db)
			defer mockDb.Close()

			mock.ExpectQuery(`SELECT (.+) FROM .users`).
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

	Context("#CreateUser", func() {
		It("should create a new user", func() {
			db, mock, _ := sqlmock.New()
			mockDb, _ := gorm.Open("mysql", db)
			defer mockDb.Close()

			mock.ExpectBegin()
			mock.ExpectExec(`INSERT INTO .users`).
				WithArgs("Fail", "Garwood", "dnlgrwd@gmail.com").
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			service := user.UserService{DB: mockDb}
			user, err := service.CreateUser(models.User{FirstName: "Daniel", LastName: "Garwood", Email: "dnlgrwd@gmail.com"})

			if err := mock.ExpectationsWereMet(); err != nil {
				Fail("Expectations were not met")
			}

			Expect(err).ToNot(HaveOccurred())
			Expect(user).To(Equal(expectedUser))
		})
	})
})
