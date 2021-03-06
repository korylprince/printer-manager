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

func testModels(t *testing.T) {
	t.Parallel()

	query := Models()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testModelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
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

	count, err := Models().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testModelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Models().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Models().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testModelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ModelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Models().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testModelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ModelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Model exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ModelExists to return true, but got false.")
	}
}

func testModelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	modelFound, err := FindModel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if modelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testModelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Models().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testModelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Models().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testModelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	modelOne := &Model{}
	modelTwo := &Model{}
	if err = randomize.Struct(seed, modelOne, modelDBTypes, false, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}
	if err = randomize.Struct(seed, modelTwo, modelDBTypes, false, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = modelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = modelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Models().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testModelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	modelOne := &Model{}
	modelTwo := &Model{}
	if err = randomize.Struct(seed, modelOne, modelDBTypes, false, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}
	if err = randomize.Struct(seed, modelTwo, modelDBTypes, false, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = modelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = modelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Models().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func modelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Model) error {
	*o = Model{}
	return nil
}

func modelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Model) error {
	*o = Model{}
	return nil
}

func modelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Model) error {
	*o = Model{}
	return nil
}

func modelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Model) error {
	*o = Model{}
	return nil
}

func modelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Model) error {
	*o = Model{}
	return nil
}

func modelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Model) error {
	*o = Model{}
	return nil
}

func modelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Model) error {
	*o = Model{}
	return nil
}

func modelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Model) error {
	*o = Model{}
	return nil
}

func modelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Model) error {
	*o = Model{}
	return nil
}

func testModelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Model{}
	o := &Model{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, modelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Model object: %s", err)
	}

	AddModelHook(boil.BeforeInsertHook, modelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	modelBeforeInsertHooks = []ModelHook{}

	AddModelHook(boil.AfterInsertHook, modelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	modelAfterInsertHooks = []ModelHook{}

	AddModelHook(boil.AfterSelectHook, modelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	modelAfterSelectHooks = []ModelHook{}

	AddModelHook(boil.BeforeUpdateHook, modelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	modelBeforeUpdateHooks = []ModelHook{}

	AddModelHook(boil.AfterUpdateHook, modelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	modelAfterUpdateHooks = []ModelHook{}

	AddModelHook(boil.BeforeDeleteHook, modelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	modelBeforeDeleteHooks = []ModelHook{}

	AddModelHook(boil.AfterDeleteHook, modelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	modelAfterDeleteHooks = []ModelHook{}

	AddModelHook(boil.BeforeUpsertHook, modelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	modelBeforeUpsertHooks = []ModelHook{}

	AddModelHook(boil.AfterUpsertHook, modelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	modelAfterUpsertHooks = []ModelHook{}
}

func testModelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Models().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testModelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(modelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Models().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testModelToManyPrinters(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Model
	var b, c Printer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, printerDBTypes, false, printerColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, printerDBTypes, false, printerColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ModelID = a.ID
	c.ModelID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Printers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ModelID == b.ModelID {
			bFound = true
		}
		if v.ModelID == c.ModelID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ModelSlice{&a}
	if err = a.L.LoadPrinters(ctx, tx, false, (*[]*Model)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Printers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Printers = nil
	if err = a.L.LoadPrinters(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Printers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testModelToManyAddOpPrinters(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Model
	var b, c, d, e Printer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, modelDBTypes, false, strmangle.SetComplement(modelPrimaryKeyColumns, modelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Printer{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, printerDBTypes, false, strmangle.SetComplement(printerPrimaryKeyColumns, printerColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Printer{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPrinters(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ModelID {
			t.Error("foreign key was wrong value", a.ID, first.ModelID)
		}
		if a.ID != second.ModelID {
			t.Error("foreign key was wrong value", a.ID, second.ModelID)
		}

		if first.R.Model != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Model != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Printers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Printers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Printers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testModelToOneManufacturerUsingManufacturer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Model
	var foreign Manufacturer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, modelDBTypes, false, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, manufacturerDBTypes, false, manufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Manufacturer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ManufacturerID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Manufacturer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ModelSlice{&local}
	if err = local.L.LoadManufacturer(ctx, tx, false, (*[]*Model)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Manufacturer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Manufacturer = nil
	if err = local.L.LoadManufacturer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Manufacturer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testModelToOneSetOpManufacturerUsingManufacturer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Model
	var b, c Manufacturer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, modelDBTypes, false, strmangle.SetComplement(modelPrimaryKeyColumns, modelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, manufacturerDBTypes, false, strmangle.SetComplement(manufacturerPrimaryKeyColumns, manufacturerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, manufacturerDBTypes, false, strmangle.SetComplement(manufacturerPrimaryKeyColumns, manufacturerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Manufacturer{&b, &c} {
		err = a.SetManufacturer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Manufacturer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Models[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ManufacturerID != x.ID {
			t.Error("foreign key was wrong value", a.ManufacturerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ManufacturerID))
		reflect.Indirect(reflect.ValueOf(&a.ManufacturerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ManufacturerID != x.ID {
			t.Error("foreign key was wrong value", a.ManufacturerID, x.ID)
		}
	}
}

func testModelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
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

func testModelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ModelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testModelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Models().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	modelDBTypes = map[string]string{`ID`: `uuid`, `ManufacturerID`: `uuid`, `Name`: `character varying`, `Driver`: `jsonb`}
	_            = bytes.MinRead
)

func testModelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(modelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(modelAllColumns) == len(modelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Models().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, modelDBTypes, true, modelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testModelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(modelAllColumns) == len(modelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Model{}
	if err = randomize.Struct(seed, o, modelDBTypes, true, modelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Models().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, modelDBTypes, true, modelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(modelAllColumns, modelPrimaryKeyColumns) {
		fields = modelAllColumns
	} else {
		fields = strmangle.SetComplement(
			modelAllColumns,
			modelPrimaryKeyColumns,
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

	slice := ModelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testModelsUpsert(t *testing.T) {
	t.Parallel()

	if len(modelAllColumns) == len(modelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Model{}
	if err = randomize.Struct(seed, &o, modelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Model: %s", err)
	}

	count, err := Models().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, modelDBTypes, false, modelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Model struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Model: %s", err)
	}

	count, err = Models().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
