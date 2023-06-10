# GetBranchStatusResponse


## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ContentType`                                                                                                | *string*                                                                                                     | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `StatusCode`                                                                                                 | *int*                                                                                                        | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `RawResponse`                                                                                                | [*http.Response](https://pkg.go.dev/net/http#Response)                                                       | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |
| `GetBranchStatus200ApplicationJSONObject`                                                                    | map[string][GetBranchStatus200ApplicationJSON](../../models/operations/getbranchstatus200applicationjson.md) | :heavy_minus_sign:                                                                                           | Returns the branch status                                                                                    |