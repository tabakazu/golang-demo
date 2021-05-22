# Example:
#   $ make status
status:
	goose -path . status

# Example:
#   $ make migrate
migrate:
	goose -path . up

# Example:
#   $ make create NAME=CreateUsersTable
create:
	goose -path . create $(NAME) sql
