===========
Development
===========

.. warning::

   These docs are WIP !

.. toctree::
   :maxdepth: 2
   :hidden:

   vendoring
   release

Internals
=========

henry is a wrapper on steroids around `Sphinx <http://www.sphinx-doc.org/en/stable/index.html>`_ and `Docker <https://www.docker.com/>`_ written in `golang <https://golang.org/>`_.

The main motivation behind henry is to make it less hassle as possible to get a visual impression
of your docs.

henry uses `Cobra <https://github.com/spf13/cobra>`_ as base.

  Cobra is a library providing a simple interface to create powerful modern CLI interfaces similar to git & go tools.

We use Cobra to run certain (Sphinx) commands in Docker.


