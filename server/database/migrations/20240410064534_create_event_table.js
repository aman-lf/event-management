/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function (knex) {
  return knex.schema.createTable('events', function (table) {
    table.increments('id').primary();
    table.string('name').notNullable();
    table.dateTime('start_date').notNullable();
    table.dateTime('end_date').notNullable();
    table.string('location').notNullable();
    table.string('type');
    table.string('description');
    table.timestamps(true, true);
    table.timestamp('deleted_at');
  });
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function (knex) {
  return knex.schema.dropTableIfExists('events');
};
