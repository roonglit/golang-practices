package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"learning/app/controllers"
	"learning/spec/factories"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GET /todos", func() {
	var (
		testServer *controllers.Server
		recorder   *httptest.ResponseRecorder
		req        *http.Request
		factory    *factories.Factory
	)

	BeforeEach(func() {
		testServer = controllers.New(
			controllers.SetConfig(testConfig),
			controllers.SetStore(testStore),
		)
		recorder = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodGet, "/todos", nil)
		factory = factories.New(testStore)
	})

	Context("Todos exist", func() {
		var (
			err error
		)
		BeforeEach(func() {
			_, err = factory.CreateTodos(context.Background(), 5, func(params map[string]interface{}, index int) {
				params["title"] = fmt.Sprintf("Todo %d", index+1)
				params["description"] = fmt.Sprintf("Description for Todo %d", index+1)
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns a list of todos", func() {
			testServer.Router.ServeHTTP(recorder, req)
			Expect(recorder.Code).To(Equal(http.StatusOK))

			var todos []controllers.Todo
			err := json.Unmarshal(recorder.Body.Bytes(), &todos)
			Expect(err).NotTo(HaveOccurred())
			Expect(todos).To(HaveLen(5))

			for i, todo := range todos {
				Expect(todo.ID).To(Equal(todos[i].ID))
				Expect(todo.Title).To(Equal(fmt.Sprintf("Todo %d", i+1)))
				Expect(todo.Description).To(Equal(fmt.Sprintf("Description for Todo %d", i+1)))
				Expect(*todo.Completed).To(BeFalse())
			}
		})
	})
})
