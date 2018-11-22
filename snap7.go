package snap7

import (
	"unsafe"
)

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lsnap7
#include "snap7.h"
*/
import "C"

//var S7AreaPE = int(C.S7AreaPE)

const (
	S7AreaPE = 0x81
	S7AreaPA = 0x82
	S7AreaMK = 0x83
	S7AreaDB = 0x84
	S7AreaCT = 0x1C
	S7AreaTM = 0x1D
)

type Snap7Client struct {
	inner C.S7Object
}

func ConnentTo(address string, rack int, slot int) (Snap7Client, error) {
	var client C.S7Object = C.Cli_Create()
	var addr *C.char = C.CString(address)

	var err C.int = C.Cli_ConnectTo(client, addr, C.int(rack), C.int(slot))
	if err != 0 {
		C.Cli_Destroy(&client)
		return Snap7Client{}, ErrIsoConnect
	}

	return Snap7Client{inner: client}, nil
}

func (client *Snap7Client) Close() {
	C.Cli_Disconnect(client.inner)
	C.Cli_Destroy(&client.inner)
}

func (client *Snap7Client) ReadArea(area int, db_number int, start int, amount int) ([]byte, error) {
	buff := make([]byte, amount)

	var err C.int = C.Cli_ReadArea(client.inner, C.int(area), C.int(db_number), C.int(start), C.int(amount), C.S7WLByte, unsafe.Pointer(&buff[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return nil, convertError(int(err))
	}

	return buff, nil
}

func (client *Snap7Client) WriteArea(area int, db_number int, start int, data []byte) error {
	data_len := len(data)

	var err C.int = C.Cli_WriteArea(client.inner, C.int(area), C.int(db_number), C.int(start), C.int(data_len), C.S7WLByte, unsafe.Pointer(&data[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return convertError(int(err))
	}

	return nil
}

func (client *Snap7Client) DBRead(db_number int, start int, size int) ([]byte, error) {
	buff := make([]byte, size)

	var err C.int = C.Cli_DBRead(client.inner, C.int(db_number), C.int(start), C.int(size), unsafe.Pointer(&buff[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return nil, convertError(int(err))
	}

	return buff, nil
}

func (client *Snap7Client) DBWrite(db_number, start int, data []byte) error {
	data_len := len(data)

	var err C.int = C.Cli_DBWrite(client.inner, C.int(db_number), C.int(start), C.int(data_len), unsafe.Pointer(&data[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return convertError(int(err))
	}

	return nil
}

func (client *Snap7Client) MBRead(start int, size int) ([]byte, error) {
	buff := make([]byte, size)

	var err C.int = C.Cli_MBRead(client.inner, C.int(start), C.int(size), unsafe.Pointer(&buff[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return nil, convertError(int(err))
	}

	return buff, nil
}

func (client *Snap7Client) MBWrite(start int, data []byte) error {
	data_len := len(data)

	var err C.int = C.Cli_MBWrite(client.inner, C.int(start), C.int(data_len), unsafe.Pointer(&data[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return convertError(int(err))
	}

	return nil
}

func (client *Snap7Client) EBRead(start int, size int) ([]byte, error) {
	buff := make([]byte, size)

	var err C.int = C.Cli_EBRead(client.inner, C.int(start), C.int(size), unsafe.Pointer(&buff[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return nil, convertError(int(err))
	}

	return buff, nil
}

func (client *Snap7Client) EBWrite(start int, data []byte) error {
	data_len := len(data)

	var err C.int = C.Cli_EBWrite(client.inner, C.int(start), C.int(data_len), unsafe.Pointer(&data[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return convertError(int(err))
	}

	return nil
}

func (client *Snap7Client) ABRead(start int, size int) ([]byte, error) {
	buff := make([]byte, size)

	var err C.int = C.Cli_ABRead(client.inner, C.int(start), C.int(size), unsafe.Pointer(&buff[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return nil, convertError(int(err))
	}

	return buff, nil
}

func (client *Snap7Client) ABWrite(start int, data []byte) error {
	data_len := len(data)

	var err C.int = C.Cli_ABWrite(client.inner, C.int(start), C.int(data_len), unsafe.Pointer(&data[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return convertError(int(err))
	}

	return nil
}

func (client *Snap7Client) TMRead(start int, amount int) ([]byte, error) {
	buff := make([]byte, amount)

	var err C.int = C.Cli_TMRead(client.inner, C.int(start), C.int(amount), unsafe.Pointer(&buff[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return nil, convertError(int(err))
	}

	return buff, nil
}

func (client *Snap7Client) TMWrite(start int, data []byte) error {
	data_len := len(data)

	var err C.int = C.Cli_TMWrite(client.inner, C.int(start), C.int(data_len), unsafe.Pointer(&data[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return convertError(int(err))
	}

	return nil
}

func (client *Snap7Client) CTRead(start int, amount int) ([]byte, error) {
	buff := make([]byte, amount)

	var err C.int = C.Cli_CTRead(client.inner, C.int(start), C.int(amount), unsafe.Pointer(&buff[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return nil, convertError(int(err))
	}

	return buff, nil
}

func (client *Snap7Client) CTWrite(start int, data []byte) error {
	data_len := len(data)

	var err C.int = C.Cli_CTWrite(client.inner, C.int(start), C.int(data_len), unsafe.Pointer(&data[0]))
	if err != 0 {
		C.Cli_Destroy(&client.inner)
		return convertError(int(err))
	}

	return nil
}

type DataTime struct {
	sec  int
	min  int
	hour int
	mday int
	mon  int
	year int
}

func (client *Snap7Client) GetPlcDateTime() (DataTime, error) {
	tm := C.tm{}

	var err C.int = C.Cli_GetPlcDateTime(client.inner, &tm)
	if err != 0 {
		return DataTime{}, convertError(int(err))
	}

	return DataTime{
		sec:  int(tm.tm_sec),
		min:  int(tm.tm_min),
		hour: int(tm.tm_hour),
		mday: int(tm.tm_mday),
		mon:  int(tm.tm_mon) + 1,
		year: int(tm.tm_year) + 1900,
	}, nil
}

func (client *Snap7Client) SetPlcDateTime(time DataTime) error {
	tm := C.tm{}

	tm.tm_sec = C.int(time.sec)
	tm.tm_min = C.int(time.min)
	tm.tm_hour = C.int(time.hour)
	tm.tm_mday = C.int(time.mday)
	tm.tm_mon = C.int(time.mon - 1)
	tm.tm_year = C.int(time.year - 1900)

	var err C.int = C.Cli_SetPlcDateTime(client.inner, &tm)
	if err != 0 {
		return convertError(int(err))
	}

	return nil
}
