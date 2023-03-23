=================
Simple golang app
=================

This is a starter project for the third laboratory assignment on containerization.

How to run
==========

You will need `golang <https://go.dev/doc/install>`_ toolchain installed in
order to compile this project.

Build the project with the following command:

.. code-block:: console

   solo@falcon ~/project $ go build -o build/fizzbuzz

When you have the project built, you can run it by invoking the binary:

.. code-block:: console

   solo@falcon ~/project $ ./build/fizzbuzz
   Usage:
     fizzbuzz [command]

   Available Commands:
     completion  Generate the autocompletion script for the specified shell
     help        Help about any command
     query       Query if the number is fizz/buzz or fizzbuzz
     serve       Run an http server to anser fizzbuzz queries

   Flags:
     -h, --help   help for fizzbuzz

   solo@falcon ~/project $ ./build/fizzbuzz serve
   Listening on http://0.0.0.0:8080

When serving a project, press Ctrl+C to sent SIGTERM signal to the running
server and bring it down.
