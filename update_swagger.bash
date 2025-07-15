BACKEND_LOCATION=./backend/swagger
FRONTEND_LOCATION=./frontend/src/api/generated

swag init -d ./backend/src -o $BACKEND_LOCATION


npx @openapitools/openapi-generator-cli generate \
  -i $BACKEND_LOCATION/swagger.json \
  -g typescript-axios \
  -o $FRONTEND_LOCATION \
