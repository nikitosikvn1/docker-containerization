from fastapi import APIRouter

router = APIRouter()


@router.get('')
def hello_world() -> dict:
    data = {
        "name": "Nikita",
        "second_name": "Petrykin",
        "age": 19,
        "hobby": ["Powerlifting", "Volleyball", "Biking"],
        "faculty": "Faculty of Informatics and Computer Engineering",
        "group": "IM-12"
    }

    return data
