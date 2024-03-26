// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package signature

import (
	"github.com/berachain/beacon-kit/config"
	"github.com/berachain/beacon-kit/primitives"
)

// ComputeDomain as defined in the Ethereum 2.0 specification.
// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#compute_domain
//
//nolint:lll
func ComputeDomain(
	domainType primitives.DomainType,
	forkVersion primitives.Version,
	genesisValidatorsRoot primitives.Root,
) (primitives.Domain, error) {
	forkDataRoot, err := ComputeForkDataRoot(forkVersion, genesisValidatorsRoot)
	if err != nil {
		return primitives.Domain{}, err
	}
	return primitives.Domain(
		append(
			domainType[:],
			forkDataRoot[:(primitives.RootLength-DomainTypeLength)]...),
	), nil
}

// GetDomain returns the domain for the DomainType and epoch.
func GetDomain(
	cfg *config.Config,
	genesisValidatorsRoot primitives.Root,
	domainType primitives.DomainType,
	epoch primitives.Epoch,
) (primitives.Domain, error) {
	return ComputeDomain(
		domainType,
		VersionFromUint32(cfg.Beacon.ActiveForkVersionByEpoch(epoch)),
		genesisValidatorsRoot,
	)
}
