=========
Vendoring
=========

.. note::

   We changed from ``govendor`` to ``dep``

Dep
===

More info about `dep <https://golang.github.io/dep/>`_.


Below are the **old** docs about ``govendor``.

Govendor
========

Use: `govendor <https://github.com/kardianos/govendor>`_.

History
-------

.. code-block:: shell

   govendor init
   govendor add +external
   govendor install +local

This in the */root* of the repo

Update All Packages
-------------------

.. code-block:: console

govendor update +v
