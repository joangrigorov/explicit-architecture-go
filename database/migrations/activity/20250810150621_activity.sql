-- Create "activities" table
CREATE TABLE "activities" (
  "id" uuid NOT NULL,
  "slug" character varying NOT NULL,
  "title" character varying NOT NULL,
  "poster_image_url" character varying NOT NULL,
  "short_description" character varying NOT NULL,
  "full_description" text NOT NULL,
  "happens_at" timestamptz NOT NULL,
  "attendants" bigint NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "activities_slug_key" to table: "activities"
CREATE UNIQUE INDEX "activities_slug_key" ON "activities" ("slug");
-- Create index "activity_attendants" to table: "activities"
CREATE INDEX "activity_attendants" ON "activities" ("attendants");
-- Create index "activity_created_at" to table: "activities"
CREATE INDEX "activity_created_at" ON "activities" ("created_at");
-- Create index "activity_deleted_at" to table: "activities"
CREATE INDEX "activity_deleted_at" ON "activities" ("deleted_at");
-- Create index "activity_happens_at" to table: "activities"
CREATE INDEX "activity_happens_at" ON "activities" ("happens_at");
-- Create index "activity_updated_at" to table: "activities"
CREATE INDEX "activity_updated_at" ON "activities" ("updated_at");
