/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function (knex) {
  return knex.schema.createTable('activities', function (table) {
    table.increments('id').primary();
    table.string('name').notNullable();
    table.dateTime('start_time').notNullable();
    table.dateTime('end_time').notNullable();
    table.string('description');
    table.integer('event_id').unsigned().notNullable();
    table.foreign('event_id').references('events.id');
    table.timestamps(true, true);
    table.timestamp('deleted_at');
  });
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function (knex) {
  return knex.schema.dropTableIfExists('activities');
};
