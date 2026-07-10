package repo

// // UserReactionPostgresRepo -.
// type UserReactionPostgresRepo struct {
// 	*postgres.Postgres
// }

// // NewUserReactionRepo -.
// func NewUserReactionRepo(pg *postgres.Postgres) *UserReactionPostgresRepo {
// 	return &UserReactionPostgresRepo{pg}
// }

// // Create -.
// func (r *UserReactionPostgresRepo) Create(ctx context.Context, ur *models.UserReaction) error {
// 	sql, args, err := r.Builder.
// 		Insert("user_reactions").
// 		Columns("user_id, post_id, reaction").
// 		Values(ur.UserID, ur.PostID, ur.Reaction).
// 		Suffix("RETURNING \"id\", \"created_at\"").
// 		ToSql()

// 	if err != nil {
// 		return fmt.Errorf("UserReactionPostgresRepo - Create - r.Builder: %w", err)
// 	}

// 	row := r.Pool.QueryRow(ctx, sql, args...)

// 	err = row.Scan(&ur.ID, &ur.CreatedAt)
// 	if err != nil {
// 		return fmt.Errorf("UserReactionPostgresRepo - Create - r.Pool.Scan: %w", err)
// 	}

// 	return nil
// }

// // Update -.
// func (r *UserReactionPostgresRepo) Update(ctx context.Context, ur *models.UserReaction) error {
// 	sql, args, err := r.Builder.
// 		Update("user_reactions").
// 		SetMap(squirrel.Eq{
// 			"reaction":   ur.Reaction,
// 			"updated_at": time.Now().UTC(),
// 		}).
// 		Where(squirrel.Eq{"user_id": ur.UserID, "post_id": ur.PostID}).
// 		ToSql()

// 	if err != nil {
// 		return fmt.Errorf("UserReactionPostgresRepo - Update - r.Builder: %w", err)
// 	}

// 	_, err = r.Pool.Exec(ctx, sql, args...)
// 	if err != nil {
// 		return fmt.Errorf("UserReactionPostgresRepo - Update - r.Pool.Exec: %w", err)
// 	}

// 	return nil
// }

// // GetPostReactions -.
// func (r *UserReactionPostgresRepo) GetPostReactions(ctx context.Context, postID int) ([]models.UserReaction, error) {
// 	builder := r.Builder.
// 		Select("id, user_id, post_id, reaction, created_at, updated_at").
// 		From("user_reactions").
// 		Where("post_id = ?", postID).
// 		OrderBy("created_at DESC")

// 	sql, args, err := builder.ToSql()
// 	if err != nil {
// 		return nil, fmt.Errorf("UserReactionPostgresRepo - GetPostReactions - r.Builder: %w", err)
// 	}

// 	rows, err := r.Pool.Query(ctx, sql, args...)
// 	if err != nil {
// 		return nil, fmt.Errorf("UserReactionPostgresRepo - GetPostReactions - r.Pool.Query: %w", err)
// 	}
// 	defer rows.Close()

// 	posts := make([]models.UserReaction, 0)
// 	for rows.Next() {
// 		p := models.UserReaction{
// 			Post: models.Post{},
// 			User: models.User{},
// 		}

// 		err = rows.Scan(
// 			&p.ID,
// 			&p.UserID,
// 			&p.PostID,
// 			&p.Reaction,
// 			&p.CreatedAt,
// 			&p.UpdatedAt,
// 		)

// 		if err != nil {
// 			return nil, fmt.Errorf("UserReactionPostgresRepo - GetPostReactions - rows.Scan: %w", err)
// 		}

// 		posts = append(posts, p)
// 	}

// 	return posts, nil
// }

// func (r *UserReactionPostgresRepo) GetReaction(ctx context.Context, userID int, postID int) (*models.UserReaction, error) {
// 	builder := r.Builder.
// 		Select("id, user_id, post_id, reaction, created_at, updated_at").
// 		From("user_reactions").
// 		Where("post_id = ? AND user_id = ?", postID, userID)

// 	sql, args, err := builder.ToSql()
// 	if err != nil {
// 		return nil, fmt.Errorf("UserReactionPostgresRepo - GetReaction - r.Builder: %w", err)
// 	}

// 	row := r.Pool.QueryRow(ctx, sql, args...)
// 	if err != nil {
// 		return nil, fmt.Errorf("UserReactionPostgresRepo - GetReaction - r.Pool.Query: %w", err)
// 	}

// 	reaction := models.UserReaction{
// 		Post: models.Post{},
// 		User: models.User{},
// 	}

// 	err = row.Scan(
// 		&reaction.ID,
// 		&reaction.UserID,
// 		&reaction.PostID,
// 		&reaction.Reaction,
// 		&reaction.CreatedAt,
// 		&reaction.UpdatedAt,
// 	)

// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			return nil, nil
// 		}

// 		return nil, fmt.Errorf("UserReactionPostgresRepo - GetReaction - rows.Scan: %w", err)
// 	}

// 	return &reaction, nil
// }

// // Delete -.
// func (r *UserReactionPostgresRepo) Delete(ctx context.Context, reactionID int) error {
// 	sql, args, err := r.Builder.
// 		Delete("user_reactions").
// 		Where("id = ?", reactionID).
// 		ToSql()

// 	if err != nil {
// 		return fmt.Errorf("UserReactionPostgresRepo - Delete - r.Builder: %w", err)
// 	}

// 	_, err = r.Pool.Exec(ctx, sql, args...)
// 	if err != nil {
// 		return fmt.Errorf("UserReactionPostgresRepo - Delete - r.Pool.Exec: %w", err)
// 	}

// 	return nil

// }
