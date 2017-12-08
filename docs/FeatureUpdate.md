# FeatureUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the feature | [optional] [default to null]
**Description** | **string** | Description of the feature and it can include HTML formatting. | [optional] [default to null]
**CreatedBy** | **string** | Email address of user that created the feature. | [optional] [default to null]
**AssignedToUser** | **string** | Email address of user that is assigned the feature. | [optional] [default to null]
**Tags** | **string** | Tags can be automatically assigned to the new feature. If more than one tag is used then tags should be separated by commas | [optional] [default to null]
**OriginalEstimateText** | **string** | Set the original estimated effort in a text format, you can use d, h, min (or &#39;p&#39; for points) to indicate the units to use. | [optional] [default to null]
**RemainingEstimateText** | **string** |  Set the remaining estimated effort in a text format, you can use d, h, min (or &#39;p&#39; for points) to indicate the units to use. | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Date that work will start on the feature in format YYYY-MM-DD. | [optional] [default to null]
**DueDate** | [**time.Time**](time.Time.md) | Date that work is due to be completed on the feature in format YYYY-MM-DD. | [optional] [default to null]
**ReleasePhase** | **string** | Name or id of release phase which the feature belongs to. | [optional] [default to null]
**Initiative** | **string** | Name or id of initiative which the feature belongs to. | [optional] [default to null]
**MasterFeature** | **string** | Name or id of master feature which the feature belongs to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


