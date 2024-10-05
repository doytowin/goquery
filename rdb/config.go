/*
 * The Clear BSD License
 *
 * Copyright (c) 2024, DoytoWin, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package rdb

import "fmt"

var Config = struct {
	TableFormat     string
	JoinIdFormat    string
	JoinTableFormat string
}{
	"t_%s",
	"%s_id",
	"a_%s_and_%s",
}

func FormatTable(domain string) string {
	return fmt.Sprintf(Config.TableFormat, domain)
}

func FormatJoinId(domain string) string {
	return fmt.Sprintf(Config.JoinIdFormat, domain)
}

func FormatJoinTable(domain1 string, domain2 string) string {
	return fmt.Sprintf(Config.JoinTableFormat, domain1, domain2)
}