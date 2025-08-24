-- Create "confirmations" table
CREATE TABLE "confirmations" (
  "id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "hmac_secret" character varying NOT NULL DEFAULT 'secret',
  "created_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "confirmation_created_at" to table: "confirmations"
CREATE INDEX "confirmation_created_at" ON "confirmations" ("created_at");
-- Create index "confirmations_user_id_key" to table: "confirmations"
CREATE UNIQUE INDEX "confirmations_user_id_key" ON "confirmations" ("user_id");
-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL,
  "email" character varying NOT NULL,
  "username" character varying NOT NULL,
  "first_name" character varying NOT NULL,
  "last_name" character varying NOT NULL,
  "role" character varying NOT NULL,
  "confirmed_at" timestamptz NULL,
  "idp_user_id" character varying NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "user_confirmed_at" to table: "users"
CREATE INDEX "user_confirmed_at" ON "users" ("confirmed_at");
-- Create index "user_created_at" to table: "users"
CREATE INDEX "user_created_at" ON "users" ("created_at");
-- Create index "user_deleted_at" to table: "users"
CREATE INDEX "user_deleted_at" ON "users" ("deleted_at");
-- Create index "user_role" to table: "users"
CREATE INDEX "user_role" ON "users" ("role");
-- Create index "user_updated_at" to table: "users"
CREATE INDEX "user_updated_at" ON "users" ("updated_at");
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create index "users_idp_user_id_key" to table: "users"
CREATE UNIQUE INDEX "users_idp_user_id_key" ON "users" ("idp_user_id");
-- Create index "users_username_key" to table: "users"
CREATE UNIQUE INDEX "users_username_key" ON "users" ("username");
