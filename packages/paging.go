package ManuUtils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func GenerateTable(recordsG any, elemPerPage int) {

	currentPage := 1

	s := reflect.ValueOf(recordsG)
	if s.Kind() != reflect.Slice {
		panic("Passed a non slice element to GenerateTable()")
	}

	if s.IsNil() {
		panic("Passed a nil slice element to GenerateTable()")
	}

	records := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		records[i] = s.Index(i).Interface()
	}

	if len(records) == 0 {
		panic("Records array lenght equals to 0")
	}

	totalPages := len(records) / elemPerPage

	var fields []string = make([]string, 0)

	tmp := reflect.VisibleFields(reflect.TypeOf(records[0]))

	for _, v := range tmp {
		fields = append(fields, v.Name)
	}

	for {
		fmt.Print(buildGraphicHeader(fields))

		var tmpRecord []any

		if currentPage == totalPages {
			tmpRecord = records[currentPage*elemPerPage:]
		} else {
			tmpRecord = records[currentPage*elemPerPage : currentPage*elemPerPage+elemPerPage]
		}

		for _, v := range tmpRecord {
			fmt.Print(buildGraphicRecord(v))
		}

		println(strconv.FormatInt(int64(currentPage), 32) + " / " + (strconv.FormatInt(int64(totalPages), 32)))

		key := GetKeyboardKey()

		deleteTable(len(tmpRecord))

		if key == 13 {
			break
		}

		if key == 65515 {
			if currentPage != 0 {
				currentPage--
			}
		}

		if key == 65514 {
			if currentPage != totalPages {
				currentPage++
			}
		}

	}

}

func buildGraphicRecord(recordG any) string {

	record := reflect.ValueOf(recordG)

	t := reflect.TypeOf(recordG)

	visibleFields := reflect.VisibleFields(t)

	values := make([]interface{}, len(visibleFields))

	for i := 0; i < len(visibleFields); i++ {
		values[i] = record.FieldByIndex(visibleFields[i].Index)
	}

	topLine := ""
	contentLine := ""
	bottomLine := ""

	for i, v := range values {

		element := ""

		element = strings.TrimSpace(fmt.Sprint(v))

		odd := ""

		if len(element) > 16 {
			element = "elementTooLong"
		}
		if len(element)%2 == 1 {
			odd = " "
		}

		stringWithPadding := ""
		if element == "elementTooLong" {
			stringWithPadding = strings.Repeat(" ", (16-len(element))/2) + "\x1b[32m" + element + "\x1b[0m" + strings.Repeat(" ", (16-len(element))/2) + odd
		} else {
			stringWithPadding = strings.Repeat(" ", (16-len(element))/2) + element + strings.Repeat(" ", (16-len(element))/2) + odd
		}

		switch i {
		case 0:
			topLine += "|                |"
			contentLine += fmt.Sprintf("|%s|", stringWithPadding)
			bottomLine += "|________________|"
		default:
			topLine += "                |"
			contentLine += fmt.Sprintf("%s|", stringWithPadding)
			bottomLine += "________________|"
		}
	}

	return topLine + "\n" + contentLine + "\n" + bottomLine + "\n"

}

func buildGraphicHeader(fields []string) string {
	//genero l'header
	headerTopLine := ""
	headerContentLine := ""
	headerBottomLine := ""

	for i, v := range fields {

		odd := ""

		if len(strings.TrimSpace(v))%2 == 1 {
			odd = " "
		}

		stringWithPadding := strings.Repeat(" ", (16-len(strings.TrimSpace(v)))/2) + strings.TrimSpace(v) + strings.Repeat(" ", (16-len(strings.TrimSpace(v)))/2) + odd

		switch i {
		case 0:
			headerTopLine += "|‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾|"
			headerContentLine += fmt.Sprintf("|%s|", stringWithPadding)
			headerBottomLine += "|________________|"
		default:
			headerTopLine += "‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾|"
			headerContentLine += fmt.Sprintf("%s|", stringWithPadding)
			headerBottomLine += "________________|"
		}

	}

	return headerTopLine + "\n" + headerContentLine + "\n" + headerBottomLine + "\n"
}

func deleteTable(elemToDel int) {
	for i := 0; i < 3*(elemToDel+1)+1; i++ {
		ClearLine()
	}
}
