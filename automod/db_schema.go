package automod

const DBSchema = `
CREATE TABLE IF NOT EXISTS automod_rulesets (
	id BIGSERIAL PRIMARY KEY,
	guild_id BIGINT NOT NULL,

	name TEXT NOT NULL,
	enabled BOOLEAN NOT NULL
);

CREATE INDEX IF NOT EXISTS automod_rulesets_guild_idx ON automod_rulesets(guild_id);


CREATE TABLE IF NOT EXISTS automod_rules (
	id BIGSERIAL PRIMARY KEY,
	guild_id BIGINT NOT NULL,
	ruleset_id BIGINT references automod_rulesets(id) ON DELETE CASCADE NOT NULL,
	name TEXT NOT NULL,
	trigger_counter BIGINT NOT NULL
);

CREATE INDEX IF NOT EXISTS automod_rules_guild_idx ON automod_rules(guild_id);

CREATE TABLE IF NOT EXISTS automod_rule_data (
	id BIGSERIAL PRIMARY KEY,
	guild_id BIGINT NOT NULL,
	rule_id BIGINT references automod_rules(id) ON DELETE CASCADE NOT NULL,

	kind int NOT NULL,
	type_id INT NOT NULL,
	settings JSONB NOT NULL
);

CREATE INDEX IF NOT EXISTS automod_rule_data_guild_idx ON automod_rule_data(guild_id);

CREATE TABLE IF NOT EXISTS automod_ruleset_conditions (
	id BIGSERIAL PRIMARY KEY,
	guild_id BIGINT NOT NULL,
	ruleset_id BIGINT references automod_rulesets(id) ON DELETE CASCADE NOT NULL,

	kind int NOT NULL,
	type_id INT NOT NULL,
	settings JSONB NOT NULL
);

CREATE INDEX IF NOT EXISTS automod_ruleset_conditions_guild_idx ON automod_ruleset_conditions(guild_id);

CREATE TABLE IF NOT EXISTS automod_violations (
	id BIGSERIAL PRIMARY KEY,
	guild_id BIGINT NOT NULL,
	user_id BIGINT NOT NULL,
	rule_id BIGINT references automod_rules(id) ON DELETE SET NULL,

	created_at TIMESTAMP WITH TIME ZONE NOT NULL,

	name TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS automod_violations_guild_idx ON automod_violations(guild_id);
CREATE INDEX IF NOT EXISTS automod_violations_user_idx ON automod_violations(user_id);
`