=======================
Simple web app (python)
=======================

This is a starter project for the third laboratory assignment on containerization.
This project is written in `Python <https://www.python.org/>`_ programming
language and uses `FastAPI <https://fastapi.tiangolo.com/>`_ framework.

How to run
==========

As the project is written in Python, you will need Python installed.
Recommended version of Python is 3.7 and above.

Firstly, prepare the environment:

.. code-block:: console

   solo@falcon ~/project $ python -m venv ./.venv

As virtual environment is created for you, enter it and install project
dependencies:

.. code-block:: console

   solo@falcon ~/project $ . ./.venv/bin/acticate
   (.venv) solo@falcon ~/project $ pip install -r requirements/backend.in
   Collecting fastapi
   ...
   Installing collected packages: sniffio, idna ...
   Successfully installed PyYAML-6.0 ...

As the dependencies are installed you can start the project using the following
command:

.. code-block:: console

   (.venv) solo@falcon ~/project $ uvicorn spaceship.main:app --host=0.0.0.0 --port=8080
   INFO:     Started server process [16689]
   INFO:     Waiting for application startup.
   INFO:     Application startup complete.

Open http://127.0.0.1:8080/ in your browser to see the main project page.
