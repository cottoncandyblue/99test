package entity

import(
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`
	Url  string `gorm:"uniqueIndex" valid:"url"`
}

func TestValidation(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("check data valid",func(t *testing.T) {

		u := Video{
			Name: "eeeee",
			Url: "http://www.youtube.com/",
		}

		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})

}

func TestName(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("Check Name be blank", func(t *testing.T) {

		u := Video{
			Name: "",
			Url: "http://www.youtube.com/",
		}

		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))
	})

}

func TestUrl(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("check Url is valid", func(t *testing.T) {

		u:= Video{
			Name: "fdvofk",
			Url: "://www.youtubesss.com/",
		}

		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Url: ://www.youtubesss.com/ does not validate as url"))
	})

}