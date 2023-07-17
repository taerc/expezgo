// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"expezgo/pkg/ent/migrate"

	"expezgo/pkg/ent/city"
	"expezgo/pkg/ent/county"
	"expezgo/pkg/ent/licence"
	"expezgo/pkg/ent/province"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// City is the client for interacting with the City builders.
	City *CityClient
	// County is the client for interacting with the County builders.
	County *CountyClient
	// Licence is the client for interacting with the Licence builders.
	Licence *LicenceClient
	// Province is the client for interacting with the Province builders.
	Province *ProvinceClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.City = NewCityClient(c.config)
	c.County = NewCountyClient(c.config)
	c.Licence = NewLicenceClient(c.config)
	c.Province = NewProvinceClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		City:     NewCityClient(cfg),
		County:   NewCountyClient(cfg),
		Licence:  NewLicenceClient(cfg),
		Province: NewProvinceClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		City:     NewCityClient(cfg),
		County:   NewCountyClient(cfg),
		Licence:  NewLicenceClient(cfg),
		Province: NewProvinceClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		City.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.City.Use(hooks...)
	c.County.Use(hooks...)
	c.Licence.Use(hooks...)
	c.Province.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.City.Intercept(interceptors...)
	c.County.Intercept(interceptors...)
	c.Licence.Intercept(interceptors...)
	c.Province.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CityMutation:
		return c.City.mutate(ctx, m)
	case *CountyMutation:
		return c.County.mutate(ctx, m)
	case *LicenceMutation:
		return c.Licence.mutate(ctx, m)
	case *ProvinceMutation:
		return c.Province.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CityClient is a client for the City schema.
type CityClient struct {
	config
}

