// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	attachment "github.com/saagie/upjet-provider-bitwarden/internal/controller/bitwarden/attachment"
	folder "github.com/saagie/upjet-provider-bitwarden/internal/controller/bitwarden/folder"
	login "github.com/saagie/upjet-provider-bitwarden/internal/controller/item/login"
	securenote "github.com/saagie/upjet-provider-bitwarden/internal/controller/item/securenote"
	providerconfig "github.com/saagie/upjet-provider-bitwarden/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachment.Setup,
		folder.Setup,
		login.Setup,
		securenote.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
