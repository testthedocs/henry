=====
henry
=====

About
=====

The main idea behind henry is to make it a bit more pleasant and maybe less hard to install `Sphinx <http://www.sphinx-doc.org/en/stable/>`_.

Installing `Sphinx <http://www.sphinx-doc.org/en/stable/>`_ can be overwhelming, especially if you do not use it all the time.

You need to install Python, figure out the whole *virtualenv* thing and such stuff.

Imagine you are a new contributor to a open source project and you worked hard on a piece of *reST* based documentation.

You followed the style-guide and you checked for typos, grammatic, tone, wording and so on, you even used the *reST* preview mode of your editor to get a first impression of how it looks.

**AWESOME !**

But, well before you finally create the *pull request* you want to make sure that it will look **OK** if you build is as HTML with the theme of your documentation.

Maybe you will have to change some headers or add a line break or do something else.

According to the docs of your project, you have to download and install Sphinx into a *virtualenv*, you have to make sure to use the right versions, you need python and so on. **HELP!!??**

"I am not a developer, I work with CSS, all I wanted to do was to improve the docs part about theming" you may think.

**You are right !**

Say hello to henry, henry is here to help you, you do not have to care about all this, henry is doing that for !

Since henry uses `Docker <https://www.docker.com/>`_, you also do not have to mess with your system or waste time on complicated setups.

After the installation of henry, you can review your docs in build HTML with **one** short command:

.. code-block:: shell

   henry html
