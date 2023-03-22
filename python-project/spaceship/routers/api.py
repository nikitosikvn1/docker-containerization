from fastapi import APIRouter

router = APIRouter()


@router.get('')
def hello_world() -> dict:
    data = {
        "name": "Nikita",
        "second_name": "Petrykin",
        "age": 19,
        "hobby": ["Powerlifting", "Volleyball", "Biking"]
    }
    
    return data
