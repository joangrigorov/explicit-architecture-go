-- Rename a column from "hashed_token" to "csrf_token"
ALTER TABLE "verifications" RENAME COLUMN "hashed_token" TO "csrf_token";
