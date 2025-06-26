package controllers_test

import (
	"bytes"
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
	})
})
