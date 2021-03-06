// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPrinters(t *testing.T) {
	t.Parallel()

	query := Printers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPrintersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Printers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPrintersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Printers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Printers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPrintersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PrinterSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Printers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPrintersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PrinterExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Printer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PrinterExists to return true, but got false.")
	}
}

func testPrintersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	printerFound, err := FindPrinter(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if printerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPrintersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Printers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPrintersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Printers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPrintersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	printerOne := &Printer{}
	printerTwo := &Printer{}
	if err = randomize.Struct(seed, printerOne, printerDBTypes, false, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}
	if err = randomize.Struct(seed, printerTwo, printerDBTypes, false, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = printerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = printerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Printers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPrintersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	printerOne := &Printer{}
	printerTwo := &Printer{}
	if err = randomize.Struct(seed, printerOne, printerDBTypes, false, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}
	if err = randomize.Struct(seed, printerTwo, printerDBTypes, false, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = printerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = printerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Printers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func printerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Printer) error {
	*o = Printer{}
	return nil
}

func printerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Printer) error {
	*o = Printer{}
	return nil
}

func printerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Printer) error {
	*o = Printer{}
	return nil
}

func printerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Printer) error {
	*o = Printer{}
	return nil
}

func printerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Printer) error {
	*o = Printer{}
	return nil
}

func printerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Printer) error {
	*o = Printer{}
	return nil
}

func printerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Printer) error {
	*o = Printer{}
	return nil
}

func printerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Printer) error {
	*o = Printer{}
	return nil
}

func printerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Printer) error {
	*o = Printer{}
	return nil
}

func testPrintersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Printer{}
	o := &Printer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, printerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Printer object: %s", err)
	}

	AddPrinterHook(boil.BeforeInsertHook, printerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	printerBeforeInsertHooks = []PrinterHook{}

	AddPrinterHook(boil.AfterInsertHook, printerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	printerAfterInsertHooks = []PrinterHook{}

	AddPrinterHook(boil.AfterSelectHook, printerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	printerAfterSelectHooks = []PrinterHook{}

	AddPrinterHook(boil.BeforeUpdateHook, printerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	printerBeforeUpdateHooks = []PrinterHook{}

	AddPrinterHook(boil.AfterUpdateHook, printerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	printerAfterUpdateHooks = []PrinterHook{}

	AddPrinterHook(boil.BeforeDeleteHook, printerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	printerBeforeDeleteHooks = []PrinterHook{}

	AddPrinterHook(boil.AfterDeleteHook, printerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	printerAfterDeleteHooks = []PrinterHook{}

	AddPrinterHook(boil.BeforeUpsertHook, printerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	printerBeforeUpsertHooks = []PrinterHook{}

	AddPrinterHook(boil.AfterUpsertHook, printerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	printerAfterUpsertHooks = []PrinterHook{}
}

func testPrintersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Printers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPrintersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(printerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Printers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPrinterToOneLocationUsingLocation(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Printer
	var foreign Location

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, printerDBTypes, false, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, locationDBTypes, false, locationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Location struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.LocationID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Location().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PrinterSlice{&local}
	if err = local.L.LoadLocation(ctx, tx, false, (*[]*Printer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Location == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Location = nil
	if err = local.L.LoadLocation(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Location == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPrinterToOneModelUsingModel(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Printer
	var foreign Model

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, printerDBTypes, false, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, modelDBTypes, false, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ModelID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Model().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PrinterSlice{&local}
	if err = local.L.LoadModel(ctx, tx, false, (*[]*Printer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Model == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Model = nil
	if err = local.L.LoadModel(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Model == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPrinterToOneSetOpLocationUsingLocation(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Printer
	var b, c Location

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, printerDBTypes, false, strmangle.SetComplement(printerPrimaryKeyColumns, printerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, locationDBTypes, false, strmangle.SetComplement(locationPrimaryKeyColumns, locationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, locationDBTypes, false, strmangle.SetComplement(locationPrimaryKeyColumns, locationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Location{&b, &c} {
		err = a.SetLocation(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Location != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Printers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.LocationID != x.ID {
			t.Error("foreign key was wrong value", a.LocationID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.LocationID))
		reflect.Indirect(reflect.ValueOf(&a.LocationID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.LocationID != x.ID {
			t.Error("foreign key was wrong value", a.LocationID, x.ID)
		}
	}
}
func testPrinterToOneSetOpModelUsingModel(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Printer
	var b, c Model

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, printerDBTypes, false, strmangle.SetComplement(printerPrimaryKeyColumns, printerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, modelDBTypes, false, strmangle.SetComplement(modelPrimaryKeyColumns, modelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, modelDBTypes, false, strmangle.SetComplement(modelPrimaryKeyColumns, modelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Model{&b, &c} {
		err = a.SetModel(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Model != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Printers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ModelID != x.ID {
			t.Error("foreign key was wrong value", a.ModelID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ModelID))
		reflect.Indirect(reflect.ValueOf(&a.ModelID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ModelID != x.ID {
			t.Error("foreign key was wrong value", a.ModelID, x.ID)
		}
	}
}

func testPrintersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPrintersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PrinterSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPrintersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Printers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	printerDBTypes = map[string]string{`ID`: `uuid`, `ModelID`: `uuid`, `LocationID`: `uuid`, `Hostname`: `character varying`, `DriverExtra`: `jsonb`}
	_              = bytes.MinRead
)

func testPrintersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(printerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(printerAllColumns) == len(printerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Printers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, printerDBTypes, true, printerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPrintersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(printerAllColumns) == len(printerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Printer{}
	if err = randomize.Struct(seed, o, printerDBTypes, true, printerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Printers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, printerDBTypes, true, printerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(printerAllColumns, printerPrimaryKeyColumns) {
		fields = printerAllColumns
	} else {
		fields = strmangle.SetComplement(
			printerAllColumns,
			printerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PrinterSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPrintersUpsert(t *testing.T) {
	t.Parallel()

	if len(printerAllColumns) == len(printerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Printer{}
	if err = randomize.Struct(seed, &o, printerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Printer: %s", err)
	}

	count, err := Printers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, printerDBTypes, false, printerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Printer struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Printer: %s", err)
	}

	count, err = Printers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
