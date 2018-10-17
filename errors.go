package snap7

import "errors"

// var ErrClientConnent = errors.New("Client cannot connent")
// var ErrRead = errors.New("Cannot read data")
// var ErrWrite = errors.New("Cannot write data")

// ISO Errors
var ErrIsoConnect = errors.New("ErrIsoConnect")                   // Connection error
var ErrIsoDisconnect = errors.New("ErrIsoDisconnect")             // Disconnect error
var ErrIsoInvalidPDU = errors.New("ErrIsoInvalidPDU")             // Bad format
var ErrIsoInvalidDataSize = errors.New("ErrIsoInvalidDataSize")   // Bad Datasize passed to send/recv buffer is invalid
var ErrIsoNullPointer = errors.New("ErrIsoNullPointer")           // Null passed as pointer
var ErrIsoShortPacket = errors.New("rrIsoShortPacket")            // A short packet received
var ErrIsoTooManyFragments = errors.New("ErrIsoTooManyFragments") // Too many packets without EoT flag
var ErrIsoPduOverflow = errors.New("ErrIsoPduOverflow")           // The sum of fragments data exceded maximum packet size
var ErrIsoSendPacket = errors.New("ErrIsoSendPacket")             // An error occurred during send
var ErrIsoRecvPacket = errors.New("ErrIsoRecvPacket")             // An error occurred during recv
var ErrIsoInvalidParams = errors.New("ErrIsoInvalidParams")       // Invalid TSAP params
var ErrIsoResvd_1 = errors.New("ErrIsoResvd_1")                   // Unassigned
var ErrIsoResvd_2 = errors.New("ErrIsoResvd_2")                   // Unassigned
var ErrIsoResvd_3 = errors.New("ErrIsoResvd_3")                   // Unassigned
var ErrIsoResvd_4 = errors.New("ErrIsoResvd_4")                   // Unassigned

// Error codes
var ErrNEgotiatingPDU = errors.New("ErrNEgotiatingPDU")
var ErrCliInvalidParams = errors.New("ErrCliInvalidParams")
var ErrCliJobPending = errors.New("ErrCliJobPending")
var ErrCliTooManyItems = errors.New("ErrCliTooManyItems")
var ErrCliInvalidWordLen = errors.New("ErrCliInvalidWordLen")
var ErrCliPartialDataWritten = errors.New("ErrCliPartialDataWritten")
var ErrCliSizeOverPDU = errors.New("ErrCliSizeOverPDU")
var ErrCliInvalidPlcAnswer = errors.New("ErrCliInvalidPlcAnswer")
var ErrCliAddressOutOfRange = errors.New("ErrCliAddressOutOfRange")
var ErrCliInvalidTransportSize = errors.New("ErrCliInvalidTransportSize")
var ErrCliWriteDataSizeMismatch = errors.New("ErrCliWriteDataSizeMismatch")
var ErrCliItemNotAvailable = errors.New("ErrCliItemNotAvailable")
var ErrCliInvalidValue = errors.New("ErrCliInvalidValue")
var ErrCliCannotStartPLC = errors.New("ErrCliCannotStartPLC")
var ErrCliAlreadyRun = errors.New("ErrCliAlreadyRun")
var ErrCliCannotStopPLC = errors.New("ErrCliCannotStopPLC")
var ErrCliCannotCopyRamToRom = errors.New("ErrCliCannotCopyRamToRom")
var ErrCliCannotCompress = errors.New("ErrCliCannotCompress")
var ErrCliAlreadyStop = errors.New("ErrCliAlreadyStop")
var ErrCliFunNotAvailable = errors.New("ErrCliFunNotAvailable")
var ErrCliUploadSequenceFailed = errors.New("ErrCliUploadSequenceFailed")
var ErrCliInvalidDataSizeRecvd = errors.New("ErrCliInvalidDataSizeRecvd")
var ErrCliInvalidBlockType = errors.New("ErrCliInvalidBlockType")
var ErrCliInvalidBlockNumber = errors.New("ErrCliInvalidBlockNumber")
var ErrCliInvalidBlockSize = errors.New("ErrCliInvalidBlockSize")
var ErrCliDownloadSequenceFailed = errors.New("ErrCliDownloadSequenceFailed")
var ErrCliInsertRefused = errors.New("ErrCliInsertRefused")
var ErrCliDeleteRefused = errors.New("ErrCliDeleteRefused")
var ErrCliNeedPassword = errors.New("ErrCliNeedPassword")
var ErrCliInvalidPassword = errors.New("ErrCliInvalidPassword")
var ErrCliNoPasswordToSetOrClear = errors.New("ErrCliNoPasswordToSetOrClear")
var ErrCliJobTimeout = errors.New("ErrCliJobTimeout")
var ErrCliPartialDataRead = errors.New("ErrCliPartialDataRead")
var ErrCliBufferTooSmall = errors.New("ErrCliBufferTooSmall")
var ErrCliFunctionRefused = errors.New("ErrCliFunctionRefused")
var ErrCliDestroying = errors.New("ErrCliDestroying")
var ErrCliInvalidParamNumber = errors.New("ErrCliInvalidParamNumber")
var ErrCliCannotChangeParam = errors.New("ErrCliCannotChangeParam")

