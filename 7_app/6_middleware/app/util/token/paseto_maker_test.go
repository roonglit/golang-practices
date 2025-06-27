package token_test

import (
	"learning/app/util/token"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Paseto Maker", func() {
	var maker token.Maker
	var err error
	var payload *token.Payload
	var t string

	BeforeEach(func() {
		maker, err = token.NewPasetoMaker("12345678901234567890123456789012")
		Expect(err).ToNot(HaveOccurred())
	})

	Context("token not created", func() {
		It("create a valid token and return payload", func() {
			token, payload, err := maker.CreateToken("john_doe", time.Minute)
			Expect(err).ToNot(HaveOccurred())
			Expect(token).ToNot(BeEmpty())
			Expect(payload).ToNot(BeNil())

			payload, err = maker.VerifyToken(token)
			Expect(err).ToNot(HaveOccurred())
			Expect(payload.ID).ToNot(BeNil())
			Expect(payload.Username).To(Equal("john_doe"))
			Expect(payload.IssueDate).To(BeTemporally("~", time.Now(), time.Second))
			Expect(payload.ExpireDate).To(BeTemporally("~", time.Now().Add(time.Minute), time.Second))
		})
	})

	Context("token created", func() {
		Context("token expired", func() {
			BeforeEach(func() {
				t, payload, err = maker.CreateToken("john_doe", -time.Minute)
				Expect(err).ToNot(HaveOccurred())
				Expect(t).ToNot(BeEmpty())
				Expect(payload).ToNot(BeNil())
			})

			It("return error", func() {
				payload, err = maker.VerifyToken(t)
				Expect(err).To(HaveOccurred())
				Expect(payload).To(BeNil())
			})
		})

		Context("token not expired", func() {
			BeforeEach(func() {
				t, payload, err = maker.CreateToken("john_doe", time.Minute)
				Expect(err).ToNot(HaveOccurred())
				Expect(t).ToNot(BeEmpty())
				Expect(payload).ToNot(BeNil())
			})

			It("return payload", func() {
				payload, err = maker.VerifyToken(t)
				Expect(err).ToNot(HaveOccurred())
				Expect(payload.ID).ToNot(BeEmpty())
				Expect(payload.Username).To(Equal("john_doe"))
				Expect(payload.IssueDate).To(BeTemporally("~", time.Now(), time.Second))
				Expect(payload.ExpireDate).To(BeTemporally("~", time.Now().Add(time.Minute), time.Second))
			})
		})
	})
})
