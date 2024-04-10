/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function (knex) {
  return knex.schema.createTable('participants', function (table) {
    table.increments('id').primary();
    table.integer('user_id').unsigned().notNullable();
    table.integer('event_id').unsigned().notNullable();
    table.foreign('user_id').references('users.id');
    table.foreign('event_id').references('events.id');
    table.string('role').notNullable();
    table.timestamps(true, true);
    table.timestamp('deleted_at');
  });
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function (knex) {
  return knex.schema.dropTableIfExists('participants');
};
