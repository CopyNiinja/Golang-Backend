/*
| Type         | Go Type                  | Ordered? | Typical Use                          |
| ------------ | ------------------------ | -------- | ------------------------------------ |
| **bson.D**   | `[]bson.E`               | Yes    | Filters, queries where order matters |
| **bson.M**   | `map[string]interface{}` |  No     | Simple inserts, updates              |
| **bson.A**   | `[]interface{}`          | N/A      | Arrays (e.g., `$in`, `$or`)          |
| **bson.E**   | struct with Key & Value  |  Yes    | One element of bson.D                |
| **bson.Raw** | `[]byte`                 | N/A      | Advanced / binary BSON data          |

*/
