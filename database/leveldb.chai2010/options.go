// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

const (
	DefaultWriteBufferSize      = 4 << 20
	DefaultMaxOpenFiles         = 1000
	DefaultBlockCacheSize       = 8 << 20
	DefaultBlockSize            = 4096
	DefaultBlockRestartInterval = 16
	DefaultCompressionType      = SnappyCompression
)

type OptionsFlag uint

const (
	// If set, the database will be created if it is missing.
	OFCreateIfMissing OptionsFlag = 1 << iota

	// If set, an error is raised if the database already exists.
	OFErrorIfExist

	// If set, the implementation will do aggressive checking of the
	// data it is processing and will stop early if it detects any
	// errors.  This may have unforeseen ramifications: for example, a
	// corruption of one DB entry may cause a large number of entries to
	// become unreadable or for the entire DB to become unopenable.
	OFStrict
)

// Database compression type
type Compression int

func (c Compression) String() string {
	switch c {
	case DefaultCompression:
		return "leveldb.Compression: default(snappy)"
	case NoCompression:
		return "leveldb.Compression: none"
	case SnappyCompression:
		return "leveldb.Compression: snappy"
	}
	return "leveldb.Compression: unknown"
}

const (
	DefaultCompression Compression = iota
	NoCompression
	SnappyCompression
	nCompression
)

// Options to control the behavior of a database (passed to DB::Open)
type Options struct {
	// Comparator used to define the order of keys in the table.
	// Default: a comparator that uses lexicographic byte-wise ordering
	//
	// REQUIRES: The client must ensure that the comparator supplied
	// here has the same name and orders keys *exactly* the same as the
	// comparator provided to previous open calls on the same DB.
	Comparator Comparator

	// Specify the database flag.
	Flag OptionsFlag

	// Amount of data to build up in memory (backed by an unsorted log
	// on disk) before converting to a sorted on-disk file.
	//
	// Larger values increase performance, especially during bulk loads.
	// Up to two write buffers may be held in memory at the same time,
	// so you may wish to adjust this parameter to control memory usage.
	// Also, a larger write buffer will result in a longer recovery time
	// the next time the database is opened.
	//
	// Default: 4MB
	WriteBufferSize int

	// Number of open files that can be used by the DB.  You may need to
	// increase this if your database has a large working set (budget
	// one open file per 2MB of working set).
	//
	// Default: 1000
	MaxOpenFiles int

	// Control over blocks (user data is stored in a set of blocks, and
	// a block is the unit of reading from disk).
	//
	// If non-nil, use the specified cache for blocks.
	// If nil, leveldb will automatically create and use an 8MB internal cache.
	//
	// Default: nil
	BlockCache *Cache

	// Approximate size of user data packed per block.  Note that the
	// block size specified here corresponds to uncompressed data.  The
	// actual size of the unit read from disk may be smaller if
	// compression is enabled.  This parameter can be changed dynamically.
	//
	// Default: 4K
	BlockSize int

	// Number of keys between restart points for delta encoding of keys.
	// This parameter can be changed dynamically.  Most clients should
	// leave this parameter alone.
	//
	// Default: 16
	BlockRestartInterval int

	// Compress blocks using the specified compression algorithm.  This
	// parameter can be changed dynamically.
	//
	// Default: kSnappyCompression, which gives lightweight but fast
	// compression.
	//
	// Typical speeds of kSnappyCompression on an Intel(R) Core(TM)2 2.4GHz:
	//    ~200-500MB/s compression
	//    ~400-800MB/s decompression
	// Note that these speeds are significantly faster than most
	// persistent storage speeds, and therefore it is typically never
	// worth switching to kNoCompression.  Even if the input data is
	// incompressible, the kSnappyCompression implementation will
	// efficiently detect that and will switch to uncompressed mode.
	Compression Compression

	// If non-NULL, use the specified filter policy to reduce disk reads.
	// Many applications will benefit from passing the result of
	// NewBloomFilterPolicy() here.
	//
	// Default: NULL
	Filter Filter
}

func (o *Options) GetComparator() Comparator {
	return o.Comparator
}
func (o *Options) SetComparator(cmp Comparator) {
	o.Comparator = cmp
}

func (o *Options) HasFlag(flag OptionsFlag) bool {
	return (o.Flag & flag) != 0
}
func (o *Options) SetFlag(flag OptionsFlag) {
	o.Flag |= flag
}
func (o *Options) ClearFlag(flag OptionsFlag) {
	o.Flag &= ^flag
}

