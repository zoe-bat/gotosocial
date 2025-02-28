/*
   GoToSocial
   Copyright (C) 2021-2022 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package config

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitViper(f *pflag.FlagSet) error {
	// environment variable stuff
	// flag 'some-flag-name' becomes env var 'GTS_SOME_FLAG_NAME'
	viper.SetEnvPrefix("gts")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	// flag stuff
	// bind all of the flags in flagset to viper so that we can retrieve their values from the viper store
	if err := viper.BindPFlags(f); err != nil {
		return err
	}

	return nil
}
