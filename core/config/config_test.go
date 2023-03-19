package config_test

import (
	"os"
	"testing"

	. "commerce-platform/core/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config Suite", func() {

	BeforeEach(func() {
		os.Setenv("GO_ENV", "")
	})

	Describe("Load function", func() {
		Context("With correct config file", func() {
			It("should run as expected", func() {
				os.Setenv("GO_ENV", "test")
				config, _ := Load()
				Expect(config).To(Equal(AppConfig{
					DBConnectionStr: "test-connection",
				}))
			})
		})
		Context("With wrong config format", func() {
			It("should run as expected", func() {
				os.Setenv("GO_ENV", "wrong")
				_, err := Load()
				Expect(err.Error()).To(Equal("unexpected end of JSON input"))
			})
		})
		Context("With not found config file", func() {
			It("should run as expected", func() {
				os.Setenv("GO_ENV", "not_found_env")
				_, err := Load()
				Expect(err.Error()).To(Equal("open config.not_found_env.json: no such file or directory"))
			})
		})
	})

	Describe("GetConfig function", func() {
		Context("Test single config", func() {
			It("should run as expected", func() {
				os.Setenv("GO_ENV", "test")
				config1 := GetConfig()
				config2 := GetConfig()
				Expect(&config1).To(Equal(&config2))
			})
		})
	})
})
