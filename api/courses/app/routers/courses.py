from fastapi import APIRouter

# import boto3
# from boto3.dynamodb.conditions import Key, Attr

router = APIRouter()
# dynamodb = boto3.resource("dynamodb")
# dictionary_table = dynamodb.Table("Dictionary")


@router.get("/courses")
async def list_courses():
    # response = dictionary_table.query(KeyConditionExpression=Key("pk").eq("country"))
    # items_raw = response["Items"]
    items = [{"id": "1234", "name": "Kubernetes course"}]

    # for item in items_raw:
    #     items.append(
    #         {
    #             "code": item["sk"],
    #             "name": item.get("name", ""),
    #             "native": item.get("native", ""),
    #             "phone": item.get("phone", ""),
    #         }
    #     )

    return items