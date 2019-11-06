/*
 * Copyright (C) 2018 The dna Authors
 * This file is part of The dna library.
 *
 * The dna is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The dna is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The dna.  If not, see <http://www.gnu.org/licenses/>.
 */
package methods

import (
	"github.com/dnaproject2/dna-tool/methods/smartcontract"
)

func init() {
	smartcontract.RegisterSmartContract()
}