// NewCityClient returns a client for the City from the given config.
func NewCityClient(c config) *CityClient {
	return &CityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `city.Hooks(f(g(h())))`.
func (c *CityClient) Use(hooks ...Hook) {
	c.hooks.City = append(c.hooks.City, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `city.Intercept(f(g(h())))`.
func (c *CityClient) Intercept(interceptors ...Interceptor) {
	c.inters.City = append(c.inters.City, interceptors...)
}

// Create returns a builder for creating a City entity.
func (c *CityClient) Create() *CityCreate {
	mutation := newCityMutation(c.config, OpCreate)
	return &CityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of City entities.
func (c *CityClient) CreateBulk(builders ...*CityCreate) *CityCreateBulk {
	return &CityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for City.
func (c *CityClient) Update() *CityUpdate {
	mutation := newCityMutation(c.config, OpUpdate)
	return &CityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CityClient) UpdateOne(ci *City) *CityUpdateOne {
	mutation := newCityMutation(c.config, OpUpdateOne, withCity(ci))
	return &CityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CityClient) UpdateOneID(id uint32) *CityUpdateOne {
	mutation := newCityMutation(c.config, OpUpdateOne, withCityID(id))
	return &CityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for City.
func (c *CityClient) Delete() *CityDelete {
	mutation := newCityMutation(c.config, OpDelete)
	return &CityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CityClient) DeleteOne(ci *City) *CityDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CityClient) DeleteOneID(id uint32) *CityDeleteOne {
	builder := c.Delete().Where(city.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CityDeleteOne{builder}
}

// Query returns a query builder for City.
func (c *CityClient) Query() *CityQuery {
	return &CityQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCity},
		inters: c.Interceptors(),
	}
}

// Get returns a City entity by its id.
func (c *CityClient) Get(ctx context.Context, id uint32) (*City, error) {
	return c.Query().Where(city.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CityClient) GetX(ctx context.Context, id uint32) *City {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CityClient) Hooks() []Hook {
	return c.hooks.City
}

// Interceptors returns the client interceptors.
func (c *CityClient) Interceptors() []Interceptor {
	return c.inters.City
}

func (c *CityClient) mutate(ctx context.Context, m *CityMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CityCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CityUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CityDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown City mutation op: %q", m.Op())
	}
}

// CountyClient is a client for the County schema.
type CountyClient struct {
	config
}

// NewCountyClient returns a client for the County from the given config.
func NewCountyClient(c config) *CountyClient {
	return &CountyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `county.Hooks(f(g(h())))`.
func (c *CountyClient) Use(hooks ...Hook) {
	c.hooks.County = append(c.hooks.County, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `county.Intercept(f(g(h())))`.
func (c *CountyClient) Intercept(interceptors ...Interceptor) {
	c.inters.County = append(c.inters.County, interceptors...)
}

// Create returns a builder for creating a County entity.
func (c *CountyClient) Create() *CountyCreate {
	mutation := newCountyMutation(c.config, OpCreate)
	return &CountyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of County entities.
func (c *CountyClient) CreateBulk(builders ...*CountyCreate) *CountyCreateBulk {
	return &CountyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for County.
func (c *CountyClient) Update() *CountyUpdate {
	mutation := newCountyMutation(c.config, OpUpdate)
	return &CountyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CountyClient) UpdateOne(co *County) *CountyUpdateOne {
	mutation := newCountyMutation(c.config, OpUpdateOne, withCounty(co))
	return &CountyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CountyClient) UpdateOneID(id uint32) *CountyUpdateOne {
	mutation := newCountyMutation(c.config, OpUpdateOne, withCountyID(id))
	return &CountyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for County.
func (c *CountyClient) Delete() *CountyDelete {
	mutation := newCountyMutation(c.config, OpDelete)
	return &CountyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CountyClient) DeleteOne(co *County) *CountyDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CountyClient) DeleteOneID(id uint32) *CountyDeleteOne {
	builder := c.Delete().Where(county.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CountyDeleteOne{builder}
}

// Query returns a query builder for County.
func (c *CountyClient) Query() *CountyQuery {
	return &CountyQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCounty},
		inters: c.Interceptors(),
	}
}

// Get returns a County entity by its id.
func (c *CountyClient) Get(ctx context.Context, id uint32) (*County, error) {
	return c.Query().Where(county.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CountyClient) GetX(ctx context.Context, id uint32) *County {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CountyClient) Hooks() []Hook {
	return c.hooks.County
}

// Interceptors returns the client interceptors.
func (c *CountyClient) Interceptors() []Interceptor {
	return c.inters.County
}

func (c *CountyClient) mutate(ctx context.Context, m *CountyMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CountyCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CountyUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CountyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CountyDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown County mutation op: %q", m.Op())
	}
}

// LicenceClient is a client for the Licence schema.
type LicenceClient struct {
	config
}

// NewLicenceClient returns a client for the Licence from the given config.
func NewLicenceClient(c config) *LicenceClient {
	return &LicenceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `licence.Hooks(f(g(h())))`.
func (c *LicenceClient) Use(hooks ...Hook) {
	c.hooks.Licence = append(c.hooks.Licence, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `licence.Intercept(f(g(h())))`.
func (c *LicenceClient) Intercept(interceptors ...Interceptor) {
	c.inters.Licence = append(c.inters.Licence, interceptors...)
}

// Create returns a builder for creating a Licence entity.
func (c *LicenceClient) Create() *LicenceCreate {
	mutation := newLicenceMutation(c.config, OpCreate)
	return &LicenceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Licence entities.
func (c *LicenceClient) CreateBulk(builders ...*LicenceCreate) *LicenceCreateBulk {
	return &LicenceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Licence.
func (c *LicenceClient) Update() *LicenceUpdate {
	mutation := newLicenceMutation(c.config, OpUpdate)
	return &LicenceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LicenceClient) UpdateOne(l *Licence) *LicenceUpdateOne {
	mutation := newLicenceMutation(c.config, OpUpdateOne, withLicence(l))
	return &LicenceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LicenceClient) UpdateOneID(id int64) *LicenceUpdateOne {
	mutation := newLicenceMutation(c.config, OpUpdateOne, withLicenceID(id))
	return &LicenceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Licence.
func (c *LicenceClient) Delete() *LicenceDelete {
	mutation := newLicenceMutation(c.config, OpDelete)
	return &LicenceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LicenceClient) DeleteOne(l *Licence) *LicenceDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LicenceClient) DeleteOneID(id int64) *LicenceDeleteOne {
	builder := c.Delete().Where(licence.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LicenceDeleteOne{builder}
}

// Query returns a query builder for Licence.
func (c *LicenceClient) Query() *LicenceQuery {
	return &LicenceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLicence},
		inters: c.Interceptors(),
	}
}

// Get returns a Licence entity by its id.
func (c *LicenceClient) Get(ctx context.Context, id int64) (*Licence, error) {
	return c.Query().Where(licence.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LicenceClient) GetX(ctx context.Context, id int64) *Licence {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LicenceClient) Hooks() []Hook {
	return c.hooks.Licence
}

// Interceptors returns the client interceptors.
func (c *LicenceClient) Interceptors() []Interceptor {
	return c.inters.Licence
}

func (c *LicenceClient) mutate(ctx context.Context, m *LicenceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LicenceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LicenceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LicenceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LicenceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Licence mutation op: %q", m.Op())
	}
}

// ProvinceClient is a client for the Province schema.
type ProvinceClient struct {
	config
}

// NewProvinceClient returns a client for the Province from the given config.
func NewProvinceClient(c config) *ProvinceClient {
	return &ProvinceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `province.Hooks(f(g(h())))`.
func (c *ProvinceClient) Use(hooks ...Hook) {
	c.hooks.Province = append(c.hooks.Province, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `province.Intercept(f(g(h())))`.
func (c *ProvinceClient) Intercept(interceptors ...Interceptor) {
	c.inters.Province = append(c.inters.Province, interceptors...)
}

// Create returns a builder for creating a Province entity.
func (c *ProvinceClient) Create() *ProvinceCreate {
	mutation := newProvinceMutation(c.config, OpCreate)
	return &ProvinceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Province entities.
func (c *ProvinceClient) CreateBulk(builders ...*ProvinceCreate) *ProvinceCreateBulk {
	return &ProvinceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Province.
func (c *ProvinceClient) Update() *ProvinceUpdate {
	mutation := newProvinceMutation(c.config, OpUpdate)
	return &ProvinceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProvinceClient) UpdateOne(pr *Province) *ProvinceUpdateOne {
	mutation := newProvinceMutation(c.config, OpUpdateOne, withProvince(pr))
	return &ProvinceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProvinceClient) UpdateOneID(id uint32) *ProvinceUpdateOne {
	mutation := newProvinceMutation(c.config, OpUpdateOne, withProvinceID(id))
	return &ProvinceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Province.
func (c *ProvinceClient) Delete() *ProvinceDelete {
	mutation := newProvinceMutation(c.config, OpDelete)
	return &ProvinceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProvinceClient) DeleteOne(pr *Province) *ProvinceDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProvinceClient) DeleteOneID(id uint32) *ProvinceDeleteOne {
	builder := c.Delete().Where(province.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProvinceDeleteOne{builder}
}

// Query returns a query builder for Province.
func (c *ProvinceClient) Query() *ProvinceQuery {
	return &ProvinceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProvince},
		inters: c.Interceptors(),
	}
}

// Get returns a Province entity by its id.
func (c *ProvinceClient) Get(ctx context.Context, id uint32) (*Province, error) {
	return c.Query().Where(province.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProvinceClient) GetX(ctx context.Context, id uint32) *Province {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ProvinceClient) Hooks() []Hook {
	return c.hooks.Province
}

// Interceptors returns the client interceptors.
func (c *ProvinceClient) Interceptors() []Interceptor {
	return c.inters.Province
}

func (c *ProvinceClient) mutate(ctx context.Context, m *ProvinceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ProvinceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ProvinceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ProvinceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ProvinceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Province mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		City, County, Licence, Province []ent.Hook
	}
	inters struct {
		City, County, Licence, Province []ent.Interceptor
	}
)
