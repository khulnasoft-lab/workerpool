// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Khulnasoft

package workerpool

// export for testing only

// TasksCap returns the tasks channel capacity.
func (wp *WorkerPool) TasksCap() int {
	return cap(wp.tasks)
}
