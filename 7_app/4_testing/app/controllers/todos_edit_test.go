package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"learning/app/controllers"
	"learning/app/models"
	"learning/spec/factories"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PUT /todos/:id", func() {
	var (
		recorder   *httptest.ResponseRecorder
		req        *http.Request
		testServer *controllers.Server
		factory    *factories.Factory
	)

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		testServer = controllers.New(
			controllers.SetConfig(testConfig),
			controllers.SetStore(testStore),
		)

		req, _ = http.NewRequest(http.MethodPut, "/todos/1", nil)

		factory = factories.New(testStore)
	})

	Context("a todo exists", func() {
		var (
			todo models.Todo
			err  error
		)
		BeforeEach(func() {
			todo, err = factory.CreateTodo(context.Background(), map[string]interface{}{
				"title":       "Test Todo",
				"description": "This is a test todo",
				"completed":   false,
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("updates an existing todo", func() {
			req.URL.Path = fmt.Sprintf("/todos/%d", todo.ID)
			requestBody := map[string]interface{}{
				"title":       "Hello Test",
				"description": "This is a test todo",
				"completed":   true,
			}
			data, err := json.Marshal(requestBody)
			Expect(err).NotTo(HaveOccurred())
			req.Body = io.NopCloser(bytes.NewBuffer(data))

			testServer.Router.ServeHTTP(recorder, req)
			Expect(recorder.Code).To(Equal(200))

			var body controllers.Todo
			err = json.Unmarshal(recorder.Body.Bytes(), &body)
			Expect(err).NotTo(HaveOccurred())
			Expect(body.ID).To(Equal(int(todo.ID)))
			Expect(body.Title).To(Equal(requestBody["title"]))
			Expect(body.Description).To(Equal(requestBody["description"]))
			Expect(*body.Completed).To(Equal(requestBody["completed"]))
		})
	})
})
