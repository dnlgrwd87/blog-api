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

var expectedUser1 = models.User{
	ID:        1,
	FirstName: "Daniel",
	LastName:  "Garwood",
	Email:     "dnlgrwd@gmail.com",
}

var _ = Describe("Handler", func() {
	Context("UserHandler", func() {
		It("should return a list of users", func() {
			userHandler := user.UserHandler{
				Service: &user.MockUserService{},
			}

			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users", nil)
			handler := http.HandlerFunc(userHandler.GetUsers)
			handler.ServeHTTP(recorder, req)

			var users []models.User
			json.NewDecoder(recorder.Body).Decode(&users)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			Expect(users).To(Equal([]models.User{expectedUser1}))
		})
	})
})
