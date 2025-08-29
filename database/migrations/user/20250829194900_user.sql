-- Create "verifications" table
CREATE TABLE "verifications" (
  "id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "token" text NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "used_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "verification_created_at" to table: "verifications"
CREATE INDEX "verification_created_at" ON "verifications" ("created_at");
-- Create index "verification_expires_at" to table: "verifications"
CREATE INDEX "verification_expires_at" ON "verifications" ("expires_at");
-- Create index "verification_used_at" to table: "verifications"
CREATE INDEX "verification_used_at" ON "verifications" ("used_at");
-- Create index "verifications_user_id_key" to table: "verifications"
CREATE UNIQUE INDEX "verifications_user_id_key" ON "verifications" ("user_id");
-- Drop "confirmations" table
DROP TABLE "confirmations";
