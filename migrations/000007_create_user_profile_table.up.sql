CREATE TABLE IF NOT EXISTS user_profile
(
    id                   bigserial PRIMARY KEY,
    created_at           timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    username             citext UNIQUE               NOT NULL,
    title                text                        NOT NULL,
    full_name            text[]                      NOT NULL,
    images               text[]                      NOT NULL,
    phone_number         text                        NOT NULL,
    email                text                        NOT NULL,
    answers              text[]                      NOT NULL,
    display_contact_info bool[]                      NOT NULL,
    user_id              bigint                      NOT NULL REFERENCES users ON DELETE CASCADE
);