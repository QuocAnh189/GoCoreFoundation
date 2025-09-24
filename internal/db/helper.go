package db

import (
	"database/sql"
	"fmt"
	"log"
	"path"
	"runtime"
	"strings"

	"github.com/QuocAnh189/GoCoreFoundation/internal/utils/colors"
)

func (*Database) logInputSQL(fgColor colors.Color, query string, args ...any) {
	var tag string
	pc, file, line, ok := runtime.Caller(2) // Get information about the caller 3 frames up
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		// Extract just the function name from the full path
		if lastDot := strings.LastIndex(funcName, "."); lastDot != -1 {
			funcName = funcName[lastDot+1:]
		}

		// Extract package/directory name
		dir := path.Dir(file)
		if lastSlash := strings.LastIndex(dir, "/"); lastSlash != -1 {
			dir = dir[lastSlash+1:]
		}

		// Extract filename
		fileName := path.Base(file)

		tag = fmt.Sprintf("# %s/%s %s:%d", dir, fileName, funcName, line) // Updated log format
	}
	log.Printf("# SQL START: %s\n", tag)

	msg := fmt.Sprintf("%s, args: %v", query, args)
	log.Printf("%s\n", colors.Colorize(fgColor, colors.BGBlack, msg))

	log.Println("# SQL END")
}

func (*Database) logQueryError(err error) {
	log.Printf("# !SQL ERROR: %s\n", colors.Colorize(colors.FGRed, colors.BGBlack, err.Error()))
}

func (*Database) logExecResult(result sql.Result) {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("# !SQL ERROR: %s\n", colors.Colorize(colors.FGRed, colors.BGBlack, err.Error()))
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Printf("# !SQL ERROR: %s\n", colors.Colorize(colors.FGRed, colors.BGBlack, err.Error()))
	}

	msg := fmt.Sprintf("rows affected: %d, last insert id: %d", rowsAffected, lastInsertId)
	log.Printf("# @SQL EXEC RESULT: %s\n", colors.Colorize(colors.FGBlue, colors.BGBlack, msg))
}

func (*Database) logQueryRowsResult(rows *sql.Rows) {
	cols, err := rows.Columns()
	if err != nil {
		log.Printf("failed to get columns: %v", err)
		return
	}
	log.Printf("# @SQL QUERY ROWS RESULT: %s\n", colors.Colorize(colors.FGBlue, colors.BGBlack, strings.Join(cols, ", ")))
}

func (*Database) logQueryRowResult(_ *sql.Row) {
	// nothing to do i guess
}
