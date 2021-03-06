package settingsutil_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/solo-io/gloo/pkg/utils/settingsutil"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

var _ = Describe("Settings", func() {

	It("should store settings in ctx", func() {
		settings := &v1.Settings{}
		ctx := context.Background()

		ctx = WithSettings(ctx, settings)

		Expect(FromContext(ctx)).To(Equal(settings))
	})

	It("should panic when no settings", func() {
		ctx := context.Background()

		Expect(func() { FromContext(ctx) }).Should(Panic())
	})

	It("should return true for IsAllNamespacesFromSettings when WatchNamespaces is empty", func() {
		settings := &v1.Settings{
			WatchNamespaces: []string{},
		}

		Expect(IsAllNamespacesFromSettings(settings)).To(BeTrue())
	})

	It("should return true for IsAllNamespacesFromSettings when WatchNamespaces only has empty namespace", func() {
		settings := &v1.Settings{
			WatchNamespaces: []string{""},
		}

		Expect(IsAllNamespacesFromSettings(settings)).To(BeTrue())
	})

	It("should return false for IsAllNamespacesFromSettings when WatchNamespaces has a namespace", func() {
		settings := &v1.Settings{
			WatchNamespaces: []string{"test"},
		}

		Expect(IsAllNamespacesFromSettings(settings)).To(BeFalse())
	})

})
