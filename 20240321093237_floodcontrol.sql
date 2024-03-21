-- +goose Up
-- +goose StatementBegin
CREATE TABLE floodcontrol (
    user_id SERIAL PRIMARY KEY,
    first_call TIMESTAMP,
    call_count INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE floodcontrol;
-- +goose StatementEnd
