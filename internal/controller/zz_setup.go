// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	attachment "saagie.io/provider-bitwarden/internal/controller/bitwarden/attachment"
	folder "saagie.io/provider-bitwarden/internal/controller/bitwarden/folder"
	login "saagie.io/provider-bitwarden/internal/controller/item/login"
	securenote "saagie.io/provider-bitwarden/internal/controller/item/securenote"
	providerconfig "saagie.io/provider-bitwarden/internal/controller/providerconfig"
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
