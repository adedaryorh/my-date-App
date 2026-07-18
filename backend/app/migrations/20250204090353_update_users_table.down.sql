ALTER TABLE "users" DROP COLUMN business_name;
ALTER TABLE "users" DROP COLUMN industry_type;
ALTER TABLE "users" ALTER COLUMN date_of_birth SET NOT NULL;
