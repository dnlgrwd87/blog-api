package user_test

import (
	"encoding/json"
	"github.com/dnlgrwd87/blog-api/models"
	"github.com/dnlgrwd87/blog-api/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Handler", func() {
	var expectedUser = models.User{
		ID:        1,
		FirstName: "Daniel",
		LastName:  "Garwood",
		Email:     "dnlgrwd@gmail.com",
	}

	Context("UserHandler", func() {
		It("should respond with a list of users", func() {
			userHandler := user.UserHandler{
				UserService: &user.MockUserService{},
			}

			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users", nil)
			handler := http.HandlerFunc(userHandler.GetUsers)
			handler.ServeHTTP(recorder, req)

			var users []models.User
			json.NewDecoder(recorder.Body).Decode(&users)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			Expect(users).To(Equal([]models.User{expectedUser}))
		})
		It("should respond with a single user with the passed in id", func() {
			userHandler := user.UserHandler{
				UserService: &user.MockUserService{},
			}

			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users/1", nil)
			handler := http.HandlerFunc(userHandler.GetUserById)
			handler.ServeHTTP(recorder, req)

			var user models.User
			json.NewDecoder(recorder.Body).Decode(&user)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			Expect(user).To(Equal(expectedUser))
		})
	})
})
