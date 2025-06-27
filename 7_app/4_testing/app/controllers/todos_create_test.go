package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"learning/app/controllers"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("POST /todos", func() {
	var (
		testServer *controllers.Server
		recorder   *httptest.ResponseRecorder
		req        *http.Request
	)

	BeforeEach(func() {
		testServer = controllers.New(
			controllers.SetConfig(testConfig),
			controllers.SetStore(testStore),
		)
		recorder = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodPost, "/todos", nil)
	})

	It("creates a new todo", func() {
		requestBody := map[string]interface{}{
			"title":       "Hello Test",
			"description": "This is a test todo",
			"completed":   false,
		}
		data, err := json.Marshal(requestBody)
		Expect(err).NotTo(HaveOccurred())
		req.Body = io.NopCloser(bytes.NewBuffer(data))

		testServer.Router.ServeHTTP(recorder, req)
		Expect(recorder.Code).To(Equal(http.StatusCreated))

		var body controllers.Todo
		err = json.Unmarshal(recorder.Body.Bytes(), &body)
		Expect(err).NotTo(HaveOccurred())

		todo, err := testStore.GetTodo(context.Background(), int32(body.ID))
		Expect(err).NotTo(HaveOccurred())
		Expect(todo.Title).To(Equal(requestBody["title"]))
		Expect(todo.Description.String).To(Equal(requestBody["description"]))
		Expect(todo.Completed.Bool).To(Equal(requestBody["completed"]))
	})
})
