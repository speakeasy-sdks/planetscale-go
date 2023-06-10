# CreateADeployRequestResponse


## Fields

| Field                                                                                                                  | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ContentType`                                                                                                          | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `StatusCode`                                                                                                           | *int*                                                                                                                  | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `RawResponse`                                                                                                          | [*http.Response](https://pkg.go.dev/net/http#Response)                                                                 | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `CreateADeployRequest200ApplicationJSONObject`                                                                         | map[string][CreateADeployRequest200ApplicationJSON](../../models/operations/createadeployrequest200applicationjson.md) | :heavy_minus_sign:                                                                                                     | Returns the created deploy request                                                                                     |