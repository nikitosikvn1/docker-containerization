=================
Simple Rust HTTP server
=================

This is a simple single-threaded http server written in Rust that returns the current time and a randomly generated password.

How to run
==========

You will need `Rust <https://www.rust-lang.org/tools/install>`_ toolchain installed in
order to compile this project.

If you are using Linux or macOS:

.. code-block:: console
  
  $ curl --proto '=https' --tlsv1.3 https://sh.rustup.rs -sSf | sh

If you are using Windows then install Linux.

Build the project with the following command:

.. code-block:: console

  $ cargo build

When you have the project built, you can run it by invoking the binary:

.. code-block:: console
  $ ./target/debug/rust-project
  
  Listening on http://0.0.0.0:8080

When serving a project, press Ctrl+C to sent SIGTERM signal to the running
server and bring it down.
