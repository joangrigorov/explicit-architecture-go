-- Create "attendances" table
CREATE TABLE "attendances" (
  "id" uuid NOT NULL,
  "attendee_id" uuid NOT NULL,
  "activity_id" uuid NOT NULL,
  "activity_slug" character varying NOT NULL,
  "activity_title" character varying NOT NULL,
  "activity_poster_image_url" character varying NOT NULL,
  "activity_short_description" character varying NOT NULL,
  "activity_happens_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "attendance_activity_happens_at" to table: "attendances"
CREATE INDEX "attendance_activity_happens_at" ON "attendances" ("activity_happens_at");
-- Create index "attendance_activity_slug" to table: "attendances"
CREATE INDEX "attendance_activity_slug" ON "attendances" ("activity_slug");
-- Create index "attendance_attendee_id_activity_id" to table: "attendances"
CREATE UNIQUE INDEX "attendance_attendee_id_activity_id" ON "attendances" ("attendee_id", "activity_id");
-- Create index "attendance_created_at" to table: "attendances"
CREATE INDEX "attendance_created_at" ON "attendances" ("created_at");
-- Create index "attendance_deleted_at" to table: "attendances"
CREATE INDEX "attendance_deleted_at" ON "attendances" ("deleted_at");
-- Create index "attendance_updated_at" to table: "attendances"
CREATE INDEX "attendance_updated_at" ON "attendances" ("updated_at");
