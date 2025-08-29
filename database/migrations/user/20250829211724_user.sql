-- Rename a column from "token" to "hashed_token"
ALTER TABLE "verifications" RENAME COLUMN "token" TO "hashed_token";
