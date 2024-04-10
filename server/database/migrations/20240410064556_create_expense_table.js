/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function (knex) {
  return knex.schema.createTable('expenses', function (table) {
    table.increments('id').primary();
    table.string('item_name').notNullable();
    table.integer('cost').notNullable();
    table.string('description');
    table.string('type').notNullable();
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
  return knex.schema.dropTableIfExists('expenses');
};
