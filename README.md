# ms utilities

This is a collection of utilities for dealing with things Microsot does differently than other standards or projects, and for dealing with Microsoft specific data. This repository is **not** owned by, endorced by, represented by, or related to Microsoft in any way.

## Package List

### guid
 
This package contains a struct to hold a Microsoft GUID, which is defined as the same 128 bit number as a UUID but stored in a binary format with Big Endian byte ordering instead of Little Endian. If you use an Active Directory or SQL Server library that does returns a raw byte slice, you might need to re-order some of the bytes.

Libraries such as [jtblin/go-ldap-client](https://github.com/jtblin/go-ldap-client) will return an object GUID in the Microsoft binary storage format. For example, if the `Attributes` field of the `LDAPClient` looks like `Attributes: []string{"sAMAccountName", "objectGUID", "userPrincipalName", "cn", "department"}` and you want to get a UUID in a library like [satori/go.uuid](https://github.com/satori/go.uuid) then you'll want to do the following:

```
ok, user, err := client.Authenticate(username, password)
if err != nil {
  log.Printf("Error authenticating user %s: %+v", username,
    strings.Trim(err.Error(), "\x00")) //trim null terminator in C-string
  return
}
defer client.Close()
if !ok {
  log.Printf("Authenticating failed for user %s", username)
  return
}

// Fix endianness
g, err := guid.ParseRawString(user["objectGUID"])
if err != nil {
  log.Printf("Could not parse guid slice from AD %v into guid object", user["objectGUID"])
  return
}
guid, err := uuid.FromBytes(g.UUIDBytes())
```

One could accomplish the same think for **SQL Server** data returning the UNIQUEIDENTIFIER in the raw binary format (the `Guid` class in C# handles this is .NET code) but I prefer to solve the issue with casting the result column to a VARCHAR if I'm able to modify the stored procedure. Again, any .NET code will be happy with either data type. Some drivers for SQL Server might eventually make this library redundant, but at the time of this writing, [denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb) does not.

## Testing

Switch to the package directory desired and enter the following:

```
go test
```

Remove build artifacts (and anything else matched by the gitignore):

```
git clean -Xf
```