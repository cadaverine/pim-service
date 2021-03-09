-- +goose Up

-- +goose StatementBegin
create or replace function trigger_set_timestamp()
returns trigger as $$
begin
  new.updated_at = now();
  return new;
end;
$$ language plpgsql;
-- +goose StatementEnd

-- +goose Down
drop function if exists trigger_set_timestamp;