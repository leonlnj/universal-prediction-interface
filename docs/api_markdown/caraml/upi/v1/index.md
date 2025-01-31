# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [caraml/upi/v1/type.proto](#caraml_upi_v1_type-proto)
    - [Type](#caraml-upi-v1-Type)
  
- [caraml/upi/v1/variable.proto](#caraml_upi_v1_variable-proto)
    - [Variable](#caraml-upi-v1-Variable)
  
- [caraml/upi/v1/observation_log.proto](#caraml_upi_v1_observation_log-proto)
    - [LogObservationsRequest](#caraml-upi-v1-LogObservationsRequest)
    - [LogObservationsResponse](#caraml-upi-v1-LogObservationsResponse)
    - [ObservationLog](#caraml-upi-v1-ObservationLog)
    - [ObservationLogKey](#caraml-upi-v1-ObservationLogKey)
  
    - [ObservationService](#caraml-upi-v1-ObservationService)
  
- [caraml/upi/v1/prediction_log.proto](#caraml_upi_v1_prediction_log-proto)
    - [ModelInput](#caraml-upi-v1-ModelInput)
    - [ModelInput.HeadersEntry](#caraml-upi-v1-ModelInput-HeadersEntry)
    - [ModelOutput](#caraml-upi-v1-ModelOutput)
    - [ModelOutput.HeadersEntry](#caraml-upi-v1-ModelOutput-HeadersEntry)
    - [PredictionLog](#caraml-upi-v1-PredictionLog)
  
- [caraml/upi/v1/table.proto](#caraml_upi_v1_table-proto)
    - [Column](#caraml-upi-v1-Column)
    - [Row](#caraml-upi-v1-Row)
    - [Table](#caraml-upi-v1-Table)
    - [Value](#caraml-upi-v1-Value)
  
- [caraml/upi/v1/upi.proto](#caraml_upi_v1_upi-proto)
    - [ModelMetadata](#caraml-upi-v1-ModelMetadata)
    - [PredictValuesRequest](#caraml-upi-v1-PredictValuesRequest)
    - [PredictValuesResponse](#caraml-upi-v1-PredictValuesResponse)
    - [RequestMetadata](#caraml-upi-v1-RequestMetadata)
    - [ResponseMetadata](#caraml-upi-v1-ResponseMetadata)
    - [TransformerInput](#caraml-upi-v1-TransformerInput)
  
    - [UniversalPredictionService](#caraml-upi-v1-UniversalPredictionService)
  
- [caraml/upi/v1/router_log.proto](#caraml_upi_v1_router_log-proto)
    - [RouterInput](#caraml-upi-v1-RouterInput)
    - [RouterInput.HeadersEntry](#caraml-upi-v1-RouterInput-HeadersEntry)
    - [RouterLog](#caraml-upi-v1-RouterLog)
    - [RouterOutput](#caraml-upi-v1-RouterOutput)
    - [RouterOutput.HeadersEntry](#caraml-upi-v1-RouterOutput-HeadersEntry)
    - [RoutingLogic](#caraml-upi-v1-RoutingLogic)
  
- [Scalar Value Types](#scalar-value-types)



<a name="caraml_upi_v1_type-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## caraml/upi/v1/type.proto


 


<a name="caraml-upi-v1-Type"></a>

### Type
Type supported by UPI

| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_DOUBLE | 1 | Double precision floating number (64-bit) |
| TYPE_INTEGER | 2 | 64-bit Integer |
| TYPE_STRING | 3 | String |


 

 

 



<a name="caraml_upi_v1_variable-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## caraml/upi/v1/variable.proto



<a name="caraml-upi-v1-Variable"></a>

### Variable
Represents a named and typed data point.
Can be used as a prediction input, output or metadata.
Oneof types are avoided as these can be difficult to handle


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name describing what the value represents. Uses include: - Ensuring ML models process columns in the correct order - Defining a Feast row entity name - Parsing metadata to apply traffic rules |
| type | [Type](#caraml-upi-v1-Type) |  | Type of the variable |
| double_value | [double](#double) |  | One of the following field will be set depending on the type |
| integer_value | [int64](#int64) |  |  |
| string_value | [string](#string) |  |  |





 

 

 

 



<a name="caraml_upi_v1_observation_log-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## caraml/upi/v1/observation_log.proto



<a name="caraml-upi-v1-LogObservationsRequest"></a>

### LogObservationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observations | [ObservationLog](#caraml-upi-v1-ObservationLog) | repeated | List of observations per request |






<a name="caraml-upi-v1-LogObservationsResponse"></a>

### LogObservationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observation_batch_id | [string](#string) |  | Unique identifier to identify records from a batch of observation logs |






<a name="caraml-upi-v1-ObservationLog"></a>

### ObservationLog
ObservationLog represents ground truth signals to be combined
with the prediction log produced by CaraML prediction service
to form data sets used for training ML models


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prediction_id | [string](#string) |  | Unique identifier of a prediction response returned by prediction service. This information is used to join the prediction to an observation. |
| row_id | [string](#string) |  | Prediction requests may contain multiple prediction instances. The row_id identifies a particular prediction instance that was used to produce an observation. This information is used to join the prediction to an observation. |
| target_name | [string](#string) |  | The name of the observation target. This information is used to join the prediction to an observation. |
| observation_values | [Variable](#caraml-upi-v1-Variable) | repeated | The ground-truth value. It can be a double, string or integer type. |
| observation_context | [Variable](#caraml-upi-v1-Variable) | repeated | A set of key-value pairs to provide additional context for the observation. |
| observation_timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Timestamp when the observation is made |






<a name="caraml-upi-v1-ObservationLogKey"></a>

### ObservationLogKey
ObservationLogKey contains necessary values for generating unique records for
downstream usages, eg. Dataset Generation Service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observation_batch_id | [string](#string) |  | Id to uniquely identify records from a batch of observation logs |
| prediction_id | [string](#string) |  | Unique identifier of a prediction response returned by prediction service. This information is used to join the prediction to an observation. |
| row_id | [string](#string) |  | Prediction requests may contain multiple prediction instances. The row_id identifies a particular prediction instance that was used to produce an observation. This information is used to join the prediction to an observation. |





 

 

 


<a name="caraml-upi-v1-ObservationService"></a>

### ObservationService
Service for logging observations

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| LogObservations | [LogObservationsRequest](#caraml-upi-v1-LogObservationsRequest) | [LogObservationsResponse](#caraml-upi-v1-LogObservationsResponse) |  |

 



<a name="caraml_upi_v1_prediction_log-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## caraml/upi/v1/prediction_log.proto



<a name="caraml-upi-v1-ModelInput"></a>

### ModelInput
Model input stores information of all input for prediction process.
The information in model input are extracted from the prediction request received by model.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| features_table | [google.protobuf.Struct](#google-protobuf-Struct) |  | JSON-representation of features table. &#34;table_schema_version&#34; field describe the encoding of this field. |
| entities_table | [google.protobuf.Struct](#google-protobuf-Struct) |  | JSON-representation of entities table. &#34;table_schema_version&#34; field describe the encoding of this field. |
| raw_features | [google.protobuf.Struct](#google-protobuf-Struct) |  | JSON-representation of raw_features table. &#34;table_schema_version&#34; field describe the encoding of this field. |
| prediction_context | [Variable](#caraml-upi-v1-Variable) | repeated | Context of the prediction request. |
| headers | [ModelInput.HeadersEntry](#caraml-upi-v1-ModelInput-HeadersEntry) | repeated | map containing request headers/metadata |






<a name="caraml-upi-v1-ModelInput-HeadersEntry"></a>

### ModelInput.HeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="caraml-upi-v1-ModelOutput"></a>

### ModelOutput
Model output stores information of all output produced from prediction process.
Model output is extracted from the prediction response sent by model.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prediction_results_table | [google.protobuf.Struct](#google-protobuf-Struct) |  | JSON-representation of prediction result table. &#34;table_schema_version&#34; field describe the encoding of this field. |
| prediction_context | [Variable](#caraml-upi-v1-Variable) | repeated | Context of the prediction response. |
| headers | [ModelOutput.HeadersEntry](#caraml-upi-v1-ModelOutput-HeadersEntry) | repeated | map containing response headers/metadata |
| status | [uint32](#uint32) |  | grpc status of the response from model (see https://grpc.github.io/grpc/core/md_doc_statuscodes.html) |
| message | [string](#string) |  | grpc message |






<a name="caraml-upi-v1-ModelOutput-HeadersEntry"></a>

### ModelOutput.HeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="caraml-upi-v1-PredictionLog"></a>

### PredictionLog
PredictionLog stores information of prediction request handled by a specific model version.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prediction_id | [string](#string) |  | Unique identifier of prediction. |
| target_name | [string](#string) |  | Target name / concept to be predicted by the prediction. |
| project_name | [string](#string) |  | Project name that host the model performing prediction. |
| model_name | [string](#string) |  | Model name performing the prediction. |
| model_version | [string](#string) |  | Model version performing the prediction. |
| input | [ModelInput](#caraml-upi-v1-ModelInput) |  | Input of the prediction process |
| output | [ModelOutput](#caraml-upi-v1-ModelOutput) |  | Output of the prediction process |
| request_timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Timestamp of the corresponding prediction request |
| table_schema_version | [uint32](#uint32) |  | Schema version of raw_features, features, entities, and prediction results fields are formatted I.e. for version 1, all of those fields will be formatted as a modified JSON SPLIT representation of a table |





 

 

 

 



<a name="caraml_upi_v1_table-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## caraml/upi/v1/table.proto



<a name="caraml-upi-v1-Column"></a>

### Column
Column represent a column definition within a table


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Column&#39;s name |
| type | [Type](#caraml-upi-v1-Type) |  | Column&#39;s type |






<a name="caraml-upi-v1-Row"></a>

### Row
Row represents list of value stored within a row of a table


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| row_id | [string](#string) |  | Id of the particular row in a table. The row id should be at least locally unique within the table. Row ID must be populated for prediction_table |
| values | [Value](#caraml-upi-v1-Value) | repeated | List of values within a row. It is table&#39;s creator responsibility to ensure that the number of entry values matches with the length of columns in the table. |






<a name="caraml-upi-v1-Table"></a>

### Table
Table represents a 2D data structure that has one or more columns 
with potentially different types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Table&#39;s name |
| columns | [Column](#caraml-upi-v1-Column) | repeated | Columns stores schema informations of all columns in the table. |
| rows | [Row](#caraml-upi-v1-Row) | repeated | Rows stores list of row values in the table. |






<a name="caraml-upi-v1-Value"></a>

### Value
Value of a cell within a table. Value is nullable.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| double_value | [double](#double) |  | One of following field will be set depending on the column&#39;s type |
| integer_value | [int64](#int64) |  |  |
| string_value | [string](#string) |  |  |
| is_null | [bool](#bool) |  | Flag to be used to signify that the value is null |





 

 

 

 



<a name="caraml_upi_v1_upi-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## caraml/upi/v1/upi.proto



<a name="caraml-upi-v1-ModelMetadata"></a>

### ModelMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Model name that produce prediction |
| version | [string](#string) |  | Model version that produce prediction |






<a name="caraml-upi-v1-PredictValuesRequest"></a>

### PredictValuesRequest
Represents a request to predict multiple values


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prediction_table | [Table](#caraml-upi-v1-Table) |  | Prediction table contains instances to be predicted. Each row in the table correspond to one prediction instance. Prediction table should contain all preprocessed feature that model use to perform prediction. The column ordering in the prediction table must be the same as feature order expected by model in the case of standard model. Prediction table can be populated via 3 ways: - By performing preprocessing in the client-side and sent as part of original request. - By transforming feature values stored in transformer_inputs. - By retrieving precomputed feature value from feature store. Row ID of the prediction_table must be populated by the client and can be used to join a row in prediction_table with another row in the prediction_result_table, and to track predictions generated by multiple models. The user is expected to include row ID (along with prediction ID) when calling the observations API so that predictions and observations can be joined. NOTE: the ordering of rows might differ in the response but the number of row must remain the same. |
| transformer_input | [TransformerInput](#caraml-upi-v1-TransformerInput) |  | Transformer input contains list of tables and variables that can be used to enrich prediction_table using transformer. Typically transformer_inputs contains: - unprocessed/raw features that requires further transformation. - list of entities for which their precomputed features are retrieved from feature store using standard transformer. |
| target_name | [string](#string) |  | Name of the concept we wish to predict. For example in context of iris classification problem it can be &#34;iris-species&#34; |
| prediction_context | [Variable](#caraml-upi-v1-Variable) | repeated | Prediction context may contain additional data applicable to all prediction instances For example it can be used to store information for traffic rules, experimentation or tracking purposes. Eg. country_code, service_type, service_area_id |
| metadata | [RequestMetadata](#caraml-upi-v1-RequestMetadata) |  | Request metadata |






<a name="caraml-upi-v1-PredictValuesResponse"></a>

### PredictValuesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prediction_result_table | [Table](#caraml-upi-v1-Table) |  | Prediction results corresponding to the prediction rows provided in the request. NOTE: the ordering of prediction_result_rows might differ with prediction_table in the request but the number of row must match with the prediction_table |
| target_name | [string](#string) |  | Target name as defined in the request metadata |
| prediction_context | [Variable](#caraml-upi-v1-Variable) | repeated | Extensible field to cover unforeseen requirements |
| metadata | [ResponseMetadata](#caraml-upi-v1-ResponseMetadata) |  | Response metadata |






<a name="caraml-upi-v1-RequestMetadata"></a>

### RequestMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prediction_id | [string](#string) |  | Unique identifier for each request. Optional. Prediction ID will generated by the platform. The user is expected include the prediction ID (along with row ID) when calling the observations API so that predictions and observations can be joined. Prediction ID is needed because row ID may not be globally unique across requests (only locally unique within each request). If there are experiments with alternative models, the mapping from prediciton ID to treatment ID will be logged by the platform |
| request_timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Timestamp of the request |






<a name="caraml-upi-v1-ResponseMetadata"></a>

### ResponseMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prediction_id | [string](#string) |  | Prediction ID generated by the platform. The user is expected include the prediction ID (along with row ID) when calling the observations API so that predictions and observations can be joined. Prediction ID is needed because row ID may not be globally unique across requests (only locally unique within each request). If there are experiments with alternative models, the mapping from prediciton ID to treatment ID will be logged by the platform |
| models | [ModelMetadata](#caraml-upi-v1-ModelMetadata) | repeated | List of model that produces the prediction This field is repeated to cater for use case such as ensembling several model production results |
| experiment_name | [string](#string) |  | Name of the experiment that is used to produce the response The value might be empty if experimentation is not involved to produce response |
| treatment_name | [string](#string) |  | Name of the treatment chosen from the experiment to produce the response The value might be empty if experimentation is not involved to produce response |






<a name="caraml-upi-v1-TransformerInput"></a>

### TransformerInput
Transformer input contains additional information that can be used to enrich prediction_table using standard transformer.
All tables and variables within transformer input will be imported to the standard transformer runtime automatically.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tables | [Table](#caraml-upi-v1-Table) | repeated | List of tables All tables must have unique name. Each table doesn&#39;t need to have same number of row. |
| variables | [Variable](#caraml-upi-v1-Variable) | repeated | List of variables |





 

 

 


<a name="caraml-upi-v1-UniversalPredictionService"></a>

### UniversalPredictionService
Service for performing model prediction

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PredictValues | [PredictValuesRequest](#caraml-upi-v1-PredictValuesRequest) | [PredictValuesResponse](#caraml-upi-v1-PredictValuesResponse) |  |

 



<a name="caraml_upi_v1_router_log-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## caraml/upi/v1/router_log.proto



<a name="caraml-upi-v1-RouterInput"></a>

### RouterInput
Input received by router.
These informations are extracted from the request received by the router.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prediction_table | [google.protobuf.Struct](#google-protobuf-Struct) |  | JSON-representation of prediction_table of the UPI request to router. &#34;table_schema_version&#34; field describe the encoding of this field. |
| transformer_tables | [google.protobuf.Struct](#google-protobuf-Struct) | repeated | List of tables in the transformer inputs in JSON format. &#34;table_schema_version&#34; field describe the encoding of this field. |
| transformer_variables | [Variable](#caraml-upi-v1-Variable) | repeated | List of variables extracted from &#34;transformer_inputs&#34; |
| prediction_context | [Variable](#caraml-upi-v1-Variable) | repeated | Context of the prediction request. |
| headers | [RouterInput.HeadersEntry](#caraml-upi-v1-RouterInput-HeadersEntry) | repeated | map containing request headers/metadata |






<a name="caraml-upi-v1-RouterInput-HeadersEntry"></a>

### RouterInput.HeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="caraml-upi-v1-RouterLog"></a>

### RouterLog
RouterLog stores information of a multi-model orchestration performed by a router.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prediction_id | [string](#string) |  | Unique identifier of prediction. |
| target_name | [string](#string) |  | Target name / concept to be predicted by the prediction. |
| project_name | [string](#string) |  | Project name that host the router. |
| router_name | [string](#string) |  | Name of the router. |
| router_version | [string](#string) |  | Router version. |
| routing_logic | [RoutingLogic](#caraml-upi-v1-RoutingLogic) |  | Routing logic describes the decision that was made within the router to produce the router output |
| router_input | [RouterInput](#caraml-upi-v1-RouterInput) |  | Input of the router |
| router_output | [RouterOutput](#caraml-upi-v1-RouterOutput) |  | Output of the router |
| request_timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Timestamp of the corresponding prediction request |
| table_schema_version | [uint32](#uint32) |  | Schema version of raw_features, features, entities, and prediction results fields are formatted I.e. for version 1, all of those fields will be formatted as a modified JSON SPLIT representation of a table |






<a name="caraml-upi-v1-RouterOutput"></a>

### RouterOutput
Output produced by router.
These informations are extracted from the response produced by the router.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prediction_results_table | [google.protobuf.Struct](#google-protobuf-Struct) |  | JSON-representation of prediction result table returned by router. &#34;table_schema_version&#34; field describe the encoding of this field. |
| prediction_context | [Variable](#caraml-upi-v1-Variable) | repeated | Context of the prediction response. |
| headers | [RouterOutput.HeadersEntry](#caraml-upi-v1-RouterOutput-HeadersEntry) | repeated | map containing response headers/metadata |
| status | [uint32](#uint32) |  | grpc status of the response from model (see https://grpc.github.io/grpc/core/md_doc_statuscodes.html) |
| message | [string](#string) |  | grpc message |






<a name="caraml-upi-v1-RouterOutput-HeadersEntry"></a>

### RouterOutput.HeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="caraml-upi-v1-RoutingLogic"></a>

### RoutingLogic
Routing logic describes the decision that was made within the router to produce the router output


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| models | [ModelMetadata](#caraml-upi-v1-ModelMetadata) | repeated | List of models that was involved in producing the router output. Router can behave as multiplexer or combiner. In multiplexer case, the router will select 1 out of many models, thus this field will only have 1 entry. In combiner case, this field can contain more than 1 entries. This field will be used to join witht the prediction log. |
| traffic_rule | [string](#string) |  | Traffic rule that was used to route the prediction request (optional). |
| experiment_name | [string](#string) |  | Experiment name that was used to handle the prediction request (optional). |
| treatment_name | [string](#string) |  | Treatment name from the experiment that was used to handle the prediction request (optional). |





 

 

 

 



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

