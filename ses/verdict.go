package ses

// sns
// Copyright (C) 2021 Maximilian Pachl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// ---------------------------------------------------------------------------------------
//  imports
// ---------------------------------------------------------------------------------------

import ()

// ---------------------------------------------------------------------------------------
//  types
// ---------------------------------------------------------------------------------------

// Verdict is a common type, which indicates the status of a test.
// https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-notifications-contents.html#receiving-email-notifications-contents-dkimverdict-object
type Verdict struct {
	Status string `json:"status"`
}

// ---------------------------------------------------------------------------------------
//  public members
// ---------------------------------------------------------------------------------------

// IsOk returns true when the verdict is not in FAIL state.
// GRAY, ord PROCESSING_FAILED counts as Ok as well.
func (v *Verdict) IsOk() bool {
	return v.Status != "FAIL"
}

// HasPassed return true when the verdict state is PASS.
func (v *Verdict) HasPassed() bool {
	return v.Status == "PASS"
}

// String returns a string representation fo this Verdict.
func (v *Verdict) String() string {
	return v.Status
}
