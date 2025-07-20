package postgres

import (
	"context"

	"1337b04rd/internal/domain"
)

func (db *poolDB) GetComments(ctx context.Context, postID uint64) ([]domain.Comment, error) {
	const query = `
		SELECT 
			c.comment_id,
			u.avatar_url,
			c.parent_comment_id,
			c.comment_content,
			c.has_image
		FROM comments c
		JOIN users u ON c.user_id = u.session_id
		WHERE c.post_id = $1
		ORDER BY c.comment_time ASC`

	rows, err := db.Pool.Query(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []domain.Comment
	for rows.Next() {
		var c domain.Comment
		err := rows.Scan(
			&c.ID,
			&c.AvatarURL,
			&c.ReplyToID,
			&c.Content,
			&c.HasImage,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
