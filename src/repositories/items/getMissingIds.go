package itemsrepository

import (
	"github.com/lib/pq"
)

func (i sItemsRepository) GetMissingIDs(ids []int) ([]int, error) {
	query := `
    WITH ids AS (
        SELECT unnest($1::int[]) AS id
    )
    SELECT ids.id
    FROM ids
    LEFT JOIN items i ON ids.id = i.id
    WHERE i.id IS NULL;
    `

	rows, err := i.DB.Query(query, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missingIDs []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		missingIDs = append(missingIDs, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return missingIDs, nil
}
