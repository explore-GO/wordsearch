package db

import (
    "database/sql"
    "time"

    "github.com/LukasJoswiak/wordsearch/models"
)

func (db *Database) GetPuzzle(url string) (*models.Puzzle, error) {
    puzzle := &models.Puzzle{}

    sqlStr := "SELECT id, view_url, data FROM puzzles WHERE url = ?"

    row := db.db.QueryRow(sqlStr, url)
    err := row.Scan(&puzzle.ID, &puzzle.ViewURL, &puzzle.Data)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        } else {
            return nil, DBError{sqlStr, err}
        }
    }

    return puzzle, nil
}

func (db *Database) GetPuzzleByViewUrl(url string) (*models.Puzzle, error) {
    puzzle := &models.Puzzle{}

    sqlStr := "SELECT id, data FROM puzzles WHERE view_url = ?"

    row := db.db.QueryRow(sqlStr, url)
    err := row.Scan(&puzzle.ID, &puzzle.Data)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        } else {
            return nil, DBError{sqlStr, err}
        }
    }

    return puzzle, nil
}

func (db *Database) CreatePuzzle(puzzle *models.Puzzle) error {
    _, err := db.db.Exec(`INSERT INTO puzzles (url, view_url, data, type, datetime) VALUES (?, ?, ?, ?, ?)`, puzzle.URL, puzzle.ViewURL, puzzle.Data, puzzle.Type, time.Now())
    return err
}

func (db *Database) UpdatePuzzle(puzzle *models.Puzzle) error {
    sql := `UPDATE puzzles
            SET data = ?
            WHERE id = ?`

    stmt, err := db.db.Prepare(sql)
    if err != nil {
        return DBError{sql, err}
    }

    _, err = stmt.Exec(puzzle.Data, puzzle.ID)
    if err != nil {
        return DBError{sql, err}
    }

    return nil
}
