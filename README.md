# ms utilities

This is a collection of utilities for dealing with things Microsot does differently than other standards or projects, and for dealing with Microsoft specific data. This repository is **not** owned by, endorced by, represented by, or related to Microsoft in any way.

## Package List

 - **guid**: This package contains a struct to hold a Microsoft GUID, which is defined as the same 128 bit number as a UUID but stored in 
 a binary format with Big Endian byte ordering instead of Little Endian. If you use an Active Directory or SQL Server library that does 
 returns a raw byte slice, you might need to re-order some of the bytes.

## Testing

Switch to the package directory desired and enter the following:

```
go test
```

Remove build artifacts (and anything else matched by the gitignore):

```
git clean -Xf
```