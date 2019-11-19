// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Manufacturer is an object representing the database table.
type Manufacturer struct {
	ID   string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *manufacturerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L manufacturerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ManufacturerColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

// Generated where

var ManufacturerWhere = struct {
	ID   whereHelperstring
	Name whereHelperstring
}{
	ID:   whereHelperstring{field: "\"manufacturer\".\"id\""},
	Name: whereHelperstring{field: "\"manufacturer\".\"name\""},
}

// ManufacturerRels is where relationship names are stored.
var ManufacturerRels = struct {
	Models string
}{
	Models: "Models",
}

// manufacturerR is where relationships are stored.
type manufacturerR struct {
	Models ModelSlice
}

// NewStruct creates a new relationship struct
func (*manufacturerR) NewStruct() *manufacturerR {
	return &manufacturerR{}
}

// manufacturerL is where Load methods for each relationship are stored.
type manufacturerL struct{}

var (
	manufacturerAllColumns            = []string{"id", "name"}
	manufacturerColumnsWithoutDefault = []string{"name"}
	manufacturerColumnsWithDefault    = []string{"id"}
	manufacturerPrimaryKeyColumns     = []string{"id"}
)

type (
	// ManufacturerSlice is an alias for a slice of pointers to Manufacturer.
	// This should generally be used opposed to []Manufacturer.
	ManufacturerSlice []*Manufacturer
	// ManufacturerHook is the signature for custom Manufacturer hook methods
	ManufacturerHook func(context.Context, boil.ContextExecutor, *Manufacturer) error

	manufacturerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	manufacturerType                 = reflect.TypeOf(&Manufacturer{})
	manufacturerMapping              = queries.MakeStructMapping(manufacturerType)
	manufacturerPrimaryKeyMapping, _ = queries.BindMapping(manufacturerType, manufacturerMapping, manufacturerPrimaryKeyColumns)
	manufacturerInsertCacheMut       sync.RWMutex
	manufacturerInsertCache          = make(map[string]insertCache)
	manufacturerUpdateCacheMut       sync.RWMutex
	manufacturerUpdateCache          = make(map[string]updateCache)
	manufacturerUpsertCacheMut       sync.RWMutex
	manufacturerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var manufacturerBeforeInsertHooks []ManufacturerHook
var manufacturerBeforeUpdateHooks []ManufacturerHook
var manufacturerBeforeDeleteHooks []ManufacturerHook
var manufacturerBeforeUpsertHooks []ManufacturerHook

var manufacturerAfterInsertHooks []ManufacturerHook
var manufacturerAfterSelectHooks []ManufacturerHook
var manufacturerAfterUpdateHooks []ManufacturerHook
var manufacturerAfterDeleteHooks []ManufacturerHook
var manufacturerAfterUpsertHooks []ManufacturerHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Manufacturer) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range manufacturerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Manufacturer) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range manufacturerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Manufacturer) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range manufacturerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Manufacturer) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range manufacturerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Manufacturer) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range manufacturerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Manufacturer) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range manufacturerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Manufacturer) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range manufacturerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Manufacturer) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range manufacturerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Manufacturer) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range manufacturerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddManufacturerHook registers your hook function for all future operations.
func AddManufacturerHook(hookPoint boil.HookPoint, manufacturerHook ManufacturerHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		manufacturerBeforeInsertHooks = append(manufacturerBeforeInsertHooks, manufacturerHook)
	case boil.BeforeUpdateHook:
		manufacturerBeforeUpdateHooks = append(manufacturerBeforeUpdateHooks, manufacturerHook)
	case boil.BeforeDeleteHook:
		manufacturerBeforeDeleteHooks = append(manufacturerBeforeDeleteHooks, manufacturerHook)
	case boil.BeforeUpsertHook:
		manufacturerBeforeUpsertHooks = append(manufacturerBeforeUpsertHooks, manufacturerHook)
	case boil.AfterInsertHook:
		manufacturerAfterInsertHooks = append(manufacturerAfterInsertHooks, manufacturerHook)
	case boil.AfterSelectHook:
		manufacturerAfterSelectHooks = append(manufacturerAfterSelectHooks, manufacturerHook)
	case boil.AfterUpdateHook:
		manufacturerAfterUpdateHooks = append(manufacturerAfterUpdateHooks, manufacturerHook)
	case boil.AfterDeleteHook:
		manufacturerAfterDeleteHooks = append(manufacturerAfterDeleteHooks, manufacturerHook)
	case boil.AfterUpsertHook:
		manufacturerAfterUpsertHooks = append(manufacturerAfterUpsertHooks, manufacturerHook)
	}
}

// One returns a single manufacturer record from the query.
func (q manufacturerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Manufacturer, error) {
	o := &Manufacturer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for manufacturer")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Manufacturer records from the query.
func (q manufacturerQuery) All(ctx context.Context, exec boil.ContextExecutor) (ManufacturerSlice, error) {
	var o []*Manufacturer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to Manufacturer slice")
	}

	if len(manufacturerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Manufacturer records in the query.
func (q manufacturerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count manufacturer rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q manufacturerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if manufacturer exists")
	}

	return count > 0, nil
}

// Models retrieves all the model's Models with an executor.
func (o *Manufacturer) Models(mods ...qm.QueryMod) modelQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"model\".\"manufacturer_id\"=?", o.ID),
	)

	query := Models(queryMods...)
	queries.SetFrom(query.Query, "\"model\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"model\".*"})
	}

	return query
}

