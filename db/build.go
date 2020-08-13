package db

import sq "github.com/Masterminds/squirrel"

// Build is a query builder that uses the dollar format (for postgres).
var Build = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
