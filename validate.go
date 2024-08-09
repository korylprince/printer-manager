package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/korylprince/printer-manager/db"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ValidatePrinter(hostRegexp *regexp.Regexp) db.PrinterHook {
	return db.PrinterHook(func(_ context.Context, _ boil.ContextExecutor, p *db.Printer) error {
		if !hostRegexp.MatchString(p.Hostname) {
			return fmt.Errorf("hostname %q doesn't match regexp %q", p.Hostname, hostRegexp.String())
		}
		return nil
	})
}
