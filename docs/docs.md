# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [pim-service/pim-service.proto](#pim-service/pim-service.proto)
    - [CategoriesTrees](#pim_service.CategoriesTrees)
    - [Category](#pim_service.Category)
    - [IDs](#pim_service.IDs)
    - [Product](#pim_service.Product)
    - [Product.Param](#pim_service.Product.Param)
    - [ProductIDs](#pim_service.ProductIDs)
    - [Products](#pim_service.Products)
    - [ScrollDescriptor](#pim_service.ScrollDescriptor)
    - [SearchRequest](#pim_service.SearchRequest)
    - [SearchRequest.Filters](#pim_service.SearchRequest.Filters)
    - [ShopID](#pim_service.ShopID)
  
    - [PimService](#pim_service.PimService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="pim-service/pim-service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pim-service/pim-service.proto



<a name="pim_service.CategoriesTrees"></a>

### CategoriesTrees
Деревья категорий товаров


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| categories | [Category](#pim_service.Category) | repeated | список категорий товаров |






<a name="pim_service.Category"></a>

### Category
Категория товаров


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int32](#int32) |  | id категории товара |
| ParentID | [int32](#int32) |  | id родительской категории (если 0 или отсутствует, то элемент корневой) |
| shopID | [int32](#int32) |  | id магазина |
| name | [string](#string) |  | название категории |
| children | [Category](#pim_service.Category) | repeated | дочерние категории |






<a name="pim_service.IDs"></a>

### IDs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int32](#int32) |  | id |
| shopID | [int32](#int32) |  | id магазина |






<a name="pim_service.Product"></a>

### Product
Товар


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int32](#int32) |  | сквозной id товара |
| itemID | [string](#string) |  | id товара в магазине |
| shopID | [int32](#int32) |  | id магазина |
| name | [string](#string) |  | наименование товара |
| available | [bool](#bool) |  | доступен ли товар |
| type | [string](#string) |  | тип поставки |
| url | [string](#string) |  | ссылка на товар в магазине |
| price | [int32](#int32) |  | цена |
| vendor | [string](#string) |  | поставщик |
| description | [string](#string) |  | описание товара |
| currencyID | [string](#string) |  | id валюты |
| params | [Product.Param](#pim_service.Product.Param) | repeated | атрибуты товара |






<a name="pim_service.Product.Param"></a>

### Product.Param
Атрибут товара


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | имя атрибута |
| type | [string](#string) |  | тип атрибута |
| value | [google.protobuf.Struct](#google.protobuf.Struct) |  | значение атрибута |






<a name="pim_service.ProductIDs"></a>

### ProductIDs
Запрос товара


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  | id товара |
| shopID | [int32](#int32) |  | id магазина |






<a name="pim_service.Products"></a>

### Products
Cписок товаров


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [Product](#pim_service.Product) | repeated | список товаров |
| meta | [ScrollDescriptor](#pim_service.ScrollDescriptor) |  | опции постраничной выдачи |






<a name="pim_service.ScrollDescriptor"></a>

### ScrollDescriptor
Опции постраничной выдачи


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) |  | максимальное количество элементов в выдаче |
| offset | [int32](#int32) |  | смещение (страница) выдачи |






<a name="pim_service.SearchRequest"></a>

### SearchRequest
Поисковый запрос


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shopID | [int32](#int32) |  | id магазина |
| filters | [SearchRequest.Filters](#pim_service.SearchRequest.Filters) |  | фильтры поиска |
| searchTerm | [string](#string) |  | поисковый запрос |
| withParams | [bool](#bool) |  | показывать ли атрибуты товаров |
| meta | [ScrollDescriptor](#pim_service.ScrollDescriptor) |  | опции постраничной выдачи |






<a name="pim_service.SearchRequest.Filters"></a>

### SearchRequest.Filters
Фильтры поиска


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| available | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | доступен ли товар |
| categoriesIDs | [int32](#int32) | repeated | id категорий товаров |






<a name="pim_service.ShopID"></a>

### ShopID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shopID | [int32](#int32) |  | id магазина |





 

 

 


<a name="pim_service.PimService"></a>

### PimService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SearchProducts | [SearchRequest](#pim_service.SearchRequest) | [Products](#pim_service.Products) | нечеткий поиск по продуктам магазина с фильтрами |
| GetAllCategoriesByShop | [ShopID](#pim_service.ShopID) | [CategoriesTrees](#pim_service.CategoriesTrees) | получить все категории товаров магазина |
| CreateProduct | [Product](#pim_service.Product) | [.google.protobuf.Empty](#google.protobuf.Empty) | создать товар для магазина |
| GetProduct | [ProductIDs](#pim_service.ProductIDs) | [Product](#pim_service.Product) | получить товар магазина по id |
| UpdateProduct | [Product](#pim_service.Product) | [.google.protobuf.Empty](#google.protobuf.Empty) | обновить товар |
| DeleteProduct | [ProductIDs](#pim_service.ProductIDs) | [.google.protobuf.Empty](#google.protobuf.Empty) | удалить товар |
| CreateCategory | [Category](#pim_service.Category) | [.google.protobuf.Empty](#google.protobuf.Empty) | создать категорию товаров для магазина |
| GetCategory | [IDs](#pim_service.IDs) | [Category](#pim_service.Category) | получить категорию по id |
| UpdateCategory | [Category](#pim_service.Category) | [.google.protobuf.Empty](#google.protobuf.Empty) | обновить категорию товаров |
| DeleteCategory | [IDs](#pim_service.IDs) | [.google.protobuf.Empty](#google.protobuf.Empty) | удалить категорию товаров |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

