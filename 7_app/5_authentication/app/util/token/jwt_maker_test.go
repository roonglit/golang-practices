package token_test

import (
	"learning/app/util/token"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JWT Maker", func() {

	Context("verify token", func() {
		var maker *token.JWTMaker
		var err error

		BeforeEach(func() {
			maker, err = token.NewJWTMaker("12345678901234567890123456789012")
			Expect(err).ToNot(HaveOccurred())
		})

		Context("token is expired", func() {
			var jwtToken *string
			var payload *token.Payload

			BeforeEach(func() {
				jwtToken, payload, err = maker.CreateToken("john_doe", time.Minute)
				Expect(err).ToNot(HaveOccurred())
				Expect(payload).ToNot(BeNil())
			})

			It("not allows user to perform", func() {
				payload, err = maker.VerifyToken(*jwtToken)
				Expect(err).ToNot(HaveOccurred())
				Expect(payload.Username).To(Equal("john_doe"))
			})
		})

		Context("token not expired", func() {
			var jwtToken *string
			var payload *token.Payload

			BeforeEach(func() {
				jwtToken, payload, err = maker.CreateToken("john_doe", -time.Minute)
				Expect(err).ToNot(HaveOccurred())
				Expect(payload).ToNot(BeNil())
			})

			It("allows user to perform", func() {
				payload, err = maker.VerifyToken(*jwtToken)
				Expect(err).To(Equal(token.ErrExpireToken))
				Expect(payload).To(BeNil())
			})
		})
	})

	Context("create token", func() {
		var maker *token.JWTMaker
		var err error

		BeforeEach(func() {
			maker, err = token.NewJWTMaker("12345678901234567890123456789012")
			Expect(err).ToNot(HaveOccurred())
		})

		It("creates a valid token and return payload", func() {
			jwtToken, payload, err := maker.CreateToken("john_doe", time.Minute)
			Expect(err).ToNot(HaveOccurred())
			Expect(jwtToken).ToNot(BeNil())
			Expect(payload).ToNot(BeNil())
			Expect(payload.Username).To(Equal("john_doe"))
		})
	})

	Context("set invalid secret key", func() {
		var err error

		It("should return error invalid key size", func() {
			_, err = token.NewJWTMaker("123456789")
			Expect(err.Error()).To(Equal("invalid key size: must be at least 32 characters"))
		})
	})
})
