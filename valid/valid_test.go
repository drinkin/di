package valid_test

import (
	"github.com/drinkin/di/random"
	"github.com/drinkin/di/valid"
	"github.com/drinkin/di/valid/gvalid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("V", func() {
	var (
		v     *valid.V
		field string
	)

	BeforeEach(func() {
		v = valid.New()
		field = random.String(10)
	})

	It("Int64 required", func() {
		v.F(field).Int64(1).Required()
		Expect(v.Check()).ToNot(HaveOccurred())

		v.F(field).Int64(0).Required()
		Expect(v.Check()).To(gvalid.MatchRequiredErr(field))
	})

	It("String required", func() {
		v.F(field).String("a").Required()
		Expect(v.Check()).ToNot(HaveOccurred())

		v.F(field).String("").Required()
		Expect(v.Check()).To(gvalid.MatchRequiredErr(field))
	})

	It("String enum", func() {
		values := []string{"a", "b"}
		v.F(field).String("a").Enum(values)
		Expect(v.Check()).ToNot(HaveOccurred())

		v.F(field).String("c").Enum(values)
		Expect(v.Check()).To(gvalid.MatchEnumErr(field))
	})
})