var ErrUnknow = errors.New("ErrUnknow")

func convertError(error_code int) error {
	switch error_code {
	case 0x00010000:
		return ErrIsoConnect
	case 0x00020000:
		return ErrIsoDisconnect
	case 0x00030000:
		return ErrIsoInvalidPDU
	case 0x00040000:
		return ErrIsoInvalidDataSize
	case 0x00050000:
		return ErrIsoNullPointer
	case 0x00060000:
		return ErrIsoShortPacket
	case 0x00070000:
		return ErrIsoTooManyFragments
	case 0x00080000:
		return ErrIsoPduOverflow
	case 0x00090000:
		return ErrIsoSendPacket
	case 0x000A0000:
		return ErrIsoRecvPacket
	case 0x000B0000:
		return ErrIsoInvalidParams
	case 0x000C0000:
		return ErrIsoResvd_1
	case 0x000D0000:
		return ErrIsoResvd_2
	case 0x000E0000:
		return ErrIsoResvd_3
	case 0x000F0000:
		return ErrIsoResvd_4
	case 0x00100000:
		return ErrNEgotiatingPDU
	case 0x00200000:
		return ErrCliInvalidParams
	case 0x00300000:
		return ErrCliJobPending
	case 0x00400000:
		return ErrCliTooManyItems
	case 0x00500000:
		return ErrCliInvalidWordLen
	case 0x00600000:
		return ErrCliPartialDataWritten
	case 0x00700000:
		return ErrCliSizeOverPDU
	case 0x00800000:
		return ErrCliInvalidPlcAnswer
	case 0x00900000:
		return ErrCliAddressOutOfRange
	case 0x00A00000:
		return ErrCliInvalidTransportSize
	case 0x00B00000:
		return ErrCliWriteDataSizeMismatch
	case 0x00C00000:
		return ErrCliItemNotAvailable
	case 0x00D00000:
		return ErrCliInvalidValue
	case 0x00E00000:
		return ErrCliCannotStartPLC
	case 0x00F00000:
		return ErrCliAlreadyRun
	case 0x01000000:
		return ErrCliCannotStopPLC
	case 0x01100000:
		return ErrCliCannotCopyRamToRom
	case 0x01200000:
		return ErrCliCannotCompress
	case 0x01300000:
		return ErrCliAlreadyStop
	case 0x01400000:
		return ErrCliFunNotAvailable
	case 0x01500000:
		return ErrCliUploadSequenceFailed
	case 0x01600000:
		return ErrCliInvalidBlockType
	case 0x01700000:
		return ErrCliInvalidBlockNumber
	case 0x01800000:
		return ErrCliInvalidBlockNumber
	case 0x01900000:
		return ErrCliInvalidBlockSize
	case 0x01A00000:
		return ErrCliDownloadSequenceFailed
	case 0x01B00000:
		return ErrCliInsertRefused
	case 0x01C00000:
		return ErrCliDeleteRefused
	case 0x01D00000:
		return ErrCliNeedPassword
	case 0x01F00000:
		return ErrCliNoPasswordToSetOrClear
	case 0x02000000:
		return ErrCliJobTimeout
	case 0x02100000:
		return ErrCliPartialDataRead
	case 0x02200000:
		return ErrCliBufferTooSmall
	case 0x02300000:
		return ErrCliFunctionRefused
	case 0x02400000:
		return ErrCliDestroying
	case 0x02500000:
		return ErrCliInvalidParamNumber
	case 0x02600000:
		return ErrCliCannotChangeParam
	default:
		return ErrUnknow
	}
}