// LoadModels allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (manufacturerL) LoadModels(ctx context.Context, e boil.ContextExecutor, singular bool, maybeManufacturer interface{}, mods queries.Applicator) error {
	var slice []*Manufacturer
	var object *Manufacturer

	if singular {
		object = maybeManufacturer.(*Manufacturer)
	} else {
		slice = *maybeManufacturer.(*[]*Manufacturer)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &manufacturerR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &manufacturerR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`model`), qm.WhereIn(`model.manufacturer_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load model")
	}

	var resultSlice []*Model
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice model")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on model")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for model")
	}

	if len(modelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Models = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &modelR{}
			}
			foreign.R.Manufacturer = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ManufacturerID {
				local.R.Models = append(local.R.Models, foreign)
				if foreign.R == nil {
					foreign.R = &modelR{}
				}
				foreign.R.Manufacturer = local
				break
			}
		}
	}

	return nil
}

// AddModels adds the given related objects to the existing relationships
// of the manufacturer, optionally inserting them as new records.
// Appends related to o.R.Models.
// Sets related.R.Manufacturer appropriately.
func (o *Manufacturer) AddModels(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Model) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ManufacturerID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"model\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"manufacturer_id"}),
				strmangle.WhereClause("\"", "\"", 2, modelPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ManufacturerID = o.ID
		}
	}

	if o.R == nil {
		o.R = &manufacturerR{
			Models: related,
		}
	} else {
		o.R.Models = append(o.R.Models, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &modelR{
				Manufacturer: o,
			}
		} else {
			rel.R.Manufacturer = o
		}
	}
	return nil
}

// Manufacturers retrieves all the records using an executor.
func Manufacturers(mods ...qm.QueryMod) manufacturerQuery {
	mods = append(mods, qm.From("\"manufacturer\""))
	return manufacturerQuery{NewQuery(mods...)}
}

// FindManufacturer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindManufacturer(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Manufacturer, error) {
	manufacturerObj := &Manufacturer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"manufacturer\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, manufacturerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from manufacturer")
	}

	return manufacturerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Manufacturer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no manufacturer provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(manufacturerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	manufacturerInsertCacheMut.RLock()
	cache, cached := manufacturerInsertCache[key]
	manufacturerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			manufacturerAllColumns,
			manufacturerColumnsWithDefault,
			manufacturerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(manufacturerType, manufacturerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(manufacturerType, manufacturerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"manufacturer\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"manufacturer\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "db: unable to insert into manufacturer")
	}

	if !cached {
		manufacturerInsertCacheMut.Lock()
		manufacturerInsertCache[key] = cache
		manufacturerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Manufacturer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Manufacturer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	manufacturerUpdateCacheMut.RLock()
	cache, cached := manufacturerUpdateCache[key]
	manufacturerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			manufacturerAllColumns,
			manufacturerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db: unable to update manufacturer, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"manufacturer\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, manufacturerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(manufacturerType, manufacturerMapping, append(wl, manufacturerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update manufacturer row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by update for manufacturer")
	}

	if !cached {
		manufacturerUpdateCacheMut.Lock()
		manufacturerUpdateCache[key] = cache
		manufacturerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q manufacturerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all for manufacturer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected for manufacturer")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ManufacturerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("db: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), manufacturerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"manufacturer\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, manufacturerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all in manufacturer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected all in update all manufacturer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Manufacturer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no manufacturer provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(manufacturerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	manufacturerUpsertCacheMut.RLock()
	cache, cached := manufacturerUpsertCache[key]
	manufacturerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			manufacturerAllColumns,
			manufacturerColumnsWithDefault,
			manufacturerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			manufacturerAllColumns,
			manufacturerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("db: unable to upsert manufacturer, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(manufacturerPrimaryKeyColumns))
			copy(conflict, manufacturerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"manufacturer\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(manufacturerType, manufacturerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(manufacturerType, manufacturerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "db: unable to upsert manufacturer")
	}

	if !cached {
		manufacturerUpsertCacheMut.Lock()
		manufacturerUpsertCache[key] = cache
		manufacturerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Manufacturer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Manufacturer) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no Manufacturer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), manufacturerPrimaryKeyMapping)
	sql := "DELETE FROM \"manufacturer\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete from manufacturer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by delete for manufacturer")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q manufacturerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db: no manufacturerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from manufacturer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for manufacturer")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ManufacturerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(manufacturerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), manufacturerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"manufacturer\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, manufacturerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from manufacturer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for manufacturer")
	}

	if len(manufacturerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Manufacturer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindManufacturer(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ManufacturerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ManufacturerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), manufacturerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"manufacturer\".* FROM \"manufacturer\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, manufacturerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in ManufacturerSlice")
	}

	*o = slice

	return nil
}

// ManufacturerExists checks if the Manufacturer row exists.
func ManufacturerExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"manufacturer\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if manufacturer exists")
	}

	return exists, nil
}