func (o *Options) GetWriteBufferSize() int {
	if o.WriteBufferSize <= 0 {
		return DefaultWriteBufferSize
	}
	return o.WriteBufferSize
}
func (o *Options) SetWriteBufferSize(size int) {
	o.WriteBufferSize = size
}

func (o *Options) GetMaxOpenFiles() int {
	if o.MaxOpenFiles <= 0 {
		return DefaultMaxOpenFiles
	}
	return o.MaxOpenFiles
}
func (o *Options) SetMaxOpenFiles(max int) {
	o.MaxOpenFiles = max
}

func (o *Options) GetBlockCache() *Cache {
	return o.BlockCache
}
func (o *Options) SetBlockCache(cache *Cache) {
	o.BlockCache = cache
}
func (o *Options) SetBlockCacheCapacity(capacity int64) {
	o.BlockCache = NewCache(capacity)
}

func (o *Options) GetBlockSize() int {
	if o.BlockSize <= 0 {
		return DefaultBlockSize
	}
	return o.BlockSize
}
func (o *Options) SetBlockSize(size int) {
	o.BlockSize = size
}

func (o *Options) GetBlockRestartInterval() int {
	if o.BlockRestartInterval <= 0 {
		return DefaultBlockRestartInterval
	}
	return o.BlockRestartInterval
}
func (o *Options) SetBlockRestartInterval(interval int) {
	o.BlockRestartInterval = interval
}

func (o *Options) GetCompression() Compression {
	if o.Compression <= DefaultCompression || o.Compression >= nCompression {
		return DefaultCompressionType
	}
	return o.Compression
}
func (o *Options) SetCompression(compression Compression) {
	o.Compression = compression
}

func (o *Options) GetFilter() Filter {
	return o.Filter
}
func (o *Options) SetFilter(p Filter) {
	o.Filter = p
}

type ReadOptionsFlag uint

const (
	// If true, all data read from underlying storage will be
	// verified against corresponding checksums.
	RFVerifyChecksums ReadOptionsFlag = 1 << iota

	// Should the data read for this iteration be cached in memory?
	// If set iteration chaching will be disabled.
	// Callers may wish to set this flag for bulk scans.
	RFDontFillCache
)

// ReadOptions represent sets of options used by LevelDB during read
// operations.
type ReadOptions struct {
	// Specify the read flag.
	Flag ReadOptionsFlag

	// If "snapshot" is non-NULL, read as of the supplied snapshot
	// (which must belong to the DB that is being read and which must
	// not have been released).  If "snapshot" is NULL, use an impliicit
	// snapshot of the state at the beginning of this read operation.
	Snapshot *Snapshot
}

func (o *ReadOptions) HasFlag(flag ReadOptionsFlag) bool {
	return (o.Flag & flag) != 0
}
func (o *ReadOptions) SetFlag(flag ReadOptionsFlag) {
	o.Flag |= flag
}
func (o *ReadOptions) ClearFlag(flag ReadOptionsFlag) {
	o.Flag &= ^flag
}

func (o *ReadOptions) GetSnapshot() *Snapshot {
	return o.Snapshot
}
func (o *ReadOptions) SetSnapshot(snap *Snapshot) {
	o.Snapshot = snap
}

type WriteOptionsFlag uint

const (
	// If set, the write will be flushed from the operating system
	// buffer cache (by calling WritableFile::Sync()) before the write
	// is considered complete.  If this flag is true, writes will be
	// slower.
	//
	// If this flag is false, and the machine crashes, some recent
	// writes may be lost.  Note that if it is just the process that
	// crashes (i.e., the machine does not reboot), no writes will be
	// lost even if sync==false.
	//
	// In other words, a DB write with sync==false has similar
	// crash semantics as the "write()" system call.  A DB write
	// with sync==true has similar crash semantics to a "write()"
	// system call followed by "fsync()".
	WFSync WriteOptionsFlag = 1 << iota
)

// WriteOptions represent sets of options used by LevelDB during write
// operations.
type WriteOptions struct {
	// Specify the write flag.
	Flag WriteOptionsFlag
}

func (o *WriteOptions) HasFlag(flag WriteOptionsFlag) bool {
	return (o.Flag & flag) != 0
}
func (o *WriteOptions) SetFlag(flag WriteOptionsFlag) {
	o.Flag |= flag
}
func (o *WriteOptions) ClearFlag(flag WriteOptionsFlag) {
	o.Flag &= ^flag
}
