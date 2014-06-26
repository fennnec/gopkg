// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package las

/*
Las Header

	Change list:
	+--------------------+-----------------------+-------------------------+
	|       1.0          |       1.1             |          1.2            |
	+--------------------+-----------------------+-------------------------+
	|        ........    |        ........       |        ........         |
	| uint32 reserved    | uint16 file_source_id | uint16 file_source_id   |
	|        --------    |                       |                         |
	|        --------    | uint16 reserved       | uint16 global_encoding  |
	|        --------    |        --------       |                         |
	|        ........    |        ........       |        ........         |
	| uint16 julian_date | uint16 day_of_year    | uint16 day_of_year      |
	|        ........    |        ........       |        ........         |
	+--------------------+-----------------------+-------------------------+
*/
type Header struct {
	FileSignature                 [4]byte // LASF
	FileSourceId                  uint16
	GlobalEncoding                uint16
	GuidData1                     uint32
	GuidData2                     uint16
	GuidData3                     uint16
	GuidData4                     [8]uint8
	MajorVersion                  int8
	MinorVersion                  int8
	SystemIdentifier              [32]byte
	GeneratingSoftware            [32]byte
	FileCreationDay               uint16
	FileCreationYear              uint16
	HeaderSize                    uint16
	OffsetToPointData             uint32
	NumberOfVariableLengthRecords uint32
	PointDataFormatId             uint8
	PointDataRecordLen            uint16
	NumberOfPointRecords          uint32
	NumberOfPointsByReturn        [5]uint32
	XScaleFactor                  float64
	YScaleFactor                  float64
	ZScaleFactor                  float64
	XOffset                       float64
	YOffset                       float64
	ZOffset                       float64
	MaxX                          float64
	MinX                          float64
	MaxY                          float64
	MinY                          float64
	MaxZ                          float64
	MinZ                          float64
}

// Las Header V14
type HeaderV14 struct {
	Header
	OffsetWaveform uint64
}

type VLRRecordIdType uint16

const (
	VLRRecordId_ClassificationLookup VLRRecordIdType = 0
	VLRRecordId_FlightLineLookup     VLRRecordIdType = 1
	VLRRecordId_Histogram            VLRRecordIdType = 2
	VLRRecordId_TextAreaDesc         VLRRecordIdType = 3
	VLRRecordId_GeoKeyDirectory      VLRRecordIdType = 34735
	VLRRecordId_GeoDoubleParam       VLRRecordIdType = 34735
	VLRRecordId_GeoAsciiParam        VLRRecordIdType = 34737
)

type VLR struct {
	Reserved                uint16
	UserId                  [16]int8
	RecordId                uint16
	RecordLengthAfterHeader uint16
	Description             [32]byte
	Data                    []byte
}

type VLR_GeoKeysEntry struct {
	KeyId           uint16
	TiffTagLocation uint16
	Count           uint16
	ValueOffset     uint16
}

type Point struct {
	X, Y, Z       float64
	Intensity     uint16
	BitFields     uint16
	ScanAngleRank int8
	UserData      uint8
	PointSourceId uint16
}
