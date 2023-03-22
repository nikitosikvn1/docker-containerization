import numpy as np
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

@router.get('/matrix')
def matrix() -> dict:
    matrix_a = np.random.rand(10, 10)
    matrix_b = np.random.rand(10, 10)

    product = np.dot(matrix_a, matrix_b)

    result = {
        "matrix_a": matrix_a.tolist(),
        "matrix_b": matrix_b.tolist(),
        "product": product.tolist()
    }

    return result
